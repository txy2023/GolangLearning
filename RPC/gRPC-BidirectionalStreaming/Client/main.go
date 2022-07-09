package main

import (
	"context"
	"io"
	"log"
	"shannont/grpc-bidirectionalStreaming/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//建立无认证的连接
	conn, _ := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()
	client := pb.NewQueryClient(conn)

	//返回GetAge方法对应的流
	log.Printf("start of stream")
	queryStream, _ := client.GetAge(context.Background())

	// 创建goroutine用来向stream中发送message
	ch := make(chan string, 2)
	go func() {
		names := []string{"foo", "bar"}
		for _, v := range names {
			log.Printf("send message wtih Name is %s\n", v)
			ch <- v
			_ = queryStream.Send(&pb.UserInfo{Name: v})
			time.Sleep(time.Second)
		}
		// 主动关闭流
		err := queryStream.CloseSend()
		if err != nil {
			log.Println(err)
		}
		close(ch)
	}()

	// 从stream中接收message
	for {
		name := <-ch
		ageinfoRecv, err := queryStream.Recv()
		if err == io.EOF {
			log.Println("end of stream")
			break
		}
		log.Printf("The age of %s is %d\n", name, ageinfoRecv.GetAge())
	}
}
