package models

import "time"

func init() {
	Tables = append(Tables, new(UserCoin))
}

type UserCoin struct {
	Id     int64
	UserId int64

	Coin int64 //金币数量 负数代表减少

	TradeType int32  //交易类型 1 上分 2 下分 3 下注 4 赢 5 返佣  6任务 7 奖励 ...
	TradeInfo string // json

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
