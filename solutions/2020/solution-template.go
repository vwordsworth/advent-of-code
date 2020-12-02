package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func getInput() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	start := time.Now()

	input := getInput()

	fmt.Println("\nTime elapsed:\t", time.Now().Sub(start))
}
