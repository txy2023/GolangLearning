package main

import (
	"context"
	"fmt"
	"log"
	"shannont/gRPC-cred/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//建立单向连接
	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	if err != nil {
		log.Panic(err)
	}
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	client := pb.NewQueryClient(conn)
	//RPC方法调用
	age, err := client.GetAge(context.Background(), &pb.UserInfo{Name: "foo"})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(age)
}
