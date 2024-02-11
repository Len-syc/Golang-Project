package main

import (
	"fmt"
	"strings"
)

func MakeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	r := MakeSuffixFunc(".bin")
	fmt.Println(r("HK4E"))
}
