package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 建立http连接
	client, _ := rpc.DialHTTP("tcp", ":1234")

	// 远程调用GetAge方法
	var res string
	err := client.Call("QueryService.GetAge", "foo", &res)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res)
	_ = client.Close()
}
