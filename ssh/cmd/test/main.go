package main

import (
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	a := "{'da':'test'}"
	fmt.Println(a)
	aa := strings.ReplaceAll(a, "'", "\"")
	fmt.Println(aa)
	// bs, _ := json.Marshal(aa)
	v := gojsonq.New().FromString(aa)
	fmt.Println(v.Find("da"))

}
