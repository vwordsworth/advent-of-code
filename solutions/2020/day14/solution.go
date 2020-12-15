package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type program struct {
	memory map[int64]int64
	instr []*instruction
}

type mask struct {
	toOne, toZero, toFloat int64
	floatLocs []int64
}

type instruction struct {
	bitmask *mask
	address, value int64
}

func newProgram(instr []string) *program {
	p := program{}
	p.memory = make(map[int64]int64)
	var currMask *mask
	for _, ins := range instr {
		if strings.Contains(ins, "mask") {
			currMask = newMask(ins)
		} else {
			p.instr = append(p.instr, newInstruction(ins, currMask))
		}
	}
	return &p
}

func (p *program) sumMemory() int64 {
	sum := int64(0)
	for _, v := range p.memory {
		sum += v
	}
	return sum
}

func (p *program) resetMemory() {
	p.memory = make(map[int64]int64)
}

func newInstruction(line string, bitMask *mask) *instruction {
	i := instruction{}
	instrRegEx := regexp.MustCompile(`mem\[(?P<addr>\d+)\] = (?P<value>\d+)`)
	matches := instrRegEx.FindAllStringSubmatch(line, 1)[0]
	addr, _ := strconv.Atoi(matches[1])
	i.address = int64(addr)
	val, _ := strconv.Atoi(matches[2])
	i.value = int64(val)
	i.bitmask = bitMask
	return &i
}

func newMask(line string) *mask {
	m := mask{}
	maskRegEx := regexp.MustCompile(`mask = (?P<value>[\dX]+)`)
	pattern := maskRegEx.FindAllStringSubmatch(line, 1)[0][1]
	toOne := ""
	toZero := ""
	toFloat := ""
	for i, dig := range pattern {
		switch string(dig) {
		case "1":
			toOne += "1"
			toZero += "1"
			toFloat += "1"
		case "0":
			toOne += "0"
			toZero += "0"
			toFloat += "1"
		case "X":
			toOne += "0"
			toZero += "1"
			toFloat += "0"
			m.floatLocs = append(m.floatLocs, (1 << (len(pattern)-1-i)))
		}
	}
	m.toOne, _ = strconv.ParseInt(toOne, 2, 64)
	m.toZero, _ = strconv.ParseInt(toZero, 2, 64)
	m.toFloat, _ = strconv.ParseInt(toFloat, 2, 64)
	return &m
}

func getPowerset(arr []int64, left, right int, sum int64) []int64{
 var sums []int64
 if left > right {
	 sums = append(sums, sum)
	 return sums
 }
 for _, elem := range getPowerset(arr, left + 1, right, sum + arr[left]) {
	 sums = append(sums, elem)
 }
 for _, elem := range getPowerset(arr, left + 1, right, sum) {
	 sums = append(sums, elem)
 }
 return sums
}

func getInput() *program {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return newProgram(lines)
}

func process(input *program) int64 {
	for _, instr := range input.instr {
		val := (instr.value | instr.bitmask.toOne) & instr.bitmask.toZero
		input.memory[instr.address] = val
	}
	return input.sumMemory()
}

func floatingProcess(input *program) int64 {
	for _, instr := range input.instr {
		addr := (instr.address | instr.bitmask.toOne) & instr.bitmask.toFloat
		n := len(instr.bitmask.floatLocs)
		for _, i := range getPowerset(instr.bitmask.floatLocs, 0, n - 1, int64(0)) {
			input.memory[i+addr] = instr.value
		}
	}
	return input.sumMemory()
}

func main() {
	input := getInput()
	fmt.Println("Sum in memory\t", process(input))
	input.resetMemory()
	fmt.Println("Floating process\t", floatingProcess(input))
}
