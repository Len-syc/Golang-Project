package main

import "fmt"

var (
	balance_new = [...]float32{1.10000, 12.234, 1234, 2345}
)

func GetBalance(balance []float32) float32 {
	var temp float32
	for i := 0; i < len(balance); i++ {
		temp += balance[i]
	}
	return temp
}

func main() {
	var (
		Get float32
	)
	Get = GetBalance(balance_new[:])
	fmt.Println(Get)
}
