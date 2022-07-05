package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"shannont/grpc-bidirectionalStreaming/pb"
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

func (q *Query) GetAge(server pb.Query_GetAgeServer) error {
	log.Println("start of stream")
	for {
		userinfoRecv, err := server.Recv()
		// 带客户端主动关闭流后，退出for循环
		if err == io.EOF {
			log.Println("end of stream")
			break
		}
		err = server.Send(&pb.AgeInfo{Age: userinfo[userinfoRecv.Name]})
		if err != nil {
			log.Panic(err)
		}
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
