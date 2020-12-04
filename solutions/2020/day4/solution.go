package main

import (
	"bytes"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type validation struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid *regexp.Regexp
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

func newFieldValidation() *validation {
	v := validation{}
	v.byr, _ = regexp.Compile("\\s*(byr:)[#\\w]+\\s*")
	v.iyr, _ = regexp.Compile("\\s*(iyr:)[#\\w]+\\s*")
	v.eyr, _ = regexp.Compile("\\s*(eyr:)[#\\w]+\\s*")
	v.hgt, _ = regexp.Compile("\\s*(hgt:)[#\\w]+\\s*")
	v.hcl, _ = regexp.Compile("\\s*(hcl:)[#\\w]+\\s*")
	v.ecl, _ = regexp.Compile("\\s*(ecl:)[#\\w]+\\s*")
	v.pid, _ = regexp.Compile("\\s*(pid:)[#\\w]+\\s*")
	return &v
}

func newFieldAndContentValidation() *validation {
	v := validation{}
	v.byr, _ = regexp.Compile("\\s*(byr:)((200)[0-2]|(19)[2-9][0-9])\\s*")
	v.iyr, _ = regexp.Compile("\\s*(iyr:)((201)[0-9]|2020)\\s*")
	v.eyr, _ = regexp.Compile("\\s*(eyr:)((202)[0-9]|2030)\\s*")
	v.hgt, _ = regexp.Compile("\\s*(hgt:)((1[5-8][0-9]|19[0-3])(cm)|(59|6[0-9]|7[0-6])(in))\\s*")
	v.hcl, _ = regexp.Compile("\\s*(hcl:)#[0-9a-f]{6}\\s*")
	v.ecl, _ = regexp.Compile("\\s*(ecl:)(amb|blu|brn|gry|grn|hzl|oth)\\s*")
	v.pid, _ = regexp.Compile("\\s*(pid:)\\d{9}($|\\s+)")
	return &v
}

func countValidPassports(passports []string, rules *validation) int {
	count := 0
	for _, pass := range passports {
		if isPassportValid(pass, rules) {
			count++
		}
	}
	return count
}

func isPassportValid(pass string, rules *validation) bool {
	return rules.byr.MatchString(pass) && rules.iyr.MatchString(pass) && rules.eyr.MatchString(pass) && rules.hgt.MatchString(pass) && rules.hcl.MatchString(pass) && rules.ecl.MatchString(pass) && rules.pid.MatchString(pass)
}

func main() {
	input := getInput()

	fieldRules := newFieldValidation()
	count := countValidPassports(input, fieldRules)
	fmt.Println("Valid fields:\t", count)

	contentRules := newFieldAndContentValidation()
	countWithRules := countValidPassports(input, contentRules)
	fmt.Println("Valid fields & contents:\t", countWithRules)
}
