package windows

import "github.com/lxn/win"

type POINT struct {
	win.POINT
}

func (wp POINT) GetAxis() (x, y int) {
	return int(wp.X), int(wp.Y)
}
