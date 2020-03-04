package cache

import (
	"fmt"
	"rpg-demo/db"
	"rpg-demo/ip"
	"time"
)

func SetToken(uid int64, token string) error {
	var userCacheKey = fmt.Sprintf("user:token:%s", token)
	var userCache = fmt.Sprintf("%d", uid)
	return db.Redis.Set(userCacheKey, userCache, 30*time.Minute).Err()
}

func GetUid(token string) (int64, error) {
	var userCacheKey = fmt.Sprintf("user:token:%s", token)
	return db.Redis.Get(userCacheKey).Int64()
}

func SetUserService(uid int64) error {
	var userServiceKey = fmt.Sprintf("user:service:%d", uid)
	return db.Redis.HSet(userServiceKey, "center", ip.ServerIP).Err()
}

func GetUserService(uid int64, service string) (string, error) {
	var userServiceKey = fmt.Sprintf("user:service:%d", uid)
	return db.Redis.HGet(userServiceKey, service).Result()
}
