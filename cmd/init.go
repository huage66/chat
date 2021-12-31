package cmd

import (
	"github.com/huage66/chat/api"
	"github.com/huage66/chat/config"
	"github.com/huage66/chat/database"
	"github.com/huage66/chat/interval/repo"
)

func Init() {
	// 初始化config
	config.Setup()
	// 连接redis
	database.NewRedis()
	// 加载群聊名称进内存
	repo.InitGroup()
	api.Server()
}
