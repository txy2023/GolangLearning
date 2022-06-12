package main

import (
	"fmt"
	"log"
	"net/http"
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
	// 绑定pattern和handler
	rpc.HandleHTTP()
	// 开启监听
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Panic(err)
	}
}
