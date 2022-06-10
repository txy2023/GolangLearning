package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 用户信息
var userinfo = map[string]int{
	"foo": 18,
	"bar": 20,
}

// 实现查询服务，结构体Query实现了GetAge方法
type Query struct {
}

func (q *Query) GetAge(req string, res *string) error {
	*res = fmt.Sprintf("The age of %s is %d", req, userinfo[req])
	return nil
}

func main() {
	if err := rpc.RegisterName("QueryService", new(Query)); err != nil {
		log.Println(err)
	}
	listener, _ := net.Listen("tcp", ":1234")
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
