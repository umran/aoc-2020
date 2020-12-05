package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func solution1(passports []map[string]string) int {
	valid := 0

	for _, passport := range passports {
		if hasRequiredFields(passport) {
			valid++
		}
	}

	return valid
}

func solution2(passports []map[string]string) int {
	valid := 0

outer:
	for _, passport := range passports {
		if !hasRequiredFields(passport) {
			continue outer
		}

		for k, v := range passport {
			if !isValidField(k, v) {
				continue outer
			}
		}

		valid++
	}

	return valid
}

func hasRequiredFields(passport map[string]string) bool {
	for _, rf := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := passport[rf]; !ok {
			return false
		}
	}

	return true
}

func isValidField(key, value string) bool {
	valueChars := []rune(value)

	switch key {
	case "byr", "iyr", "eyr":
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if key == "byr" && (year < 1920 || year > 2002) {
			return false
		}
		if key == "iyr" && (year < 2010 || year > 2020) {
			return false
		}
		if key == "eyr" && (year < 2020 || year > 2030) {
			return false
		}
	case "hgt":
		unitStartIndex := (len(valueChars) - 2)
		if unitStartIndex-1 < 0 {
			return false
		}
		magnitude, err := strconv.Atoi(string(valueChars[:unitStartIndex]))
		if err != nil {
			return false
		}

		units := value[unitStartIndex:]
		switch units {
		case "cm":
			if magnitude < 150 || magnitude > 193 {
				return false
			}
		case "in":
			if magnitude < 59 || magnitude > 76 {
				return false
			}
		default:
			return false
		}
	case "hcl":
		if valueChars[0] != '#' {
			return false
		}

		information := valueChars[1:]
		if len(information) != 6 {
			return false
		}

		if match, err := regexp.Match("[0-9a-f]{6}", []byte(string(information))); err != nil || !match {
			return false
		}
	case "ecl":
		if len(valueChars) != 3 {
			return false
		}

		if match, err := regexp.Match("amb|blu|brn|gry|grn|hzl|oth", []byte(value)); err != nil || !match {
			return false
		}
	case "pid":
		if len(valueChars) != 9 {
			return false
		}

		if match, err := regexp.Match("[0-9]{9}", []byte(value)); err != nil || !match {
			return false
		}
	case "cid":
		// noop
	default:
		// unknown whether additional fields are invalid, leaving it as restrictive as possible for now
		return false
	}

	return true
}

func parsePassport(data string) map[string]string {
	var fields []string
	fieldSets := strings.Split(data, "\n")
	for _, fieldSet := range fieldSets {
		for _, f := range strings.Split(fieldSet, " ") {
			fields = append(fields, f)
		}
	}

	passport := make(map[string]string)
	for _, field := range fields {
		keyValue := strings.Split(field, ":")
		passport[keyValue[0]] = keyValue[1]
	}

	return passport
}

func parseInput(filename string) []map[string]string {
	bytes, _ := ioutil.ReadFile(filename)

	text := string(bytes)
	rawPassports := strings.Split(text, "\n\n")
	passports := make([]map[string]string, len(rawPassports))
	for _, rawPassport := range rawPassports {
		passports = append(passports, parsePassport(rawPassport))
	}

	return passports
}

func main() {
	passports := parseInput("input.txt")
	fmt.Println(solution1(passports))
	fmt.Println(solution2(passports))
}
