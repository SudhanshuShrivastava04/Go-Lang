package main

import (
	"fmt"
)

var print = fmt.Printf

func main() {
	arrayNums := []int{11, 12, 23, 64, 35}
	for i, nums := range arrayNums {
		print("%d : %d\n", i, nums)
	}
}
