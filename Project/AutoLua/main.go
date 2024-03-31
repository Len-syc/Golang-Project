package main

import (
	"fmt"
	"os"
)

func ShowMenu() {
	fmt.Println("--------AutoLua--------")
	fmt.Println("1.编译复制")
	fmt.Println("2.编译")
	fmt.Println("3.退出系统")
}
func main() {
	for {
		var input int
		//1.打印菜单
		//2.等待用户选择要执行的选项
		//3.执行用户选择选项
		ShowMenu()
		fmt.Print("请输入你要操作的序号：")
		fmt.Scanf("%d\n", &input)
		//fmt.Println("用户输入的是：", input)
		switch input {
		case 1:
		case 2:
		case 3:
			//退出系统
			os.Exit(0)
		}
	}
}
