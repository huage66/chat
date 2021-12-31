package msg

const (
	ConnSuccess = "连接成功, 开始聊天"
	ConnFailG   = "连接失败, 群聊不存在"
	QuitSuccess = "退出成功"

	CreateGroupFail = "创建群聊失败, 请重新创建"
	GroupExists     = "群聊已存在, 请勿重复创建"
	GroupNotExists  = "群聊不存在, 请确认群聊名称"
	SendMsgFail     = "消息发送失败, 请重新尝试"
)
