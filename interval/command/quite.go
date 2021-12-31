package command

import (
	"github.com/huage66/chat/interval/msg"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
)

type Quite struct {
	ReceiveType string
	Ip          string
	SentTo      string
}

func (q *Quite) Name() string {
	return "quite"
}

func (q *Quite) Run() model.ReceiveMessage {
	if q.ReceiveType == vars.GroupType {
		// 更新内存数据
		maps, ok := vars.ChatMap.Load(q.SentTo)
		if ok {
			if m, ok := maps.(map[string]bool); ok {
				delete(m, q.Ip)
			}
		}
	}
	return model.ReceiveMessage{
		ReceiveMessage: msg.QuitSuccess,
		Success:        true,
	}
}
