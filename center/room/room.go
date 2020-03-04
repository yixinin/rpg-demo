package room

import (
	"errors"
	"rpg-demo/enums"
)

var ddzGameTypes map[int32]string
var DDZGameId int32

func init() {
	ddzGameTypes = make(map[int32]string, 5)
}

type DDZRoom struct {
	RoomId   int64
	Name     string
	GameId   int32
	GameType int32 //初级-中级-高级
	Kind     int8  //房间类型 匹配房 好友房

	Round int64 //局数

	Status    int8 //房间状态 1=Idle 2=队列中 3=游戏中 4=解散
	MaxPlayer int
	MinPlayer int

	Players []int64
	Owner   int64 // 0表示系统创建
}

func CreateDDZRoom(p int64, gameType int32, name string) *DDZRoom {
	var r = &DDZRoom{
		GameType:  gameType,
		Kind:      enums.RoomKindMatch,
		Status:    enums.RoomStatusIdle,
		MaxPlayer: 3,
		MinPlayer: 3,
		Players:   make([]int64, 0, 3),
	}
	fullName, ok := ddzGameTypes[gameType]
	if !ok {
		return nil
	}

	if name == "" {
		r.Name = fullName
	} else {
		r.Name = name
	}
	if p != 0 {
		r.Owner = p
		r.Players = append(r.Players, p)
	}

	return r
}

func (r *DDZRoom) GetRoomId() int64 {
	return r.RoomId
}
func (r *DDZRoom) GetGameId() int32 {
	return r.GameId
}
func (r *DDZRoom) GetGameType() int32 {
	return r.GameType
}
func (r *DDZRoom) GetPlayer(i int) int64 {
	return r.Players[i]
}
func (r *DDZRoom) GetPlayers() []int64 {
	return r.Players
}

func (r *DDZRoom) JoinRoom(p int64) error {
	if r.GetIsFull() {
		return errors.New("is full")
	}
	r.Players = append(r.Players, p)

	return nil
}
func (r *DDZRoom) StartGame() error {
	if r.GetIsFull() {
		return errors.New("not ready")
	}
	r.Status = enums.RoomStatusGaming
	return nil
}
func (r *DDZRoom) ExitRoom(p int64) error {
	return nil
}
func (r *DDZRoom) GetStatus() int8 {
	return r.Status
}
func (r *DDZRoom) GetMaxPlayer() int {
	return r.MaxPlayer
}
func (r *DDZRoom) GetMinPlayer() int {
	return r.MinPlayer
}
func (r *DDZRoom) GetIsFull() bool {
	return len(r.Players) == r.MaxPlayer
}
func (r *DDZRoom) GetNeedFull() bool {
	return r.MaxPlayer == r.MinPlayer
}
func (r *DDZRoom) GetRound() int64 {
	return r.Round
}
func (r *DDZRoom) Discard() {

}
