package models

import "time"

func init() {
	Tables = append(Tables, new(GameRecord))
}

type GameRecord struct {
	Id             int64 `xorm:"pk autoincr"`
	GameId         int32
	GameTypeId     int32
	GameCateId     int32
	GamePlatformId int32

	RoundNo string
	Record  string //json格式

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
