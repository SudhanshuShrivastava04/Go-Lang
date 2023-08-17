package main

import (
	"fmt"
	"log"
)

func main() {
	var size int

	fmt.Print("Enter the size of the array: ")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("You entered size: %d\n", size)

	myArray := make([]int, size)

	for i := 0; i < size; i++ {
		_, err = fmt.Scan(&myArray[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Array elements:")
	for _, value := range myArray {
		fmt.Println(value)
	}
}
