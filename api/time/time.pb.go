// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time/time.proto

package time

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"

	common "github.com/talos-systems/talos/api/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The response message containing the ntp server
type TimeRequest struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRequest) Reset()         { *m = TimeRequest{} }
func (m *TimeRequest) String() string { return proto.CompactTextString(m) }
func (*TimeRequest) ProtoMessage()    {}
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ed1ef5b20ef4ce, []int{0}
}

func (m *TimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRequest.Unmarshal(m, b)
}

func (m *TimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRequest.Marshal(b, m, deterministic)
}

func (m *TimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRequest.Merge(m, src)
}

func (m *TimeRequest) XXX_Size() int {
	return xxx_messageInfo_TimeRequest.Size(m)
}

func (m *TimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRequest proto.InternalMessageInfo

func (m *TimeRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

// The response message containing the ntp server, time, and offset
type TimeReply struct {
	Response             []*TimeResponse `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TimeReply) Reset()         { *m = TimeReply{} }
func (m *TimeReply) String() string { return proto.CompactTextString(m) }
func (*TimeReply) ProtoMessage()    {}
func (*TimeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ed1ef5b20ef4ce, []int{1}
}

func (m *TimeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeReply.Unmarshal(m, b)
}

func (m *TimeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeReply.Marshal(b, m, deterministic)
}

func (m *TimeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeReply.Merge(m, src)
}

func (m *TimeReply) XXX_Size() int {
	return xxx_messageInfo_TimeReply.Size(m)
}

func (m *TimeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeReply.DiscardUnknown(m)
}

var xxx_messageInfo_TimeReply proto.InternalMessageInfo

func (m *TimeReply) GetResponse() []*TimeResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type TimeResponse struct {
	Metadata             *common.ResponseMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Server               string                   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Localtime            *timestamp.Timestamp     `protobuf:"bytes,3,opt,name=localtime,proto3" json:"localtime,omitempty"`
	Remotetime           *timestamp.Timestamp     `protobuf:"bytes,4,opt,name=remotetime,proto3" json:"remotetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TimeResponse) Reset()         { *m = TimeResponse{} }
func (m *TimeResponse) String() string { return proto.CompactTextString(m) }
func (*TimeResponse) ProtoMessage()    {}
func (*TimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ed1ef5b20ef4ce, []int{2}
}

func (m *TimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeResponse.Unmarshal(m, b)
}

func (m *TimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeResponse.Marshal(b, m, deterministic)
}

func (m *TimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeResponse.Merge(m, src)
}

func (m *TimeResponse) XXX_Size() int {
	return xxx_messageInfo_TimeResponse.Size(m)
}

func (m *TimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimeResponse proto.InternalMessageInfo

func (m *TimeResponse) GetMetadata() *common.ResponseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TimeResponse) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *TimeResponse) GetLocaltime() *timestamp.Timestamp {
	if m != nil {
		return m.Localtime
	}
	return nil
}

func (m *TimeResponse) GetRemotetime() *timestamp.Timestamp {
	if m != nil {
		return m.Remotetime
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeRequest)(nil), "time.TimeRequest")
	proto.RegisterType((*TimeReply)(nil), "time.TimeReply")
	proto.RegisterType((*TimeResponse)(nil), "time.TimeResponse")
}

func init() { proto.RegisterFile("time/time.proto", fileDescriptor_e7ed1ef5b20ef4ce) }

var fileDescriptor_e7ed1ef5b20ef4ce = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xe9, 0x7f, 0x63, 0xff, 0xed, 0xdd, 0x60, 0x18, 0x61, 0x94, 0x7a, 0x70, 0x0c, 0xc4,
	0x5d, 0x4c, 0x60, 0x7a, 0x10, 0x3d, 0x39, 0xf1, 0x28, 0x48, 0xf1, 0xe4, 0x2d, 0xab, 0xaf, 0x5b,
	0xb0, 0x59, 0x62, 0x93, 0x0a, 0xfd, 0x9a, 0x7e, 0x22, 0x49, 0xd2, 0xcd, 0xe0, 0x0e, 0x5e, 0x9a,
	0xe6, 0x7d, 0x7e, 0x4f, 0xfb, 0xf0, 0x24, 0x30, 0xb6, 0x42, 0x22, 0x73, 0x0f, 0xaa, 0x2b, 0x65,
	0x15, 0xe9, 0xba, 0xf7, 0xec, 0x64, 0xad, 0xd4, 0xba, 0x44, 0xe6, 0x67, 0xab, 0xfa, 0x8d, 0xa1,
	0xd4, 0xb6, 0x09, 0x48, 0x76, 0xfa, 0x5b, 0x74, 0x16, 0x63, 0xb9, 0xd4, 0x2d, 0x70, 0x5c, 0x28,
	0x29, 0xd5, 0x96, 0x85, 0x25, 0x0c, 0x67, 0x67, 0x30, 0x7c, 0x16, 0x12, 0x73, 0xfc, 0xa8, 0xd1,
	0x58, 0x32, 0x81, 0x9e, 0xc1, 0xea, 0x13, 0xab, 0x34, 0x99, 0x26, 0xf3, 0x41, 0xde, 0xee, 0x66,
	0xb7, 0x30, 0x08, 0x98, 0x2e, 0x1b, 0x42, 0xa1, 0x5f, 0xa1, 0xd1, 0x6a, 0x6b, 0x30, 0x4d, 0xa6,
	0x9d, 0xf9, 0x70, 0x41, 0xa8, 0xcf, 0x1a, 0x90, 0xa0, 0xe4, 0x7b, 0x66, 0xf6, 0x95, 0xc0, 0x28,
	0x96, 0xc8, 0x15, 0xf4, 0x25, 0x5a, 0xfe, 0xca, 0x2d, 0xf7, 0xff, 0x19, 0x2e, 0x52, 0xda, 0xa6,
	0xda, 0x31, 0x8f, 0xad, 0x9e, 0xef, 0xc9, 0x28, 0xdb, 0xbf, 0x38, 0x1b, 0xb9, 0x86, 0x41, 0xa9,
	0x0a, 0x5e, 0xba, 0x08, 0x69, 0xc7, 0x7f, 0x2e, 0xa3, 0xa1, 0x0c, 0xba, 0x2b, 0xc3, 0x47, 0xf3,
	0x65, 0xe4, 0x3f, 0x30, 0xb9, 0x01, 0xa8, 0x50, 0x2a, 0x8b, 0xde, 0xda, 0xfd, 0xd3, 0x1a, 0xd1,
	0x8b, 0x0d, 0x74, 0x9d, 0x40, 0x58, 0xbb, 0x4e, 0x0e, 0x7c, 0x0f, 0xee, 0x70, 0xb2, 0x71, 0x5c,
	0x8d, 0x6b, 0x8f, 0x85, 0x2a, 0xef, 0x37, 0x58, 0xbc, 0x93, 0xa3, 0x58, 0xf5, 0x47, 0x70, 0x60,
	0x58, 0x2e, 0x61, 0x54, 0x28, 0x19, 0xa6, 0x5c, 0x8b, 0xe5, 0x7f, 0x27, 0xdd, 0x69, 0xf1, 0x94,
	0xbc, 0x9c, 0xaf, 0x85, 0xdd, 0xd4, 0x2b, 0x57, 0x1e, 0xb3, 0xbc, 0x54, 0xe6, 0xc2, 0x34, 0xc6,
	0xa2, 0x34, 0x61, 0xc7, 0xb8, 0x16, 0xfe, 0x1a, 0xac, 0x7a, 0x3e, 0xd5, 0xe5, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0x43, 0x34, 0x6f, 0x59, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimeClient is the client API for Time service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeClient interface {
	Time(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TimeReply, error)
	TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeReply, error)
}

type timeClient struct {
	cc *grpc.ClientConn
}

func NewTimeClient(cc *grpc.ClientConn) TimeClient {
	return &timeClient{cc}
}

func (c *timeClient) Time(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TimeReply, error) {
	out := new(TimeReply)
	err := c.cc.Invoke(ctx, "/time.Time/Time", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeClient) TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeReply, error) {
	out := new(TimeReply)
	err := c.cc.Invoke(ctx, "/time.Time/TimeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServer is the server API for Time service.
type TimeServer interface {
	Time(context.Context, *empty.Empty) (*TimeReply, error)
	TimeCheck(context.Context, *TimeRequest) (*TimeReply, error)
}

func RegisterTimeServer(s *grpc.Server, srv TimeServer) {
	s.RegisterService(&_Time_serviceDesc, srv)
}

func _Time_Time_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServer).Time(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/time.Time/Time",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServer).Time(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Time_TimeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServer).TimeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/time.Time/TimeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServer).TimeCheck(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Time_serviceDesc = grpc.ServiceDesc{
	ServiceName: "time.Time",
	HandlerType: (*TimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Time",
			Handler:    _Time_Time_Handler,
		},
		{
			MethodName: "TimeCheck",
			Handler:    _Time_TimeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time/time.proto",
}
