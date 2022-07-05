package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"shannont/grpc-ClientStreming/pb"
)

// 用户信息
var userinfo = map[string]int32{
	"foo": 18,
	"bar": 20,
}

// Query 结构体，实现QueryServer接口
// type QueryServer interface {
//  GetAge(context.Context, *UserInfo) (*AgeInfo, error)
//	mustEmbedUnimplementedQueryServer()
// }
type Query struct {
	pb.UnimplementedQueryServer // 涉及版本兼容
}

func (q *Query) GetAge(serverStream pb.Query_GetAgeServer) error {
	log.Println("start of stream")
	var names_received []*pb.UserInfo
	for {
		userinfoRecv, err := serverStream.Recv()
		// 待客户端主动关闭流后，退出for循环
		if err == io.EOF {
			log.Println("end of stream")
			break
		}
		names_received = append(names_received, userinfoRecv)
	}
	// 统计年龄和
	var ages_sum int32
	for _, v := range names_received {
		ages_sum += userinfo[v.GetName()]
	}
	// 返回message
	err := serverStream.SendAndClose(&pb.AgeInfo{Age: ages_sum})
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func main() {
	// 创建socket监听器
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panic(err)
	}
	// new一个gRPC服务器，用来注册服务
	grpcserver := grpc.NewServer()
	// 注册服务方法
	pb.RegisterQueryServer(grpcserver, new(Query))
	// 开启gRPC服务
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}
