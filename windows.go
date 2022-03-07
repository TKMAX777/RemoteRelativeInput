//go:build windows
// +build windows

package relative_input

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
	"github.com/TKMAX777/RemoteRelativeInput/winapi"
	"github.com/TKMAX777/RemoteRelativeInput/windows"
	"github.com/lxn/win"
)

func StartClient() {
	win.MessageBox(win.HWND(winapi.NULL), winapi.MustUTF16PtrFromString("Click to start client"), winapi.MustUTF16PtrFromString("Confirmation"), win.MB_OK)

	var rHandler = remote_send.New(os.Stdout)
	var wHandler = windows.New(rHandler)

	wHandler.SetLogger(log.New(os.Stderr, "", 10))
	var windowName = winapi.MustUTF16PtrFromString(os.Getenv("CLIENT_NAME"))
	var rdHwnd = winapi.FindWindow(nil, windowName)

	_, err := wHandler.CreateWindow(rdHwnd)
	if err != nil {
		panic(err)
	}

	// err = wHandler.SendCursor(hwnd)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Fprintln(os.Stderr, "Ready for sending messages")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	wHandler.Close()
}

func StartServer() {
	// TODO: make windows relative input server
}
