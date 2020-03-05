package ddz

type RoomConfig struct {
	seat     int // 3/4
	timeout  int //出牌读秒
	passtime int //无法管牌读秒

	boom      int   //炸弹倍数
	jokerBoom int   //王炸倍数
	base      int64 //底数
	top       int64 //封顶
	minCoin   int64 //最低进入
	maxCoin   int64 //最高进入
}

type RobotConfig struct {
}

var (
	default3RoomConfig *RoomConfig
	default4RoomConfig *RoomConfig
)

func init() {
	default3RoomConfig = &RoomConfig{
		seat:      3,
		timeout:   30,
		passtime:  5,
		boom:      2,
		jokerBoom: 4,
		base:      100,
		top:       10000,
		minCoin:   2000,
		maxCoin:   200000,
	}
	default4RoomConfig = &RoomConfig{
		seat:      4,
		timeout:   30,
		passtime:  5,
		boom:      2,
		jokerBoom: 4,
		base:      100,
		top:       10000,
		minCoin:   2000,
		maxCoin:   200000,
	}
}
