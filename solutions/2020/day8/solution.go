package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	opCode, modifiedOpCode string
	offset int
}

type program struct {
	instructions []*instruction
	seenInstr map[int]bool
	accum, instrPtr int
}

func getInput() []*instruction {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []*instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, newInstruction(scanner.Text()))
	}
	return lines
}

func newInstruction(line string) *instruction {
	parts := strings.Fields(line)
	opCode := parts[0]
	arg := parts[1]

	var offset int
	if string(arg[0]) == "+" {
		offset, _ = strconv.Atoi(strings.Split(arg, "+")[1])
	} else {
		offset, _ = strconv.Atoi(strings.Split(arg, "-")[1])
		offset *= -1
	}

	i := instruction{}
	i.opCode = opCode
	i.offset = offset
	return &i
}

func newProgram(instructions []*instruction) *program{
	p := program{}
	p.accum = 0
	p.instrPtr = 0
	p.instructions = instructions
	p.seenInstr = make(map[int]bool)
	return &p
}

func (p *program) runInstruction() {
	instr := p.instructions[p.instrPtr]
	opCode := instr.opCode
	if modifiedOp := instr.modifiedOpCode; len(modifiedOp) > 0 {
		opCode = modifiedOp
	}

	p.seenInstr[p.instrPtr] = true

	if opCode == "nop" {
		p.instrPtr++
	} else if opCode == "jmp" {
		p.instrPtr += instr.offset
	} else if opCode == "acc" {
		p.accum += instr.offset
		p.instrPtr++
	}
}

func getAccumAtRepeat(input []*instruction) int {
	program := newProgram(input)
	for !program.seenInstr[program.instrPtr] {
		program.runInstruction()
	}
	return program.accum
}

func getTerminatingAccum(input []*instruction) int {
	modIndex := 0

	for modIndex < len(input) {
		if input[modIndex].opCode == "nop" {
			input[modIndex].modifiedOpCode = "jmp"
		} else if input[modIndex].opCode == "jmp" {
			input[modIndex].modifiedOpCode = "nop"
		}

		if accum := getAccumIfTerminates(input); accum != math.MaxInt32 {
			return accum
		}

		input[modIndex].modifiedOpCode = ""
		modIndex++
	}

	return math.MaxInt32
}

func getAccumIfTerminates(input []*instruction) int {
	desiredPtr := len(input)
	program := newProgram(input)

	for program.instrPtr != desiredPtr && !program.seenInstr[program.instrPtr] {
		program.runInstruction()
	}

	if program.instrPtr == desiredPtr {
		return program.accum
	}

	return math.MaxInt32
}

func main() {
	input := getInput()
	fmt.Println("Accumulator at repeat:\t", getAccumAtRepeat(input))
	fmt.Println("Accumulator at end:\t", getTerminatingAccum(input))
}
