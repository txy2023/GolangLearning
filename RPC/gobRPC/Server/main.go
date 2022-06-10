package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Query struct {
}

func (q *Query) GetAge(req string, res *string) error {
	*res = fmt.Sprintf("The age of %s is %d", req, 18)
	return nil
}

func main() {
	if err := rpc.RegisterName("QueryService", new(Query)); err != nil {
		log.Println(err)
	}
	listener, _ := net.Listen("tcp", ":1234")
	rpc.Accept(listener)
}
