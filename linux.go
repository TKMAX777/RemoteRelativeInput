//go:build linux
// +build linux

package relative_input

import (
	"bufio"
	"os"
	"strings"

	"github.com/TKMAX777/RemoteRelativeInput/linuxapi"
)

func StartClient() {
	// TODO: make linux relative input client
}

func StartServer() {
	var display = linuxapi.GetDisplay()

	var eventType, eventInput, eventValue1, eventValue2 string

	var xdot = linuxapi.NewXdotool(display)

	scanner := bufio.NewScanner(os.Stdin)

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
				xdot.MouseMoveRelative(eventValue1, eventValue2)
			case "AbsXY":
				xdot.MouseMoveAbsolute(eventValue1, eventValue2)
			case "right":
				switch eventValue1 {
				case "down":
					xdot.MouseDown(linuxapi.XdotoolMouseClickRight)
				case "up":
					xdot.MouseUp(linuxapi.XdotoolMouseClickRight)
				}
			case "left":
				switch eventValue1 {
				case "down":
					xdot.MouseDown(linuxapi.XdotoolMouseClickLeft)
				case "up":
					xdot.MouseUp(linuxapi.XdotoolMouseClickLeft)
				}
			case "middle":
				switch eventValue1 {
				case "down":
					xdot.MouseDown(linuxapi.XdotoolMouseClickMiddle)
				case "up":
					xdot.MouseUp(linuxapi.XdotoolMouseClickMiddle)
				}
			case "wheel":
				switch eventValue1 {
				case "down":
					xdot.WheelDown()
				case "up":
					xdot.WheelUp()
				}
			}
		case "EV_KEY":
			switch eventValue1 {
			case "down":
				xdot.KeyDown(eventInput)
			case "up":
				xdot.KeyUp(eventInput)
			}
		}
	}
}
