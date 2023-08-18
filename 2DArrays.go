package main

import (
	"fmt"
	"log"
)

var print = fmt.Printf

func main() {
	var col int
	var rows int
	print("enter columns \n")
	_, err := fmt.Scan(&col)
	if err != nil {
		log.Fatal(err)
	}
	print("enter rows \n")
	_, err01 := fmt.Scan(&rows)
	if err01 != nil {
		log.Fatal(err01)
	}
	newArray := make([][]rune, rows)
	for i := range newArray {
		newArray[i] = make([]rune, col)
		for j := range newArray[i] {
			newArray[i][j] = '*'
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			print("%c", newArray[i][j])
		}
		print("\n")
	}
}
