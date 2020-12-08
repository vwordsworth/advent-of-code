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
	i := instruction{}
	i.opCode = parts[0]
	i.offset, _ = strconv.Atoi(parts[1])
	return &i
}

func (i *instruction) modifyOp() {
	if i.opCode == "nop" {
		i.modifiedOpCode = "jmp"
	} else if i.opCode == "jmp" {
		i.modifiedOpCode = "nop"
	}
}

func (i *instruction) resetOp() {
	i.modifiedOpCode = ""
}

func newProgram(instructions []*instruction) *program{
	p := program{}
	p.accum = 0
	p.instrPtr = 0
	p.instructions = instructions
	p.seenInstr = make(map[int]bool)
	return &p
}

func (p *program) runProgramUntilRepeat() {
	for !p.seenInstr[p.instrPtr] {
		p.runInstruction()
	}
}

func (p *program) doesProgramTerminate() bool {
	for p.instrPtr != len(p.instructions) && !p.seenInstr[p.instrPtr] {
		p.runInstruction()
	}
	return p.instrPtr == len(p.instructions)
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
	program.runProgramUntilRepeat()
	return program.accum
}

func getTerminatingAccum(input []*instruction) int {
	modIndex := 0
	for modIndex < len(input) {
		input[modIndex].modifyOp()

		program := newProgram(input)
		if program.doesProgramTerminate() {
			return program.accum
		}

		input[modIndex].resetOp()
		modIndex++
	}
	return math.MaxInt32
}

func main() {
	input := getInput()
	fmt.Println("Accumulator at repeat:\t", getAccumAtRepeat(input))
	fmt.Println("Accumulator at end:\t", getTerminatingAccum(input))
}
