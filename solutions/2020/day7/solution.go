package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type inputLine struct {
	bagName string
	contentNames []string
	quantities []int
}

type bag struct {
	name      string
	contents     []*inventory
}

type inventory struct {
	bagStock *bag
	quantity int
}

func newInputLine(line string) *inputLine {
	i := inputLine{}
	nameRegEx := regexp.MustCompile(`(?P<name>[a-zA-Z\s]+) bags contain `)
	i.bagName = nameRegEx.FindStringSubmatch(line)[1]

	contentRegEx := regexp.MustCompile(`(?P<quantity>\d)\s(?P<content>[a-zA-Z\s]+) bags?`)
	for _, content := range contentRegEx.FindAllStringSubmatch(line, -1) {
		quantity, _ := strconv.Atoi(content[1])
		i.quantities = append(i.quantities, quantity)
		i.contentNames = append(i.contentNames, content[2])
	}
	return &i
}

func newBag(name string, contents []*inventory) *bag {
	b := bag{}
	b.name = name
	b.contents = contents
	return &b
}

func newInventory(bagStock *bag, quantity int) *inventory {
	i := inventory{}
	i.bagStock = bagStock
	i.quantity = quantity
	return &i
}

func getInput() map[string]*bag {
	file, _ := os.Open("input.txt")
	defer file.Close()

	bags := make(map[string]*bag)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := newInputLine(scanner.Text())

		var childInventory []*inventory
		for i, childName := range input.contentNames {
			var b *bag
			if _, exists := bags[childName]; exists {
				b = bags[childName]
			} else {
				b = newBag(childName, nil)
				bags[childName] = b
			}
			inv := newInventory(b, input.quantities[i])
			childInventory = append(childInventory, inv)
		}

		if existingBag, hasBag := bags[input.bagName]; hasBag {
			existingBag.contents = append(existingBag.contents, childInventory...)
		} else {
			bags[input.bagName] = newBag(input.bagName, childInventory)
		}
	}
	return bags
}

func numberContainingColor(arrangement map[string]*bag, color string) int {
	searched := make(map[string]bool)
	count := 0
	for name, b := range arrangement {
		doesContain := containsColor(b, color)
		if _, isPresent := searched[name]; doesContain && !isPresent {
			count++
			searched[name] = true
		}
	}
	return count
}

func containsColor(b *bag, color string) bool {
	if b.contents != nil {
		for _, child := range b.contents {
			if child.bagStock.name == color || containsColor(child.bagStock, color) {
				return true
			}
		}
	}
	return false
}

func countUnderColor(arrangement map[string]*bag, root string) int {
	return countTree(arrangement[root])-1
}

func countTree(b *bag) int {
	count := 1
	for _, child := range b.contents {
		count += countTree(child.bagStock)*child.quantity
	}
	return count
}

func main() {
	input := getInput()
	fmt.Println("Bags containing color:\t", numberContainingColor(input, "shiny gold"))
	fmt.Println("Bags under color:\t", countUnderColor(input, "shiny gold"))
}
