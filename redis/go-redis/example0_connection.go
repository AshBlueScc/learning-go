package redis

import (
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init(){
	if Client == nil {
		Client = redis.NewClient(&redis.Options{
			Network:            "tcp",
			Addr:               "127.0.0.1:6379",
			Password:           "",
			DB:                 1,
		})
	}
}