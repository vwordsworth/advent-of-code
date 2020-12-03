package main

import (
	"bufio"
	"fmt"
	"os"
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
	input := getInput()
	fmt.Println(input)
}
