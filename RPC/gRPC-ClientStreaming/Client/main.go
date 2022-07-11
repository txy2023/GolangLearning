package main

import (
	"context"
	"fmt"
	"log"
	"shannont/grpc-ClientStreming/pb"
	"time"

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
	//返回GetAge方法对应的流
	queryStream, _ := client.GetAge(context.Background())
	// 向stream中发送message
	_ = queryStream.Send(&pb.UserInfo{Name: "foo"})
	time.Sleep(time.Second)
	_ = queryStream.Send(&pb.UserInfo{Name: "bar"})
	time.Sleep(time.Second)

	// 发送两次数据后主动关闭流并等待接收来自server端的message
	ages_sum, err := queryStream.CloseAndRecv()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("The total of ages of foo and bar is %d", ages_sum.GetAge())
}
