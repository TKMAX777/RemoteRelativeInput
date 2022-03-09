package windows

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/TKMAX777/RemoteRelativeInput/keymap"
	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
	"github.com/TKMAX777/RemoteRelativeInput/winapi"
	"github.com/pkg/errors"

	"github.com/lxn/win"
)

type Handler struct {
	metrics SystemMetrics
	options option
	remote  *remote_send.Handler

	logger logger
	debug  bool
}

type ToggleType int

const (
	ToggleTypeOnce ToggleType = iota + 1
	ToggleTypeAlive
)

type option struct {
	toggleKey  string
	toggleType ToggleType
}

type SystemMetrics struct {
	FrameWidthX int32
	FrameWidthY int32
	TitleHeight int32
}

func New(r *remote_send.Handler) *Handler {
	return &Handler{
		remote: r,
		options: option{
			// set default options
			toggleKey:  "F8",
			toggleType: ToggleTypeAlive,
		},
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

func (h *Handler) SetToggleKey(k string) error {
	if k == "" {
		return errors.New("NotSpecified")
	}
	h.options.toggleKey = k
	return nil
}

func (h *Handler) SetToggleType(t ToggleType) {
	h.options.toggleType = t
}

func (h Handler) getWindowCenterPos(rect win.RECT) win.POINT {
	var windowCenterPosition win.POINT
	windowCenterPosition.X = int32(rect.Left+rect.Right) / 2
	windowCenterPosition.Y = int32(rect.Top+rect.Bottom) / 2

	return windowCenterPosition
}

// Create window on remote desktop client
// rdClientHwnd must be remote desktop client hwnd, and toggleKey is a keyname for toggle wrapper mode
func (h Handler) StartClient(rdClientHwnd win.HWND) (win.HWND, error) {
	if rdClientHwnd == win.HWND(winapi.NULL) {
		return win.HWND(winapi.NULL), errors.New("NilWindowHandler")
	}

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

		// get window proc
		var windowProc = h.getWindowProc(rdClientHwnd)

		// lock os thread to avoid hanging GetMessage
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		var wClass win.WNDCLASSEX
		wClass = win.WNDCLASSEX{
			CbSize:    uint32(unsafe.Sizeof(wClass)),
			HInstance: hInstance,

			// redraw when the window size changed
			Style: win.CS_HREDRAW | win.CS_VREDRAW,

			// set window background color to white
			HbrBackground: win.HBRUSH(win.GetStockObject(win.WHITE_BRUSH)),
			LpszClassName: className,
			LpfnWndProc:   syscall.NewCallback(windowProc),
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
			return
		}

		winapi.ShowWindow(hwnd, win.SW_SHOW)
		winapi.UpdateWindow(hwnd)

		result <- resultAttr{hwnd, nil}

		for {
			var msg win.MSG
			switch win.GetMessage(&msg, hwnd, 0, 0) {
			case 0:
				h.Debugln("Quit")
				return
			case -1:
				os.Exit(0)
				return
			}

			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
			h.Output(10, fmt.Sprintf("disp: %+v \n", msg))
		}
	}()

	var res = <-result
	close(result)

	return res.hwnd, res.err
}

func (h Handler) initWindowAndCursor(hwnd, rdClientHwnd win.HWND) error {
	// Set window on RDP client window
	if !win.SetForegroundWindow(hwnd) {
		return errors.New("SetForegroundWindow: failed to get foreground permission")
	}

	var crectAbs win.RECT
	if !winapi.GetWindowRect(rdClientHwnd, &crectAbs) {
		return errors.New("GetWindowRectError")
	}

	if !win.SetWindowPos(hwnd, 0,
		crectAbs.Left, crectAbs.Top, crectAbs.Right-crectAbs.Left, crectAbs.Bottom-crectAbs.Top,
		win.SWP_SHOWWINDOW,
	) {
		return errors.New("SetWindowPos: failed to set window pos")
	}

	// get remote desktop client rect
	var rect win.RECT
	if !winapi.GetWindowRect(rdClientHwnd, &rect) {
		return errors.New("GetWindowRectError")
	}

	// get remote desktop client center position
	var windowCenterPosition = h.getWindowCenterPos(rect)

	// set cursot to center of the remote desktop client window
	if !win.SetCursorPos(windowCenterPosition.X, windowCenterPosition.Y) {
		return errors.New("Error in set cursor position")
	}

	// then clip cursor
	okInt, _ := winapi.ClipCursor(&rect)
	if okInt != 1 {
		return errors.Errorf("Error in clip cursor: code: %d\n", win.GetLastError())
	}

	winapi.ShowCursor(false)

	return nil
}

func (h Handler) getWindowProc(rdClientHwnd win.HWND) func(hwnd win.HWND, uMsg uint32, wParam uintptr, lParam uintptr) uintptr {
	var pos POINT

	// get remote desktop client rect
	var rect win.RECT
	if !winapi.GetWindowRect(rdClientHwnd, &rect) {
		fmt.Fprintf(os.Stderr, "getWindowProc: GetWindowRectError")
	}

	var isRelativeMode = true

	// get remote desktop client center position
	var windowCenterPosition = h.getWindowCenterPos(rect)

	var currentPosition = windowCenterPosition

	return func(hwnd win.HWND, uMsg uint32, wParam uintptr, lParam uintptr) uintptr {
		var send = func(key keymap.WindowsKey, state InputType) {
			// toggle window mode
			if key.EventInput == h.options.toggleKey {
				switch state {
				case KeyDown:
					isRelativeMode = !isRelativeMode || h.options.toggleType == ToggleTypeOnce
				case KeyUp:
					if h.options.toggleType == ToggleTypeOnce {
						isRelativeMode = false
					}
				}

				if isRelativeMode {
					h.initWindowAndCursor(hwnd, rdClientHwnd)
				} else {
					var crectAbs win.RECT
					if !winapi.GetWindowRect(rdClientHwnd, &crectAbs) {
						fmt.Fprintf(os.Stderr, "GetWindowRectError")
					}

					// show window title only
					if !win.SetWindowPos(hwnd, 0,
						crectAbs.Left, crectAbs.Top, crectAbs.Right-crectAbs.Left, h.metrics.TitleHeight+h.metrics.FrameWidthY*2,
						win.SWP_SHOWWINDOW,
					) {
						fmt.Fprintf(os.Stderr, "SetWindowPos: failed to set window pos")
					}

					winapi.ShowCursor(true)
					winapi.ClipCursor(nil)
				}
			}
			if isRelativeMode {
				h.remote.SendInput(KeyInput{key, state})
			}
		}

		h.Output(10, fmt.Sprintf("%X(%d) ", uMsg, uMsg))

		switch uMsg {
		case win.WM_CREATE:
			winapi.SetLayeredWindowAttributes(hwnd, 0x0000FF, byte(1), winapi.LWA_COLORKEY)
			h.initWindowAndCursor(hwnd, rdClientHwnd)
			win.UpdateWindow(hwnd)

			return winapi.NULL
		case win.WM_PAINT:
			var ps = new(win.PAINTSTRUCT)
			var hdc = win.BeginPaint(hwnd, ps)
			var hBrush = winapi.CreateSolidBrush(0x000000FF)

			win.SelectObject(hdc, hBrush)
			winapi.ExtFloodFill(hdc, 1, 1, 0xFFFFFF, winapi.FLOODFILLSURFACE)
			win.DeleteObject(hBrush)
			win.EndPaint(hwnd, ps)

			if isRelativeMode {
				winapi.SetLayeredWindowAttributes(hwnd, 0x0000FF, byte(1), winapi.LWA_COLORKEY)

				// get remote desktop client rect
				var rect win.RECT
				if !winapi.GetWindowRect(rdClientHwnd, &rect) {
					fmt.Fprintf(os.Stderr, "getWindowProc: GetWindowRectError")
				}

				// get remote desktop client center position
				windowCenterPosition = h.getWindowCenterPos(rect)
			}
			return winapi.NULL
		case win.WM_MOUSEMOVE:
			if isRelativeMode {
				ok := winapi.GetCursorPos(&currentPosition)
				if !ok {
					log.Printf("Error in get cursor position\n")
				}
				// Relative position mode
				pos.POINT = win.POINT{X: currentPosition.X - windowCenterPosition.X, Y: currentPosition.Y - windowCenterPosition.Y}
				if pos.X != 0 || pos.Y != 0 {
					h.Debugf("X: %4d Y: %4d\n", pos.X, pos.Y)
					h.remote.SendRelativeCursor(pos)

					ok = win.SetCursorPos(windowCenterPosition.X, windowCenterPosition.Y)
					if !ok {
						log.Printf("Error in set cursor position")
					}
				}
			}

			return winapi.NULL
		case win.WA_CLICKACTIVE:
			h.Debugf("WA_CLICKACTIVE\n")
			return winapi.NULL
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
				h.Output(4, fmt.Sprintf("WM_SYSKEYDOWN: wParam: %v lParam: %v\n", key.Constant, lParam))
			}
			return winapi.NULL
		case win.WM_SYSKEYUP:
			key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
			if err == nil {
				send(*key, KeyUp)
				h.Output(4, fmt.Sprintf("WM_SYSKEYUP: wParam: %v lParam: %v\n", key.Constant, lParam))
			} else {
				h.Debugf("WM_SYSKEYUP: GetKeyError: %d\n", wParam)
			}
			return winapi.NULL
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
			return winapi.NULL
		case win.WM_KEYUP:
			key, err := keymap.GetWindowsKeyDetail(uint32(wParam))
			if err == nil {
				send(*key, KeyUp)
				h.Output(4, fmt.Sprintf("WM_KEYUP: wParam: %v lParam: %v\n", key.Constant, lParam))
			} else {
				h.Debugf("WM_KEYUP: GetKeyError: %d\n", wParam)
			}
			return winapi.NULL
		case win.WM_MOUSEWHEEL:
			if int16(wParam>>16) > 0 {
				send(keymap.WindowsKey{Constant: "WheelUp", EventType: "EV_MOUSE", EventInput: "wheel"}, KeyUp)
				h.Output(4, "WM_MOUSEWHEEL UP")
			} else {
				send(keymap.WindowsKey{Constant: "WheelDown", EventType: "EV_MOUSE", EventInput: "wheel"}, KeyDown)
				h.Output(4, "WM_MOUSEWHEEL DOWN")
			}
			return winapi.NULL
		case win.WM_LBUTTONUP:
			key, _ := keymap.GetWindowsKeyDetail(0x01)
			send(*key, KeyUp)
			h.Output(4, "WM_LBUTTONUP UP\n")
			return winapi.NULL
		case win.WM_LBUTTONDOWN:
			key, _ := keymap.GetWindowsKeyDetail(0x01)
			send(*key, KeyDown)
			h.Output(4, "WM_LBUTTONUP DOWN\n")
			return winapi.NULL
		case win.WM_RBUTTONUP:
			key, _ := keymap.GetWindowsKeyDetail(0x02)
			send(*key, KeyUp)
			h.Output(4, "WM_RBUTTONUP UP\n")
			return winapi.NULL
		case win.WM_RBUTTONDOWN:
			key, _ := keymap.GetWindowsKeyDetail(0x02)
			send(*key, KeyDown)
			h.Output(4, "WM_RBUTTONDOWN DOWN\n")
			return winapi.NULL
		case win.WM_MBUTTONUP:
			key, _ := keymap.GetWindowsKeyDetail(0x04)
			send(*key, KeyUp)
			h.Output(4, "WM_RBUTTONUP UP\n")
			return winapi.NULL
		case win.WM_MBUTTONDOWN:
			key, _ := keymap.GetWindowsKeyDetail(0x02)
			send(*key, KeyDown)
			h.Output(4, "WM_MBUTTONDOWN DOWN\n")
			return winapi.NULL
		default:
			return win.DefWindowProc(hwnd, uMsg, wParam, lParam)
		}
	}
}

func (h Handler) Close() {
	winapi.ClipCursor(nil)
	winapi.ShowCursor(true)
}
