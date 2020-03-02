package db

import (
	"fmt"
	"rpg-demo/config"

	"github.com/go-redis/redis/v7"
)

var Redis *redis.Client

func InitRedis() {
	var conf = config.DefaultConfig.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		// Password: conf., // no password set
		DB: 0, // use default DB
	})
}
