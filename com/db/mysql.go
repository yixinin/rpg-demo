package db

import (
	"com/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

var Mysql *xorm.Engine

func InitMysql(conf *MysqlConfig) {
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
