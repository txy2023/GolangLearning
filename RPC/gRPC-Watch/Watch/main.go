package main

import (
	"context"
	"io"
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
	stream, _ := client.Watch(context.Background(), &pb.WatchTime{Time: 10})
	for {
		userInfoRecv, err := stream.Recv()
		if err == io.EOF {
			log.Println("end of watch")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		log.Printf("The name of %s is updated\n", userInfoRecv.GetName())
	}
}
