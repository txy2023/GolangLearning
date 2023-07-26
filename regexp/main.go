package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("./hello.txt")
	contents, _ := ioutil.ReadAll(file)
	wanted := regexp.MustCompile(`(?m:^hello.*)`)
	res := wanted.FindAllString(string(contents), -1)
	fmt.Println(res)
	log.Println(res)
	fmt.Println(len(res))
	fmt.Println(res[0])
}
