package server

import (
	"encoding/json"
	"fmt"
	"github.com/huage66/chat/interval/command"
	"github.com/huage66/chat/interval/msg"
	"github.com/huage66/chat/interval/repo"
	"github.com/huage66/chat/interval/utils"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
	"io"
	"net"
	"strings"
)

// 单聊广播
//func SingleWork(chat model.ChatMessage) error {
//
//}

// 群聊广播
func GroupWork(conn net.Conn, chat model.ChatMessage) {
	rec := model.ReceiveMessage{
		SendTo:         chat.SendTo,
		ReceiveMessage: fmt.Sprintf("%s\t%s", chat.Ip, chat.Message),
		Success:        true,
	}
	userInfo, ok := vars.UserMap.Load(rec.SendTo)
	if ok {
		userItem, ok := userInfo.(*model.UserInfo)
		if ok && len(userItem.Name) > 0 {
			rec.ReceiveMessage = fmt.Sprintf("%s\t%s", userItem.Name, chat.Message)
		}
	}
	load, ok := vars.ChatMap.Load(chat.SendTo)
	if !ok {
		write(conn, msg.GroupNotExists)
		return
	}

	if m, ok := load.(map[string]bool); ok {
		for name, _ := range m {
			if name == chat.Ip {
				continue
			}
			if userInterface, ok := vars.UserMap.Load(name); ok {
				if info, ok := userInterface.(*model.UserInfo); ok {
					if info.Conn == nil {
						vars.UserMap.Delete(info.IP)
						delete(m, info.IP)
						vars.ChatMap.Store(chat.SendTo, m)
						continue
					}
					go WriteReceive(info.Conn, rec)
				}
			}
		}
	}

}

// 命令管理, 命令解析,
func CmdWork(conn net.Conn, chat model.ChatMessage) {
	switch chat.ChatType {
	case vars.Select:
		c := command.Select{
			Ip:  chat.Ip,
			Msg: chat.Message,
			Arg: chat.Arg,
		}
		WriteReceive(conn, c.Run())
	case vars.Quite:
		c := command.Quite{
			ReceiveType: chat.ChatType,
			Ip:          chat.Ip,
			SentTo:      chat.SendTo,
		}
		WriteReceive(conn, c.Run())
	case vars.MakeGroup:
		c := command.Make{Ip: chat.Ip, Msg: chat.Message}
		WriteReceive(conn, c.Run())
	case vars.Register:
		if len(strings.TrimSpace(chat.Message)) < 1 {
			write(conn, "名字不符合规范, 请重新命名")
			return
		}
		c := command.Register{Ip: chat.Ip, Username: chat.Message}
		WriteReceive(conn, c.Run())
	case vars.Rename:
		c := command.Rename{Ip: chat.Ip, Username: chat.Message}
		WriteReceive(conn, c.Run())
	}
}

func Handler(conn net.Conn) {
	chatMessage := make(chan model.ChatMessage)
	go read(conn, chatMessage)
	for chat := range chatMessage {
		// 判断登录处理
		if chat.ChatType != vars.Register && chat.ChatType != vars.Rename {
			if !repo.Login(chat, conn) {
				write(conn, "连接失败, 新用户,请注册之后在登录, 使用register命令来注册新用户")
				continue
			}
		}

		switch chat.CmdType {
		case vars.CmdType:
			CmdWork(conn, chat)
		case vars.GroupType:
			GroupWork(conn, chat)
		}
	}
}

func read(conn net.Conn, c chan model.ChatMessage) {
	var chat model.ChatMessage
	for {
		bArr := make([]byte, 1024)
		if conn == nil {
			break
		}
		n, err := conn.Read(bArr)
		if err != nil {
			bArr = bArr[:0]
			fmt.Println("enter err", err)
			return
		}
		if err == io.EOF {
			conn.Close()
			return
		}

		err = json.Unmarshal(bArr[:n], &chat)
		bArr = bArr[:0]
		if err != nil {
			write(conn, msg.SendMsgFail)
			continue
		}
		c <- chat
	}
}

func WriteReceive(conn net.Conn, receive model.ReceiveMessage) {
	marshal, _ := json.Marshal(receive)
	_, err := conn.Write(marshal)
	if err != nil {
		utils.AddError("write message fail", err)
	}
}

func write(conn net.Conn, msg string) {
	receive := model.ReceiveMessage{
		ReceiveMessage: msg,
	}
	data, _ := json.Marshal(receive)
	_, err := conn.Write(data)
	if err != nil {
		utils.AddError("receive message fail", err)
	}
}
