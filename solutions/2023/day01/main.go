package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTwoDigits(line string) int {
	var first, last *int
	splitLine := strings.Split(line, "")
	end := len(splitLine) - 1

	for i, c := range splitLine {
		c_end := splitLine[end-i]
		if c_int, err := strconv.Atoi(c); err == nil && first == nil {
			first = &c_int
		}

		if c_end_int, err := strconv.Atoi(c_end); err == nil && last == nil {
			last = &c_end_int
		}

		if first != nil && last != nil {
			break
		}
	}
	return (*first * 10) + (*last)
}

func GetListSum(digitList []int) (sum int) {
	for _, n := range digitList {
		sum += n
	}
	return
}

func ReadInputFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReplaceWordsWithDigits(line string) string {
	words := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var result bytes.Buffer
	var sub bytes.Buffer
	for _, c := range strings.Split(line, "") {
		sub.WriteString(c)

		s := sub.String()
		for k, v := range words {
			s = strings.Replace(s, k, v, 1)
		}

		if sub.String() != s {
			result.WriteString(s)
			sub.Reset()
		}
	}
	result.Write(sub.Bytes())
	return result.String()
}

func main() {
	input, err := ReadInputFromFile("./data/input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error reading input: %w", err))
	}

	var result1 []int
	for _, r := range input {
		result1 = append(result1, GetTwoDigits(r))
	}

	answer1 := GetListSum(result1)
	fmt.Printf("Part 1: %d\n", answer1)

	var result2 []int
	for _, r := range input {
		result2 = append(result2, GetTwoDigits(ReplaceWordsWithDigits(r)))
	}

	answer2 := GetListSum(result2)
	fmt.Printf("Part 2: %d\n", answer2)
}
