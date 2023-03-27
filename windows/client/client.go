package client

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	relative_input "github.com/TKMAX777/RemoteRelativeInput"
	"github.com/TKMAX777/RemoteRelativeInput/debug"
	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
	"github.com/TKMAX777/RemoteRelativeInput/winapi"
	"github.com/lxn/win"
)

func StartClient() {
	defer os.Stdout.Write([]byte("CLOSE\n"))
	win.MessageBox(win.HWND(winapi.NULL), winapi.MustUTF16PtrFromString("Click to start client"), winapi.MustUTF16PtrFromString("Confirmation"), win.MB_OK|win.MB_ICONINFORMATION)

	debug.Debugln("==== START CLIENT APPLICATION ====")
	debug.Debugln("ServerProtocolVersion:", relative_input.PROTOCOL_VERSION)

	debug.Debugf("Wait for client headers...")

	var rHandler = remote_send.New(os.Stdout)
	var wHandler = New(rHandler)

	var toggleKey = os.Getenv("RELATIVE_INPUT_TOGGLE_KEY")
	if toggleKey == "" {
		toggleKey = "F8"
	}

	var toggleType = os.Getenv("RELATIVE_INPUT_TOGGLE_TYPE")
	switch toggleType {
	case "ONCE":
		wHandler.SetToggleType(ToggleTypeOnce)
	default:
		wHandler.SetToggleType(ToggleTypeAlive)
	}

	wHandler.SetToggleKey(toggleKey)

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
