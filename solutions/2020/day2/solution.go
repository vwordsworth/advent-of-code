package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

/*
* Initial submission
 */
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

/*
* Refactor
 */
type entry struct {
	value      string
	letter     byte
	int1, int2 int
}

func newEntry(line string) *entry {
	words := strings.Fields(line)
	bounds := strings.Split(words[0], "-")

	pwdEntry := entry{}
	pwdEntry.int1, _ = strconv.Atoi(bounds[0])
	pwdEntry.int2, _ = strconv.Atoi(bounds[1])
	pwdEntry.letter = words[1][0]
	pwdEntry.value = words[2]

	return &pwdEntry
}

func refactorIsPasswordCountValid(line *entry) bool {
	count := strings.Count(line.value, string(line.letter))
	return count >= line.int1 && count <= line.int2
}

func refactorIsPositionValid(line *entry) bool {
	return (line.value[line.int1-1] == line.letter) != (line.value[line.int2-1] == line.letter)
}

func main() {
	passwords := getInput()

	validCountPasswords := 0
	validPositionPasswords := 0

	for _, line := range passwords {
		pwdEntry := newEntry(line)

		if refactorIsPasswordCountValid(pwdEntry) {
			validCountPasswords++
		}

		if refactorIsPositionValid(pwdEntry) {
			validPositionPasswords++
		}
	}

	fmt.Println("Valid count passwords:\t", validCountPasswords)
	fmt.Println("Valid position passwords:\t", validPositionPasswords)
}
