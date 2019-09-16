// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871) }

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 63 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x15,
	0xcc, 0x31, 0x62, 0xe7, 0x82, 0x30, 0x9c, 0x38, 0xa3, 0xd8, 0xa1, 0xca, 0x92, 0xd8, 0xc0, 0x2a,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80, 0x7b, 0xe0, 0xb0, 0x3e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetClient is the client API for Greet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetClient interface {
}

type greetClient struct {
	cc *grpc.ClientConn
}

func NewGreetClient(cc *grpc.ClientConn) GreetClient {
	return &greetClient{cc}
}

// GreetServer is the server API for Greet service.
type GreetServer interface {
}

// UnimplementedGreetServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServer struct {
}

func RegisterGreetServer(s *grpc.Server, srv GreetServer) {
	s.RegisterService(&_Greet_serviceDesc, srv)
}

var _Greet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.greet",
	HandlerType: (*GreetServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "greet/greetpb/greet.proto",
}
