package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
	"wanshantian/grpc-consul/pb"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

// 用户信息
var userinfo = map[string]int32{
	"foo": 18,
	"bar": 20,
}
var grpcIp = flag.String("address", "127.0.0.1", "The address of gRPC Server")
var grpcPort = flag.Int("port", 1234, "The port of gRPC Server")

// Query 结构体，实现QueryServer接口
// type QueryServer interface {
//  GetAge(context.Context, *UserInfo) (*AgeInfo, error)
//	mustEmbedUnimplementedQueryServer()
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
	flag.Parse()

	// Consul Client
	consulConfig := api.Config{
		Address: "192.168.101.108:8500",
	}
	registry, err := api.NewClient(&consulConfig)
	if err != nil {
		log.Fatalln(err)
	}
	// 注册到 Consul，包含地址、端口信息，以及健康检查
	uuid := uuid.New().String()

	log.Println(*grpcIp, *grpcPort)

	err = registry.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid,
		Name:    "QueryServer",
		Port:    *grpcPort,
		Address: *grpcIp,
		Check: &api.AgentServiceCheck{
			TTL:     (31 * time.Second).String(),
			Timeout: time.Minute.String(),
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		checkid := "service:" + uuid
		for range time.Tick(time.Duration(30)) {
			err := registry.Agent().PassTTL(checkid, "")
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
	go func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for {
			select {
			case <-time.After(time.Duration(30 * time.Second)):
				log.Println("chaoshi")
				return nil
			case <-c:
				registry.Agent().ServiceDeregister(uuid)
			}
		}
	}()
	// 创建socket监听器
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
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
