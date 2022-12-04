//go:build linux
// +build linux

package relative_input

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/TKMAX777/RemoteRelativeInput/keymap"
	"github.com/TKMAX777/RemoteRelativeInput/linuxapi"
	"github.com/TKMAX777/RemoteRelativeInput/remote_send"
)

func StartClient() {
	// TODO: make linux relative input client
}

func StartServer() {
	var display = linuxapi.GetDisplay()
	var xdot = linuxapi.NewXdotool(display)

	scanner := bufio.NewScanner(os.Stdin)

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
				xdot.MouseMoveRelative(strconv.Itoa(int(eventValue1)), strconv.Itoa(int(eventValue2)))
			case uint32(remote_send.MouseMoveTypeAbsolute):
				xdot.MouseMoveAbsolute(strconv.Itoa(int(eventValue1)), strconv.Itoa(int(eventValue2)))
			}
		case keymap.EV_TYPE_MOUSE:
			switch eventInput {
			// Mouse Right
			case 0x02:
				switch uint32(eventValue1) {
				case uint32(remote_send.KeyDown):
					xdot.MouseDown(linuxapi.XdotoolMouseClickRight)
				case uint32(remote_send.KeyUp):
					xdot.MouseUp(linuxapi.XdotoolMouseClickRight)
				}
			// Mouse Left
			case 0x01:
				switch uint32(eventValue1) {
				case uint32(remote_send.KeyDown):
					xdot.MouseDown(linuxapi.XdotoolMouseClickLeft)
				case uint32(remote_send.KeyUp):
					xdot.MouseUp(linuxapi.XdotoolMouseClickLeft)
				}
			// Mouse Middle
			case 0x04:
				switch uint32(eventValue1) {
				case uint32(remote_send.KeyDown):
					xdot.MouseDown(linuxapi.XdotoolMouseClickMiddle)
				case uint32(remote_send.KeyUp):
					xdot.MouseUp(linuxapi.XdotoolMouseClickMiddle)
				}
			}
		case keymap.EV_TYPE_WHEEL:
			switch uint32(eventValue1) {
			case uint32(remote_send.KeyDown):
				xdot.WheelDown()
			case uint32(remote_send.KeyUp):
				xdot.WheelUp()
			}
		case keymap.EV_TYPE_KEY:
			key, err := keymap.GetWindowsKeyDetail(uint32(eventInput))
			if err != nil || key.EventInput == "" {
				continue
			}

			switch uint32(eventValue1) {
			case uint32(remote_send.KeyDown):
				xdot.KeyDown(key.EventInput)
			case uint32(remote_send.KeyUp):
				xdot.KeyUp(key.EventInput)
			}
		}
	}
}
