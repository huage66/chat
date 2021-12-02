package command

import (
	"context"
	"fmt"
	"github.com/huage66/chat/interval/keys"
	"github.com/huage66/chat/interval/msg"
	"github.com/huage66/chat/interval/orm"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
)

type Make struct {
	Ip  string
	Msg string
}

func (r *Make) GetName() string {
	return "make room"
}

func (r *Make) Run() model.ReceiveMessage {
	var rec model.ReceiveMessage
	rec.ReceiveType = vars.CmdType
	ctx := context.Background()
	groupKey := fmt.Sprintf("%s:%s", keys.GroupPrefix, r.Msg)
	data := orm.RedisClient.Exists(ctx, groupKey)
	if data.Err() != nil {
		rec.ReceiveMessage = msg.CreateGroupFail
		return rec
	}
	if data.Val() == 0 {
		rec.ReceiveMessage = msg.GroupExists
	}

	if err := orm.RedisClient.Set(ctx, groupKey, r.Msg, 0).Err(); err != nil {
		rec.ReceiveMessage = msg.CreateGroupFail
		return rec
	}

	// 内存中创建chatMap
	maps := make(map[string]bool)
	maps[r.Ip] = true
	vars.ChatMap.Store(r.Msg, maps)
	rec.ReceiveType = vars.GroupType
	rec.ReceiveMessage = "群聊创建成功, 快邀请小伙伴加入群聊"
	rec.Success = true
	return rec
}
