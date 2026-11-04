package rt

import (
	"fmt"
	"os"
	"runtime"
)

func PrintStack(header string) {
	os.Stdout.WriteString(header + "\n")
	stackBuf := make([]byte, 1024)
	bytesWritten := runtime.Stack(stackBuf, false)
	os.Stdout.Write(stackBuf[:bytesWritten])
}

func PrintCallerStack(f *os.File) {
	callers := make([]uintptr, 32)
	callersNum := runtime.Callers(2, callers)
	if callersNum == 0 {
		return
	}

	frames := runtime.CallersFrames(callers[:callersNum])
	for {
		frame, more := frames.Next()

		fmt.Fprintf(f, "%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
}
