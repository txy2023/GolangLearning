package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"wanshantian/grpc/pb"
)

var userinfo = map[string]int32{
	"foo": 18,
	"bar": 20,
}

type Query struct {
	pb.UnimplementedQueryServer
}

func (q *Query) GetAge(ctx context.Context, info *pb.UserInfo) (*pb.AgeInfo, error) {
	age := userinfo[info.GetName()]
	var res = new(pb.AgeInfo)
	res.Age = age
	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panic(err)
	}
	grpcserver := grpc.NewServer()
	pb.RegisterQueryServer(grpcserver, new(Query))
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}
