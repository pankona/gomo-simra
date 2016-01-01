package peer

import (
	"fmt"
	"runtime"
)

func printLog(tag string, message string) {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	fmt.Printf("[%s][%s:%d] %s\n", tag, f.Name(), line, message)
}

func LogDebug(message string) {
	printLog("DEBUG", message)

}

func LogError(message string) {
	printLog("ERROR", message)
}
