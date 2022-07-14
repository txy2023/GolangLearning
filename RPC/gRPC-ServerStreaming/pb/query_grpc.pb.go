// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: query.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// 服务端流模式
	GetAge(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (Query_GetAgeClient, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetAge(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (Query_GetAgeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Query_ServiceDesc.Streams[0], "/pb.Query/GetAge", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryGetAgeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_GetAgeClient interface {
	Recv() (*AgeInfo, error)
	grpc.ClientStream
}

type queryGetAgeClient struct {
	grpc.ClientStream
}

func (x *queryGetAgeClient) Recv() (*AgeInfo, error) {
	m := new(AgeInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// 服务端流模式
	GetAge(*UserInfo, Query_GetAgeServer) error
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) GetAge(*UserInfo, Query_GetAgeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAge not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_GetAge_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).GetAge(m, &queryGetAgeServer{stream})
}

type Query_GetAgeServer interface {
	Send(*AgeInfo) error
	grpc.ServerStream
}

type queryGetAgeServer struct {
	grpc.ServerStream
}

func (x *queryGetAgeServer) Send(m *AgeInfo) error {
	return x.ServerStream.SendMsg(m)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAge",
			Handler:       _Query_GetAge_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "query.proto",
}
