package config

import (
	"com/db"
	"com/ip"
	"com/lib/bus"
	"com/utils"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Etcd  []string `yaml:"etcd"`
	Grpc  string   `yaml:"grpc"` //监听地址 0.0.0.0:9999
	Mysql *db.MysqlConfig
	Redis *db.RedisConfig
	Nats  *bus.Config
}

// var DefaultConfig Config

func ParseConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)
	ip.ServerIP = fmt.Sprintf("%s:%s", utils.IPAddress(), conf.Grpc)

	return &conf, err
}
