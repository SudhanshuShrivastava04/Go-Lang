package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	x := 1
	for x <= 10 {
		print(2 * x)
		x++
	}
}
