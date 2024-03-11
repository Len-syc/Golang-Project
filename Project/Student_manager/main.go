package main

import (
	"fmt"
	"os"
)

// 需求:
// 1.添加学员信息
// 2.编辑学员信息
// 3.展示所有学员信息
func ShowMenu() {
	fmt.Println("--------学生管理系统--------")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}
func GetUserInfo() *Student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Print("请输入学生ID：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生名字：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生班级：")
	fmt.Scanf("%d\n", &class)
	Stu := NewStudent(id, name, class)
	return Stu

}
func main() {

	sm := NewStudentMar()
	for {
		var input int
		//1.打印菜单
		//2.等待用户选择要执行的选项
		//3.执行用户选择选项
		ShowMenu()
		fmt.Print("请输入你要操作的序号：")
		fmt.Scanf("%d\n", &input)
		//fmt.Println("用户输入的是：", input)
		switch input {
		case 1:
			//添加学员信息
			sm.AddStudent(GetUserInfo())
		case 2:
			//编辑学员信息
			sm.EditStudent(GetUserInfo())
		case 3:
			//展示所有学员信息
			sm.ShowAll()
		case 4:
			//退出系统
			os.Exit(0)
		}
	}
}
