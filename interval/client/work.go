package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
	"net"
	"strings"
)

var sendTo string
func write(conn net.Conn) {
	var rec model.ChatMessage
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if scanner.Err() != nil {
			conn.Close()
			break
		}
		text := scanner.Text()
		space := strings.TrimSpace(text)
		data := strings.Split(space, " ")
		args := make([]string, 0, 3)
		for _, ele := range data {
			if len(ele) > 0 {
				args = append(args, ele)
			}
		}

		isCmd := false
		if len(args) > 0 {
			switch args[0] {
			case vars.Select:
				if len(args) <= 2 {
					fmt.Println("命令错误, 是否应该select -g [groupName]")
					return
				}
				rec.Arg = args[1]
				rec.Message = args[2]
				rec.ChatType = vars.Select
				isCmd = true
			case vars.MakeGroup:
				if len(args) < 2 {
					fmt.Println("命令错误, 是否应该make [groupName]")
					return
				}
				rec.Message = args[1]
				rec.ChatType = vars.MakeGroup
				isCmd = true
			case vars.Quite:
				rec.ChatType = vars.Quite
				isCmd = true
			}
		}

		if isCmd {
			send(conn, rec)
		}else {
			rec.ChatType = vars.GroupType
			rec.SendTo = sendTo
			rec.Message = space
			send(conn, rec)
		}
	}
}

func read(conn net.Conn) {
	var rec model.ReceiveMessage
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if scanner.Err() != nil {
			conn.Close()
			return
		}
		err := json.Unmarshal(scanner.Bytes(), &rec)
		if err != nil {
			continue
		}
		sendTo = rec.SendTo
		fmt.Println(rec.ReceiveMessage)
	}
}

func send(conn net.Conn, chat model.ChatMessage) {
	data, err := json.Marshal(chat)
	if err != nil {
		fmt.Println("消息发送失败, 请重新发送")
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("消息发送失败, 请重新发送")
		return
	}
}

func getLocalIp() string {
	return ""
}
