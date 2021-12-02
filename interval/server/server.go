package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/huage66/chat/interval/command"
	"github.com/huage66/chat/interval/msg"
	"github.com/huage66/chat/interval/repo"
	"github.com/huage66/chat/interval/utils"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
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
		ReceiveType:    chat.ChatType,
		SendTo:         chat.SendTo,
		ReceiveMessage: fmt.Sprintf("%s\t%s", chat.Ip, chat.Message),
		Success:        true,
	}
	load, ok := vars.ChatMap.Load(chat.SendTo)
	if !ok {
		write(conn, msg.GroupNotExists)
		return
	}

	if m, ok := load.(map[string]bool); ok {
		for name, _ := range m {
			if userInterface, ok := vars.UserMap.Load(name); ok {
				if info, ok := userInterface.(model.UserInfo); ok {
					go WriteReceive(info.Conn, rec)
				}
			}
		}
	}

}

// 命令管理, 命令解析,
func CmdWork(conn net.Conn, chat model.ChatMessage) {
	data := strings.Split(chat.Message, "|")
	switch data[0] {
	case vars.Select:
		c := command.Select{
			Ip:  chat.Ip,
			Msg: chat.Message,
			Arg: data[1],
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
		c := command.Make{Ip: chat.Ip, Msg: data[1]}
		WriteReceive(conn, c.Run())
	}
}

func Handler(conn net.Conn) {
	chatMessage := make(chan model.ChatMessage)
	go read(conn, chatMessage)
	for chat := range chatMessage {
		// 判断登录处理
		if !repo.Login(chat, conn) {
			write(conn, "连接失败, 请重新尝试")
			conn.Close()
			return
		}

		switch chat.ChatType {
		case vars.CmdType:
			CmdWork(conn, chat)
		case vars.GroupType:
			GroupWork(conn, chat)
		}
	}
}

func read(conn net.Conn, c chan model.ChatMessage) {
	var chat model.ChatMessage
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if scanner.Err() != nil {
			conn.Close()
			break
		}
		err := json.Unmarshal(scanner.Bytes(), &chat)
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
