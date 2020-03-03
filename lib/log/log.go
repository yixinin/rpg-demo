package log

import (
	"fmt"
	"strings"
)

type Logger interface {
	Panic(format string, args ...interface{})
	Error(format string, args ...interface{})
	Wran(format string, args ...interface{})
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
}

var logger Logger

func init() {

}

func Init(lg Logger) {
	logger = lg
}

func Panicf(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Panic(format, args...)
}

func Panic(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	var format string
	if s, ok := args[0].(string); ok {
		format = s
		args = args[1:]
	} else {
		format = "debug info: %v"
	}
	Panicf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Error(format, args...)
}

func Error(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	var format string
	if s, ok := args[0].(string); ok {
		format = s
		args = args[1:]
	} else {
		format = "debug info: %v"
	}
	Errorf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Wran(format, args...)
}

func Warn(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	var format string
	if s, ok := args[0].(string); ok {
		format = s
		args = args[1:]
	} else {
		format = "debug info: %v"
	}
	Warnf(format, args...)
}

func Infof(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Info(format, args...)
}
func Info(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	var format string
	if s, ok := args[0].(string); ok {
		format = s
		args = args[1:]
	} else {
		format = "debug info: %v"
	}
	Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Debug(format, args...)
}

func Debug(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	var format string
	if s, ok := args[0].(string); ok {
		format = s
		args = args[1:]
	} else {
		format = "debug info: %v"
	}
	Debugf(format, args...)
}
