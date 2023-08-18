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
	var Ptr *string = &a //address of a
	print("before :", a, Ptr)
	change(&a)
	print("after :", a, Ptr)
}
