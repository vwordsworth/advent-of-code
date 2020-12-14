package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	EAST = 0
	NORTH = 90
	WEST = 180
	SOUTH = 270
)

type instruction struct {
	value int
	action string
}

type ship struct {
	NS, EW int
	direction int
}

type waypoint struct {
	NS, EW int
}

func newInstruction(line string) *instruction {
	i := instruction{}
	instrRegEx := regexp.MustCompile(`(?P<action>[NSEWLRF])(?P<value>\d+)`)
	matches := instrRegEx.FindAllStringSubmatch(line, 1)[0]
	i.action = matches[1]
	i.value, _ = strconv.Atoi(matches[2])
	return &i
}

func newShip(x, y int) *ship {
	s := ship{}
	s.EW = x
	s.NS = y
	s.direction = EAST
	return &s
}

func newWaypoint(x, y int) *waypoint {
	w := waypoint{}
	w.EW = x
	w.NS = y
	return &w
}

func (s *ship) north(n int) {
	s.NS += n
}

func (s *ship) east(n int) {
	s.EW += n
}

func (s *ship) south(n int) {
	s.NS -= n
}

func (s *ship) west(n int) {
	s.EW -= n
}

func (s *ship) forward(n int) {
	switch s.direction {
	case NORTH:
		s.north(n)
	case EAST:
		s.east(n)
	case SOUTH:
		s.south(n)
	case WEST:
		s.west(n)
	}
}

func (s *ship) rotateLeft(angle int) {
	s.direction = (s.direction + angle) % 360
}

func (s *ship) rotateRight(angle int) {
	s.direction = (s.direction - angle) % 360
	if s.direction < 0 {
		s.direction += 360
	}
}

func (s *ship) shipToWaypointNTimes(w *waypoint, n int) {
	s.EW += n*w.EW
	s.NS += n*w.NS
}

func (s *ship) manhattanDistance() int {
	return abs(s.EW) + abs(s.NS)
}

func (w *waypoint) north(n int) {
	w.NS += n
}

func (w *waypoint) east(n int) {
	w.EW += n
}

func (w *waypoint) south(n int) {
	w.NS -= n
}

func (w *waypoint) west(n int) {
	w.EW -= n
}

func (w *waypoint) rotate(angle int) {
	switch angle % 360 {
	case NORTH:
		w.EW, w.NS = -1*w.NS, w.EW
	case WEST:
		w.EW, w.NS = -1*w.EW, -1*w.NS
	case SOUTH:
		w.EW, w.NS = w.NS, -1*w.EW
	}
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

func processInstructions(input []*instruction) int {
	ship := newShip(0, 0)

	for _, instr := range input {
		switch instr.action {
		case "N":
			ship.north(instr.value)
		case "E":
			ship.east(instr.value)
		case "S":
			ship.south(instr.value)
		case "W":
			ship.west(instr.value)
		case "F":
			ship.forward(instr.value)
		case "L":
			ship.rotateLeft(instr.value)
		case "R":
			ship.rotateRight(instr.value)
		}
	}
	return ship.manhattanDistance()
}

func processInstructionsWithWaypoint(input []*instruction) int {
	ship := newShip(0, 0)
	waypt :=  newWaypoint(10, 1)

	for _, instr := range input {
		switch instr.action {
		case "N":
			waypt.north(instr.value)
		case "E":
			waypt.east(instr.value)
		case "S":
			waypt.south(instr.value)
		case "W":
			waypt.west(instr.value)
		case "F":
			ship.shipToWaypointNTimes(waypt, instr.value)
		case "L":
			waypt.rotate(instr.value)
		case "R":
			waypt.rotate(360 - instr.value)
		}
	}
	return ship.manhattanDistance()
}

func manhattanDistance(x, y int) int {
	return abs(x) + abs(y)
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func main() {
	input := getInput()
	fmt.Println("Manhattan distance:\t", processInstructions(input))
	fmt.Println("Manhattan distance:\t", processInstructionsWithWaypoint(input))
}
