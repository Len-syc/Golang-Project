package main

import "fmt"

var (
	slice = make([]int, 3, 10)
)

func main() {
	var (
		temp int
	)
	slice = []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(slice); i++ {
		temp += slice[i]
	}
	if temp <= 21 {
		slice = append(slice[0:1], slice[2:3]...)
		fmt.Print(slice)
	}
}
