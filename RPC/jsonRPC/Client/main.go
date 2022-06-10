package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, _ := net.Dial("tcp", ":1234")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var res string
	_ = client.Call("QueryService.GetAge", "huan", &res)
	fmt.Println(res)
}
