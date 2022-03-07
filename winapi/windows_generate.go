// Code generated by 'go generate'; DO NOT EDIT.

package winapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modGdi32    = windows.NewLazySystemDLL("Gdi32.dll")
	modMmdevapi = windows.NewLazySystemDLL("Mmdevapi.dll")
	moduser32   = windows.NewLazySystemDLL("user32.dll")

	procCreateDIBSection            = modGdi32.NewProc("CreateDIBSection")
	procCreatePen                   = modGdi32.NewProc("CreatePen")
	procCreateRectRgnIndirect       = modGdi32.NewProc("CreateRectRgnIndirect")
	procCreateSolidBrush            = modGdi32.NewProc("CreateSolidBrush")
	procExtFloodFill                = modGdi32.NewProc("ExtFloodFill")
	procPolyDraw                    = modGdi32.NewProc("PolyDraw")
	procActivateAudioInterfaceAsync = modMmdevapi.NewProc("ActivateAudioInterfaceAsync")
	procClipCursor                  = moduser32.NewProc("ClipCursor")
	procEnumDesktopWindows          = moduser32.NewProc("EnumDesktopWindows")
	procFillRect                    = moduser32.NewProc("FillRect")
	procFindWindowExW               = moduser32.NewProc("FindWindowExW")
	procGetClassNameW               = moduser32.NewProc("GetClassNameW")
	procGetWindowTextW              = moduser32.NewProc("GetWindowTextW")
	procInvalidateRect              = moduser32.NewProc("InvalidateRect")
	procSetLayeredWindowAttributes  = moduser32.NewProc("SetLayeredWindowAttributes")
	procSetWindowRgn                = moduser32.NewProc("SetWindowRgn")
	procSetWindowTextW              = moduser32.NewProc("SetWindowTextW")
	procShowCursor                  = moduser32.NewProc("ShowCursor")
	procUpdateLayeredWindow         = moduser32.NewProc("UpdateLayeredWindow")
)

func createDIBSection(hdc uintptr, pbmi uintptr, usage uint, ppvBits uintptr, hSection uintptr, offset uint32) (hBitMap uintptr) {
	r0, _, _ := syscall.Syscall6(procCreateDIBSection.Addr(), 6, uintptr(hdc), uintptr(pbmi), uintptr(usage), uintptr(ppvBits), uintptr(hSection), uintptr(offset))
	hBitMap = uintptr(r0)
	return
}

func createPen(iStyle int, cWidth int, color uint32) (hpen uintptr) {
	r0, _, _ := syscall.Syscall(procCreatePen.Addr(), 3, uintptr(iStyle), uintptr(cWidth), uintptr(color))
	hpen = uintptr(r0)
	return
}

func createRectRgnIndirect(rect uintptr) (rgn uintptr) {
	r0, _, _ := syscall.Syscall(procCreateRectRgnIndirect.Addr(), 1, uintptr(rect), 0, 0)
	rgn = uintptr(r0)
	return
}

func createSolidBrush(color uint32) (hbrush uintptr) {
	r0, _, _ := syscall.Syscall(procCreateSolidBrush.Addr(), 1, uintptr(color), 0, 0)
	hbrush = uintptr(r0)
	return
}

func extFloodFill(hdc uintptr, x int, y int, color uint32, opType uint32) (ok bool) {
	r0, _, _ := syscall.Syscall6(procExtFloodFill.Addr(), 5, uintptr(hdc), uintptr(x), uintptr(y), uintptr(color), uintptr(opType), 0)
	ok = r0 != 0
	return
}

func polyDraw(hdc uintptr, apt uintptr, aj uintptr, cpt int) (ok bool) {
	r0, _, _ := syscall.Syscall6(procPolyDraw.Addr(), 4, uintptr(hdc), uintptr(apt), uintptr(aj), uintptr(cpt), 0, 0)
	ok = r0 != 0
	return
}

func activateAudioInterfaceAsync(deviceInterfacePath *uint16, riid uintptr, activationParams uintptr, completionHandler uintptr, createAsync uintptr) (hresult int32) {
	r0, _, _ := syscall.Syscall6(procActivateAudioInterfaceAsync.Addr(), 5, uintptr(unsafe.Pointer(deviceInterfacePath)), uintptr(riid), uintptr(activationParams), uintptr(completionHandler), uintptr(createAsync), 0)
	hresult = int32(r0)
	return
}

func clipCursor(rect uintptr) (ok int, err error) {
	r0, _, e1 := syscall.Syscall(procClipCursor.Addr(), 1, uintptr(rect), 0, 0)
	ok = int(r0)
	if ok == 0 {
		err = errnoErr(e1)
	}
	return
}

func enumDesktopWindows(hDesktop uintptr, lpEnumFunc uintptr, lParam uintptr) (ok bool) {
	r0, _, _ := syscall.Syscall(procEnumDesktopWindows.Addr(), 3, uintptr(hDesktop), uintptr(lpEnumFunc), uintptr(lParam))
	ok = r0 != 0
	return
}

func fillRect(hdc uintptr, rect uintptr, hbr uintptr) (ok bool) {
	r0, _, _ := syscall.Syscall(procFillRect.Addr(), 3, uintptr(hdc), uintptr(rect), uintptr(hbr))
	ok = r0 != 0
	return
}

func findWindowEx(hwndParent uintptr, hwndChildAfter uintptr, lpszClass *uint16, lpszWindow *uint16) (hwnd uintptr) {
	r0, _, _ := syscall.Syscall6(procFindWindowExW.Addr(), 4, uintptr(hwndParent), uintptr(hwndChildAfter), uintptr(unsafe.Pointer(lpszClass)), uintptr(unsafe.Pointer(lpszWindow)), 0, 0)
	hwnd = uintptr(r0)
	return
}

func getClassName(hwnd uintptr, lpClassName uintptr, nMax int) (length int) {
	r0, _, _ := syscall.Syscall(procGetClassNameW.Addr(), 3, uintptr(hwnd), uintptr(lpClassName), uintptr(nMax))
	length = int(r0)
	return
}

func getWindowText(hwnd uintptr, lpString uintptr, nMax int) (length int) {
	r0, _, _ := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(lpString), uintptr(nMax))
	length = int(r0)
	return
}

func invalidateRect(hwnd uintptr, rect uintptr, bErase bool) (ok bool) {
	var _p0 uint32
	if bErase {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procInvalidateRect.Addr(), 3, uintptr(hwnd), uintptr(rect), uintptr(_p0))
	ok = r0 != 0
	return
}

func setLayeredWindowAttributes(hwnd uintptr, color uint32, bAlpha byte, dwFlags uint32) (ok bool) {
	r0, _, _ := syscall.Syscall6(procSetLayeredWindowAttributes.Addr(), 4, uintptr(hwnd), uintptr(color), uintptr(bAlpha), uintptr(dwFlags), 0, 0)
	ok = r0 != 0
	return
}

func setWindowRgn(hwnd uintptr, hRgn uintptr, bRedraw bool) (ok bool) {
	var _p0 uint32
	if bRedraw {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procSetWindowRgn.Addr(), 3, uintptr(hwnd), uintptr(hRgn), uintptr(_p0))
	ok = r0 != 0
	return
}

func setWindowText(hwnd uintptr, lpString *uint16) (ok bool) {
	r0, _, _ := syscall.Syscall(procSetWindowTextW.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(lpString)), 0)
	ok = r0 != 0
	return
}

func showCursor(state bool) (counter int) {
	var _p0 uint32
	if state {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procShowCursor.Addr(), 1, uintptr(_p0), 0, 0)
	counter = int(r0)
	return
}

func updateLayeredWindow(hwnd uintptr, hdcDst uintptr, pptDst uintptr, psize uintptr, hdcSrc uintptr, pptSrc uintptr, crKey uint32, pblend uintptr, dwFlags uint32) (ok bool) {
	r0, _, _ := syscall.Syscall9(procUpdateLayeredWindow.Addr(), 9, uintptr(hwnd), uintptr(hdcDst), uintptr(pptDst), uintptr(psize), uintptr(hdcSrc), uintptr(pptSrc), uintptr(crKey), uintptr(pblend), uintptr(dwFlags))
	ok = r0 != 0
	return
}