package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Class string
}

func NewStudent(id int, name string, class string) *Student {
	return &Student{
		ID:    id,
		Name:  name,
		Class: class,
	}
}

type StudentMar struct {
	AllStudent []*Student
}

// NewStudentMar的构造函数
func NewStudentMar() *StudentMar {
	return &StudentMar{
		AllStudent: make([]*Student, 0, 100),
	}
}

// 添加学生
func (s *StudentMar) AddStudent(new *Student) {
	s.AllStudent = append(s.AllStudent, new)
}

// 编辑学生
func (s *StudentMar) EditStudent(new *Student) {
	for i, v := range s.AllStudent {
		if new.ID == v.ID {
			s.AllStudent[i] = new
			return // 找到学生并编辑完成，直接返回
		}
	}
	fmt.Printf("输入的学生信息有误，该ID：%d不在系统内\n", new.ID)
}

// 展示学生
func (s *StudentMar) ShowAll() {
	for _, v := range s.AllStudent {
		fmt.Println("-----------展示学生-----------")
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.ID, v.Name, v.Class)
	}
}
