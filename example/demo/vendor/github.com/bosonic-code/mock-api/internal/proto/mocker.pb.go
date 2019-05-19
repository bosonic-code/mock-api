// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mocker.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	mocker.proto

It has these top-level messages:
	HandleRequest
	HandleResponse
	Request
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HandleRequest struct {
	Request  *Request  `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
}

func (m *HandleRequest) Reset()                    { *m = HandleRequest{} }
func (m *HandleRequest) String() string            { return proto1.CompactTextString(m) }
func (*HandleRequest) ProtoMessage()               {}
func (*HandleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HandleRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *HandleRequest) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type HandleResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *HandleResponse) Reset()                    { *m = HandleResponse{} }
func (m *HandleResponse) String() string            { return proto1.CompactTextString(m) }
func (*HandleResponse) ProtoMessage()               {}
func (*HandleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HandleResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Request struct {
	Method string            `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string            `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Query  map[string]string `protobuf:"bytes,3,rep,name=query" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetQuery() map[string]string {
	if m != nil {
		return m.Query
	}
	return nil
}

type Response struct {
	Status int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Body   string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto1.RegisterType((*HandleRequest)(nil), "proto.HandleRequest")
	proto1.RegisterType((*HandleResponse)(nil), "proto.HandleResponse")
	proto1.RegisterType((*Request)(nil), "proto.Request")
	proto1.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Mocker service

type MockerClient interface {
	Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*HandleResponse, error)
}

type mockerClient struct {
	cc *grpc.ClientConn
}

func NewMockerClient(cc *grpc.ClientConn) MockerClient {
	return &mockerClient{cc}
}

func (c *mockerClient) Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*HandleResponse, error) {
	out := new(HandleResponse)
	err := grpc.Invoke(ctx, "/proto.Mocker/Handle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mocker service

type MockerServer interface {
	Handle(context.Context, *HandleRequest) (*HandleResponse, error)
}

func RegisterMockerServer(s *grpc.Server, srv MockerServer) {
	s.RegisterService(&_Mocker_serviceDesc, srv)
}

func _Mocker_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockerServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Mocker/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockerServer).Handle(ctx, req.(*HandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mocker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Mocker",
	HandlerType: (*MockerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Mocker_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mocker.proto",
}

func init() { proto1.RegisterFile("mocker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x8d, 0xd9, 0xb6, 0x53, 0xad, 0x32, 0xd4, 0x12, 0x7b, 0x2a, 0x39, 0x48, 0x40, 0x88,
	0x10, 0x41, 0x8b, 0x37, 0x0f, 0x82, 0x17, 0x0f, 0xee, 0x3f, 0x48, 0xcd, 0x48, 0xa1, 0x6d, 0x36,
	0xdd, 0xdd, 0x08, 0xf9, 0x37, 0xfe, 0x54, 0xd9, 0x8f, 0x44, 0xda, 0x53, 0xde, 0x7b, 0xfb, 0x86,
	0xf7, 0x32, 0x03, 0x17, 0x7b, 0xf1, 0xb5, 0x25, 0x99, 0xd5, 0x52, 0x68, 0x81, 0x91, 0xfd, 0x24,
	0xdf, 0x70, 0xf9, 0x5e, 0x54, 0xe5, 0x8e, 0x38, 0x1d, 0x1a, 0x52, 0x1a, 0x53, 0x18, 0x4a, 0x07,
	0xe3, 0x60, 0x19, 0xa4, 0x93, 0x7c, 0xea, 0x06, 0x32, 0x6f, 0xe0, 0xdd, 0x33, 0xde, 0xc3, 0x48,
	0x92, 0xaa, 0x45, 0xa5, 0x28, 0x1e, 0x58, 0xeb, 0x55, 0x6f, 0x75, 0x32, 0xef, 0x0d, 0xc9, 0x1d,
	0x4c, 0xbb, 0x1c, 0xa7, 0xe0, 0x0c, 0x22, 0x92, 0x52, 0x48, 0x1b, 0x33, 0xe6, 0x8e, 0x24, 0xbf,
	0x01, 0x0c, 0xbb, 0x2a, 0x73, 0x60, 0x7b, 0xd2, 0x1b, 0x51, 0x7a, 0x8b, 0x67, 0x88, 0x70, 0x5e,
	0x17, 0x7a, 0x63, 0x43, 0xc7, 0xdc, 0x62, 0x7c, 0x80, 0xe8, 0xd0, 0x90, 0x6c, 0xe3, 0x70, 0x19,
	0xa6, 0x93, 0xfc, 0xf6, 0xb8, 0x74, 0xf6, 0x69, 0xde, 0xde, 0x2a, 0x2d, 0x5b, 0xee, 0x7c, 0x8b,
	0x15, 0xc0, 0xbf, 0x88, 0xd7, 0x10, 0x6e, 0xa9, 0xf5, 0x39, 0x06, 0x9a, 0x7a, 0x3f, 0xc5, 0xae,
	0x21, 0x9f, 0xe2, 0xc8, 0xcb, 0x60, 0x15, 0x24, 0x4f, 0x30, 0xea, 0x7f, 0x62, 0x0e, 0x4c, 0xe9,
	0x42, 0x37, 0xca, 0x8e, 0x46, 0xdc, 0x33, 0x53, 0x71, 0x2d, 0xca, 0xb6, 0xab, 0x68, 0x70, 0xfe,
	0x0a, 0xec, 0xc3, 0x5e, 0x00, 0x9f, 0x81, 0xb9, 0x65, 0xe0, 0xcc, 0xf7, 0x3c, 0xba, 0xc1, 0xe2,
	0xe6, 0x44, 0xf5, 0x3b, 0x3c, 0x5b, 0x33, 0xab, 0x3f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0x72, 0xe5, 0xf9, 0xcb, 0x01, 0x00, 0x00,
}