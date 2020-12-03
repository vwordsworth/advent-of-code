package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func getInput() [][]string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return lines
}

func countTreesOnSlope(input [][]string, dx int, dy int) int {
	numRows := len(input)
	widthPattern := len(input[0])
	count :=0

	for y := 0; y < numRows; y += dy {
		x := ((dx*y)/dy) % widthPattern
		if input[y][x] == "#" {
			count++
		}
	}

	return count
}


func main() {
	start := time.Now()
	input := getInput()

	numTrees := countTreesOnSlope(input, 3, 1)
	fmt.Println("Tree count for slope right 3, down 1:\t", numTrees)

	slopes := [4][2]int{[2]int{1, 1}, [2]int{5, 1}, [2]int{7, 1}, [2]int{1, 2}}
	for _, slope := range slopes {
		numTrees *= countTreesOnSlope(input, slope[0], slope[1])
	}

	fmt.Println("Tree count product for all slopes:\t", numTrees)
	fmt.Println("\nTime elapsed:\t", time.Now().Sub(start))
}
