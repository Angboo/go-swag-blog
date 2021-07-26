package vxlog

import (
	"fmt"
	"log"
	"runtime"
)

func init() {
	log.SetFlags(0)
}

func Printf(format string, v ...interface{}) {
	name, file, line := getCallerInfo(3)
	s := fmt.Sprintf("[caller]: %s [file]: %s [line]: %d -> ", name, file, line)
	log.Printf(s+format, v...)
}

func getCallerInfo(skips ...int) (name string, file string, line int) {
	skip := 2
	if len(skips) > 0 {
		skip = skips[0]
	}
	fpcs := make([]uintptr, 1)
	if runtime.Callers(skip, fpcs) == 0 {
		return
	}
	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller != nil {
		file, line = caller.FileLine(fpcs[0] - 1)
		name = caller.Name()
	}
	return name, file, line
}
