package simlog

import (
	"fmt"
	"runtime"
)

// Logger represents log interface for simra
type Logger interface {
	printLog(tag string, format string, a ...interface{})
}

type logger struct {
}

var simlog Logger = &logger{}

// FuncIn shows log for function in
func FuncIn() {
	simlog.printLog("[DEBUG]", "%s", "IN")
}

// FuncOut shows log for function out
func FuncOut() {
	simlog.printLog("[DEBUG]", "%s", "OUT")
}

// Debugf shows debug log with specified format and arguments
func Debugf(format string, a ...interface{}) {
	simlog.printLog("[DEBUG]", format, a...)
}

// Errorf shows error log with specified format and arguments
func Errorf(format string, a ...interface{}) {
	simlog.printLog("[ERROR]", format, a...)
}

// Debug shows debug log with specified object
func Debug(i interface{}) {
	switch v := i.(type) {
	case error:
		simlog.printLog("[DEBUG]", v.Error())
	case fmt.Stringer:
		simlog.printLog("[DEBUG]", v.String())
	case string:
		simlog.printLog("[DEBUG]", v)
	default:
		panic("can't print input!")
	}
}

// Error shows error log with specified object
func Error(i interface{}) {
	switch v := i.(type) {
	case error:
		simlog.printLog("[ERROR]", v.Error())
	case fmt.Stringer:
		simlog.printLog("[ERROR]", v.String())
	default:
		panic("can't print input!")
	}
}

func (l *logger) printLog(tag string, format string, a ...interface{}) {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
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
