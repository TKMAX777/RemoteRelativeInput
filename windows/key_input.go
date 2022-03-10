package client

import (
	"fmt"

	"github.com/TKMAX777/RemoteRelativeInput/keymap"
)

type KeyInput struct {
	Key   keymap.WindowsKey
	State InputType
}

type InputType int

const (
	KeyDown InputType = iota
	KeyUp
)

func (k KeyInput) Transrate() string {
	var state string
	if k.State == KeyUp {
		state = "up"
	} else {
		state = "down"
	}

	if k.Key.EventInput == "" {
		return ""
	}

	return fmt.Sprintf("%s;%s;%s;%s", k.Key.EventType, k.Key.EventInput, state, "")
}
