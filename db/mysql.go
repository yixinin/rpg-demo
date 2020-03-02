package db

import (
	"fmt"
	"rpg-demo/center/models"
	"rpg-demo/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Mysql *xorm.Engine

func InitMysql() {
	var conf = config.DefaultConfig.Mysql
	var err error
	Mysql, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.User, conf.Password, conf.Host, conf.Port, conf.DB))
	if err != nil {
		panic(err)
	}
	SyncTable()
}

func SyncTable() {
	Mysql.Sync2(models.Tables...)
}
