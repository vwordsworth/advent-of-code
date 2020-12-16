package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() []int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, num := range strings.Split(scanner.Text(), ",") {
			intNum, _ := strconv.Atoi(num)
			nums = append(nums, intNum)
		}

	}
	return nums
}

func findNthNum(startNums []int, n int) int {
	seen := make(map[int]int)
	for i, startN := range startNums {
		seen[startN] = i
	}

	last := 0
	for round := len(startNums); round < n-1; round++ {
		var diff int
		if lastRound, wasSeen := seen[last]; wasSeen {
			diff = round - lastRound
		} else {
			diff = 0
		}
		seen[last] = round
		last = diff
	}
	return last
}

func main() {
	input := getInput()
	fmt.Println("2020th digit:\t", findNthNum(input, 2020))
	fmt.Println("30000000th digit:\t", findNthNum(input, 30000000))
}
