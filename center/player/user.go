package player

type User struct {
	UserId   int64
	NickName string
	Avatar   string
	IsFamale bool
	Coin     int64

	RoomId int64 //房间号
}

func (u *User) GetPlayerId() int64 {
	return u.UserId
}
func (u *User) GetRoomId() int64 {
	return u.RoomId
}
func (u *User) GetCoin() int64 {
	return u.Coin
}
func (u *User) AddCoin(v int64) { //+/-金币 {
	u.Coin += v
}
func (u *User) SetRoom(roomid int64) { //加入房间{
	u.RoomId = roomid
}
