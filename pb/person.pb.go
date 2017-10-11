// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/person.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/person.proto

It has these top-level messages:
	Person
	GreetingRequest
	GreetingReply
*/
package pb

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person struct {
	Firstname string `protobuf:"bytes,1,opt,name=firstname" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname" json:"lastname,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Person) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type GreetingRequest struct {
	Person string `protobuf:"bytes,1,opt,name=person" json:"person,omitempty"`
}

func (m *GreetingRequest) Reset()                    { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()               {}
func (*GreetingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetingRequest) GetPerson() string {
	if m != nil {
		return m.Person
	}
	return ""
}

type GreetingReply struct {
	Person  *Person `protobuf:"bytes,1,opt,name=person" json:"person,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *GreetingReply) Reset()                    { *m = GreetingReply{} }
func (m *GreetingReply) String() string            { return proto.CompactTextString(m) }
func (*GreetingReply) ProtoMessage()               {}
func (*GreetingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GreetingReply) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *GreetingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*GreetingRequest)(nil), "pb.GreetingRequest")
	proto.RegisterType((*GreetingReply)(nil), "pb.GreetingReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeting service

type GreetingClient interface {
	SayHi(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingReply, error)
}

type greetingClient struct {
	cc *grpc.ClientConn
}

func NewGreetingClient(cc *grpc.ClientConn) GreetingClient {
	return &greetingClient{cc}
}

func (c *greetingClient) SayHi(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingReply, error) {
	out := new(GreetingReply)
	err := grpc.Invoke(ctx, "/pb.Greeting/SayHi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeting service

type GreetingServer interface {
	SayHi(context.Context, *GreetingRequest) (*GreetingReply, error)
}

func RegisterGreetingServer(s *grpc.Server, srv GreetingServer) {
	s.RegisterService(&_Greeting_serviceDesc, srv)
}

func _Greeting_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Greeting/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).SayHi(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Greeting",
	HandlerType: (*GreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Greeting_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/person.proto",
}

func init() { proto.RegisterFile("pb/person.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x48, 0xd2, 0x2f,
	0x48, 0x2d, 0x2a, 0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52,
	0x72, 0xe2, 0x62, 0x0b, 0x00, 0x8b, 0x09, 0xc9, 0x70, 0x71, 0xa6, 0x65, 0x16, 0x15, 0x97, 0xe4,
	0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0xa4, 0xb8, 0x38,
	0x72, 0x12, 0xa1, 0x92, 0x4c, 0x60, 0x49, 0x38, 0x5f, 0x49, 0x93, 0x8b, 0xdf, 0xbd, 0x28, 0x35,
	0xb5, 0x24, 0x33, 0x2f, 0x3d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0x0d,
	0x62, 0x15, 0xd4, 0x24, 0x28, 0x4f, 0xc9, 0x97, 0x8b, 0x17, 0xa1, 0xb4, 0x20, 0xa7, 0x52, 0x48,
	0x09, 0x45, 0x21, 0xb7, 0x11, 0x97, 0x5e, 0x41, 0x92, 0x1e, 0xc4, 0x45, 0x30, 0x4d, 0x42, 0x12,
	0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0xab, 0x61, 0x5c, 0x23, 0x5b, 0x2e, 0x0e,
	0x98, 0x71, 0x42, 0x86, 0x5c, 0xac, 0xc1, 0x89, 0x95, 0x1e, 0x99, 0x42, 0xc2, 0x20, 0x23, 0xd0,
	0x1c, 0x24, 0x25, 0x88, 0x2a, 0x58, 0x90, 0x53, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x07, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xaa, 0xda, 0x60, 0x1a, 0x01, 0x00, 0x00,
}