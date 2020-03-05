package ddz

import (
	"rpg-demo/game/ddz/card"
	"sync"
)

type Room struct {
	sync.Mutex
	id     int32
	status int8

	players []*Player
	pokers  []*card.Poker
	discard []*DiscardPoker //已经打出去的牌
	banker  int             //庄家座位号

	points       int   //番数
	playerPoints []int //加倍

	robotConfig *RobotConfig
	config      *RoomConfig
}

type DiscardPoker struct {
	seat   int        //座位号
	poker  card.Poker //出牌
	points int8       //番数
}

func LoadRoom(roomid int32) *Room {
	var c = default3RoomConfig
	var room = &Room{
		id:      roomid,
		players: make([]*Player, c.seat),
		config:  c,
	}
	return room
}

func CreateRoom(seat int) *Room {
	var c = default3RoomConfig
	if seat == 4 {
		c = default4RoomConfig
	}
	var room = &Room{
		players: make([]*Player, seat),
		config:  c,
	}
	return room
}

func (r *Room) JoinRoom(p ...Player) {

}

func (r *Room) ExitRoom() {

}

func (r *Room) StartGame() {

}

func (r *Room) EndGame() {

}

//游戏

//洗牌
func (r *Room) shuffle() {

}

//发牌
func (r *Room) dealCards() {

}

//出牌
func (r *Room) play() {

}
