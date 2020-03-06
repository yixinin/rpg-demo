package db

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var Redis *redis.Client

func InitRedis(conf *RedisConfig) {
	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		// Password: conf., // no password set
		DB: 0, // use default DB
	})
}
