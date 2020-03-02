// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client-account.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LoginReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	UserName             string     `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ValidCode            string     `protobuf:"bytes,4,opt,name=validCode,proto3" json:"validCode,omitempty"`
	ValidId              string     `protobuf:"bytes,5,opt,name=validId,proto3" json:"validId,omitempty"`
	PhoneNumber          string     `protobuf:"bytes,6,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	DeviceCode           string     `protobuf:"bytes,7,opt,name=DeviceCode,proto3" json:"DeviceCode,omitempty"`
	InviteCode           string     `protobuf:"bytes,8,opt,name=InviteCode,proto3" json:"InviteCode,omitempty"`
	LoginType            int32      `protobuf:"varint,9,opt,name=loginType,proto3" json:"loginType,omitempty"`
	Register             bool       `protobuf:"varint,10,opt,name=register,proto3" json:"register,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eb8cd11a0a5d184, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LoginReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginReq) GetValidCode() string {
	if m != nil {
		return m.ValidCode
	}
	return ""
}

func (m *LoginReq) GetValidId() string {
	if m != nil {
		return m.ValidId
	}
	return ""
}

func (m *LoginReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *LoginReq) GetDeviceCode() string {
	if m != nil {
		return m.DeviceCode
	}
	return ""
}

func (m *LoginReq) GetInviteCode() string {
	if m != nil {
		return m.InviteCode
	}
	return ""
}

func (m *LoginReq) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *LoginReq) GetRegister() bool {
	if m != nil {
		return m.Register
	}
	return false
}

type LoginAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token                string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	DeviceCode           string     `protobuf:"bytes,3,opt,name=DeviceCode,proto3" json:"DeviceCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginAck) Reset()         { *m = LoginAck{} }
func (m *LoginAck) String() string { return proto.CompactTextString(m) }
func (*LoginAck) ProtoMessage()    {}
func (*LoginAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eb8cd11a0a5d184, []int{1}
}

func (m *LoginAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAck.Unmarshal(m, b)
}
func (m *LoginAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAck.Marshal(b, m, deterministic)
}
func (m *LoginAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAck.Merge(m, src)
}
func (m *LoginAck) XXX_Size() int {
	return xxx_messageInfo_LoginAck.Size(m)
}
func (m *LoginAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAck.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAck proto.InternalMessageInfo

func (m *LoginAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LoginAck) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginAck) GetDeviceCode() string {
	if m != nil {
		return m.DeviceCode
	}
	return ""
}

type LogoutReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LogoutReq) Reset()         { *m = LogoutReq{} }
func (m *LogoutReq) String() string { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()    {}
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eb8cd11a0a5d184, []int{2}
}

func (m *LogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutReq.Unmarshal(m, b)
}
func (m *LogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutReq.Marshal(b, m, deterministic)
}
func (m *LogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutReq.Merge(m, src)
}
func (m *LogoutReq) XXX_Size() int {
	return xxx_messageInfo_LogoutReq.Size(m)
}
func (m *LogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutReq proto.InternalMessageInfo

func (m *LogoutReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type LogoutAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LogoutAck) Reset()         { *m = LogoutAck{} }
func (m *LogoutAck) String() string { return proto.CompactTextString(m) }
func (*LogoutAck) ProtoMessage()    {}
func (*LogoutAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eb8cd11a0a5d184, []int{3}
}

func (m *LogoutAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutAck.Unmarshal(m, b)
}
func (m *LogoutAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutAck.Marshal(b, m, deterministic)
}
func (m *LogoutAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutAck.Merge(m, src)
}
func (m *LogoutAck) XXX_Size() int {
	return xxx_messageInfo_LogoutAck.Size(m)
}
func (m *LogoutAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutAck.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutAck proto.InternalMessageInfo

func (m *LogoutAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "protocol.LoginReq")
	proto.RegisterType((*LoginAck)(nil), "protocol.LoginAck")
	proto.RegisterType((*LogoutReq)(nil), "protocol.LogoutReq")
	proto.RegisterType((*LogoutAck)(nil), "protocol.LogoutAck")
}

func init() { proto.RegisterFile("client-account.proto", fileDescriptor_8eb8cd11a0a5d184) }

var fileDescriptor_8eb8cd11a0a5d184 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x6a, 0xf3, 0x30,
	0x14, 0x85, 0x71, 0xf2, 0x27, 0xb1, 0x6f, 0xfe, 0x49, 0xcd, 0x20, 0x42, 0x29, 0x26, 0x93, 0xa1,
	0x34, 0x43, 0xbb, 0x74, 0x0d, 0xed, 0xd0, 0x40, 0xc8, 0x20, 0xfa, 0x02, 0x8e, 0x7c, 0x49, 0x84,
	0x6d, 0xc9, 0x91, 0x65, 0x97, 0x3e, 0x5f, 0x5f, 0xac, 0x58, 0xb2, 0xe3, 0xe0, 0xa1, 0x90, 0xc9,
	0x9c, 0xf3, 0xdd, 0x6b, 0xee, 0x39, 0x82, 0x05, 0xcf, 0x04, 0x4a, 0xf3, 0x14, 0x73, 0xae, 0x2a,
	0x69, 0xd6, 0x85, 0x56, 0x46, 0x11, 0xdf, 0x7e, 0xb8, 0xca, 0x96, 0xff, 0x4f, 0x18, 0x27, 0xa8,
	0x9d, 0xbf, 0xfa, 0x19, 0x81, 0xbf, 0x53, 0x47, 0x21, 0x19, 0x9e, 0xc9, 0x23, 0x4c, 0x1d, 0xa4,
	0x5e, 0xe8, 0x45, 0xf3, 0xe7, 0xbb, 0x75, 0xb7, 0xb5, 0x66, 0x78, 0xfe, 0xb0, 0x88, 0xb5, 0x23,
	0x64, 0x09, 0x7e, 0x55, 0xa2, 0xde, 0xc7, 0x39, 0xd2, 0x51, 0xe8, 0x45, 0x01, 0xbb, 0xe8, 0x86,
	0x15, 0x71, 0x59, 0x7e, 0x29, 0x9d, 0xd0, 0xb1, 0x63, 0x9d, 0x26, 0xf7, 0x10, 0xd4, 0x71, 0x26,
	0x92, 0x37, 0x95, 0x20, 0xfd, 0x67, 0x61, 0x6f, 0x10, 0x0a, 0x33, 0x2b, 0xb6, 0x09, 0x9d, 0x58,
	0xd6, 0x49, 0x12, 0xc2, 0xbc, 0x38, 0x29, 0x89, 0xfb, 0x2a, 0x3f, 0xa0, 0xa6, 0x53, 0x4b, 0xaf,
	0x2d, 0xf2, 0x00, 0xf0, 0x8e, 0xb5, 0xe0, 0x68, 0x7f, 0x3d, 0xb3, 0x03, 0x57, 0x4e, 0xc3, 0xb7,
	0xb2, 0x16, 0xc6, 0x71, 0xdf, 0xf1, 0xde, 0x69, 0x2e, 0xcb, 0x9a, 0x2a, 0x3e, 0xbf, 0x0b, 0xa4,
	0x41, 0xe8, 0x45, 0x13, 0xd6, 0x1b, 0x4d, 0x26, 0x8d, 0x47, 0x51, 0x1a, 0xd4, 0x14, 0x42, 0x2f,
	0xf2, 0xd9, 0x45, 0xaf, 0xf2, 0xb6, 0xc4, 0x0d, 0x4f, 0xff, 0x2a, 0x71, 0xc3, 0xd3, 0x41, 0x89,
	0x0b, 0x98, 0x18, 0x95, 0xa2, 0x6c, 0x1b, 0x74, 0x62, 0x10, 0x64, 0x3c, 0x0c, 0xb2, 0x7a, 0x85,
	0x60, 0xa7, 0x8e, 0xaa, 0x32, 0xb7, 0x3e, 0x5a, 0xbf, 0x79, 0xeb, 0xa5, 0x87, 0xa9, 0x65, 0x2f,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xf0, 0x05, 0xe3, 0x5f, 0x02, 0x00, 0x00,
}