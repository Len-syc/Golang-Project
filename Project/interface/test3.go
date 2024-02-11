package main

import "fmt"

// Calculator 计算器接口：设计一个接口 Calculator，包括 Add(a, b float64) float64、Subtract(a, b float64) float64、Multiply(a, b float64) float64 和 Divide(a, b float64) (float64, error)
// 方法，分别用于加法、减法、乘法和除法运算。然后编写一个结构体 BasicCalculator 来实现该接口，实现基本的四则运算功能。
type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
}
type BasicCalculator struct{}

func (value *BasicCalculator) Add(a, b float64) float64 {
	return a + b
}
func (value *BasicCalculator) Subtract(a, b float64) float64 {
	return a - b
}
func (value *BasicCalculator) Multiply(a, b float64) float64 {
	return a * b
}
func (value *BasicCalculator) Divide(a, b float64) float64 {
	return a / b
}
func main() {
	test := BasicCalculator{}
	result := test.Add(10, 20)
	fmt.Println(result)
}
