package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 建立socket连接
	conn, _ := net.Dial("tcp", ":1234")
	// 指定json作编解码器
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	// 远程调用方法
	var res string
	_ = client.Call("QueryService.GetAge", "huan", &res)
	fmt.Println(res)
	_ = client.Close()
}
