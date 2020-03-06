package models

import "time"

func init() {
	Tables = append(Tables, new(Room))
}

type Room struct {
	Id   int64  `xorm:"pk autoincr"`
	Kind int8   //房间类型  1 百人房间（不解散） 2 匹配房间 3 组队房间 4 房卡房间
	Name string //房间名称

	GameId int32 //游戏

	GameRound  int64     //游戏局数
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
