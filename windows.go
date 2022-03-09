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
	"github.com/lxn/win"
)

func StartClient() {
	win.MessageBox(win.HWND(winapi.NULL), winapi.MustUTF16PtrFromString("Click to start client"), winapi.MustUTF16PtrFromString("Confirmation"), win.MB_OK)

	var rHandler = remote_send.New(os.Stdout)
	var wHandler = windows.New(rHandler)

	var toggleKey = os.Getenv("TOGGLE_KEY")
	if toggleKey == "" {
		toggleKey = "F8"
	}

	var toggleType = os.Getenv("ToggleType")
	switch toggleType {
	case "TOGGLE_TYPE_ONCE":
		wHandler.SetToggleType(windows.ToggleTypeOnce)
	default:
		wHandler.SetToggleType(windows.ToggleTypeAlive)
	}

	wHandler.SetToggleKey(toggleKey)
	// wHandler.SetLogger(log.New(os.Stderr, "", 10))

	var windowName = winapi.MustUTF16PtrFromString(os.Getenv("CLIENT_NAME"))
	var rdHwnd = winapi.FindWindow(nil, windowName)

	_, err := wHandler.StartClient(rdHwnd)
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
