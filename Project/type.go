package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		test = "ggg"
	)
	num, err := strconv.Atoi(test)
	/*strconv.iota - int - str
	strconv.ParseFloat - str - float
	strconv.FormatFloat- float - str



	*/
	if err != nil {
		fmt.Println("转化失败", err)
	} else {
		fmt.Printf("success!:%v", num)
	}
}
