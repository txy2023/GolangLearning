package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
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
	// 注册服务方法
	if err := rpc.RegisterName("QueryService", new(Query)); err != nil {
		log.Println(err)
	}
	// 开启监听，接受来自rpc客户端的请求
	listener, _ := net.Listen("tcp", ":1234")
	rpc.Accept(listener)
}
