package iface

type Room interface {
	GetRoomId() int64
	GetGameId() int32
	GetGameType() int32

	GetPlayer(i int) int64
	GetPlayers() []int64
	// GetRealPlayers() []int64
	// GetRobots() []int64

	JoinRoom(p int64) error
	ExitRoom(p int64) error
	StartGame() error

	GetStatus() int8
	GetIsFull() bool
	GetNeedFull() bool //是否需要满人才能开始游戏
	GetMaxPlayer() int
	GetMinPlayer() int
	GetRound() int64

	Discard() //解散
}
