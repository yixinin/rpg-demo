// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client-notify.proto

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

type MessageNotify struct {
	Header               *NotiHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body                 string      `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	MType                int32       `protobuf:"varint,3,opt,name=mType,proto3" json:"mType,omitempty"`
	Kind                 int32       `protobuf:"varint,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Extra                string      `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	UserId               int64       `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             int64       `protobuf:"varint,7,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MessageNotify) Reset()         { *m = MessageNotify{} }
func (m *MessageNotify) String() string { return proto.CompactTextString(m) }
func (*MessageNotify) ProtoMessage()    {}
func (*MessageNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_855697b8a11839a4, []int{0}
}

func (m *MessageNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageNotify.Unmarshal(m, b)
}
func (m *MessageNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageNotify.Marshal(b, m, deterministic)
}
func (m *MessageNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageNotify.Merge(m, src)
}
func (m *MessageNotify) XXX_Size() int {
	return xxx_messageInfo_MessageNotify.Size(m)
}
func (m *MessageNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageNotify.DiscardUnknown(m)
}

var xxx_messageInfo_MessageNotify proto.InternalMessageInfo

func (m *MessageNotify) GetHeader() *NotiHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MessageNotify) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MessageNotify) GetMType() int32 {
	if m != nil {
		return m.MType
	}
	return 0
}

func (m *MessageNotify) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *MessageNotify) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *MessageNotify) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MessageNotify) GetUserName() int64 {
	if m != nil {
		return m.UserName
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageNotify)(nil), "protocol.MessageNotify")
}

func init() { proto.RegisterFile("client-notify.proto", fileDescriptor_855697b8a11839a4) }

var fileDescriptor_855697b8a11839a4 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x31, 0x6d, 0x42, 0xb9, 0xd2, 0x81, 0xa3, 0x02, 0x2b, 0x53, 0xd4, 0x29, 0x03, 0x64,
	0x28, 0x23, 0x13, 0x62, 0x01, 0x09, 0x2a, 0x14, 0x78, 0x01, 0x37, 0x39, 0x68, 0xd4, 0x34, 0xae,
	0x6c, 0x17, 0x91, 0x07, 0xe4, 0xbd, 0x90, 0xcf, 0x81, 0xaa, 0x93, 0xbf, 0x3f, 0xbf, 0x93, 0xfc,
	0xc1, 0x45, 0xd9, 0xd4, 0xd4, 0xba, 0x9b, 0x56, 0xbb, 0xfa, 0xa3, 0xcb, 0xb7, 0x46, 0x3b, 0x8d,
	0x23, 0x7e, 0x4a, 0xdd, 0x24, 0x67, 0x2b, 0x52, 0x15, 0x99, 0x90, 0xcf, 0x7e, 0x04, 0x4c, 0x5e,
	0xc8, 0x5a, 0xf5, 0x49, 0x0b, 0xe6, 0xf1, 0x1a, 0xe2, 0x40, 0x48, 0x91, 0x8a, 0x6c, 0x3c, 0x9f,
	0xe6, 0x7f, 0xa7, 0xb9, 0x27, 0x1e, 0xb9, 0x2b, 0x7a, 0x06, 0x11, 0x86, 0x4b, 0x5d, 0x75, 0xf2,
	0x38, 0x15, 0xd9, 0x69, 0xc1, 0x1a, 0xa7, 0x10, 0x6d, 0xde, 0xbb, 0x2d, 0xc9, 0x41, 0x2a, 0xb2,
	0xa8, 0x08, 0xc6, 0x93, 0xeb, 0xba, 0xad, 0xe4, 0x90, 0x43, 0xd6, 0x9e, 0xa4, 0x6f, 0x67, 0x94,
	0x8c, 0xf8, 0x3c, 0x18, 0xbc, 0x84, 0x78, 0x67, 0xc9, 0x3c, 0x55, 0x32, 0x4e, 0x45, 0x36, 0x28,
	0x7a, 0x87, 0x09, 0x8c, 0xbc, 0x5a, 0xa8, 0x0d, 0xc9, 0x13, 0x6e, 0xfe, 0xfd, 0xfc, 0x19, 0x26,
	0xe1, 0xff, 0x6f, 0x64, 0xbe, 0xea, 0x92, 0xf0, 0x0e, 0xc6, 0xaf, 0x3b, 0xbb, 0xea, 0xb7, 0xe1,
	0xd5, 0x7e, 0xc5, 0xc1, 0xdc, 0xe4, 0x7c, 0x5f, 0x3c, 0xa8, 0xa6, 0xb9, 0x2f, 0xd7, 0xb3, 0xa3,
	0x65, 0xcc, 0xd9, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x4b, 0xc7, 0xe5, 0x4b, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotifyServiceClient is the client API for NotifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifyServiceClient interface {
	PushMessage(ctx context.Context, in *MessageNotify, opts ...grpc.CallOption) (*CallAck, error)
}

type notifyServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotifyServiceClient(cc *grpc.ClientConn) NotifyServiceClient {
	return &notifyServiceClient{cc}
}

func (c *notifyServiceClient) PushMessage(ctx context.Context, in *MessageNotify, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.NotifyService/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServiceServer is the server API for NotifyService service.
type NotifyServiceServer interface {
	PushMessage(context.Context, *MessageNotify) (*CallAck, error)
}

// UnimplementedNotifyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotifyServiceServer struct {
}

func (*UnimplementedNotifyServiceServer) PushMessage(ctx context.Context, req *MessageNotify) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessage not implemented")
}

func RegisterNotifyServiceServer(s *grpc.Server, srv NotifyServiceServer) {
	s.RegisterService(&_NotifyService_serviceDesc, srv)
}

func _NotifyService_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageNotify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServiceServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.NotifyService/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServiceServer).PushMessage(ctx, req.(*MessageNotify))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotifyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.NotifyService",
	HandlerType: (*NotifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMessage",
			Handler:    _NotifyService_PushMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client-notify.proto",
}
