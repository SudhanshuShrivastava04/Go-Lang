package main

import (
	"fmt"
	"log"
)

var print = fmt.Printf

func getSum(x int, y int) int {
	var sum int
	sum = x + y
	return sum
}
func main() {
	print("enter x : \n")
	var x int
	_, err := fmt.Scan(&x)
	print("enter y : \n")
	var y int
	_, err0 := fmt.Scan(&y)
	if err != nil || err0 != nil {
		log.Fatal(err, err0)
	}
	print("Sum : %d \n", getSum(x, y))
}
