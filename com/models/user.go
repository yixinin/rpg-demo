package models

import "time"

func init() {
	Tables = append(Tables, new(User), new(UserLog), new(Account))
}

type User struct {
	//账号信息
	Id            int64  `xorm:"pk autoincr"`
	UserName      string `xorm:"unique"`
	PhoneNumber   string `xorm:"unique"`
	Password      string
	DeviceCode    string `xorm:"unique"`
	WxOpenId      string `xorm:"unique"`
	QqOpenId      string `xorm:"unique"`
	RegTime       time.Time
	LastLoginTime time.Time //最后一次登录

	//个人信息
	NickName string
	Avatar   string
	IsFamale bool

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

type UserLog struct {
	Id         int32 `xorm:"pk autoincr"`
	Date       int32
	UserId     int64
	LoginTimes int32
	OnlineTime int32
	GameTime   int32
}

type Account struct {
	UserId int64 `xorm:"pk autoincr"`

	//金钱
	Coin    int64 `xorm:"default(0)"`
	SafeBox int64 `xorm:"default(0)"`

	//等级
	Level int32 `xorm:"default(1)"`
	Group int32 `xorm:"default(0)"`

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
