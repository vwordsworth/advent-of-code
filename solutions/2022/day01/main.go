package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"

	maxheap "day1/heap"
)

func getInput(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}
	return lines
}

func getElfSumsHeap(input []int) heap.Interface {
	output := &maxheap.MaxHeap{}
	heap.Init(output)
	runningSum := 0
	for _, val := range input {
		runningSum += val
		if val == 0 {
			heap.Push(output, runningSum)
			runningSum = 0
		}
	}
	return output
}

func getTopElves(input heap.Interface, count int) []int {
	result := []int{}
	for i := 0; i < count; i++ {
		result = append(result, heap.Pop(input).(int))
	}
	return result
}

func main() {
	p1 := getInput("data/data.txt")
	heap := getElfSumsHeap(p1)
	tops := getTopElves(heap, 3)
	fmt.Printf("1. Max elf calories: %d\n", tops[0])
	fmt.Printf("2. Sum of top three elf calories: %d\n", tops[0]+tops[1]+tops[2])
}
