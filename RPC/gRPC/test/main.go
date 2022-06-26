package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"

	"wanshantian/grpc/pb"
)

func main() {
	tian := pb.UserInfo{
		Name: "tian",
	}
	enc, _ := proto.Marshal(&tian)
	fmt.Println(enc)

	new := pb.UserInfo{}
	_ = proto.Unmarshal(enc, &new)
	fmt.Println(new)

}
