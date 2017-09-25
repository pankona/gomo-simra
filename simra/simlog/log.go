package simlog

import (
	"fmt"
	"runtime"

	"github.com/pankona/gomo-simra/simra/config"
)

// SimLogger represents log interface for simra
type SimLogger interface {
	FuncIn()
	FuncOut()
	Debugf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Debug(i interface{})
	Error(i interface{})
}

type logger struct {
	isDebug bool
}

var simlog SimLogger = &logger{
	isDebug: config.DEBUG,
}

// FuncIn shows log for function in
func FuncIn() {
	simlog.FuncIn()
}

// FuncOut shows log for function out
func FuncOut() {
	simlog.FuncOut()
}

// Debugf shows debug log with specified format and arguments
func Debugf(format string, a ...interface{}) {
	simlog.Debugf(format, a...)
}

// Errorf shows error log with specified format and arguments
func Errorf(format string, a ...interface{}) {
	simlog.Errorf(format, a...)
}

// Debug shows debug log with specified object
func Debug(i interface{}) {
	simlog.Debug(i)
}

// Error shows error log with specified object
func Error(i interface{}) {
	simlog.Error(i)
}

// FuncIn shows log for function in
func (l *logger) FuncIn() {
	if !l.isDebug {
		return
	}
	l.printLog("[DEBUG]", "%s", "IN")
}

// FuncOut shows log for function out
func (l *logger) FuncOut() {
	if !l.isDebug {
		return
	}
	l.printLog("[DEBUG]", "%s", "OUT")
}

// Debugf shows debug log with specified format and arguments
func (l *logger) Debugf(format string, a ...interface{}) {
	if !l.isDebug {
		return
	}
	l.printLog("[DEBUG]", format, a...)
}

// Errorf shows error log with specified format and arguments
func (l *logger) Errorf(format string, a ...interface{}) {
	l.printLog("[ERROR]", format, a...)
}

// Debug shows debug log with specified object
func (l *logger) Debug(i interface{}) {
	if !l.isDebug {
		return
	}
	switch v := i.(type) {
	case error:
		l.printLog("[DEBUG]", v.Error())
	case fmt.Stringer:
		l.printLog("[DEBUG]", v.String())
	case string:
		l.printLog("[DEBUG]", v)
	default:
		panic("can't print input!")
	}
}

// Error shows error log with specified object
func (l *logger) Error(i interface{}) {
	switch v := i.(type) {
	case error:
		l.printLog("[ERROR]", v.Error())
	case fmt.Stringer:
		l.printLog("[ERROR]", v.String())
	default:
		panic("can't print input!")
	}
}

func (l *logger) printLog(tag string, format string, a ...interface{}) {
	pc := make([]uintptr, 10)
	runtime.Callers(4, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	var buf string
	buf += fmt.Sprintf("[%s][%s:%d] ", tag, f.Name(), line)
	if len(a) == 0 {
		buf += format
	} else {
		buf += fmt.Sprintf(format, a)
	}
	fmt.Println(buf)
}
