package main

import (
	"fmt"
)

var print = fmt.Println

func double(arr *[4]int) {
	for i := 0; i < 4; i++ {
		arr[i] *= 2
	}
}
func main() {
	array := [4]int{1, 2, 3, 4}
	double(&array)
	for _, val := range array {
		print("", val)
	}
}
