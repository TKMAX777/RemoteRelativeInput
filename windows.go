//go:build windows
// +build windows

package relative_input

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
	"github.com/TKMAX777/RemoteRelativeInput/winapi"
	"github.com/TKMAX777/RemoteRelativeInput/windows"
)

func StartClient() {
	var rHandler = remote_send.New(os.Stdout)
	var wHandler = windows.New(rHandler)

	// wHandler.SetLogger(log.New(os.Stdout, "", 6))
	var windowName = winapi.MustUTF16PtrFromString("192.168.100.82:10061 - リモート デスクトップ接続")
	var rdHwnd = winapi.FindWindow(nil, windowName)

	hwnd, err := wHandler.CreateWindow(rdHwnd)
	if err != nil {
		panic(err)
	}

	err = wHandler.SendCursor(hwnd)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stderr, "Ready for sending messages")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	wHandler.Close()
}

func StartServer() {
	// TODO: make windows relative input server
}
