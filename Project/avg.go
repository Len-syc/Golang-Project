package main

import "fmt"

func main() {
	var number = [5]int{1000, 2, 3, 17, 50}
	var avg float32
	avg = GetNumber(number, 5)
	fmt.Printf("平均值为: %f ", avg)
}
func GetNumber(arr [5]int, size int) float32 {
	var (
		i, sum int
		avg    float32
	)
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}
