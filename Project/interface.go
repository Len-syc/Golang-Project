package main

import "fmt"

type dog struct{}
type cat struct{}
type people struct {
	name string
}

func (d dog) say() {
	fmt.Println("wawa~")
}
func (c cat) say() {
	fmt.Println("喵喵喵~")
}
func (p people) say() {
	fmt.Println("呐呐~")
}

type sayer interface {
	say()
}

func da(arg sayer) {
	arg.say()
}

func main() {
	c1 := cat{}
	da(c1)
	d2 := dog{}
	da(d2)
	p1 := people{"Len"}
	da(p1)
}
