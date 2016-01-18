package peer

import (
	"fmt"
	"runtime"

	"github.com/pankona/gomo-simra/simra/config"
)

func printLog(tag string, format string, a ...interface{}) {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	fmt.Printf("[%s][%s:%d] ", tag, f.Name(), line)
	if len(a) == 0 {
		fmt.Print(format)
	} else {
		fmt.Printf(format, a)
	}
	fmt.Printf("\n")
}

func LogDebug(format string, a ...interface{}) {
	if config.DEBUG {
		printLog("DEBUG", format, a...)
	}

}

func LogError(format string, a ...interface{}) {
	printLog("ERROR", format, a...)
}
