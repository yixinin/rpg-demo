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

func Panic(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Panic(format, args...)
}
func Error(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Error(format, args...)
}
func Wran(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Wran(format, args...)
}
func Info(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Info(format, args...)
}
func Debug(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if logger == nil {
		fmt.Printf(format, args...)
		return
	}
	logger.Debug(format, args...)
}
