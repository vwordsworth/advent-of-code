package main

import (
	"bufio"
	"fmt"
	"math"
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

func parseLine(line string) (string, int) {
	instr := strings.Fields(line)
	op := instr[0]
	arg := instr[1]

	var offset int
	if string(arg[0]) == "+" {
		offset, _ = strconv.Atoi(strings.Split(arg, "+")[1])
	} else {
		offset, _ = strconv.Atoi(strings.Split(arg, "-")[1])
		offset *= -1
	}

	return op, offset
}

func q1(input []string) int {
	accum := 0
	instrPtr := 0

	var instructionLines []bool
	for i:=0; i<len(input); i++ {
		instructionLines = append(instructionLines, false)
	}

	for !instructionLines[instrPtr] {
		op, offset := parseLine(input[instrPtr])

		instructionLines[instrPtr] = true

		if op == "nop" {
			instrPtr++
		} else if op == "jmp" {
			instrPtr += offset
		} else if op == "acc" {
			accum += offset
			instrPtr++
		}
	}

	return accum
}

func q2(input []string) int {
	modIndex := 0

	for modIndex < len(input) {
		var temp = make([]string, len(input))
		copy(temp, input)

		if strings.Contains(temp[modIndex], "nop") {
			temp[modIndex] = strings.Replace(temp[modIndex],"nop","jmp",1)
		} else if strings.Contains(temp[modIndex], "jmp") {
			temp[modIndex] = strings.Replace(temp[modIndex],"jmp","nop",1)
		}

		modIndex++

		if accum := runCheckTerminates(temp); accum != math.MaxInt32 {
			return accum
		}
	}

	return math.MaxInt32
}

func runCheckTerminates(input []string) int {
	desiredPtr := len(input)

	accum := 0
	instrPtr := 0

	var instructionLines []bool
	for i:=0; i<len(input); i++ {
		instructionLines = append(instructionLines, false)
	}

	for instrPtr != desiredPtr && !instructionLines[instrPtr] {
		op, offset := parseLine(input[instrPtr])

		instructionLines[instrPtr] = true

		if op == "nop" {
			instrPtr++
		} else if op == "jmp" {
			instrPtr += offset
		} else if op == "acc" {
			accum += offset
			instrPtr++
		}

	}

	if instrPtr == desiredPtr {
		return accum
	}

	return math.MaxInt32
}

func main() {
	input := getInput()

	fmt.Println(q1(input))

	fmt.Println(q2(input))
}
