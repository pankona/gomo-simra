package peer

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/pankona/gomo-simra/simra/config"
)

var printLogMutex sync.Mutex

func printLog(tag string, format string, a ...interface{}) {
	printLogMutex.Lock()
	defer printLogMutex.Unlock()

	pc := make([]uintptr, 10)
	runtime.Callers(4, pc)
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

// LogDebug prints logs.
// This is disabled at Release Build.
func LogDebug(format string, a ...interface{}) {
	if config.DEBUG {
		printLog("DEBUG", format, a...)
	}

}

// LogError prints logs.
// This is never disabled even for Release build.
func LogError(format string, a ...interface{}) {
	printLog("ERROR", format, a...)
}
