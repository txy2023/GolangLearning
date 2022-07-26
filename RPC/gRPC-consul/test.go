package main

import (
	"fmt"
	"os"
	"os/signal"
)

// var a = flag.Int("port", 1234, "test")

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := <-c
	fmt.Println("Got signal:", s)
}
