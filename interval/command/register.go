package command

import (
	"context"
	"github.com/huage66/chat/interval/keys"
	"github.com/huage66/chat/interval/orm"
	"github.com/huage66/chat/model"
)

type Register struct {
	Username string
	Ip       string
}

func (r *Register) Name() string {
	return "register"
}

func (r *Register) Run() model.ReceiveMessage {
	user := model.User{
		Name: r.Username,
		IP:   r.Ip,
	}
	var rec model.ReceiveMessage
	if err := orm.RedisClient.Set(context.Background(), keys.UserPrefix+r.Ip, &user, 0).Err(); err != nil {
		rec.ReceiveMessage = "注册失败"
		return rec
	}
	rec.Success = true
	rec.ReceiveMessage = "注册成功, 请使用select -g [groupName]来聊天吧"
	return rec
}
