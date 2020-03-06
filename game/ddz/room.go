package ddz

import (
	"game/ddz/poker"
	"sync"
)

type Room struct {
	sync.Mutex
	roomId int64
	status int8

	groups        [][]*Player
	players       []*Player
	playerNum     int //房间人数
	maxPlayer     int //房间最大人数
	groupNum      int //队伍数
	maxGroup      int //最大队伍数
	grouPlayerNum int //每队人数
	pokers        []poker.Poker
	discard       []*DiscardPoker //已经打出去的牌
	banker        int             //庄家座位号

	points       int   //番数
	playerPoints []int //玩家加倍

	robotConfig *RobotConfig
	config      *RoomConfig
}

type DiscardPoker struct {
	seat   int         //座位号
	poker  poker.Poker //出牌
	points int8        //番数
}

func (g *DDZGame) LoadRoom(roomid int64) *Room {
	var c = default3RoomConfig
	var room = &Room{
		roomId:       roomid,
		players:      make([]*Player, c.seat),
		playerPoints: make([]int, c.seat),
		config:       c,
	}
	for i := range room.playerPoints {
		room.playerPoints[i] = 1
	}
	return room
}

func (g *DDZGame) CreateRoom(seat int) *Room {
	var c = default3RoomConfig
	if seat == 4 {
		c = default4RoomConfig
	}
	var room = &Room{
		roomId:       g.game.RoomIdPool.GetOne(),
		players:      make([]*Player, seat),
		playerPoints: make([]int, c.seat),
		config:       c,
	}
	for i := range room.playerPoints {
		room.playerPoints[i] = 1
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
