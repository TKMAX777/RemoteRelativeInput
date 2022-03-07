package remote_send

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Handler struct {
	writer io.Writer
}

type EventInputSender interface {
	Transrate() string
}

type EventCursorSender interface {
	GetAxis() (x, y int)
}

func New(w io.Writer) *Handler {
	return &Handler{w}
}

func (h Handler) ParseMousePosition(out string) (x, y int32) {
	pointOut := strings.Split(out, ";")
	if len(pointOut) < 4 {
		return
	}
	if pointOut[0] != "POS" {
		return
	}

	X, _ := strconv.Atoi(pointOut[1])
	x = int32(X)

	Y, _ := strconv.Atoi(pointOut[2])
	y = int32(Y)

	return
}

func (h Handler) SendRelativeCursor(ecs EventCursorSender) {
	var x, y = ecs.GetAxis()
	fmt.Fprintf(h.writer, "EV_MOUSE;RelXY;%d;%d\n", x, y)
}

func (h Handler) SendAbsoluteCursor(ecs EventCursorSender) {
	var x, y = ecs.GetAxis()
	fmt.Fprintf(h.writer, "EV_MOUSE;AbsXY;%d;%d\n", x, y)
}

func (h Handler) SendInput(eis EventInputSender) {
	var send = eis.Transrate()
	if send == "" {
		return
	}
	fmt.Fprintln(h.writer, send)
}
