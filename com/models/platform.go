package models

func init() {
	Tables = append(Tables, new(Platform))
}

type Platform struct {
	Id     int32 `xorm:"pk autoincr"`
	Name   string
	Active bool //是否开启
}
