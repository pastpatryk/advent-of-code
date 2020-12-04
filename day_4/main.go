package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Passport holds all document fileds
type Passport map[string]string

// AddField adds field to pasport after parsing it
func (p *Passport) AddField(field string) {
	fieldDef := strings.Split(field, ":")
	fields := map[string]string(*p)
	fields[fieldDef[0]] = fieldDef[1]
}

// IsValid checks if passport has all required fields
func (p *Passport) IsValid() bool {
	if p.checkNumberField("byr", 1920, 2002) &&
		p.checkNumberField("iyr", 2010, 2020) &&
		p.checkNumberField("eyr", 2020, 2030) &&
		p.checkRegexField("hcl", "^#[a-f0-9]{6}$") &&
		p.checkRegexField("ecl", "^(amb|blu|brn|gry|grn|hzl|oth)$") &&
		p.checkRegexField("pid", "^[0-9]{9}$") &&
		p.checkHeight() {
		return true
	}

	return false
}

func (p *Passport) checkNumberField(field string, min, max int) bool {
	fields := map[string]string(*p)
	stringValue, ok := fields[field]
	if !ok {
		return false
	}
	value, err := strconv.Atoi(stringValue)
	if err != nil {
		return false
	}
	return min <= value && value <= max
}

func (p *Passport) checkRegexField(field, regex string) bool {
	fields := map[string]string(*p)
	value, ok := fields[field]
	if !ok {
		return false
	}
	ok, err := regexp.MatchString(regex, value)
	return ok && (err == nil)
}

func (p *Passport) checkHeight() bool {
	fields := map[string]string(*p)
	stringHgt, ok := fields["hgt"]
	if !ok {
		return false
	}
	isCm := strings.HasSuffix(stringHgt, "cm")
	stringHgt = strings.TrimSuffix(strings.TrimSuffix(stringHgt, "cm"), "in")
	hgt, err := strconv.Atoi(stringHgt)
	if err != nil {
		return false
	}
	if isCm {
		return 150 <= hgt && hgt <= 193
	}
	return 59 <= hgt && hgt <= 76
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	passports := readPassports(file)

	validPassports := 0
	for _, passport := range passports {
		if passport.IsValid() {
			validPassports++
		}
	}
	fmt.Printf("Valid passports: %d\n", validPassports)
}

func readPassports(reader io.Reader) []Passport {
	scanner := bufio.NewScanner(reader)

	var passports []Passport
	newPassport := Passport{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, newPassport)
			newPassport = Passport{}
		} else {
			parsePassportFields(line, &newPassport)
		}
	}
	passports = append(passports, newPassport)

	return passports
}

func parsePassportFields(line string, passport *Passport) {
	fields := strings.Split(line, " ")
	for _, field := range fields {
		passport.AddField(field)
	}
}
