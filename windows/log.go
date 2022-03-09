package windows

import "fmt"

// logger is a logger interface compatible with both stdlib and some
// 3rd party loggers.
type logger interface {
	Output(int, string) error
}

func (h *Handler) Debugf(format string, v ...interface{}) {
	if h.debug {
		h.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (h *Handler) Debugln(v ...interface{}) {
	if h.debug {
		h.logger.Output(2, fmt.Sprintln(v...))
	}
}

func (h *Handler) Output(level int, v ...interface{}) {
	if h.debug {
		h.logger.Output(level, fmt.Sprintln(v...))
	}
}
