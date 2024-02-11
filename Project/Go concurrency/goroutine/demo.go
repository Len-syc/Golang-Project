package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//	func hello(i int) {
//		fmt.Println("hello world", i)
//		wg.Done()
//	}
func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("hello world", i)
			wg.Done()
		}()
	}
	fmt.Println("hello main")
	wg.Wait()
}
