package winapi

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

func MustUTF16PtrFromString(str string) *uint16 {
	ptr, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		panic(err)
	}
	return ptr
}

func MustUTF16FromString(str string) []uint16 {
	ptr, err := syscall.UTF16FromString(str)
	if err != nil {
		panic(err)
	}
	return ptr
}

func EnumChildWindows(hwnd win.HWND) []win.HWND {
	var res = []win.HWND{}

	var chwnd win.HWND
	for chwnd = FindWindowEx(hwnd, chwnd, nil, nil); chwnd != win.HWND(NULL); chwnd = FindWindowEx(hwnd, chwnd, nil, nil) {
		res = append(res, chwnd)
	}

	return res
}

func FindChildWindowsFromWindowText(parentHWND win.HWND, lpszClass *uint16, lpszWindow *uint16, windowText string) win.HWND {
	var chwnd win.HWND
	for chwnd = FindWindowEx(parentHWND, chwnd, nil, nil); chwnd != win.HWND(NULL); chwnd = FindWindowEx(parentHWND, chwnd, nil, nil) {
		var UTF16name = make([]uint16, 1000)
		GetWindowText(chwnd, uintptr(unsafe.Pointer(&UTF16name[0])), 1000)
		if syscall.UTF16ToString(UTF16name) == windowText {
			return chwnd
		}
	}
	return win.HWND(NULL)
}

func SendMessage(hwnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	return win.SendMessage(hwnd, msg, wParam, lParam)
}
