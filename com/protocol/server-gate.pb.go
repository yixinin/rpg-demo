// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server-gate.proto

package protocol

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

type LinkGame4GateReq struct {
	Uid                  []int64  `protobuf:"varint,1,rep,packed,name=uid,proto3" json:"uid,omitempty"`
	Server               string   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkGame4GateReq) Reset()         { *m = LinkGame4GateReq{} }
func (m *LinkGame4GateReq) String() string { return proto.CompactTextString(m) }
func (*LinkGame4GateReq) ProtoMessage()    {}
func (*LinkGame4GateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36caabf931980ba, []int{0}
}

func (m *LinkGame4GateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkGame4GateReq.Unmarshal(m, b)
}
func (m *LinkGame4GateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkGame4GateReq.Marshal(b, m, deterministic)
}
func (m *LinkGame4GateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkGame4GateReq.Merge(m, src)
}
func (m *LinkGame4GateReq) XXX_Size() int {
	return xxx_messageInfo_LinkGame4GateReq.Size(m)
}
func (m *LinkGame4GateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkGame4GateReq.DiscardUnknown(m)
}

var xxx_messageInfo_LinkGame4GateReq proto.InternalMessageInfo

func (m *LinkGame4GateReq) GetUid() []int64 {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *LinkGame4GateReq) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type NotifyMessageReq struct {
	Uids                 []int64  `protobuf:"varint,1,rep,packed,name=uids,proto3" json:"uids,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyMessageReq) Reset()         { *m = NotifyMessageReq{} }
func (m *NotifyMessageReq) String() string { return proto.CompactTextString(m) }
func (*NotifyMessageReq) ProtoMessage()    {}
func (*NotifyMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36caabf931980ba, []int{1}
}

func (m *NotifyMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyMessageReq.Unmarshal(m, b)
}
func (m *NotifyMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyMessageReq.Marshal(b, m, deterministic)
}
func (m *NotifyMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyMessageReq.Merge(m, src)
}
func (m *NotifyMessageReq) XXX_Size() int {
	return xxx_messageInfo_NotifyMessageReq.Size(m)
}
func (m *NotifyMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyMessageReq proto.InternalMessageInfo

func (m *NotifyMessageReq) GetUids() []int64 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func (m *NotifyMessageReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*LinkGame4GateReq)(nil), "protocol.LinkGame4GateReq")
	proto.RegisterType((*NotifyMessageReq)(nil), "protocol.NotifyMessageReq")
}

func init() { proto.RegisterFile("server-gate.proto", fileDescriptor_f36caabf931980ba) }

var fileDescriptor_f36caabf931980ba = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x4d, 0x4f, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00,
	0x53, 0xc9, 0xf9, 0x39, 0x52, 0xdc, 0xa5, 0x79, 0x99, 0x25, 0x95, 0x10, 0x61, 0x29, 0x9e, 0x8c,
	0xd4, 0xc4, 0x94, 0xd4, 0x22, 0x08, 0x4f, 0xc9, 0x86, 0x4b, 0xc0, 0x27, 0x33, 0x2f, 0xdb, 0x3d,
	0x31, 0x37, 0xd5, 0xc4, 0x3d, 0xb1, 0x24, 0x35, 0x28, 0xb5, 0x50, 0x48, 0x80, 0x8b, 0xb9, 0x34,
	0x33, 0x45, 0x82, 0x51, 0x81, 0x59, 0x83, 0x39, 0x08, 0xc4, 0x14, 0x12, 0xe3, 0x62, 0x83, 0x98,
	0x2f, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x29, 0x59, 0x71, 0x09, 0xf8, 0xe5, 0x97,
	0x64, 0xa6, 0x55, 0xfa, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0x83, 0x75, 0x0b, 0x71, 0xb1, 0x94, 0x66,
	0xa6, 0x14, 0x43, 0xb5, 0x83, 0xd9, 0x20, 0xb1, 0x94, 0xc4, 0x92, 0x44, 0xb0, 0x6e, 0x9e, 0x20,
	0x30, 0xdb, 0xe8, 0x0c, 0x23, 0x17, 0x37, 0xc8, 0xc6, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54,
	0x21, 0x33, 0x2e, 0x0e, 0xef, 0xcc, 0xe4, 0xec, 0xd0, 0xe2, 0xd4, 0x22, 0x21, 0x51, 0x3d, 0x98,
	0xdb, 0xf5, 0x60, 0x62, 0x41, 0xa9, 0x85, 0x52, 0x82, 0x08, 0x61, 0xe7, 0xc4, 0x9c, 0x1c, 0xc7,
	0xe4, 0x6c, 0x25, 0x06, 0x21, 0x67, 0x2e, 0x3e, 0x98, 0x0f, 0x82, 0xc1, 0xae, 0x12, 0x92, 0x42,
	0x28, 0x43, 0xf7, 0x1b, 0x56, 0x23, 0x34, 0x18, 0x85, 0x1c, 0xb8, 0x78, 0x51, 0x3c, 0x82, 0x6c,
	0x06, 0xba, 0x0f, 0xb1, 0x9a, 0x91, 0xc4, 0x06, 0x16, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x82, 0x38, 0xe4, 0x89, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GateServiceClient is the client API for GateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GateServiceClient interface {
	KickUser(ctx context.Context, in *KickUserReq, opts ...grpc.CallOption) (*CallAck, error)
	//连接Game服务器
	LinkGameServer(ctx context.Context, opts ...grpc.CallOption) (GateService_LinkGameServerClient, error)
	//主动推送消息
	NotifyMessage(ctx context.Context, in *NotifyMessageReq, opts ...grpc.CallOption) (*CallAck, error)
}

type gateServiceClient struct {
	cc *grpc.ClientConn
}

func NewGateServiceClient(cc *grpc.ClientConn) GateServiceClient {
	return &gateServiceClient{cc}
}

func (c *gateServiceClient) KickUser(ctx context.Context, in *KickUserReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GateService/KickUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateServiceClient) LinkGameServer(ctx context.Context, opts ...grpc.CallOption) (GateService_LinkGameServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GateService_serviceDesc.Streams[0], "/protocol.GateService/LinkGameServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &gateServiceLinkGameServerClient{stream}
	return x, nil
}

type GateService_LinkGameServerClient interface {
	Send(*LinkGame4GateReq) error
	CloseAndRecv() (*CallAck, error)
	grpc.ClientStream
}

type gateServiceLinkGameServerClient struct {
	grpc.ClientStream
}

func (x *gateServiceLinkGameServerClient) Send(m *LinkGame4GateReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gateServiceLinkGameServerClient) CloseAndRecv() (*CallAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CallAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gateServiceClient) NotifyMessage(ctx context.Context, in *NotifyMessageReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GateService/NotifyMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateServiceServer is the server API for GateService service.
type GateServiceServer interface {
	KickUser(context.Context, *KickUserReq) (*CallAck, error)
	//连接Game服务器
	LinkGameServer(GateService_LinkGameServerServer) error
	//主动推送消息
	NotifyMessage(context.Context, *NotifyMessageReq) (*CallAck, error)
}

// UnimplementedGateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGateServiceServer struct {
}

func (*UnimplementedGateServiceServer) KickUser(ctx context.Context, req *KickUserReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickUser not implemented")
}
func (*UnimplementedGateServiceServer) LinkGameServer(srv GateService_LinkGameServerServer) error {
	return status.Errorf(codes.Unimplemented, "method LinkGameServer not implemented")
}
func (*UnimplementedGateServiceServer) NotifyMessage(ctx context.Context, req *NotifyMessageReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyMessage not implemented")
}

func RegisterGateServiceServer(s *grpc.Server, srv GateServiceServer) {
	s.RegisterService(&_GateService_serviceDesc, srv)
}

func _GateService_KickUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServiceServer).KickUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GateService/KickUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServiceServer).KickUser(ctx, req.(*KickUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateService_LinkGameServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GateServiceServer).LinkGameServer(&gateServiceLinkGameServerServer{stream})
}

type GateService_LinkGameServerServer interface {
	SendAndClose(*CallAck) error
	Recv() (*LinkGame4GateReq, error)
	grpc.ServerStream
}

type gateServiceLinkGameServerServer struct {
	grpc.ServerStream
}

func (x *gateServiceLinkGameServerServer) SendAndClose(m *CallAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gateServiceLinkGameServerServer) Recv() (*LinkGame4GateReq, error) {
	m := new(LinkGame4GateReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GateService_NotifyMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServiceServer).NotifyMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GateService/NotifyMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServiceServer).NotifyMessage(ctx, req.(*NotifyMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.GateService",
	HandlerType: (*GateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KickUser",
			Handler:    _GateService_KickUser_Handler,
		},
		{
			MethodName: "NotifyMessage",
			Handler:    _GateService_NotifyMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LinkGameServer",
			Handler:       _GateService_LinkGameServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "server-gate.proto",
}
