package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseLine(line string) (int, int, byte, string) {
	words := strings.Fields(line)

	bounds := strings.Split(words[0], "-")
	min, _ := strconv.Atoi(bounds[0])
	max, _ := strconv.Atoi(bounds[1])

	letter := words[1][0]

	password := words[2]

	return min, max, letter, password
}

func isPasswordCountValid(password string, letter byte, min int, max int) bool {
	count := 0

	for _, c := range password {
		if byte(c) == letter {
			count++
		}
	}

	return count >= min && count <= max
}

func isPasswordPositionValid(password string, letter byte, pos1 int, pos2 int) bool {
	return (password[pos1-1] == letter) != (password[pos2-1] == letter)
}

func main() {
	start := time.Now()
	passwords := getInput()

	validCountPasswords := 0
	validPositionPasswords := 0

	for _, line := range passwords {
		int1, int2, letter, pwd := parseLine(line)

		if isPasswordCountValid(pwd, letter, int1, int2) {
			validCountPasswords++
		}

		if isPasswordPositionValid(pwd, letter, int1, int2) {
			validPositionPasswords++
		}
	}

	fmt.Println("Valid count passwords:\t", validCountPasswords)
	fmt.Println("Valid position passwords:\t", validPositionPasswords)
	fmt.Println("\nTime elapsed:\t", time.Now().Sub(start))
}
