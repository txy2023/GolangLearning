package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"wanshantian/grpc/pb"
)

func main() {
	conn, _ := grpc.Dial(":1234", grpc.WithInsecure())
	client := pb.NewQueryClient(conn)
	age, _ := client.GetAge(context.Background(), &pb.UserInfo{Name: "foo"})
	fmt.Println(age)
}
