package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("net", "49.234.14.141:8080")
	if err != nil {
		fmt.Println("连接失败, 请重新尝试")
		return
	}
	c := make(chan os.Signal)

	go write(dial)
	go read(dial)

	<- c
	fmt.Println("系统退出")
}
