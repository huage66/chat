package cmd

import (
	"github.com/huage66/chat/api"
	"github.com/huage66/chat/config"
	"github.com/huage66/chat/database"
)

func Init() {
	// 初始化config
	config.Setup()
	// 连接redis
	database.NewRedis()

	api.Server()
}
