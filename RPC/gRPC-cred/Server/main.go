package main

import (
	"context"
	"log"
	"net"
	"shannont/gRPC-cred/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// 用户信息
var userinfo = map[string]int32{
	"foo": 18,
	"bar": 20,
}

// Query 结构体，实现QueryServer接口
// type QueryServer interface {
//   GetAge(context.Context, *UserInfo) (*AgeInfo, error)
//   mustEmbedUnimplementedQueryServer()
// }
type Query struct {
	pb.UnimplementedQueryServer // 涉及版本兼容
}

func (q *Query) GetAge(ctx context.Context, info *pb.UserInfo) (*pb.AgeInfo, error) {
	age := userinfo[info.GetName()]
	var res = new(pb.AgeInfo)
	res.Age = age
	return res, nil
}

func main() {
	// 创建socket监听器
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panic(err)
	}
	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Panic(err)
	}
	// new一个gRPC服务器，用来注册服务
	grpcserver := grpc.NewServer(grpc.Creds(creds))
	// 注册服务方法
	pb.RegisterQueryServer(grpcserver, new(Query))
	// 开启gRPC服务
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}
