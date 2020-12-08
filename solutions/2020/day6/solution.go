package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput() [][]string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines [][]string
	var group []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		if len(currentLine) == 0 {
			lines = append(lines, group)
			group = nil
		} else {
			group = append(group, currentLine)
		}
	}

	if len(group) > 0 {
		lines = append(lines, group)
	}

	return lines
}

func countGroupTotals(groups [][]string) int {
	total := 0

	for _, group := range groups {
		groupCount := 0
		set := make(map[rune]bool)

		for _, person := range group {
			for _, letter := range person {
				if _, seen := set[letter]; !seen {
					set[letter] = true
					groupCount++
				}
			}
		}
		total += groupCount
	}

	return total
}

func countCommonGroupTotals(groups [][]string) int {
	total := 0

	for _, group := range groups {
		firstSet := getFirstPersonMap(group[0])

		for _, person := range group[1:] {
			currentSet := make(map[rune]bool)

			for _, letter := range person {
				currentSet[letter] = true

				if _, inFirst := firstSet[letter]; !inFirst {
					delete(firstSet, letter)
					delete(currentSet, letter)
				}
			}

			for letter, _ := range firstSet {
				if _, inSecond := currentSet[letter]; !inSecond {
					delete(firstSet, letter)
					delete(currentSet, letter)
				}
			}
		}

		total += len(firstSet)
	}

	return total
}

func getFirstPersonMap(person string) map[rune]bool {
	set := make(map[rune]bool)

	for _, letter := range person {
		if _, seen := set[letter]; !seen {
			set[letter] = true
		}
	}

	return set
}

func main() {
	input := getInput()
	fmt.Println("Total questions answered yes:\t", countGroupTotals(input))
	fmt.Println("Common questions answered yes:\t", countCommonGroupTotals(input))
}
