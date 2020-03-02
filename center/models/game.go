package models

import "time"

func init() {
	Tables = append(Tables, new(Game), new(GameType), new(GameCate))
}

type Game struct {
	Id             int32 `xorm:"pk autoincr"`
	Name           string
	Active         bool //是否开启
	TopupInGame    bool //游戏中上分
	WithdrawInGame bool //游戏中下分

	Icon     string //游戏图标
	Banner   string //游戏封面
	AdBanner string //广告图

	PlatformId int32
	GameCateId int32

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

type GameType struct {
	Id       int32 `xorm:"pk autoincr"`
	GameId   int32
	GameType int32
	FullName string
	Active   bool //是否开启
	Config   string
}

type GameCate struct {
	Id     int32 `xorm:"pk autoincr"`
	Name   string
	Active bool //是否开启
}
