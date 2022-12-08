package main

import (
	"bufio"
	"fmt"
	"os"

	"day3/bag"
)

func getInput(path string) []bag.Bag {
	file, _ := os.Open(path)
	defer file.Close()

	var bags []bag.Bag
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr := bag.NewBag(scanner.Text())
		bags = append(bags, curr)
	}
	return bags
}

func assignGroups(input []bag.Bag) []bag.Group {
	groups := []bag.Group{}
	for i := 0; i < len(input); i += 3 {
		groups = append(groups, bag.NewGroup(input[i], input[i+1], input[i+2]))
	}
	return groups
}

func sumAllSingleBagRepeatItemPriorities(bags []bag.Bag) int {
	sum := 0
	for _, bag := range bags {
		sum += bag.GetSingleBagRepeatedRunePriority()
	}
	return sum
}

func sumAllGroupRepeatItemPriorities(groups []bag.Group) int {
	sum := 0
	for _, group := range groups {
		sum += group.GetGroupRepeatedRunePriority()
	}
	return sum
}

func main() {
	bags := getInput("data/data.txt")
	fmt.Printf("1. Sum of all repeated item priorities: %d\n", sumAllSingleBagRepeatItemPriorities(bags))
	groups := assignGroups(bags)
	fmt.Printf("2. Sum of all group repeated item priorities: %d\n", sumAllGroupRepeatItemPriorities(groups))
}
