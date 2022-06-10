package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 使用默认的gob作为编解码器建立与RPC服务端的连接
	client, _ := rpc.Dial("tcp", ":1234")
	// 调用方法
	var res string
	err := client.Call("QueryService.GetAge", "bar", &res)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}
