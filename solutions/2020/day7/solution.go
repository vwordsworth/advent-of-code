package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type bag struct {
	name      string
	contents     []*inventory
}

type inventory struct {
	bagStock *bag
	quantity int
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
		bagName, contentNames, quantities := parseBagLine(scanner.Text())
		var contents []*inventory
		for i, content := range contentNames {
			var b *bag
			if _, exists := bags[content]; exists {
				b = bags[content]
			} else {
				b = newBag(content, nil)
				bags[content] = b
			}
			inv := newInventory(b, quantities[i])
			contents = append(contents, inv)
		}

		if _, hasBag := bags[bagName]; hasBag {
			b := bags[bagName]
			b.contents = append(b.contents, contents...)
		} else {
			bags[bagName] = newBag(bagName, contents)
		}
	}
	return bags
}

func parseBagLine(line string) (string, []string, []int) {
	nameRegEx := regexp.MustCompile(`(?P<name>[a-zA-Z\s]+) bags contain `)
	bagName := nameRegEx.FindStringSubmatch(line)[1]

	contentRegEx := regexp.MustCompile(`(?P<quantity>\d)\s(?P<content>[a-zA-Z\s]+) bags?`)
	var contentNames []string
	var quantities []int

	for _, content := range contentRegEx.FindAllStringSubmatch(line, -1) {
		quantity, _ := strconv.Atoi(content[1])
		quantities = append(quantities, quantity)
		contentNames = append(contentNames, content[2])
	}

	return bagName, contentNames, quantities
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
	if b == nil || b.contents == nil {
		return false
	} else {
		for _, child := range b.contents {
			if child.bagStock.name == color || containsColor(child.bagStock, color) {
				return true
			}
		}
		return false
	}
}

func countUnderColor(arrangement map[string]*bag, root string) int {
	return countTree(arrangement[root])-1
}

func countTree(b *bag) int {
	if len(b.contents) == 0 {
		return 1
	} else {
		childCount := 0
		for _, child := range b.contents {
			childCount += countTree(child.bagStock)*child.quantity
		}
		return 1 + childCount
	}
}

func main() {
	input := getInput()
	fmt.Println("Bags containing color:\t", numberContainingColor(input, "shiny gold"))
	fmt.Println("Bags under color:\t", countUnderColor(input, "shiny gold"))
}
