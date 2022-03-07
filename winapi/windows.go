package winapi

const NULL uintptr = 0

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output windows_generate.go windows.go
//sys clipCursor(rect uintptr)(ok int, err error) = user32.ClipCursor
//sys showCursor(state bool) (counter int) = user32.ShowCursor
//sys enumDesktopWindows(hDesktop uintptr, lpEnumFunc uintptr, lParam uintptr) (ok bool) = user32.EnumDesktopWindows
//sys setLayeredWindowAttributes(hwnd uintptr, color uint32, bAlpha byte, dwFlags uint32) (ok bool) = user32.SetLayeredWindowAttributes
//sys fillRect(hdc uintptr, rect uintptr, hbr uintptr) (ok bool) = user32.FillRect
//sys setWindowRgn(hwnd uintptr, hRgn uintptr, bRedraw bool) (ok bool) = user32.SetWindowRgn
//sys updateLayeredWindow(hwnd uintptr, hdcDst uintptr, pptDst uintptr, psize uintptr, hdcSrc uintptr, pptSrc uintptr, crKey uint32, pblend uintptr, dwFlags uint32) (ok bool) = user32.UpdateLayeredWindow
//sys findWindowEx(hwndParent uintptr, hwndChildAfter uintptr, lpszClass *uint16, lpszWindow *uint16) (hwnd uintptr) = user32.FindWindowExW
//sys getWindowText(hwnd uintptr, lpString uintptr, nMax int) (length int) = user32.GetWindowTextW
//sys getClassName(hwnd uintptr, lpClassName uintptr, nMax int) (length int) = user32.GetClassNameW
//sys setWindowText(hwnd uintptr, lpString *uint16) (ok bool) = user32.SetWindowTextW
//sys invalidateRect(hwnd uintptr, rect uintptr, bErase bool) (ok bool) = user32.InvalidateRect

//sys createSolidBrush(color uint32) (hbrush uintptr) = Gdi32.CreateSolidBrush
//sys createPen(iStyle int, cWidth int, color uint32) (hpen uintptr) = Gdi32.CreatePen
//sys polyDraw(hdc uintptr, apt uintptr, aj uintptr, cpt int) (ok bool) = Gdi32.PolyDraw
//sys createRectRgnIndirect(rect uintptr) (rgn uintptr) = Gdi32.CreateRectRgnIndirect
//sys createDIBSection(hdc uintptr, pbmi uintptr, usage uint, ppvBits uintptr, hSection uintptr, offset uint32) (hBitMap uintptr) = Gdi32.CreateDIBSection
//sys extFloodFill(hdc uintptr, x int, y int, color uint32, opType uint32) (ok bool) = Gdi32.ExtFloodFill

//sys activateAudioInterfaceAsync(deviceInterfacePath *uint16, riid uintptr, activationParams uintptr, completionHandler uintptr, createAsync uintptr) (hresult int32) = Mmdevapi.ActivateAudioInterfaceAsync
