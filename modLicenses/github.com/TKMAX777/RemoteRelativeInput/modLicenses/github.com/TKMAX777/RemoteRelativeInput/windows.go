//go:build windows
// +build windows

package relative_input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"github.com/TKMAX777/RemoteRelativeInput/debug"
	"github.com/TKMAX777/RemoteRelativeInput/keymap"
	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
	"github.com/TKMAX777/RemoteRelativeInput/winapi"
	"github.com/TKMAX777/RemoteRelativeInput/windows_client/client"
	"github.com/lxn/win"
	"github.com/natefinch/npipe"
	"github.com/pkg/errors"
)

func StartClient() {
	debug.Debugln("==== START CLIENT APPLICATION ====")
	win.MessageBox(win.HWND(winapi.NULL), winapi.MustUTF16PtrFromString("Click to start client"), winapi.MustUTF16PtrFromString("Confirmation"), win.MB_OK|win.MB_ICONINFORMATION)

	var rHandler = remote_send.New(os.Stdout)
	var wHandler = client.New(rHandler)

	var toggleKey = os.Getenv("TOGGLE_KEY")
	if toggleKey == "" {
		toggleKey = "F8"
	}

	var toggleType = os.Getenv("TOGGLE_TYPE")
	switch toggleType {
	case "ONCE":
		wHandler.SetToggleType(client.ToggleTypeOnce)
	default:
		wHandler.SetToggleType(client.ToggleTypeAlive)
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

func StartServer() {
	userinfo, err := user.Current()
	if err != nil {
		panic(err)
	}

	conn, err := npipe.Dial(`\\.\pipe\RemoteRelativeInput\` + userinfo.Uid)
	if err != nil {
		panic(err)
	}

	go func() {
		var worker = bufio.NewScanner(conn)
		for worker.Scan() {
			fmt.Println(worker.Text())
		}
	}()

	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
}

func StartWorker() {
	userinfo, err := user.Current()
	if err != nil {
		panic(err)
	}

	ln, err := npipe.Listen(`\\.\pipe\RemoteRelativeInput\` + userinfo.Uid)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(errors.Wrap(err, "Accept"))
			continue
		}

		var res = win.MessageBox(win.HWND(winapi.NULL), winapi.MustUTF16PtrFromString("Click to start client"), winapi.MustUTF16PtrFromString("Confirmation"), win.MB_YESNO|win.MB_ICONWARNING)
		if res == win.IDNO {
			continue
		}

		fmt.Fprintln(conn, "Start Connection")
		fmt.Fprintln(os.Stderr, "Start Connection")

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			var augs = strings.Split(scanner.Text(), " ")
			if len(augs) < 4 {
				continue
			}

			eventType, err := strconv.ParseUint(augs[0], 10, 32)
			if err != nil {
				continue
			}
			eventInput, err := strconv.ParseUint(augs[1], 10, 32)
			if err != nil {
				continue
			}
			eventValue1, err := strconv.ParseInt(augs[2], 10, 32)
			if err != nil {
				continue
			}
			eventValue2, err := strconv.ParseInt(augs[3], 10, 32)
			if err != nil {
				continue
			}

			switch keymap.EV_TYPE(eventType) {
			case keymap.EV_TYPE_MOUSE_MOVE:
				switch uint32(eventInput) {
				case uint32(remote_send.MouseMoveTypeRelative):
					var dx = eventValue1
					var dy = eventValue2

					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
						Mi: win.MOUSEINPUT{
							Dx:      int32(dx),
							Dy:      int32(dy),
							DwFlags: win.MOUSEEVENTF_MOVE,
						},
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
					debug.Debugf("SendInput: MouseREL: dx: %d dy: %d\n", dx, dy)
				case uint32(remote_send.MouseMoveTypeAbsolute):
					var x = eventValue1
					var y = eventValue2

					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
						Mi: win.MOUSEINPUT{
							Dx:      int32(x),
							Dy:      int32(y),
							DwFlags: win.MOUSEEVENTF_MOVE | win.MOUSEEVENTF_ABSOLUTE,
						},
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
					debug.Debugf("SendInput: MouseREL: x: %d y: %d\n", x, y)
				}
			case keymap.EV_TYPE_MOUSE:
				switch eventInput {
				// Mouse Right
				case 0x02:
					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
					}

					switch uint32(eventValue1) {
					case uint32(remote_send.KeyDown):
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_RIGHTDOWN,
						}
						debug.Debugln("SendInput: MouseRightDown")
					case uint32(remote_send.KeyUp):
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_RIGHTUP,
						}
						debug.Debugln("SendInput: MouseRightUp")
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))

				// Mouse Left
				case 0x01:
					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
					}

					switch uint32(eventValue1) {
					case uint32(remote_send.KeyDown):
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_LEFTDOWN,
						}
						debug.Debugln("SendInput: MouseLeftDown")
					case uint32(remote_send.KeyUp):
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_LEFTUP,
						}
						debug.Debugln("SendInput: MouseLeftUp")
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))

				// Mouse Middle
				case 0x04:
					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
					}

					switch uint32(eventValue1) {
					case uint32(remote_send.KeyDown):
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_MIDDLEDOWN,
						}
						debug.Debugln("SendInput: MouseMiddleDown")
					case uint32(remote_send.KeyUp):
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_MIDDLEUP,
						}
						debug.Debugln("SendInput: MouseMiddleUp")
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
				}
			case keymap.EV_TYPE_WHEEL:
				var mouseInput = win.MOUSE_INPUT{
					Type: win.INPUT_MOUSE,
					Mi: win.MOUSEINPUT{
						DwFlags: win.MOUSEEVENTF_WHEEL,
					},
				}
				switch uint32(eventValue1) {
				case uint32(remote_send.KeyDown):
					mouseInput.Mi.MouseData = ^uint32(120) + 1
					debug.Debugln("SendInput: MouseMiddleDown")
				case uint32(remote_send.KeyUp):
					mouseInput.Mi.MouseData = 120
					debug.Debugln("SendInput: MouseMiddleUp")
				}

				win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))

			case keymap.EV_TYPE_KEY:
				var mappedKey = winapi.MapVirtualKey(uint32(eventInput), winapi.MAPVK_VK_TO_VSC)

				var keyInput = win.KEYBD_INPUT{
					Type: win.INPUT_KEYBOARD,
					Ki: win.KEYBDINPUT{
						WScan: uint16(mappedKey),
					},
				}

				switch uint32(eventValue1) {
				case uint32(remote_send.KeyDown):
					keyInput.Ki.DwFlags = win.KEYEVENTF_SCANCODE
					win.SendInput(1, unsafe.Pointer(&keyInput), int32(unsafe.Sizeof(win.KEYBD_INPUT{})))
				case uint32(remote_send.KeyUp):
					keyInput.Ki.DwFlags = win.KEYEVENTF_KEYUP | win.KEYEVENTF_SCANCODE
					win.SendInput(1, unsafe.Pointer(&keyInput), int32(unsafe.Sizeof(win.KEYBD_INPUT{})))
				}
			}
		}
		conn.Close()
	}
}
