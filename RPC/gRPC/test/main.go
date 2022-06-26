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
	tianmess := tian.ProtoReflect()


	newname := pb.UserInfo{}
	_ = proto.Unmarshal(enc, &newname)
	fmt.Println(newname)

}
