package api

import (
	"fmt"
	"github.com/huage66/chat/config"
	"github.com/huage66/chat/interval/server"
	"net"
)

func Server() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.ApplicationConf.Host, config.ApplicationConf.Port))
	if err != nil {
		fmt.Printf("Server: net Listen err = %v\n", err)
		panic(err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Printf("连接失败")
			continue
		}
		// 处理发送的数据
		go server.Handler(accept)
	}
}
