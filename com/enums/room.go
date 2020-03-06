package enums

var (
	RoomKindMatch int8 = 1 //匹配房
	RoomKindOwn   int8 = 2 //好友放
	RoomKindCard  int8 = 3 //房卡房
)

var (
	RoomStatusIdle     int8 = 1 //空闲状态
	RoomStatusMatching int8 = 2 //匹配状态 Moba
	RoomStatusGaming   int8 = 3 //游戏中
	RoomStatusSettle   int8 = 4 //结算
	RoomStatusDiscard  int8 = 5 //解散
)
