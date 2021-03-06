// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server-game.proto

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

type Player struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Coin                 int64    `protobuf:"varint,4,opt,name=coin,proto3" json:"coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cbf3ce83a4d353, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *Player) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Player) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Player) GetCoin() int64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

type UpdateUserInfoReq struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Avatar               string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfoReq) Reset()         { *m = UpdateUserInfoReq{} }
func (m *UpdateUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfoReq) ProtoMessage()    {}
func (*UpdateUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cbf3ce83a4d353, []int{1}
}

func (m *UpdateUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfoReq.Unmarshal(m, b)
}
func (m *UpdateUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfoReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfoReq.Merge(m, src)
}
func (m *UpdateUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfoReq.Size(m)
}
func (m *UpdateUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfoReq proto.InternalMessageInfo

func (m *UpdateUserInfoReq) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UpdateUserInfoReq) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type CloseGameReq struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseGameReq) Reset()         { *m = CloseGameReq{} }
func (m *CloseGameReq) String() string { return proto.CompactTextString(m) }
func (*CloseGameReq) ProtoMessage()    {}
func (*CloseGameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cbf3ce83a4d353, []int{2}
}

func (m *CloseGameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseGameReq.Unmarshal(m, b)
}
func (m *CloseGameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseGameReq.Marshal(b, m, deterministic)
}
func (m *CloseGameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseGameReq.Merge(m, src)
}
func (m *CloseGameReq) XXX_Size() int {
	return xxx_messageInfo_CloseGameReq.Size(m)
}
func (m *CloseGameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseGameReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseGameReq proto.InternalMessageInfo

func (m *CloseGameReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type AddCoinReq struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Coin                 int64    `protobuf:"varint,2,opt,name=coin,proto3" json:"coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddCoinReq) Reset()         { *m = AddCoinReq{} }
func (m *AddCoinReq) String() string { return proto.CompactTextString(m) }
func (*AddCoinReq) ProtoMessage()    {}
func (*AddCoinReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cbf3ce83a4d353, []int{3}
}

func (m *AddCoinReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCoinReq.Unmarshal(m, b)
}
func (m *AddCoinReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCoinReq.Marshal(b, m, deterministic)
}
func (m *AddCoinReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCoinReq.Merge(m, src)
}
func (m *AddCoinReq) XXX_Size() int {
	return xxx_messageInfo_AddCoinReq.Size(m)
}
func (m *AddCoinReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCoinReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddCoinReq proto.InternalMessageInfo

func (m *AddCoinReq) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *AddCoinReq) GetCoin() int64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

type AddPlayerReq struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPlayerReq) Reset()         { *m = AddPlayerReq{} }
func (m *AddPlayerReq) String() string { return proto.CompactTextString(m) }
func (*AddPlayerReq) ProtoMessage()    {}
func (*AddPlayerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cbf3ce83a4d353, []int{4}
}

func (m *AddPlayerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPlayerReq.Unmarshal(m, b)
}
func (m *AddPlayerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPlayerReq.Marshal(b, m, deterministic)
}
func (m *AddPlayerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPlayerReq.Merge(m, src)
}
func (m *AddPlayerReq) XXX_Size() int {
	return xxx_messageInfo_AddPlayerReq.Size(m)
}
func (m *AddPlayerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPlayerReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddPlayerReq proto.InternalMessageInfo

func (m *AddPlayerReq) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

//进入游戏 展示房间列表
type EnterGameReq struct {
	GameId               int32    `protobuf:"varint,1,opt,name=gameId,proto3" json:"gameId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterGameReq) Reset()         { *m = EnterGameReq{} }
func (m *EnterGameReq) String() string { return proto.CompactTextString(m) }
func (*EnterGameReq) ProtoMessage()    {}
func (*EnterGameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cbf3ce83a4d353, []int{5}
}

func (m *EnterGameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterGameReq.Unmarshal(m, b)
}
func (m *EnterGameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterGameReq.Marshal(b, m, deterministic)
}
func (m *EnterGameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterGameReq.Merge(m, src)
}
func (m *EnterGameReq) XXX_Size() int {
	return xxx_messageInfo_EnterGameReq.Size(m)
}
func (m *EnterGameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterGameReq.DiscardUnknown(m)
}

var xxx_messageInfo_EnterGameReq proto.InternalMessageInfo

func (m *EnterGameReq) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

type LeaveGameReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveGameReq) Reset()         { *m = LeaveGameReq{} }
func (m *LeaveGameReq) String() string { return proto.CompactTextString(m) }
func (*LeaveGameReq) ProtoMessage()    {}
func (*LeaveGameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_84cbf3ce83a4d353, []int{6}
}

func (m *LeaveGameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveGameReq.Unmarshal(m, b)
}
func (m *LeaveGameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveGameReq.Marshal(b, m, deterministic)
}
func (m *LeaveGameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveGameReq.Merge(m, src)
}
func (m *LeaveGameReq) XXX_Size() int {
	return xxx_messageInfo_LeaveGameReq.Size(m)
}
func (m *LeaveGameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveGameReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveGameReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Player)(nil), "protocol.Player")
	proto.RegisterType((*UpdateUserInfoReq)(nil), "protocol.UpdateUserInfoReq")
	proto.RegisterType((*CloseGameReq)(nil), "protocol.CloseGameReq")
	proto.RegisterType((*AddCoinReq)(nil), "protocol.AddCoinReq")
	proto.RegisterType((*AddPlayerReq)(nil), "protocol.AddPlayerReq")
	proto.RegisterType((*EnterGameReq)(nil), "protocol.EnterGameReq")
	proto.RegisterType((*LeaveGameReq)(nil), "protocol.LeaveGameReq")
}

func init() { proto.RegisterFile("server-game.proto", fileDescriptor_84cbf3ce83a4d353) }

var fileDescriptor_84cbf3ce83a4d353 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xb5, 0x80, 0x15, 0x86, 0x86, 0xc8, 0x46, 0x49, 0x53, 0x2f, 0xa4, 0x07, 0xc3, 0x45, 0x0e,
	0x68, 0x0c, 0x07, 0x2f, 0x48, 0x0c, 0x21, 0x1a, 0x63, 0x6a, 0xf8, 0x01, 0x6b, 0x3b, 0x6a, 0x43,
	0xdb, 0x85, 0xa5, 0x36, 0xe1, 0x17, 0xf9, 0x37, 0x4d, 0xb7, 0xdb, 0xed, 0x12, 0xe9, 0x69, 0x77,
	0x3e, 0xde, 0x9b, 0x79, 0xf3, 0xa0, 0xbf, 0x43, 0x9e, 0x21, 0xbf, 0xf9, 0xa2, 0x31, 0x8e, 0x37,
	0x9c, 0xa5, 0x8c, 0xb4, 0xc5, 0xe3, 0xb3, 0xc8, 0xe9, 0xfe, 0x24, 0x61, 0xba, 0x2f, 0xd2, 0x8e,
	0xf5, 0x8d, 0x34, 0x40, 0x5e, 0x44, 0x6e, 0x04, 0xe6, 0x5b, 0x44, 0xf7, 0xc8, 0x89, 0x03, 0xed,
	0x8d, 0xf8, 0x2d, 0x03, 0xdb, 0x18, 0x1a, 0xa3, 0xa6, 0xa7, 0xe2, 0xbc, 0x96, 0x84, 0xfe, 0xfa,
	0x95, 0xc6, 0x68, 0x37, 0x86, 0xc6, 0xa8, 0xe3, 0xa9, 0x98, 0x0c, 0xc0, 0xa4, 0x19, 0x4d, 0x29,
	0xb7, 0x9b, 0xa2, 0x22, 0x23, 0x42, 0xa0, 0xe5, 0xb3, 0x30, 0xb1, 0x5b, 0x82, 0x4b, 0xfc, 0xdd,
	0x05, 0xf4, 0x57, 0x9b, 0x80, 0xa6, 0xb8, 0xda, 0x21, 0x5f, 0x26, 0x9f, 0xcc, 0xc3, 0xed, 0x01,
	0xb9, 0x51, 0x4b, 0xde, 0xd0, 0xc9, 0xdd, 0x6b, 0xb0, 0xe6, 0x11, 0xdb, 0xe1, 0x82, 0xc6, 0x98,
	0x73, 0x0c, 0xc0, 0xe4, 0x8c, 0xc5, 0x6a, 0x75, 0x19, 0xb9, 0x0f, 0x00, 0xb3, 0x20, 0x98, 0xb3,
	0x30, 0x91, 0x93, 0x6a, 0x25, 0x96, 0xeb, 0x36, 0xb4, 0x75, 0xa7, 0x60, 0xcd, 0x82, 0xa0, 0xb8,
	0x4f, 0x8e, 0x1f, 0x81, 0x59, 0xf4, 0x0b, 0x74, 0x77, 0x72, 0x3e, 0x2e, 0x4f, 0x3c, 0x96, 0x4d,
	0xb2, 0x9e, 0xef, 0xf7, 0x94, 0xa4, 0xc8, 0xb5, 0xfd, 0x72, 0x67, 0xe4, 0xdc, 0x53, 0x4f, 0x46,
	0x6e, 0x0f, 0xac, 0x17, 0xa4, 0x59, 0xa9, 0x63, 0xf2, 0xdb, 0x84, 0x6e, 0xfe, 0x7f, 0x47, 0x9e,
	0x85, 0x3e, 0x92, 0x7b, 0x68, 0x3f, 0x87, 0xfe, 0x3a, 0x3f, 0x17, 0xb9, 0xac, 0xa6, 0x95, 0x39,
	0x0f, 0xb7, 0x4e, 0xbf, 0x4a, 0xcf, 0x69, 0x14, 0xcd, 0xfc, 0xb5, 0x7b, 0x42, 0x1e, 0xa1, 0x77,
	0x78, 0x68, 0x72, 0x55, 0xb5, 0xfd, 0xb3, 0xe0, 0x38, 0xc7, 0x14, 0x3a, 0x4a, 0x03, 0x19, 0x54,
	0x1d, 0xba, 0xb0, 0x5a, 0xa4, 0x52, 0xa5, 0x23, 0x75, 0xa9, 0xb5, 0x48, 0xe5, 0xab, 0x8e, 0xd4,
	0xcd, 0x3e, 0x8e, 0xbc, 0x83, 0x33, 0xe9, 0x34, 0xb9, 0xa8, 0xea, 0x95, 0xf9, 0xb5, 0xf3, 0x94,
	0xc3, 0xfa, 0x3c, 0xdd, 0xf6, 0xa3, 0xc8, 0x0f, 0x53, 0xe4, 0x6e, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x26, 0xb2, 0x2d, 0x79, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	KickUser(ctx context.Context, in *KickUserReq, opts ...grpc.CallOption) (*CallAck, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*CallAck, error)
	EnterGame(ctx context.Context, in *EnterGameReq, opts ...grpc.CallOption) (*CallAck, error)
	LeaveGame(ctx context.Context, in *LeaveGameReq, opts ...grpc.CallOption) (*CallAck, error)
	//强制中断游戏
	CloseGame(ctx context.Context, in *CloseGameReq, opts ...grpc.CallOption) (*CallAck, error)
	//上/下分
	AddCoin(ctx context.Context, in *AddCoinReq, opts ...grpc.CallOption) (*CallAck, error)
	//中途加入游戏/ 断线重连
	AddPlayer(ctx context.Context, in *AddPlayerReq, opts ...grpc.CallOption) (*CallAck, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) KickUser(ctx context.Context, in *KickUserReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GameService/KickUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GameService/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) EnterGame(ctx context.Context, in *EnterGameReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GameService/EnterGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) LeaveGame(ctx context.Context, in *LeaveGameReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GameService/LeaveGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) CloseGame(ctx context.Context, in *CloseGameReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GameService/CloseGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) AddCoin(ctx context.Context, in *AddCoinReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GameService/AddCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) AddPlayer(ctx context.Context, in *AddPlayerReq, opts ...grpc.CallOption) (*CallAck, error) {
	out := new(CallAck)
	err := c.cc.Invoke(ctx, "/protocol.GameService/AddPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	KickUser(context.Context, *KickUserReq) (*CallAck, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*CallAck, error)
	EnterGame(context.Context, *EnterGameReq) (*CallAck, error)
	LeaveGame(context.Context, *LeaveGameReq) (*CallAck, error)
	//强制中断游戏
	CloseGame(context.Context, *CloseGameReq) (*CallAck, error)
	//上/下分
	AddCoin(context.Context, *AddCoinReq) (*CallAck, error)
	//中途加入游戏/ 断线重连
	AddPlayer(context.Context, *AddPlayerReq) (*CallAck, error)
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) KickUser(ctx context.Context, req *KickUserReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickUser not implemented")
}
func (*UnimplementedGameServiceServer) UpdateUserInfo(ctx context.Context, req *UpdateUserInfoReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (*UnimplementedGameServiceServer) EnterGame(ctx context.Context, req *EnterGameReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnterGame not implemented")
}
func (*UnimplementedGameServiceServer) LeaveGame(ctx context.Context, req *LeaveGameReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGame not implemented")
}
func (*UnimplementedGameServiceServer) CloseGame(ctx context.Context, req *CloseGameReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseGame not implemented")
}
func (*UnimplementedGameServiceServer) AddCoin(ctx context.Context, req *AddCoinReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoin not implemented")
}
func (*UnimplementedGameServiceServer) AddPlayer(ctx context.Context, req *AddPlayerReq) (*CallAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlayer not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_KickUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).KickUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GameService/KickUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).KickUser(ctx, req.(*KickUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GameService/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_EnterGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).EnterGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GameService/EnterGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).EnterGame(ctx, req.(*EnterGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_LeaveGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).LeaveGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GameService/LeaveGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).LeaveGame(ctx, req.(*LeaveGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_CloseGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CloseGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GameService/CloseGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CloseGame(ctx, req.(*CloseGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_AddCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).AddCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GameService/AddCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).AddCoin(ctx, req.(*AddCoinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_AddPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).AddPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GameService/AddPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).AddPlayer(ctx, req.(*AddPlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KickUser",
			Handler:    _GameService_KickUser_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _GameService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "EnterGame",
			Handler:    _GameService_EnterGame_Handler,
		},
		{
			MethodName: "LeaveGame",
			Handler:    _GameService_LeaveGame_Handler,
		},
		{
			MethodName: "CloseGame",
			Handler:    _GameService_CloseGame_Handler,
		},
		{
			MethodName: "AddCoin",
			Handler:    _GameService_AddCoin_Handler,
		},
		{
			MethodName: "AddPlayer",
			Handler:    _GameService_AddPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server-game.proto",
}
