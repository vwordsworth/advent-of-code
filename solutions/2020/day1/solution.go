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
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}
	return lines
}

/*
* Initial submission
 */
func simpleTwoProduct(nums []int) int {
	for i, num1 := range nums {
		for _, num2 := range nums[i+1:] {
			if num1+num2 == 2020 {
				return num1 * num2
			}
		}
	}
	return -1
}

func simpleThreeProduct(nums []int) int {
	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			for _, num3 := range nums[j+1:] {
				if num1+num2+num3 == 2020 {
					return num1 * num2 * num3
				}
			}
		}
	}
	return -1
}

/*
* Refactor
 */
func refactorTwoProduct(nums []int, goal int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		need := goal - num
		if _, seen := set[need]; seen {
			return need * num
		}
		set[num] = true
	}
	return -1
}

func refactorThreeProduct(nums []int, goal int) int {
	for i, num := range nums {
		need := goal - num
		if result := refactorTwoProduct(nums[:i], need); result > -1 {
			return result * num
		}
	}
	return -1
}

func main() {
	numbers := getInput()

	twoProduct := refactorTwoProduct(numbers, 2020)
	fmt.Println("Two product:\t", twoProduct)

	threeProduct := refactorThreeProduct(numbers, 2020)
	fmt.Println("Three product:\t", threeProduct)
}
