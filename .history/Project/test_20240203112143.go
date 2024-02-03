package main

import (
	"fmt"
)

func GetPlayerTokenRsp() (string, int) {
	return "Player", 10020
}
func main() {
	player, Token := GetPlayerTokenRsp()
	fmt.Println("hello world!")
	tests := "hello world"
	fmt.Printf("type of test %T\n", tests)
	fmt.Println("player:{}Token:{}", player, Token)
	println(Token)
}
