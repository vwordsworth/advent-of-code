package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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

func main() {
	start := time.Now()
	numbers := getInput()

	twoProduct := simpleTwoProduct(numbers)
	fmt.Println("Two product:\t", twoProduct)

	threeProduct := simpleThreeProduct(numbers)
	fmt.Println("Three product:\t", threeProduct)

	elapsed := time.Now().Sub(start)
	fmt.Println("\nTime elapsed:\t", elapsed)
}
