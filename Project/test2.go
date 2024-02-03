package main

import "fmt"

func GetMap(x, y int) string {
	var (
		xy_map string
		//result string
	)
	switch {
	case x == 0 && y == 0:
		xy_map = "原点"
	case x == 0 && y != 0:
		xy_map = "y轴"
	case x != 0 && y == 0:
		xy_map = "x轴"
	case x > 0 && y > 0:
		xy_map = "第一象限"
	case x < 0 && y > 0:
		xy_map = "第二象限"
	case x < 0 && y < 0:
		xy_map = "第三象限"
	case x > 0 && y < 0:
		xy_map = "第四象限"
	}
	//result = fmt.Sprintf("处于:%s,坐标轴(%d,%d)", xy_map, x, y)
	return fmt.Sprintf("处于:%s,坐标轴(%d,%d)", xy_map, x, y)
}

func main() {
	var (
		x    int
		y    int
		ret  string
		test string
	)
	fmt.Print("x:")
	fmt.Scanln(&x)
	fmt.Print("y:")
	fmt.Scanln(&y)
	ret = GetMap(x, y)
	fmt.Println(ret)
	fmt.Println("hello world")
	fmt.Print("程序结束")
	fmt.Scanln(&test)
}
