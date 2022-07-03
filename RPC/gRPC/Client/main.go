package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"wanshantian/grpc/pb"
)

func main() {
	//建立无认证的连接
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	client := pb.NewQueryClient(conn)
	//RPC方法调用
	age, _ := client.GetAge(context.Background(), &pb.UserInfo{Name: "foo"})
	fmt.Println(age)
}
