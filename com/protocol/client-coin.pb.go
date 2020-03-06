// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client-coin.proto

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

//上分
type UpCoinReq struct {
	Header               *ReqHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpCoinReq) Reset()         { *m = UpCoinReq{} }
func (m *UpCoinReq) String() string { return proto.CompactTextString(m) }
func (*UpCoinReq) ProtoMessage()    {}
func (*UpCoinReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_992fecc01331446d, []int{0}
}

func (m *UpCoinReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpCoinReq.Unmarshal(m, b)
}
func (m *UpCoinReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpCoinReq.Marshal(b, m, deterministic)
}
func (m *UpCoinReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpCoinReq.Merge(m, src)
}
func (m *UpCoinReq) XXX_Size() int {
	return xxx_messageInfo_UpCoinReq.Size(m)
}
func (m *UpCoinReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpCoinReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpCoinReq proto.InternalMessageInfo

func (m *UpCoinReq) GetHeader() *ReqHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type DownCoinAck struct {
	Header               *AckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DownCoinAck) Reset()         { *m = DownCoinAck{} }
func (m *DownCoinAck) String() string { return proto.CompactTextString(m) }
func (*DownCoinAck) ProtoMessage()    {}
func (*DownCoinAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_992fecc01331446d, []int{1}
}

func (m *DownCoinAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownCoinAck.Unmarshal(m, b)
}
func (m *DownCoinAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownCoinAck.Marshal(b, m, deterministic)
}
func (m *DownCoinAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownCoinAck.Merge(m, src)
}
func (m *DownCoinAck) XXX_Size() int {
	return xxx_messageInfo_DownCoinAck.Size(m)
}
func (m *DownCoinAck) XXX_DiscardUnknown() {
	xxx_messageInfo_DownCoinAck.DiscardUnknown(m)
}

var xxx_messageInfo_DownCoinAck proto.InternalMessageInfo

func (m *DownCoinAck) GetHeader() *AckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*UpCoinReq)(nil), "protocol.UpCoinReq")
	proto.RegisterType((*DownCoinAck)(nil), "protocol.DownCoinAck")
}

func init() { proto.RegisterFile("client-coin.proto", fileDescriptor_992fecc01331446d) }

var fileDescriptor_992fecc01331446d = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x4d, 0xce, 0xcf, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00,
	0x53, 0xc9, 0xf9, 0x39, 0x52, 0x3c, 0x19, 0xa9, 0x89, 0x29, 0xa9, 0x45, 0x10, 0x71, 0x25, 0x0b,
	0x2e, 0xce, 0xd0, 0x02, 0xe7, 0xfc, 0xcc, 0xbc, 0xa0, 0xd4, 0x42, 0x21, 0x6d, 0x2e, 0x36, 0x88,
	0xa4, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xb0, 0x1e, 0x4c, 0x97, 0x5e, 0x50, 0x6a, 0xa1,
	0x07, 0x58, 0x2a, 0x08, 0xaa, 0x44, 0xc9, 0x8a, 0x8b, 0xdb, 0x25, 0xbf, 0x3c, 0x0f, 0xa4, 0xd7,
	0x31, 0x39, 0x1b, 0x9f, 0x5e, 0xc7, 0xe4, 0x6c, 0x54, 0xbd, 0x49, 0x6c, 0x60, 0x39, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x69, 0x40, 0xdf, 0xa9, 0x00, 0x00, 0x00,
}