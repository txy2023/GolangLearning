package main

import (
	"context"
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
	//先获取更新前的年龄
	age, _ := client.GetAge(ctx, &pb.UserInfo{Name: "foo"})
	log.Printf("Before updating, the age is %d\n", age.GetAge())
	//更新年龄
	log.Println("updating")
	client.Update(ctx, &pb.UserInfo{Name: "foo"})
	//再获取更新后的年龄
	age, _ = client.GetAge(ctx, &pb.UserInfo{Name: "foo"})
	log.Printf("After updating, the age is %d\n", age.GetAge())
}
