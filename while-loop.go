package main

import (
	"fmt"
	"log"
	"math/rand"
)

var print = fmt.Println

func main() {
	randNum := rand.Intn(50) + 1
	print("Choose a random num b/w 1 to 50")
	print("random number is :", randNum)
	for true {
		// reader := bufio.NewReader(os.Stdin)
		// guess, err := reader.ReadString('\n')
		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil {
			log.Fatal(err)
		}
		if guess > randNum {
			print("Pick a lower value")
		} else if guess < randNum {
			print("Pick a higher value")
		} else {
			print("you guessed it!")
			break
		}
	}
}
