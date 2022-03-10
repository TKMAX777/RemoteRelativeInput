package relative_input

import "fmt"

// logger is a logger interface compatible with both stdlib and some
// 3rd party loggers.
type logger interface {
	Output(int, string) error
}

var debug struct {
	logger logger
	debug  bool
}

func Debugf(format string, v ...interface{}) {
	if debug.debug {
		debug.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Debugln(v ...interface{}) {
	if debug.debug {
		debug.logger.Output(2, fmt.Sprintln(v...))
	}
}

func Output(level int, v ...interface{}) {
	if debug.debug {
		debug.logger.Output(level, fmt.Sprintln(v...))
	}
}
