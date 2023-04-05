package main

import (
	"log"
	"time"
)

var count int

func main() {

	// 3s后向ch中传值
	go func() {
		time.Sleep(time.Second * 3)
		count = 1
	}()
	// 循环检索ch
	timeout := time.After(time.Second * 2)
	i := 0
	for {
		select {
		case <-timeout:
			log.Panicln("timeout")
		default:
			if count == 1 {
				log.Printf("end")
				return
			} else {
				time.Sleep(time.Second)
				i++
				log.Printf("time elapsed: %d", i)

			}
		}
	}
}
