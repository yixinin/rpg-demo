package player

type Robot struct {
	RobotId  int64
	NickName string
	Avatar   string
	IsFamale bool
	Coin     int64
	RoomId   int64
}

func (u *Robot) GetPlayerId() int64 {
	return u.RobotId
}
func (u *Robot) GetRoomId() int64 {
	return u.RoomId
}
func (u *Robot) GetCoin() int64 {
	return u.Coin
}
func (u *Robot) AddCoin(v int64) { //+/-金币 {
	u.Coin += v
}
func (u *Robot) SetRoom(roomid int64) { //加入房间{
	u.RoomId = roomid
}
