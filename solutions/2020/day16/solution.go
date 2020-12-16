package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

type paper struct {
	rules []*rule
	myTickets, otherTickets, validTickets []*ticket
}

type ticket struct {
	vals []int
}

type rule struct {
	name string
	ruleCount int
	bounds []*bound
}

type bound struct {
	low, high int
}

func newPaper() *paper {
	p := paper{}
	p.rules = []*rule{}
	p.myTickets = []*ticket{}
	p.otherTickets = []*ticket{}
	return &p
}

func (p *paper) addRule(r *rule) {
	p.rules = append(p.rules, r)
}

func (p *paper) addMyTicket(t *ticket) {
	p.myTickets = append(p.myTickets, t)
	p.validTickets = append(p.validTickets, t)
}

func (p *paper) addOtherTicket(t *ticket) {
	p.otherTickets = append(p.otherTickets, t)
}

func (p *paper) addValidTicket(t *ticket) {
	p.validTickets = append(p.validTickets, t)
}

func (p *paper) isValueValid(val int) bool {
	for _, r := range p.rules {
		if r.isInRange(val) {
			return true
		}
	}
	return false
}

func (p *paper) printPaper() {
	for _, r := range p.rules {
		fmt.Println(r.name)
		for _, b := range r.bounds {
			fmt.Println(b.low, b.high)
		}
	}
	for _, r := range p.myTickets {
		fmt.Println(r.vals)
	}
	for _, r := range p.otherTickets {
		fmt.Println(r.vals)
	}
}

func newTicket(vals []int) *ticket {
	t := ticket{}
	t.vals = vals
	return &t
}

func newRule(c int, name string) *rule {
	r := rule{}
	r.name = name
	r.bounds = []*bound{}
	r.ruleCount = c
	return &r
}

func (r *rule) addBounds(b *bound) {
	r.bounds = append(r.bounds, b)
}

func (r *rule) isInRange(val int) bool {
	for _, b := range r.bounds {
		if (val >= b.low) && (val <= b.high) {
			return true
		}
	}
	return false
}

func newBound(low, high int) *bound {
	b := bound{}
	b.low = low
	b.high = high
	return &b
}

func getInput() *paper {
	file, _ := os.Open("input.txt")
	defer file.Close()

	isMyTickets := false
	isOtherTickets := false
	p := newPaper()
	ruleCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ruleRegEx := regexp.MustCompile(`(?P<name>\w+): (?P<low1>\d+)-(?P<high1>\d+) or (?P<low2>\d+)-(?P<high2>\d+)`)
		if ruleRegEx.MatchString(line) {
			matches := ruleRegEx.FindAllStringSubmatch(line, 1)[0]
			low1, _ := strconv.Atoi(matches[2])
			high1, _ := strconv.Atoi(matches[3])
			low2, _ := strconv.Atoi(matches[4])
			high2, _ := strconv.Atoi(matches[5])

			r := newRule(ruleCount, matches[1])
			r.addBounds(newBound(low1, high1))
			r.addBounds(newBound(low2, high2))
			p.addRule(r)
			ruleCount++
		} else if line == "your ticket:" {
			isMyTickets = true
		} else if isMyTickets {
			var vals []int
			for _, n := range strings.Split(scanner.Text(), ",") {
					intN, _ := strconv.Atoi(n)
					vals = append(vals, intN)
			}
			p.addMyTicket(newTicket(vals))
			isMyTickets = false
		} else if line == "nearby tickets:" {
			isOtherTickets = true
		} else if isOtherTickets {
			var vals []int
			for _, n := range strings.Split(line, ",") {
					intN, _ := strconv.Atoi(n)
					vals = append(vals, intN)
			}
			p.addOtherTicket(newTicket(vals))
		}
	}
	return p
}

func countInvalidOtherTickets(p *paper) int {
	invalidSum := 0
	for _, t := range p.otherTickets {
		isTicketValid := true
		for _, val := range t.vals {
			if !p.isValueValid(val) {
				isTicketValid = false
				invalidSum += val
			}
		}
		if isTicketValid {
			p.addValidTicket(t)
		}
	}
	return invalidSum
}

func calculateDepartureProduct(p *paper) int {
	// map position in ticket value to possible rules
	possibilities := make(map[*rule]map[int]bool)
	for _, rule := range p.rules {
		possibilities[rule] = make(map[int]bool)
		for pos := 0; pos < len(p.validTickets[0].vals); pos++ {
			isPositionEligible := true
			for _, t := range p.validTickets {
				if !rule.isInRange(t.vals[pos]) {
					isPositionEligible = false
				}
			}
			if isPositionEligible {
				possibilities[rule][pos] = true
			}
		}
	}

	// iterate until all rules are assigned to one position
	determined := make(map[*rule]int)
	for len(determined) < len(p.validTickets[0].vals){
		for class, possMap := range possibilities {
			if _, det := determined[class]; !det && len(possMap) == 1 {
				var key int
				for k,_ := range possMap {
					key = k
				}
				determined[class] = key
				for c, p := range possibilities {
					if _, contains := p[key]; contains && c != class {
						delete(p, key)
					}
				}
			}
		}
	}

	prod := 1
	for k, v := range determined {
		if k.ruleCount < 6 {
			prod *= p.myTickets[0].vals[v]
		}
	}


	return prod
}

func main() {
	input := getInput()
	fmt.Println("Invalid ticket count:\t", countInvalidOtherTickets(input))
	fmt.Println("Departure product:\t", calculateDepartureProduct(input))
}
