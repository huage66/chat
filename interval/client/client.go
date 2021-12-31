package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接失败, 请重新尝试")
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	fmt.Println("使用list可以查看命令列表")
	go write(dial, c)
	go read(dial, c)
	<- c
	fmt.Println("系统退出")
}