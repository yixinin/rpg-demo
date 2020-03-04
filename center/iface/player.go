package iface

type Player interface {
	GetPlayerId() int64
	GetRoomId() int64
	GetCoin() int64
	AddCoin(v int64)      //+/-金币
	SetRoom(roomid int64) //加入房间
}
