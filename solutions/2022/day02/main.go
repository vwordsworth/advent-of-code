package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"day2/move"
)

func getInput(path string, constructor func([]string) move.Action) []move.Action {
	file, _ := os.Open(path)
	defer file.Close()

	var actions []move.Action
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr := constructor(strings.Split(scanner.Text(), " "))
		actions = append(actions, curr)
	}
	return actions
}

func getScore(actions []move.Action) int {
	score := 0
	for _, act := range actions {
		score += act.CalculateScore()
	}
	return score
}

func main() {
	moves := getInput("data/data.txt", move.NewMove)
	fmt.Printf("1. Total Score (Assuming Move Col): %d\n", getScore(moves))
	outcomes := getInput("data/data.txt", move.NewOutcome)
	fmt.Printf("2. Total Score (Assuming Outcome Col): %d\n", getScore(outcomes))
}
