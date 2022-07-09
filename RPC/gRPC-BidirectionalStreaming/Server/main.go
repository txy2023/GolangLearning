package main

import (
	"io"
	"log"
	"net"
	"shannont/grpc-bidirectionalStreaming/pb"

	"google.golang.org/grpc"
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

func (q *Query) GetAge(ServerStream pb.Query_GetAgeServer) error {
	log.Println("start of stream")
	for {
		// 接受message
		userinfoRecv, err := ServerStream.Recv()
		// 待客户端主动关闭流后，退出for循环
		if err == io.EOF {
			log.Println("end of stream")
			break
		}
		log.Printf("The name of user received is %s", userinfoRecv.GetName())
		// 返回响应message
		log.Printf("send message about the age of %s", userinfoRecv.GetName())
		err = ServerStream.Send(&pb.AgeInfo{Age: userinfo[userinfoRecv.Name]})
		if err != nil {
			log.Panic(err)
		}
	}
	return nil
}

func main() {
	// 创建socket监听器
	listener, _ := net.Listen("tcp", ":1234")
	// new一个gRPC服务器，用来注册服务
	grpcserver := grpc.NewServer()
	// 注册服务方法
	pb.RegisterQueryServer(grpcserver, new(Query))
	// 开启gRPC服务
	_ = grpcserver.Serve(listener)
}
