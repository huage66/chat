package model

// 群聊, 群聊信息
type GroupChat struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"`
	Members   []User `json:"members"`
}

// 发送方消息组成
type ChatMessage struct {
	CmdType  string `json:"cmd_type"`  // 命令类型
	ChatType string `json:"chat_type"` // 聊天类型
	SendTo   string `json:"send_to"`   // 当聊天类型为Group,这个为群聊名称,当为个人时,为个人IP
	Ip       string `json:"ip"`        // 自己的Ip地址
	Arg      string `json:"arg"`       // 当为命令类型时, 这个为命令参数
	Message  string `json:"message"`   // 消息
}

// 接收方消息组成 []byte,解析为这个样子
type ReceiveMessage struct {
	SendTo         string `json:"receive_to"`      // 之后需要发送的对象, 可能是群聊名称,可能是单人Ip
	ReceiveMessage string `json:"receive_message"` // 接收的消息
	Success        bool   `json:"success"`         // 是否成功
}
