package tool

import (
	"fmt"
	"runtime"
)

func CallersTest() {
	pc := make([]uintptr, 32)
	n := runtime.Callers(0, pc)
	cf := runtime.CallersFrames(pc[:n])
	for {
		frame, more := cf.Next()
		if !more {
			break
		} else {
			fmt.Println(frame)
		}
	}

}
