package keymap

import "github.com/pkg/errors"

func GetWindowsKeyDetail(windowsVirtualKeyCode uint32) (*WindowsKey, error) {
	w, ok := windowsKeys[windowsVirtualKeyCode]
	if !ok {
		return nil, errors.New("NotFound")
	}
	return &w, nil
}
