package main

import (
	"fmt"
)

// ConsoleLogger 简单日志记录接口：定义一个接口 Logger，包括 Log(message string) 方法，用于记录日志消息。然后编写一个结构体 ConsoleLogger 来实现该接口，将日志消息输出到控制台
type ConsoleLogger struct {
}
type Logger interface {
	Log(message string)
}

func (Value *ConsoleLogger) Log(message string) {
	fmt.Println("Log:", message)

}

func main() {
	test := &ConsoleLogger{}
	test.Log("Hello world")
	var (
		demo int
	)
	fmt.Print("value:")
	fmt.Scanln(&demo)
	if demo < 10 {
		test.Log("demo<10")
	} else {
		test.Log("demo>10")
	}
}
