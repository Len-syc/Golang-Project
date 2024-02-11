package main

import "fmt"

var (
	input     int
	UserSlice = make([]int, 0)
	temp      int
)

func main() {
	for {
		fmt.Print("输入数字，-1为结束:")
		fmt.Scanln(&input)
		if input == -1 {
			break
		}
		UserSlice = append(UserSlice, input)
	}
	fmt.Println(UserSlice)
	for i := 0; i < len(UserSlice); i++ {
		fmt.Println(UserSlice[i])
	}
}
