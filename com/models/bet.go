package models

import "time"

func init() {
	Tables = append(Tables, new(BetLog))
}

//下注记录
type BetLog struct {
	Id             int64 `xorm:"pk autoincr"`
	UserId         int64
	GameId         int32
	RoundNo        string //局号
	GameTypeId     int32
	GameCateId     int32
	GamePlatformId int32

	Bet      int64 //投注
	ValidBet int64 //有效投注
	BetTime  time.Time

	Win       int64     //负数代表输
	CloseTime time.Time //开奖时间

	Settled     bool      //是否结算
	SettledTime time.Time //派奖时间

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
