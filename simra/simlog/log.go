package simlog

import (
	"fmt"
	"runtime"
)

type Logger interface {
	FuncIn()
	FuncOut()
	Debug(format string, a ...interface{})
	Error(format string, a ...interface{})
}

type logger struct {
}

var simlog = &logger{}

func FuncIn() {
	simlog.printLog("[DEBUG]", "%s", "IN")
}

func FuncOut() {
	simlog.printLog("[DEBUG]", "%s", "OUT")
}

func Debugf(format string, a ...interface{}) {
	simlog.printLog("[DEBUG]", format, a...)
}

func Errorf(format string, a ...interface{}) {
	simlog.printLog("[ERROR]", format, a...)
}

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
