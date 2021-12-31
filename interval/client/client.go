package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func main() {
	dial, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("连接失败, 请重新尝试")
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	fmt.Println("使用list可以查看命令列表")
	go write(dial)
	go read(dial)
	<- c
	fmt.Println("系统退出")
}