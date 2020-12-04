package main

import (
	"bytes"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type entry struct {
	value      string
	letter     byte
	int1, int2 int
}

func getInput() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []string
	var line bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if len(currentLine) == 0 {
			lines = append(lines, line.String())
			line.Reset()
		} else {
			line.WriteString(" ")
			line.WriteString(currentLine)
		}
	}

	if len(line.String()) > 0 {
		lines = append(lines, line.String())
	}

	return lines
}

func processPassports(passports []string) (int, int) {
	count := 0
	countWithRules := 0
	for _, pass := range passports {
		if isPassportValid(pass) {
			count++
			if isPassportValidWithRules(pass) {
				countWithRules++
			}
		}
	}

	return count, countWithRules
}

func isPassportValid(pass string) bool {
	byr, _ := regexp.Compile("\\s*(byr:)[#\\w]+\\s*")
	iyr, _ := regexp.Compile("\\s*(iyr:)[#\\w]+\\s*")
	eyr, _ := regexp.Compile("\\s*(eyr:)[#\\w]+\\s*")
	hgt, _ := regexp.Compile("\\s*(hgt:)[#\\w]+\\s*")
	hcl, _ := regexp.Compile("\\s*(hcl:)[#\\w]+\\s*")
	ecl, _ := regexp.Compile("\\s*(ecl:)[#\\w]+\\s*")
	pid, _ := regexp.Compile("\\s*(pid:)[#\\w]+\\s*")

	return byr.MatchString(pass) && iyr.MatchString(pass) && eyr.MatchString(pass) && hgt.MatchString(pass) && hcl.MatchString(pass) && ecl.MatchString(pass) && pid.MatchString(pass)
}

func isPassportValidWithRules(pass string) bool {
	byr, _ := regexp.Compile("\\s*(byr:)((200)[0-2]|(19)[2-9][0-9])\\s*")
	iyr, _ := regexp.Compile("\\s*(iyr:)((201)[0-9]|2020)\\s*")
	eyr, _ := regexp.Compile("\\s*(eyr:)((202)[0-9]|2030)\\s*")
	hgt, _ := regexp.Compile("\\s*(hgt:)((1[5-8][0-9]|19[0-3])(cm)|(59|6[0-9]|7[0-6])(in))\\s*")
	hcl, _ := regexp.Compile("\\s*(hcl:)#[0-9a-f]{6}\\s*")
	ecl, _ := regexp.Compile("\\s*(ecl:)(amb|blu|brn|gry|grn|hzl|oth)\\s*")
	pid, _ := regexp.Compile("\\s*(pid:)\\d{9}($|\\s+)")

	return byr.MatchString(pass) && iyr.MatchString(pass) && eyr.MatchString(pass) && hgt.MatchString(pass) && hcl.MatchString(pass) && ecl.MatchString(pass) && pid.MatchString(pass)
}

func main() {
	input := getInput()

	count, countWithRules := processPassports(input)
	fmt.Println("Passport count:\t", count)
	fmt.Println("Passport count with rules:\t", countWithRules)
}
