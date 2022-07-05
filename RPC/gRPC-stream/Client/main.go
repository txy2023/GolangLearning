package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"shannont/grpc-stream/pb"
	"time"
)

func main() {
	//建立无认证的连接
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	client := pb.NewQueryClient(conn)
	//返回GetAge方法对应的流
	queryStream, _ := client.GetAge(context.Background())
	// 创建goroutine用来向stream中发送message
	go func() {
		for i := 0; i < 3; i++ {
			_ = queryStream.Send(&pb.UserInfo{Name: "foo"})
			time.Sleep(time.Second)
		}
		// 调用指定次数后主动关闭流
		err := queryStream.CloseSend()
		if err != nil {
			log.Println(err)
		}
	}()
	// 从stream中接收message
	for {
		ageinfoRecv, err := queryStream.Recv()
		if err == io.EOF {
			log.Println("end of stream")
			break
		}
		fmt.Println(ageinfoRecv)
	}
}
