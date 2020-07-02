// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/grpc.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingStruct struct {
	Req                  string   `protobuf:"bytes,1,opt,name=Req,proto3" json:"Req,omitempty"`
	Resp                 string   `protobuf:"bytes,2,opt,name=Resp,proto3" json:"Resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingStruct) Reset()         { *m = PingStruct{} }
func (m *PingStruct) String() string { return proto.CompactTextString(m) }
func (*PingStruct) ProtoMessage()    {}
func (*PingStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac2853ef6ba3257, []int{0}
}

func (m *PingStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingStruct.Unmarshal(m, b)
}
func (m *PingStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingStruct.Marshal(b, m, deterministic)
}
func (m *PingStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingStruct.Merge(m, src)
}
func (m *PingStruct) XXX_Size() int {
	return xxx_messageInfo_PingStruct.Size(m)
}
func (m *PingStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_PingStruct.DiscardUnknown(m)
}

var xxx_messageInfo_PingStruct proto.InternalMessageInfo

func (m *PingStruct) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

func (m *PingStruct) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

func init() {
	proto.RegisterType((*PingStruct)(nil), "protos.PingStruct")
}

func init() { proto.RegisterFile("protos/grpc.proto", fileDescriptor_8ac2853ef6ba3257) }

var fileDescriptor_8ac2853ef6ba3257 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0xb3, 0x85, 0xd8, 0x20, 0x42, 0x4a, 0x46,
	0x5c, 0x5c, 0x01, 0x99, 0x79, 0xe9, 0xc1, 0x25, 0x45, 0xa5, 0xc9, 0x25, 0x42, 0x02, 0x5c, 0xcc,
	0x41, 0xa9, 0x85, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x10, 0x17, 0x4b,
	0x50, 0x6a, 0x71, 0x81, 0x04, 0x13, 0x58, 0x08, 0xcc, 0x36, 0x72, 0xe1, 0xe2, 0x06, 0xeb, 0x49,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32, 0xe5, 0xe2, 0x04, 0x71, 0x9d, 0x33, 0x52, 0x93, 0xb3,
	0x85, 0x84, 0x20, 0xe6, 0x17, 0xeb, 0x21, 0x4c, 0x95, 0xc2, 0x22, 0xa6, 0xc4, 0x90, 0x04, 0x71,
	0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x62, 0x20, 0x17, 0x14, 0x9d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingServiceClient interface {
	PingCheck(ctx context.Context, in *PingStruct, opts ...grpc.CallOption) (*PingStruct, error)
}

type pingServiceClient struct {
	cc *grpc.ClientConn
}

func NewPingServiceClient(cc *grpc.ClientConn) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) PingCheck(ctx context.Context, in *PingStruct, opts ...grpc.CallOption) (*PingStruct, error) {
	out := new(PingStruct)
	err := c.cc.Invoke(ctx, "/protos.PingService/PingCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServiceServer is the server API for PingService service.
type PingServiceServer interface {
	PingCheck(context.Context, *PingStruct) (*PingStruct, error)
}

// UnimplementedPingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPingServiceServer struct {
}

func (*UnimplementedPingServiceServer) PingCheck(ctx context.Context, req *PingStruct) (*PingStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingCheck not implemented")
}

func RegisterPingServiceServer(s *grpc.Server, srv PingServiceServer) {
	s.RegisterService(&_PingService_serviceDesc, srv)
}

func _PingService_PingCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingStruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).PingCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PingService/PingCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).PingCheck(ctx, req.(*PingStruct))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingCheck",
			Handler:    _PingService_PingCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/grpc.proto",
}
