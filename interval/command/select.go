package command

import (
	"github.com/huage66/chat/interval/msg"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
)
type Select struct {
	Ip string
	Arg string
	Msg string
}

func (s *Select) GetName() string {
	return "select"
}

// -s 单聊模式 选择用户 -g 群聊模式
func (s *Select) Run() model.ReceiveMessage {
	rec := model.ReceiveMessage{}
	switch s.Arg {
	// 目前单聊模式在命令行版本暂时无法实现
	//case "-s":
	//	load, ok := vars.UserMap.Load(s.Msg)
	//	if !ok {
	//
	//		break
	//	}
	//	rec.ReceiveType = vars.SingleType
	//	if ok {
	//		rec.SendTo = info.IP
	//		rec.ReceiveMessage = "连接成功, 开始聊天"
	//		rec.Success = true
	//	} else {
	//		rec.ReceiveMessage = msg.ConnFailS
	//	}
	case "-g":
		load, ok := vars.ChatMap.Load(s.Msg)
		rec.ReceiveType = vars.GroupType
		if ok {
			m, ok := load.(map[string]bool)
			if !ok {
				rec.ReceiveMessage = msg.ConnFailG
				break
			}
			m[s.Ip] = true
			rec.SendTo = s.Msg
			rec.ReceiveMessage = msg.ConnSuccess
			rec.Success = true
		} else {
			rec.ReceiveMessage = msg.ConnFailG
		}
	}

	return rec
}
