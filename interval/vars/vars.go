package vars

import (
	"sync"
)

const (
	// chat type
	CmdType   = "0"
	GroupType = "2"

	// 命令模式
	Select    = "select"   // 选择聊天对象 -s 选择单聊 -g 选择群聊
	Quite     = "quit"     // 退出聊天
	MakeGroup = "make"     // 创建聊天室
	Register  = "register" // 注册用户
	Rename    = "rename"   // 更改用户名
)

var (
	UserMap     sync.Map
	ChatMap     sync.Map
	OnLineCount int
)
