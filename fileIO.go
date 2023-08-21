package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	//opening the file
	inputFileName := "input.txt"
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer inputFile.Close()
	//creating the output file
	outputFileName := "output.txt"
	outputFile, err0 := os.Create(outputFileName)
	if err0 != nil {
		log.Fatal(err0)
		return
	}
	defer outputFile.Close()
	// reading the input file
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := strings.ToUpper(line)
		print(processedLine + "\n")
		//writing the file
		_, err01 := outputFile.WriteString(processedLine + "\n")
		if err01 != nil {
			log.Fatal(err01)
		}
	}
	if err := scanner.Err(); err != nil {
		print("Error reading input file:", err)
	}
}
