package tool

import (
	"bytes"
	"fmt"
	"runtime"
)

func CallerTest() {
	for i := 0; i <= 4; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if ok {
			fmt.Printf("当i:=%d时:\n调用者的pc:%v\n调用者的函数名:%s\n"+
				"调用者所在file:%s\n被调用者在调用者中的line:%d\n", i, pc, runtime.FuncForPC(pc).Name(), file, line)
			fmt.Println(string(bytes.Repeat([]byte("*"), 10)))
		}
	}
}
