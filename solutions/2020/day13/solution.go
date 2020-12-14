package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type schedule struct {
	arrivalTime int
	inServiceBuses []int
	allBuses []int
}

func newSchedule(time string, buses string) *schedule {
	s := schedule{}
	s.arrivalTime, _ = strconv.Atoi(time)
	for _, bus := range strings.Split(buses, ",") {
		if bus != "x" {
			id, _ := strconv.Atoi(bus)
			s.inServiceBuses = append(s.inServiceBuses, id)
			s.allBuses = append(s.allBuses, id)
		} else {
			s.allBuses = append(s.allBuses, -1)
		}
	}
	return &s
}

func (sched *schedule) busIdMultipliedByWaitTime() int {
	minBus := -1
	minWaitTime := math.MaxInt32
	for _, bus := range sched.inServiceBuses {
		nextBus := int(math.Ceil(float64(sched.arrivalTime)/float64(bus)))*bus
		if nextBus - sched.arrivalTime < minWaitTime {
			minWaitTime = nextBus - sched.arrivalTime
			minBus = bus
		}
	}
	return minBus * minWaitTime
}

func (sched *schedule) getFirstTimeT() *big.Int {
	var nums, remainder []*big.Int
	product := big.NewInt(1)
	for i, bus := range sched.allBuses {
		if bus != -1 {
			product.Mul(product, big.NewInt(int64(bus)))
			nums = append(nums, big.NewInt(int64(bus)))
			remainder = append(remainder, big.NewInt(int64(bus-i)))
		}
	}

	var result, pp, s, z big.Int
	for i, num := range nums {
		pp.Div(product, num)
	  z.GCD(nil, &s, num, &pp)
  	result.Add(&result, s.Mul(remainder[i], s.Mul(&s, &pp)))
 	}
	return result.Mod(&result, product)
}

func getInput() *schedule {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return newSchedule(lines[0], lines[1])
}

func main() {
	schedule := getInput()
	fmt.Println(schedule.busIdMultipliedByWaitTime())
	fmt.Println(schedule.getFirstTimeT())
}
