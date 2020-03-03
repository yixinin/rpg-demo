package config

import (
	"io/ioutil"
	"rpg-demo/lib/bus"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	EtcdAddress []string `yaml:"etcd"`
	Grpc        string   `yaml:"grpc"` //监听地址 0.0.0.0:9999
	Mysql       MysqlConfig
	Redis       RedisConfig
	Nats        *bus.Config
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var DefaultConfig Config

func Init(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &DefaultConfig)
	if err != nil {
		return err
	}
	return nil
}
