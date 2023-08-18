package main

import (
	"fmt"
)

var print = fmt.Printf

func main() {
	String := "hello world"
	array := []rune(String)
	for _, value := range array {
		print("Rune Array Elements : %c \n", value)
	}
}
