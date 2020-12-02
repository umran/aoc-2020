package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (min, max int, character rune, pw []rune) {
	ruleRest := strings.Split(line, ":")
	pw = []rune(strings.Split(ruleRest[1], " ")[1])
	rngeChar := strings.Split(ruleRest[0], " ")
	rnge := strings.Split(rngeChar[0], "-")
	min = intFromString(rnge[0])
	max = intFromString(rnge[1])
	character = []rune(rngeChar[1])[0]
	return
}

func solution1(input []string) int {
	valid := 0
	for _, line := range input {
		min, max, character, pw := parseLine(line)
		characterCount := 0
		for _, c := range pw {
			if c == character {
				characterCount++
			}
		}

		if characterCount >= min && characterCount <= max {
			valid++
		}
	}

	return valid
}

func solution2(input []string) int {
	valid := 0
	for _, line := range input {
		min, max, character, pw := parseLine(line)
		if pw[min-1] == character && pw[max-1] == character {
			continue
		}

		if pw[min-1] == character || pw[max-1] == character {
			valid++
		}
	}

	return valid
}

func main() {
	var input []string
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(solution1(input))
	fmt.Println(solution2(input))
}

func intFromString(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
