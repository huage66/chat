package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/huage66/chat/config"
	"github.com/huage66/chat/interval/orm"
)

func NewRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisConf.Host, config.RedisConf.Port),
		//Username: config.RedisConf.Username,
		Password: config.RedisConf.Password,
		DB:       config.RedisConf.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		fmt.Sprintf("NewRedis: ping redis fail, err = %v\n", err)
		panic(err)
	}

	orm.RedisClient = client
}
