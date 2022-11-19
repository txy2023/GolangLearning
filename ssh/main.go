package main

import (
	"fmt"
	"log"
	"shannont/ssh/cmd"
)

func main() {
	user1 := &cmd.LoginInfo{
		User:     "tian",
		Ip:       "192.168.1.3",
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
	stream, err := client.NewStream()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(stream.Run("pwd"))
	stream.UpdateReadUntilExpect("Password:")
	fmt.Println(stream.Run("su"))
	stream.UpdateReadUntilExpect("]#")
	fmt.Println(stream.Run("tian"))
	fmt.Println(stream.Run("cd /root"))
	fmt.Println(stream.Run("./test.sh"))
	// stream.Run("pwd")
	// stream.Run("ls")

	// fmt.Println(stream.Run("whoami"))
	// fmt.Println(stream.Run(""))
	// fmt.Println(stream.Run(""))
	// stream.Close()
}
