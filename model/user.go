package model

import (
	"encoding/json"
	"net"
)

type User struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type UserInfo struct {
	User
	Conn   net.Conn `json:"conn"`
	OnLine bool     `json:"on_line"`
}
