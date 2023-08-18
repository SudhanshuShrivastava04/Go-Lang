package main

import (
	"fmt"
)

var print = fmt.Println

func change(ptr *string) {
	*ptr = "hello world"
}
func main() {
	a := "hello"
	print("before :", a)
	change(&a)
	print("after :", a)
}
