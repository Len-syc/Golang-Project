//package main
//
//import "fmt"
//
//func main() {
//	//var (
//	//	a  = 20 /* 声明实际变量 */
//	//	ip *int /* 声明指针变量 */
//	//)
//	//ip = &a                          /* 指针变量的存储地址 */
//	//fmt.Printf("a 变量的地址是: %x\n", &a) //%x16进制 %d十进制
//	///* 指针变量的存储地址 */
//	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
//	///* 使用指针访问值 */
//	//fmt.Printf("*ip 变量的值: %d\n", *ip)
//	var (
//		test    int
//		Pointer *int
//	)
//	fmt.Print("输入值,为test:")
//	fmt.Scanln(&test)
//	fmt.Printf("[Debug]adress(初始值:test)[10]:%d\n", test)
//	fmt.Printf("[Debug]adress(初始值:test)[16]:%x\n", &test)
//
//	fmt.Printf("[Debug]adress(初始值:pointer)[10]:%d\n", Pointer)
//	fmt.Printf("[Debug]adress(初始值:pointer)[16]:%x\n", &Pointer)
//	Pointer = &test
//	fmt.Print("指针-------\n")
//	fmt.Printf("[Debug]adress(指针:pointer)[10]:%d\n", *Pointer)
//	fmt.Printf("[Debug]adress(指针:pointer)[16]:%x\n", &Pointer)
//
//}

package main

import "fmt"

func main() {
	var ptr *int

	fmt.Printf("ptr 的值为 : %v\n", ptr)
	fmt.Printf("ptr 的值为 : %d\n", ptr)
}
