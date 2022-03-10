package winapi

import (
	"unsafe"

	"github.com/lxn/win"
)

const (
	LWA_COLORKEY uint32 = 1 + iota
	LWA_ALPHA
)

const (
	MAPVK_VK_TO_VSC uint32 = iota
	MAPVK_VSC_TO_VK
	MAPVK_VK_TO_CHAR
	MAPVK_VSC_TO_VK_EX
	MAPVK_VK_TO_VSC_EX
)

func ClipCursor(rect *win.RECT) (ok int, err error) {
	if rect == nil {
		return clipCursor(NULL)
	}
	return clipCursor(uintptr(unsafe.Pointer(&rect.Left)))
}

func EnumDesktopWindows(hDesktop win.HANDLE, lpEnumFunc uintptr, lParam uintptr) (ok bool) {
	return enumDesktopWindows(uintptr(hDesktop), lpEnumFunc, lParam)
}

func FillRect(hdc win.HDC, rect win.RECT, hbr win.HBRUSH) (ok bool) {
	return fillRect(uintptr(hdc), uintptr(unsafe.Pointer(&rect.Left)), uintptr(hbr))
}

func FindWindow(lpClassName, lpWindowName *uint16) win.HWND {
	return win.FindWindow(lpClassName, lpWindowName)
}

func FindWindowEx(hwndParent win.HWND, hwndChildAfter win.HWND, lpszClass *uint16, lpszWindow *uint16) (hwnd win.HWND) {
	return win.HWND(findWindowEx(uintptr(hwndParent), uintptr(hwndChildAfter), lpszClass, lpszWindow))
}

func GetClassName(hwnd win.HWND, lpClassName uintptr, nMax int) (length int) {
	return getClassName(uintptr(hwnd), lpClassName, nMax)
}

func GetWindowText(hwnd win.HWND, lpString uintptr, nMax int) (length int) {
	return getWindowText(uintptr(hwnd), lpString, nMax)
}

func InvalidateRect(hwnd win.HWND, rect win.RECT, bErase bool) (ok bool) {
	return invalidateRect(uintptr(hwnd), uintptr(unsafe.Pointer(&rect.Left)), bErase)
}

func SetLayeredWindowAttributes(hwnd win.HWND, color uint32, bAlpha byte, dwFlags uint32) (ok bool) {
	return setLayeredWindowAttributes(uintptr(hwnd), color, bAlpha, dwFlags)
}

func SetWindowRgn(hwnd win.HWND, hRgn win.HRGN, bRedraw bool) (ok bool) {
	return setWindowRgn(uintptr(hwnd), uintptr(hRgn), bRedraw)
}

func SetWindowText(hwnd win.HWND, lpString *uint16) (ok bool) {
	return setWindowText(uintptr(hwnd), lpString)
}

func ShowCursor(state bool) (counter int) {
	return showCursor(state)
}

func ShowWindow(hWnd win.HWND, nCmdShow int32) bool {
	return win.ShowWindow(hWnd, nCmdShow)
}

func UpdateLayeredWindow(hwnd win.HWND, hdcDst win.HDC, pptDst win.POINT, psize uintptr, hdcSrc win.HDC, pptSrc win.POINT, crKey uint32, pblend win.BLENDFUNCTION, dwFlags uint32) (ok bool) {
	return updateLayeredWindow(uintptr(hwnd), uintptr(hdcDst), uintptr(unsafe.Pointer(&pptDst.X)), psize, uintptr(hdcSrc), uintptr(unsafe.Pointer(&pptSrc.X)), crKey, uintptr(unsafe.Pointer(&pblend.BlendOp)), dwFlags)
}

func UpdateWindow(hwnd win.HWND) bool {
	return win.UpdateWindow(hwnd)
}

func GetWindowRect(hwnd win.HWND, rect *win.RECT) bool {
	return win.GetWindowRect(hwnd, rect)
}

func GetCursorPos(lpPoint *win.POINT) bool {
	return win.GetCursorPos(lpPoint)
}

func SetForegroundWindow(hWnd win.HWND) bool {
	return win.SetForegroundWindow(hWnd)
}

func MapVirtualKey(uCode uint32, uMapType uint32) (code uint32) {
	return mapVirtualKey(uCode, uMapType)
}
