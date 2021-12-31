package command

import (
	"context"
	"github.com/huage66/chat/interval/keys"
	"github.com/huage66/chat/interval/orm"
	"github.com/huage66/chat/model"
)

type Rename struct {
	Username string
	Ip       string
}

func (r *Rename) Name() string {
	return "rename"
}

func (r *Rename) Run() model.ReceiveMessage {
	var (
		user model.User
		rec  model.ReceiveMessage
	)
	if err := orm.RedisClient.Get(context.Background(), keys.UserPrefix+r.Ip).Scan(&user); err != nil {
		rec.ReceiveMessage = "重命名失败"
		return rec
	}
	rec.Success = true
	rec.ReceiveMessage = "重命名成功, 请使用select -g [groupName]来和群友展示一下你炫酷的名字"
	return rec
}
