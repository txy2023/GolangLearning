package main

import (
	"flag"
	"os"
)

var (
	name = flag.String("name", "tian", "default name")
	age  = flag.String("age", "10", "default age")
	sex  = flag.String("sex", "male", "default sex")
)

func main() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()
}
