package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//建立连接
	connent, err := net.Dial("tcp", "127.0.0.1:2888")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.利用该connect连接，进行数据的发送和接受
	for {
		input := bufio.NewReader(os.Stdin)
		for {
			s, _ := input.ReadString('\n')
			s = strings.TrimSpace(s)
			if strings.ToUpper(s) == "Q" {
				return
			}
			///发消息
			_, err := connent.Write([]byte(s))
			if err != nil {
				fmt.Printf("错误，client:29.error:%v", err)
				return
			}
			//从服务端接收回复的消息
			var buf [1024]byte
			n, err := connent.Read(buf[:])
			if err != nil {
				fmt.Printf("错误，client:34.error:%v", err)
				return
			}
			fmt.Printf("\n收到服务端回复:", string(buf[:n]))
			var (
				test string
			)
			fmt.Scanln(&test)
		}
	}

}
