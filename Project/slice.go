package main

import "fmt"

// 创建一个空切片
var (
// slice1 []int

)

func main() {
	//// 使用make函数创建切片，指定长度和容量
	//slice2 := make([]int, 3)    // 创建一个长度为3的切片，初始值都为0
	//slice3 := make([]int, 3, 5) // 创建一个长度为3，容量为5的切片
	var (
		demoPoin *[]int
	)
	test := make([]int, 0, 10)
	test = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	demoPoin = &test
	*demoPoin = append(*demoPoin, 11)
	//*demoPoin = append(*demoPoin, 12)
	//new := make([]int, len(*demoPoin))
	//copy(new, *demoPoin)
	//fmt.Print(*demoPoin, "\n")
	//fmt.Print(len(*demoPoin), "\n")
	//fmt.Print(&test, "\n")
	//fmt.Print(demoPoin, "\n")
	//fmt.Print(&new)
	newTest := cap(test)
	fmt.Print(newTest)
}
