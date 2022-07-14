package main

import (
	"context"
	"fmt"
	"log"
	"wanshantian/grpc-watch/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	ctx := context.Background()
	age, _ := client.GetAge(ctx, &pb.UserInfo{Name: "foo"})
	fmt.Println(age)
	client.Update(ctx, &pb.UserInfo{Name: "foo"})
}
