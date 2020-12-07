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

func main() {
	input := getInput()

	// Part 1
	max := 0

	//Part 2
	var ids [1024]int
	for i := 0; i < 1024; i++ {
		ids[i] = -1
	}

	for _, row := range input {
		// fmt.Println(row)
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")
		row = strings.ReplaceAll(row, "L", "0")
		row = strings.ReplaceAll(row, "R", "1")

		planeRow, _ := strconv.ParseInt(row[:7], 2, 64)
		planeCol, _ := strconv.ParseInt(row[7:], 2, 64)

		val := int(planeRow * 8 + planeCol)
		ids[val] = 1

		if val > max {
			max = val
		}
	}

	// Part 1
	fmt.Println(max)

	// Part 2
	for i:=0; i < len(ids); i++ {
		if ids[i] == -1 && ((i > 1) && (ids[i-1] != -1)) && ((i < 1023) && (ids[i+1] != -1)){
			fmt.Println(i)
		}
	}
}
