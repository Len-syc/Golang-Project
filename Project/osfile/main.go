package main

import (
	"log"
	"os"
)

var (
	newfile *os.File
	err     error
)

func main() {

	// 创建文件
	newfile, err = os.Create("./file/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created file:", newfile.Name())

	// 关闭文件
	err = newfile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
