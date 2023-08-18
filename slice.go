package main

import (
	"fmt"
)

var print = fmt.Printf

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:5]
	print("first 3 : %d \n", slice[:3])
	print("last 3 : %d \n", slice[2:])
	print("slice : %d \n", slice)
	slice[0] = 10
	print("modified slice : %d \n", slice)
	print("array : %d \n", array)
	slice = append(slice, 20)
	print("print new slice : %d \n", slice)
}
