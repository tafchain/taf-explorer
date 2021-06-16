package db

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redisAddr"),
		Password: beego.AppConfig.String("redisPassword"),
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func RedisSession() *redis.Client {
	return rdb
}
