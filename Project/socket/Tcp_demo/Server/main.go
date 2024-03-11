package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() //处理好后，关闭连接
	// 针对当前的连接做数据的发送和接收操作
	reader := bufio.NewReader(conn)
	var buf [128]byte
	n, err := reader.Read(buf[:])
	if err != nil {
		fmt.Printf("reader:%v", err)
	}
	recv := string(buf[:n])
	fmt.Println("接受到的数据:", recv)
	conn.Write(buf[:]) //把收到的数据返回给客户端
}
func main() {
	//1.建立连接
	listen, err := net.Listen("tcp", "127.0.0.1:2888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("server running in:127.0.0.1:2888")
	for {
		//等待Client
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//启动进程
		go process(connect)
	}
}
