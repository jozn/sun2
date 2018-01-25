// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_rpc_other.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PB_OtherParam_Echo struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *PB_OtherParam_Echo) Reset()                    { *m = PB_OtherParam_Echo{} }
func (m *PB_OtherParam_Echo) String() string            { return proto.CompactTextString(m) }
func (*PB_OtherParam_Echo) ProtoMessage()               {}
func (*PB_OtherParam_Echo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *PB_OtherParam_Echo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type PB_OtherResponse_Echo struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *PB_OtherResponse_Echo) Reset()                    { *m = PB_OtherResponse_Echo{} }
func (m *PB_OtherResponse_Echo) String() string            { return proto.CompactTextString(m) }
func (*PB_OtherResponse_Echo) ProtoMessage()               {}
func (*PB_OtherResponse_Echo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *PB_OtherResponse_Echo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*PB_OtherParam_Echo)(nil), "PB_OtherParam_Echo")
	proto.RegisterType((*PB_OtherResponse_Echo)(nil), "PB_OtherResponse_Echo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPC_Other service

type RPC_OtherClient interface {
	Echo(ctx context.Context, in *PB_OtherParam_Echo, opts ...grpc.CallOption) (*PB_OtherResponse_Echo, error)
}

type rPC_OtherClient struct {
	cc *grpc.ClientConn
}

func NewRPC_OtherClient(cc *grpc.ClientConn) RPC_OtherClient {
	return &rPC_OtherClient{cc}
}

func (c *rPC_OtherClient) Echo(ctx context.Context, in *PB_OtherParam_Echo, opts ...grpc.CallOption) (*PB_OtherResponse_Echo, error) {
	out := new(PB_OtherResponse_Echo)
	err := grpc.Invoke(ctx, "/RPC_Other/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPC_Other service

type RPC_OtherServer interface {
	Echo(context.Context, *PB_OtherParam_Echo) (*PB_OtherResponse_Echo, error)
}

func RegisterRPC_OtherServer(s *grpc.Server, srv RPC_OtherServer) {
	s.RegisterService(&_RPC_Other_serviceDesc, srv)
}

func _RPC_Other_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PB_OtherParam_Echo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPC_OtherServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC_Other/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPC_OtherServer).Echo(ctx, req.(*PB_OtherParam_Echo))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPC_Other_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RPC_Other",
	HandlerType: (*RPC_OtherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _RPC_Other_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_rpc_other.proto",
}

func init() { proto.RegisterFile("pb_rpc_other.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0x8a, 0x2f,
	0x2a, 0x48, 0x8e, 0xcf, 0x2f, 0xc9, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2,
	0x2b, 0x48, 0x8a, 0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0x86, 0xf0, 0x95, 0x34, 0xb8, 0x84, 0x02, 0x9c,
	0xe2, 0xfd, 0x41, 0x2a, 0x02, 0x12, 0x8b, 0x12, 0x73, 0xe3, 0x5d, 0x93, 0x33, 0xf2, 0x85, 0x84,
	0xb8, 0x58, 0x42, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25,
	0x6d, 0x2e, 0x51, 0x98, 0xca, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x9c, 0x8a, 0x8d,
	0x1c, 0xb8, 0x38, 0x83, 0x02, 0x9c, 0x21, 0xaa, 0x85, 0x8c, 0xb9, 0x58, 0xc0, 0x0a, 0x85, 0xf5,
	0x30, 0xad, 0x92, 0x12, 0xd3, 0xc3, 0x6a, 0xaa, 0x93, 0x20, 0x17, 0x47, 0x66, 0x91, 0x1e, 0xc8,
	0x9d, 0x49, 0x1e, 0xcc, 0x01, 0x8c, 0x51, 0x8c, 0x15, 0x49, 0x6c, 0x60, 0x27, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x9e, 0xad, 0xd1, 0xf3, 0xd8, 0x00, 0x00, 0x00,
}
