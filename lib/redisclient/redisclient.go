package redisclient

import (
	conf "github.com/myproject-0722/mn-hosted/conf"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(
		&redis.Options{
			Addr: conf.RedisIP,
			DB:   9,
		},
	)

	_, err := Client.Ping().Result()
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
