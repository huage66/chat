package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/huage66/chat/interval/vars"
	"github.com/huage66/chat/model"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
)

var sendTo string
var status int

func write(conn net.Conn, c chan os.Signal) {
	var rec model.ChatMessage
	rec.Ip = getLocalIp()
	scanner := bufio.NewReader(os.Stdin)
	for {
		if conn == nil {
			fmt.Println("连接关闭")
			signal.Stop(c)
			return
		}
		text, err := scanner.ReadString('\n')
		if err != nil {
			continue
		}
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
					continue
				}
				status = 1
				rec.Arg = args[1]
				rec.Message = args[2]
				rec.ChatType = vars.Select
				isCmd = true
			case vars.MakeGroup:
				if len(args) < 2 {
					fmt.Println("命令错误, 是否应该make [groupName]")
					continue
				}
				rec.Message = args[1]
				rec.ChatType = vars.MakeGroup
				isCmd = true
			case vars.Register:
				if len(args) < 2 {
					fmt.Println("命令错误, 是否应该register [username]")
					continue
				}
				rec.Message = args[1]
				rec.ChatType = vars.Register
				isCmd = true
			case vars.Rename:
				if len(args) < 2 {
					fmt.Println("命令错误, 是否应该rename [username]")
					continue
				}
				rec.Message = args[1]
				rec.ChatType = vars.Rename
				isCmd = true
			case vars.Quite:
				status = 0
				rec.ChatType = vars.Quite
				isCmd = true
			case "list":
				fmt.Printf("register [username]\n\t--注册用户\nrename [username]\n\t--重命名名称\nselect -g [groupName]\n\t--选择群聊\nmake [groupName]\n\t--创建群聊\nquit\n\t--退出群聊\n")
			}
		}

		if isCmd {
			rec.CmdType = vars.CmdType
			send(conn, rec)
		} else if status == 1 {
			rec.CmdType = vars.GroupType
			rec.SendTo = sendTo
			rec.Message = space
			send(conn, rec)
		}
	}
}

func read(conn net.Conn, c chan os.Signal) {
	var rec model.ReceiveMessage
	for {
		bArr := make([]byte, 1024)
		if conn == nil {
			return
		}
		n, err := conn.Read(bArr)
		if err == io.EOF {
			bArr = bArr[:0]
			fmt.Println("连接关闭")
			conn.Close()
			signal.Stop(c)
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.Unmarshal(bArr[:n], &rec)
		bArr = bArr[:0]
		if err != nil {
			continue
		}
		if len(rec.SendTo) > 0 {
			sendTo = rec.SendTo
		}
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
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		fmt.Println("Get local IP addr failed!!!")
		return "localhost"
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
