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

	"github.com/TKMAX777/RemoteRelativeInput/keymap"
	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
	"github.com/TKMAX777/RemoteRelativeInput/winapi"
	client "github.com/TKMAX777/RemoteRelativeInput/windows"
	"github.com/lxn/win"
	"github.com/natefinch/npipe"
	"github.com/pkg/errors"
)

func StartClient() {
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
	// wHandler.SetLogger(log.New(os.Stderr, "", 6))

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

	// debug.debug = true
	// debug.logger = log.New(os.Stderr, "", 5)

	var eventType, eventInput, eventValue1, eventValue2 string

	fmt.Println("Start Worker")

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
			var augs = strings.Split(scanner.Text(), ";")
			if len(augs) < 4 {
				continue
			}
			eventType, eventInput, eventValue1, eventValue2 = augs[0], augs[1], augs[2], augs[3]
			switch eventType {
			case "EV_MOUSE":
				switch eventInput {
				case "RelXY":
					dx, err := strconv.Atoi(eventValue1)
					if err != nil {
						log.Println(errors.Wrap(err, "ParseMouseXerror"))
						continue
					}

					dy, err := strconv.Atoi(eventValue2)
					if err != nil {
						log.Println(errors.Wrap(err, "ParseMouseYerror"))
						continue
					}

					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
						Mi: win.MOUSEINPUT{
							Dx:      int32(dx),
							Dy:      int32(dy),
							DwFlags: win.MOUSEEVENTF_MOVE,
						},
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
					debug.logger.Output(2, fmt.Sprintf("SendInput: MouseREL: dx: %d dy: %d", dx, dy))
				case "AbsXY":
					x, err := strconv.Atoi(eventValue1)
					if err != nil {
						log.Println(errors.Wrap(err, "ParseMouseXerror"))
						continue
					}

					y, err := strconv.Atoi(eventValue2)
					if err != nil {
						log.Println(errors.Wrap(err, "ParseMouseYerror"))
						continue
					}

					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
						Mi: win.MOUSEINPUT{
							Dx:      int32(x),
							Dy:      int32(y),
							DwFlags: win.MOUSEEVENTF_MOVE | win.MOUSEEVENTF_ABSOLUTE,
						},
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
					debug.logger.Output(2, fmt.Sprintf("SendInput: MouseREL: x: %d y: %d", x, y))
				case "right":
					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
					}

					switch eventValue1 {
					case "down":
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_RIGHTDOWN,
						}
						debug.logger.Output(2, "SendInput: MouseRightDown")
					case "up":
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_RIGHTUP,
						}
						debug.logger.Output(2, "SendInput: MouseRightUp")
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
				case "left":
					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
					}

					switch eventValue1 {
					case "down":
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_LEFTDOWN,
						}
						debug.logger.Output(2, "SendInput: MouseLeftDown")
					case "up":
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_LEFTUP,
						}
						debug.logger.Output(2, "SendInput: MouseLeftUp")
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
				case "middle":
					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
					}

					switch eventValue1 {
					case "down":
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_MIDDLEDOWN,
						}
						debug.logger.Output(2, "SendInput: MouseMiddleDown")
					case "up":
						mouseInput.Mi = win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_MIDDLEUP,
						}
						debug.logger.Output(2, "SendInput: MouseMiddleUp")
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
				case "wheel":
					var mouseInput = win.MOUSE_INPUT{
						Type: win.INPUT_MOUSE,
						Mi: win.MOUSEINPUT{
							DwFlags: win.MOUSEEVENTF_WHEEL,
						},
					}
					switch eventValue1 {
					case "down":
						mouseInput.Mi.MouseData = ^uint32(120) + 1
						debug.logger.Output(2, "SendInput: MouseMiddleDown")
					case "up":
						mouseInput.Mi.MouseData = 120
						debug.logger.Output(2, "SendInput: MouseMiddleUp")
					}

					win.SendInput(1, unsafe.Pointer(&mouseInput), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
				}
			case "EV_KEY":
				key, err := keymap.GetWindowsKeyDetailFromEventInput(eventInput)
				if err != nil {
					log.Println(errors.New("KeyNotFound"))
					continue
				}

				var mappedKey = winapi.MapVirtualKey(key.Value, winapi.MAPVK_VK_TO_VSC)

				var keyInput = win.KEYBD_INPUT{
					Type: win.INPUT_KEYBOARD,
					Ki: win.KEYBDINPUT{
						WScan: uint16(mappedKey),
					},
				}

				switch eventValue1 {
				case "down":
					keyInput.Ki.DwFlags = win.KEYEVENTF_SCANCODE
					win.SendInput(1, unsafe.Pointer(&keyInput), int32(unsafe.Sizeof(win.KEYBD_INPUT{})))
					debug.logger.Output(2, fmt.Sprintf("SendInput: KeyDown: %+v", key))
				case "up":
					keyInput.Ki.DwFlags = win.KEYEVENTF_KEYUP | win.KEYEVENTF_SCANCODE
					win.SendInput(1, unsafe.Pointer(&keyInput), int32(unsafe.Sizeof(win.KEYBD_INPUT{})))
					debug.logger.Output(2, fmt.Sprintf("SendInput: KeyUp: %+v", key))
				}
			}
		}

		fmt.Println("Connection closed")
		conn.Close()
	}
}
