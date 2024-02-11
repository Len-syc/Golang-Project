package main

import (
	"fmt"
)

func a() {
	fmt.Println("a")
}
func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error in b")
		}
	}()
	panic("b")
}
func c() {
	fmt.Println("c")
}
func main() {
	a()
	b()
	c()
}
