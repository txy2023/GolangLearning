package main

import (
	"fmt"
	"log"
	"shannont/ssh/cmd"
	"sync"
)

func main() {
	user1 := &cmd.LoginInfo{
		User:     "tian",
		Ip:       "192.168.101.108",
		Port:     22,
		Password: "tian",
	}
	client, _ := cmd.NewClient(user1)

	s, err := client.NewStream()
	if err != nil {
		log.Panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		s.Run("sleep 5")
		fmt.Println("test1")
		wg.Done()
	}()
	go func() {
		s.Run("echo hello")
		fmt.Println("test2")
		wg.Done()
	}()
	wg.Wait()
}
