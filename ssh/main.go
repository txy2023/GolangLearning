package main

import (
	"fmt"
	"log"
	"shannont/ssh/cmd"
)

func main() {
	user1 := &cmd.LoginInfo{
		User:     "tian",
		Ip:       "192.168.101.108",
		Port:     22,
		Password: "tian",
	}
	client, err := cmd.NewClient(user1)
	if err != nil {
		log.Panic(err)
	}
	out := client.Run("whoami")
	fmt.Println(out)
	fmt.Println(client.Run("pwd"))
	stream, err := client.NewStreamPipe()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(stream.Run("pwd"))
	fmt.Println(stream.Run("ls"))
	stream.Run("pwd")
	stream.Run("ls")
	fmt.Println(stream.Run("su"))
	fmt.Println(stream.Run("tian"))
	fmt.Println(stream.Run("whoami"))
	fmt.Println(stream.Run(""))
	fmt.Println(stream.Run(""))
	stream.Close()
}
