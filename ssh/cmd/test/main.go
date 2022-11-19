package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBufferString("123456789")
	a := make([]byte, 2)
	// a := bytes.NewBuffer(make([]byte, 1))
	buf.Read(a)
	fmt.Println(a)
	// buf.Reset()
	buf.Read(a)
	fmt.Println(a)
	fmt.Println(buf.String())
	// strings.Contains()
	buf.Read(a)
	fmt.Println(a)

}
