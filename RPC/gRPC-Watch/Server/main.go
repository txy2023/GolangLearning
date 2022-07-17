package main

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"
	"time"
	"wanshantian/grpc-watch/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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
	mu                          sync.Mutex
	ch                          chan string
	pb.UnimplementedQueryServer // 涉及版本兼容
}

func (q *Query) GetAge(ctx context.Context, info *pb.UserInfo) (*pb.AgeInfo, error) {
	age := userinfo[info.GetName()]
	var res = new(pb.AgeInfo)
	res.Age = age
	return res, nil
}

func (q *Query) Update(ctx context.Context, info *pb.UserInfo) (*emptypb.Empty, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	name := info.GetName()
	userinfo[name] += 1
	if q.ch != nil {
		q.ch <- name
	}
	return &emptypb.Empty{}, nil
}

func (q *Query) Watch(timeSpecify *pb.WatchTime, serverStream pb.Query_WatchServer) error {
	if q.ch != nil {
		return errors.New("Watching is running, please stop first")
	}
	q.ch = make(chan string, 1)
	for {
		select {
		case <-time.After(time.Second * time.Duration(timeSpecify.GetTime())):
			close(q.ch)
			q.ch = nil
			return nil
		case nameModify := <-q.ch:
			log.Printf("The name of %s is updated\n", nameModify)
			serverStream.Send(&pb.UserInfo{Name: nameModify})
		}
	}
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
	pb.RegisterQueryServer(grpcserver, &Query{})
	// 开启gRPC服务
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Panic(err)
	}
}
