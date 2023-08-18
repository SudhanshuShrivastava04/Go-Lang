package main

import (
	"fmt"
	"log"
)

var print = fmt.Println

func getSumOfArray(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func main() {
	var size int
	print("enter the size : ")
	_, err := fmt.Scan(&size)
	if err != nil {
		log.Fatal(err)
	}
	array := make([]int, size)
	print("Enter nums")
	for i, _ := range array {
		_, err = fmt.Scan(&array[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	print("Sum :", getSumOfArray(array...))
}
