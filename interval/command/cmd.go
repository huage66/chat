package command

import "github.com/huage66/chat/model"

type Cmd interface {
	GetName() string
	Run() model.ReceiveMessage
}
