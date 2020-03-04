package main

type Room struct {
	RoomId int32

	Players []Player
	Pokers  [][]int
}

func LoadRoom() *Room {
	var room = &Room{}
	return room
}

func CreateRoom() *Room {
	var room = &Room{}
	return room
}

func (r *Room) JoinRoom(team Team) {

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
