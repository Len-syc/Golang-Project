package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   int
	Name string
}
type class struct {
	Title    string
	Students []Student
}

func newstudents(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}

}
func main() {
	test1 := class{
		"101",
		make([]Student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		temstu := newstudents(i, fmt.Sprintf("stu%02d", i))
		test1.Students = append(test1.Students, temstu)
	}
	//fmt.Printf("%#v\n", test1)
	data, err := json.Marshal(test1)
	if err != nil {
		fmt.Println("marshal failed err:", err)
		return
	}
	fmt.Printf("%s\n", data)
}
