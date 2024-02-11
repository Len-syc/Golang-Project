package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, value := range slice {
		if value > 8 {
			fmt.Print(index)
		}
	}
}
