package main

import "fmt"

func main() {
	TestMap := map[string]int{
		"User":     1,
		"UID":      100001,
		"Password": 6230,
	}
	TestMap["User"] = 102
	fmt.Println(TestMap)
	fmt.Println(TestMap["UID"])
	demo, v1 := TestMap["UID"]
	fmt.Printf("demo:%v\nv1:%v\n", demo, v1)
	for i, w := range TestMap {
		fmt.Printf("i:%v,w:%v\n", i, w)
	}
	delete(TestMap, "UID")
	fmt.Println(TestMap)
}
