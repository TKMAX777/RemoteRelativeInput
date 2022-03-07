package windows

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/TKMAX777/RemoteRelativeInput/keymap"
	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
	"github.com/TKMAX777/RemoteRelativeInput/winapi"
	"github.com/pkg/errors"

	"github.com/lxn/win"
)

type Handler struct {
	metrics SystemMetrics

	remote *remote_send.Handler

	logger logger
	debug  bool
}

type SystemMetrics struct {
	FrameWidthX int32
	FrameWidthY int32
	TitleHeight int32
}

func New(r *remote_send.Handler) *Handler {
	return &Handler{
		remote: r,
		metrics: SystemMetrics{
			FrameWidthX: win.GetSystemMetrics(win.SM_CXSIZEFRAME),
			FrameWidthY: win.GetSystemMetrics(win.SM_CYSIZEFRAME),
			TitleHeight: win.GetSystemMetrics(win.SM_CYCAPTION),
		},
	}
}

func (h *Handler) SetLogger(l logger) {
	h.debug = true
	h.logger = l
}

// Lock cursor to the window center and send cursor movement
func (h *Handler) SendCursor(hwnd win.HWND) error {
	if hwnd == win.HWND(winapi.NULL) {
		return errors.New("NilWindowHandler")
	}

	var rect win.RECT
	if !winapi.GetWindowRect(hwnd, &rect) {
		return errors.New("GetWindowRectError")
	}

	var windowCenterPosition = h.getWindowCenterPos(rect)

	if !win.SetCursorPos(windowCenterPosition.X, windowCenterPosition.Y) {
		return errors.New("Error in set cursor position")
	}

	okInt, _ := winapi.ClipCursor(rect)
	if okInt != 1 {
		return errors.Errorf("Error in clip cursor: code: %d\n", win.GetLastError())
	}

	go h.pointLoop(windowCenterPosition)

	return nil
}

func (h *Handler) getWindowCenterPos(rect win.RECT) win.POINT {
	var windowCenterPosition win.POINT
	windowCenterPosition.X = int32(rect.Left+rect.Right) / 2
	windowCenterPosition.Y = int32(rect.Top+rect.Bottom) / 2

	return windowCenterPosition
}

// lock cursor and send cursor movement to the channel
func (h *Handler) pointLoop(windowCenterPosition win.POINT) {
	var pos POINT
	var currentPosition = windowCenterPosition

	for {
		ok := winapi.GetCursorPos(&currentPosition)
		if !ok {
			log.Printf("Error in get cursor position\n")
		}

		// Relative position mode
		pos.POINT = win.POINT{X: currentPosition.X - windowCenterPosition.X, Y: currentPosition.Y - windowCenterPosition.Y}
		if pos.X != 0 || pos.Y != 0 {
			h.Debugf("X: %4d Y: %4d\n", pos.X, pos.Y)
			h.remote.SendRelativeCursor(pos)
		}

		ok = win.SetCursorPos(windowCenterPosition.X, windowCenterPosition.Y)
		if !ok {
			log.Printf("Error in set cursor position")
		}

		time.Sleep(time.Millisecond * 1)
	}
}

// Create window on remote desktop client
func (h *Handler) CreateWindow(rdClientHwnd win.HWND) (win.HWND, error) {
	type resultAttr struct {
		hwnd win.HWND
		err  error
	}

	var result = make(chan resultAttr)

	go func() {
		const windowName = "RDP Input Wrapper"

		// make win main function
		var hInstance = win.GetModuleHandle(nil)
		if hInstance == win.HINSTANCE(0) {
			result <- resultAttr{win.HWND(winapi.NULL), errors.Errorf("GetModuleHandle: Failed to get handler: %d\n", win.GetLastError())}
		}

		var className = winapi.MustUTF16PtrFromString(windowName)

		var wClass win.WNDCLASSEX
		wClass = win.WNDCLASSEX{
			CbSize:    uint32(unsafe.Sizeof(wClass)),
			HInstance: hInstance,

			// redraw when the window size changed
			Style: win.CS_HREDRAW | win.CS_VREDRAW,

			// set window background color to white
			HbrBackground: win.HBRUSH(win.GetStockObject(win.WHITE_BRUSH)),
			LpszClassName: className,
			LpfnWndProc:   syscall.NewCallback(h.windowProc(rdClientHwnd)),
			LpszMenuName:  nil,

			CbClsExtra: 0,
			CbWndExtra: 0,

			HIcon:   win.HICON(winapi.NULL),
			HCursor: win.HCURSOR(winapi.NULL),
			HIconSm: win.HICON(winapi.NULL),
		}

		if win.RegisterClassEx(&wClass) == 0 {
			result <- resultAttr{win.HWND(winapi.NULL), errors.Errorf("RegisterClassEx: Failed to make window class %v\n", win.GetLastError())}
		}

		var windowNameUTF16 = winapi.MustUTF16PtrFromString(windowName)
		var hwnd = win.CreateWindowEx(
			win.WS_EX_OVERLAPPEDWINDOW|win.WS_EX_TOPMOST|win.WS_EX_LAYERED,
			className,
			windowNameUTF16,
			win.WS_OVERLAPPEDWINDOW,
			win.CW_USEDEFAULT, win.CW_USEDEFAULT, int32(100), int32(100),
			win.HWND(winapi.NULL), win.HMENU(winapi.NULL), hInstance, unsafe.Pointer(nil),
		)

		if hwnd == win.HWND(winapi.NULL) {
			result <- resultAttr{hwnd, errors.Errorf("CreateWindowEx: Failed to make window")}
		}

		winapi.ShowWindow(hwnd, win.SW_SHOW)
		winapi.UpdateWindow(hwnd)

		var windowProc = h.getMessage()

		result <- resultAttr{hwnd, nil}

		for {
			var msg win.MSG
			switch win.GetMessage(&msg, hwnd, 0, 0) {
			case 0:
				h.Debugln("Quit")
				return
			case -1:
				fmt.Fprintln(os.Stderr, "GetMessageErrorOccured")
				return
			}

			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
			h.Output(10, fmt.Sprintf("disp: %+v \n", msg))
			windowProc(msg)
		}
	}()

	var res = <-result

	close(result)

	return res.hwnd, res.err
}

func (h Handler) getMessage() func(msg win.MSG) error {
	// var send = func(key keymap.WindowsKey, state InputType) {
	// 	fmt.Printf("send: %s\n", key.EventInput)
	// 	// input <- KeyInput{key, state}
	// 	fmt.Printf("OK: %s\n", key.EventInput)
	// }

	var send = func(key keymap.WindowsKey, state InputType) {
		h.remote.SendInput(KeyInput{key, state})
	}

	return func(msg win.MSG) error {
		var wParam = msg.WParam
		var lParam = msg.LParam

		h.Output(10, fmt.Sprintf("%X(%d) ", msg.Message, msg.Message))
		h.Output(10, fmt.Sprintf("%+v\n", msg))

		switch msg.Message {
		case win.WA_CLICKACTIVE:
			h.Debugf("WA_CLICKACTIVE\n")
			return nil
		case win.WM_PAINT:
			// var ps = new(win.PAINTSTRUCT)
			// var hdc = win.BeginPaint(hwnd, ps)
			// var hBrush = winapi.CreateSolidBrush(0x00FFFFFF)

			// win.SelectObject(hdc, hBrush)
			// winapi.ExtFloodFill(hdc, 1, 1, 0xFFFFFF, winapi.FLOODFILLSURFACE)
			// win.DeleteObject(hBrush)
			// win.EndPaint(hwnd, ps)
			return nil
		case win.WM_SYSKEYDOWN:
			if lParam>>30&1 == 0 {
				key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
				if err == nil {
					send(*key, KeyDown)
				}
			}
			key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
			if err != nil {
				h.Debugf("WM_SYSKEYDOWN: GetKeyError: %d\n", wParam)
			} else {
				h.logger.Output(4, fmt.Sprintf("WM_SYSKEYDOWN: wParam: %v lParam: %v\n", key.Constant, lParam))
			}
			return nil
		case win.WM_SYSKEYUP:
			key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
			if err == nil {
				send(*key, KeyUp)
				h.Output(4, fmt.Sprintf("WM_SYSKEYUP: wParam: %v lParam: %v\n", key.Constant, lParam))
			} else {
				h.Debugf("WM_SYSKEYUP: GetKeyError: %d\n", wParam)
			}
			return nil
		case win.WM_KEYDOWN:
			if lParam>>30&1 == 0 {
				key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
				if err == nil {
					send(*key, KeyDown)
				}
			}
			key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
			if err != nil {
				h.Debugf("WM_KEYDOWN: GetKeyError: %d\n", wParam)
			} else {
				h.Output(4, fmt.Sprintf("WM_KEYDOWN: wParam: %v lParam: %v\n", key.Constant, lParam))
			}
			return nil
		case win.WM_KEYUP:
			key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
			if err == nil {
				send(*key, KeyUp)
				h.Output(4, fmt.Sprintf("WM_KEYUP: wParam: %v lParam: %v\n", key.Constant, lParam))
			} else {
				h.Debugf("WM_KEYUP: GetKeyError: %d\n", wParam)
			}
			return nil
		case win.WM_MOUSEWHEEL:
			if int16(wParam>>16) > 0 {
				send(keymap.WindowsKey{Constant: "WheelUp", EventType: "EV_MOUSE", EventInput: "wheel"}, KeyUp)
				h.Output(4, "WM_MOUSEWHEEL UP")
			} else {
				send(keymap.WindowsKey{Constant: "WheelDown", EventType: "EV_MOUSE", EventInput: "wheel"}, KeyDown)
				h.Output(4, "WM_MOUSEWHEEL DOWN")
			}
			return nil
		case win.WM_LBUTTONUP:
			key, _ := keymap.GetWindowsKeyDetail(0x01)
			send(*key, KeyUp)
			h.Output(4, "WM_LBUTTONUP UP\n")
			return nil
		case win.WM_LBUTTONDOWN:
			key, _ := keymap.GetWindowsKeyDetail(0x01)
			send(*key, KeyDown)
			h.Output(4, "WM_LBUTTONUP DOWN\n")
			return nil
		case win.WM_RBUTTONUP:
			key, _ := keymap.GetWindowsKeyDetail(0x02)
			send(*key, KeyUp)
			h.Output(4, "WM_RBUTTONUP UP\n")
			return nil
		case win.WM_RBUTTONDOWN:
			key, _ := keymap.GetWindowsKeyDetail(0x02)
			send(*key, KeyDown)
			h.Output(4, "WM_RBUTTONDOWN DOWN\n")
			return nil
		case win.WM_MBUTTONUP:
			key, _ := keymap.GetWindowsKeyDetail(0x04)
			send(*key, KeyUp)
			h.Output(4, "WM_RBUTTONUP UP\n")
			return nil
		case win.WM_MBUTTONDOWN:
			key, _ := keymap.GetWindowsKeyDetail(0x02)
			send(*key, KeyDown)
			h.Output(4, "WM_MBUTTONDOWN DOWN\n")
			return nil
		default:
			return nil
		}
	}
}

func (h Handler) windowProc(rdClientHwnd win.HWND) func(hwnd win.HWND, uMsg uint32, wParam uintptr, lParam uintptr) uintptr {
	return func(hwnd win.HWND, uMsg uint32, wParam uintptr, lParam uintptr) uintptr {
		switch uMsg {
		case win.WM_CREATE:
			// Show window on RDP client with transparent style
			winapi.SetLayeredWindowAttributes(hwnd, 0xFFFFFF, byte(1), winapi.LWA_ALPHA)
			var err = h.PutWindowOnAnotherWindow(hwnd, rdClientHwnd)
			if err != nil {
				h.Debugf("%s", errors.Wrap(err, "windowProc"))
			}
			// winapi.ShowCursor(false)
			return winapi.NULL
		}

		return win.DefWindowProc(hwnd, uMsg, wParam, lParam)
	}
}

func (h *Handler) PutWindowOnAnotherWindow(hwnd win.HWND, otherHWND win.HWND) error {
	if !win.SetForegroundWindow(hwnd) {
		return errors.New("SetForegroundWindow: failed to get foreground permission")
	}

	var crectAbs win.RECT
	if !winapi.GetWindowRect(otherHWND, &crectAbs) {
		return errors.New("GetWindowRectError")
	}

	if !win.SetWindowPos(hwnd, 0,
		crectAbs.Left, crectAbs.Top, crectAbs.Right-crectAbs.Left, crectAbs.Bottom-crectAbs.Top,
		win.SWP_SHOWWINDOW,
	) {
		return errors.New("SetWindowPos: failed to set window pos")
	}

	return nil
}

func (h *Handler) Close() {
	winapi.ClipCursor(win.RECT{})
	winapi.ShowCursor(true)
}
