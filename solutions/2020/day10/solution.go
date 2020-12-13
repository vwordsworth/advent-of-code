package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func getInput() map[int][]int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	adapters := make(map[int][]int)
	maxJolt := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		adapters[num] = []int{}

		// Add existing 1-3 greater adapaters as possible next adapters for this
		for i := num+1; i <= num+3; i++ {
			if _, seen := adapters[i]; seen {
				adapters[num] = append(adapters[num], i)
			}
		}
		// Add this to possible next adapters for existing 1-3 smaller adapaters
		for i := num-3; i <= num-1; i++ {
			if _, seen := adapters[i]; seen {
				adapters[i] = append(adapters[i], num)
			}
		}

		maxJolt = max(maxJolt, num)
	}

	// Add origin (plug) and its possible next adapters after processing all lines
	adapters[0] = []int{}
	for i := 1; i <= 3; i++ {
		if _, seen := adapters[i]; seen {
			adapters[0] = append(adapters[0], i)
		}
	}

	// Add destination (device) and add to max adapter's next adapter
	adapters[maxJolt] = []int{maxJolt+3}
	adapters[maxJolt+3] = []int{}

	return adapters
}

func processAdapters(adapters map[int][]int, startJolt int) int {
	jump1Jolt := 0
	jump3Jolts := 0

	currentJolt := startJolt

	for true {
		// Return product for problem because must be device (no next adapters)
		if len(adapters[currentJolt]) == 0 {
			return jump1Jolt * jump3Jolts
		}

		// Take smallest next adapter because must use all adapters
		sort.Ints(adapters[currentJolt])
		nextJolt := adapters[currentJolt][0]

		diff := nextJolt - currentJolt
		if diff == 1 {
			jump1Jolt++
		} else if diff == 3 {
			jump3Jolts++
		}

		currentJolt = nextJolt
	}
	return -1
}

func countDistinctArrangements(adapters map[int][]int) int {
	// Initialize at least 1 arrangement with all present (from part 1)
	options := 1

	// Sort adapters to make it easier to determine which can be removed
	keys := sortMapKeys(adapters)
	isInBlock := false
	startIdx := 0

	// Find potentially removable adapters and multiply all valid permutations
	for i, isContender := range getRemovalContenders(keys) {
		if isContender && !isInBlock {
			// Current adapter can be removed, start tracking a removable block
			isInBlock = true
			startIdx = i
		} else if !isContender && isInBlock {
			// Current adapter can't be removed, account for preceeding block
			isInBlock = false
			options *= countPermutations(keys[startIdx:i], keys[startIdx-1], keys[i])
		}
	}
	return options
}

func getRemovalContenders(keys []int) []bool {
	// Set adapter as a contender if the one before and after are still within 3
	contenders := make([]bool, len(keys))
	for i := 1; i < len(keys) - 1; i++ {
		if keys[i+1] - keys[i-1] <= 3 {
			contenders[i] = true
		}
	}
	return contenders
}

func countPermutations(block []int, before int, after int) int {
	// Each adapter is either present or not, 2^length permutations
	permutations := intPow(2, len(block))

	// All permutations are possible
	if len(block) == 1 || (after-before) <= 3 {
		return permutations
	}

	// Remove invalid permutations (boundaries need at least 1 adapter to remain)
	for i := 0; i < len(block); i++ {
		if (block[i] - before) >= 3 {
			return permutations - (len(block) - i)
		}
	}
	return -1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func intPow(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

func sortMapKeys(m map[int][]int) []int {
  keys := make([]int, len(m))
  i := 0
  for k := range m {
      keys[i] = k
      i++
  }
  sort.Ints(keys)
  return keys
}

func main() {
	input := getInput()
	fmt.Println("Product of 1 and 3 jolt jumps:\t", processAdapters(input, 0))
	fmt.Println("All distinct arrangements:\t", countDistinctArrangements(input))
}
