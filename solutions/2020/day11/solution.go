package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type fn func([][]string, [][]string)

func getInput() [][]string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	return lines
}

func simulateSeating(chart [][]string, simulate fn) int {
	prevSeating := chart

	for true {
		currentSeating := makeSeatingCopy(prevSeating)
		simulate(prevSeating, currentSeating)

		if areSeatingChartsEqual(prevSeating, currentSeating) {
			return countOccupiedSeats(currentSeating)
		}

		prevSeating = currentSeating
	}
	return -1
}

func simulateAdjacentSeats(prev, curr [][]string) {
	for i, row := range prev {
		for j, s := range row {
			if isEmptySeat(s) && allAdjacentSeatsAreEmpty(prev, i, j) {
				setOccupiedSeat(curr, i, j)
			}
			if isOccupiedSeat(s) && fourOrMoreAdjacentSeatsAreOccupied(prev, i, j) {
				setEmptySeat(curr, i, j)
			}
		}
	}
}

func simulateNearestSeats(prev, curr [][]string) {
	for i, row := range prev {
		for j, s := range row {
			if isEmptySeat(s) && allNearestSeatsAreEmpty(prev, i, j) {
				setOccupiedSeat(curr, i, j)
			}
			if isOccupiedSeat(s) && fiveOrMoreNearestSeatsAreOccupied(prev, i, j) {
				setEmptySeat(curr, i, j)
			}
		}
	}
}

func countOccupiedSeats(chart [][]string) int {
	count := 0
	for _, row := range chart {
		for _, seat := range row {
			if isOccupiedSeat(seat) {
				count++
			}
		}
	}
	return count
}

func fourOrMoreAdjacentSeatsAreOccupied(chart [][]string, x, y int) bool {
	occupiedCount := 0
	for _, seat := range getAdjacentSeats(chart, x, y) {
		if isOccupiedSeat(seat) {
			occupiedCount++
		}
	}
	return occupiedCount >= 4
}

func fiveOrMoreNearestSeatsAreOccupied(chart [][]string, x, y int) bool {
	occupiedCount := 0
	for _, seat := range getFirstSurroundingSeats(chart, x, y) {
		if isOccupiedSeat(seat) {
			occupiedCount++
		}
	}
	return occupiedCount >= 5
}

func allAdjacentSeatsAreEmpty(chart [][]string, x, y int) bool {
	for _, seat := range getAdjacentSeats(chart, x, y) {
		if isOccupiedSeat(seat) {
			return false
		}
	}
	return true
}

func allNearestSeatsAreEmpty(chart [][]string, x, y int) bool {
	for _, seat := range getFirstSurroundingSeats(chart, x, y) {
		if isOccupiedSeat(seat) {
			return false
		}
	}
	return true
}

func getAdjacentSeats(chart [][]string, x, y int) []string {
	height := len(chart)
	width := len(chart[0])
	var seats []string

	if x > 0 {
		seats = append(seats, chart[x-1][y])
		if y > 0 {
			seats = append(seats, chart[x-1][y-1])
		}
		if y < (width-1) {
			seats = append(seats, chart[x-1][y+1])
		}
	}

	if y > 0 {
		seats = append(seats, chart[x][y-1])
	}
	if y < (width-1) {
		seats = append(seats, chart[x][y+1])
	}

	if x < (height-1) {
		seats = append(seats, chart[x+1][y])
		if y > 0 {
			seats = append(seats, chart[x+1][y-1])
		}
		if y < (width-1) {
			seats = append(seats, chart[x+1][y+1])
		}
	}

	return seats
}

func getFirstSurroundingSeats(chart [][]string, x, y int) []string {
	var seats []string
	if n := northSeat(chart, x, y); n != "" {
		seats = append(seats, n)
	}
	if ne := northEastSeat(chart, x, y); ne != "" {
		seats = append(seats, ne)
	}
	if e := eastSeat(chart, x, y); e != "" {
		seats = append(seats, e)
	}
	if se := southEastSeat(chart, x, y); se != "" {
		seats = append(seats, se)
	}
	if s := southSeat(chart, x, y); s != "" {
		seats = append(seats, s)
	}
	if sw := southWestSeat(chart, x, y); sw != "" {
		seats = append(seats, sw)
	}
	if w := westSeat(chart, x, y); w != "" {
		seats = append(seats, w)
	}
	if nw := northWestSeat(chart, x, y); nw != "" {
		seats = append(seats, nw)
	}
	return seats
}

func northSeat(chart [][]string, x, y int) string {
	x--
	for x >= 0 {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		x--
	}
	return ""
}

func northEastSeat(chart [][]string, x, y int) string {
	x--
	y++
	for x >= 0 && y < len(chart[0]) {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		x--
		y++
	}
	return ""
}

func eastSeat(chart [][]string, x, y int) string {
	y++
	for y < len(chart[0]) {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		y++
	}
	return ""
}

func southEastSeat(chart [][]string, x, y int) string {
	x++
	y++
	for x < len(chart) && y < len(chart[0]) {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		x++
		y++
	}
	return ""
}

func southSeat(chart [][]string, x, y int) string {
	x++
	for x < len(chart) {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		x++
	}
	return ""
}

func southWestSeat(chart [][]string, x, y int) string {
	x++
	y--
	for x < len(chart) && y >= 0 {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		x++
		y--
	}
	return ""
}

func westSeat(chart [][]string, x, y int) string {
	y--
	for y >= 0 {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		y--
	}
	return ""
}

func northWestSeat(chart [][]string, x, y int) string {
	x--
	y--
	for x >= 0 && y >= 0 {
		if !isFloor(chart[x][y]) {
			return chart[x][y]
		}
		x--
		y--
	}
	return ""
}

func setEmptySeat(chart [][]string, x, y int) {
	chart[x][y] = "L"
}

func setOccupiedSeat(chart [][]string, x, y int) {
	chart[x][y] = "#"
}

func isEmptySeat(val string) bool {
	return val == "L"
}

func isOccupiedSeat(val string) bool {
	return val == "#"
}

func isFloor(val string) bool {
	return val == "."
}

func makeSeatingCopy(chart [][]string) [][]string {
	seatingCopy := make([][]string, len(chart))
	for i, row := range chart {
		seatingCopy[i] = make([]string, len(row))
		copy(seatingCopy[i], row)
	}
	return seatingCopy
}

func areSeatingChartsEqual(c1, c2 [][]string) bool {
	for i, row1 := range c1 {
		for j, seat1 := range row1 {
			if seat1 != c2[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	input := getInput()
	fmt.Println("Adj Seats:\t", simulateSeating(input, simulateAdjacentSeats))
	fmt.Println("Near Seats:\t", simulateSeating(input, simulateNearestSeats))
}
