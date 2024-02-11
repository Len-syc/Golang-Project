package main

import "fmt"

type Player struct {
	name string
	UID  int
}

func (P *Player) Get() {
	fmt.Println("this is Get")
	fmt.Printf("PlayerName:%s\n", P.name)
	fmt.Printf("PlayerUID:%v\n", P.UID)
}
func (P *Player) demo() {
	fmt.Println("this is demo")
	fmt.Printf("[%s]hello world!", P.name)
}

type GetHttp interface {
	Get()
}
type Info interface {
	GetHttp
	demo()
}

func main() {
	var (
		test  GetHttp
		test2 Info
	)
	demo := Player{
		name: "LenSyc",
		UID:  100000,
	}
	demo2 := Player{
		name: "test",
		UID:  1000001,
	}
	test = &demo
	test2 = &demo2
	test.Get()
	test2.demo()
	test2.Get()
}
