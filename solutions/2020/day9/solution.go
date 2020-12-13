package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput() []int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines
}

func findFirstInvalidNum(nums []int, preambleLen int) int {
	prevStart := 0
	prevEnd := preambleLen

	for _, num := range nums[preambleLen:] {
		if !isNumberValid(num, nums[prevStart:prevEnd]) {
			return num
		}
		prevStart++
		prevEnd++
	}
	return 0
}

func isNumberValid(num int, prevNums []int) bool {
	for i, prev1 := range prevNums {
		for _, prev2 := range prevNums[i:] {
			if (prev1 + prev2) == num {
				return true
			}
		}
	}
	return false
}

func findSumOfRangeBounds(input []int, num int) int {
	rangeStartIndex := 0
	rangeEndIndex := 1

	rangeMin := input[rangeStartIndex]
	rangeMax := input[rangeStartIndex]
	sum := input[rangeStartIndex]

	for rangeEndIndex < len(input) {
		sum += input[rangeEndIndex]
		rangeMin = min(rangeMin, input[rangeEndIndex])
		rangeMax = max(rangeMax, input[rangeEndIndex])

		if sum == num {
			return rangeMin + rangeMax
		} else if sum < num {
			rangeEndIndex++
		} else {
			rangeStartIndex++
			rangeEndIndex = rangeStartIndex + 1
			sum = input[rangeStartIndex]
			rangeMin = input[rangeStartIndex]
			rangeMax = input[rangeStartIndex]
		}
	}
	return 0
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	input := getInput()

	invalidNum := findFirstInvalidNum(input, 25)
	sumRangeBounds := findSumOfRangeBounds(input, invalidNum)
	fmt.Println("First invalid num:\t", invalidNum)
	fmt.Println("Sum of range bounds:\t", sumRangeBounds)
}
