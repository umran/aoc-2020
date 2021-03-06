package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagDirectory map[string]map[string]int

func (directory bagDirectory) leadsToTarget(bag, target string) bool {
	for candidate := range directory[bag] {
		if candidate == target {
			return true
		}
		if directory.leadsToTarget(candidate, target) {
			return true
		}
	}

	return false
}

func (directory bagDirectory) countInnerBags(bag string) int {
	total := 0
	for innerBag, amount := range directory[bag] {
		total += amount + amount*directory.countInnerBags(innerBag)
	}
	return total
}

func parseBags(input []string) bagDirectory {
	directory := make(bagDirectory)
	for _, ruleDot := range input {
		rule := strings.TrimSuffix(ruleDot, ".")
		parts := strings.Split(rule, "contain")
		bagKey := strings.TrimSuffix(parts[0], " ")
		spaceContents := strings.Split(parts[1], ",")
		directory[bagKey] = make(map[string]int)
		for _, spaceContent := range spaceContents {
			content := strings.TrimPrefix(spaceContent, " ")
			if content == "no other bags" {
				break
			}
			contentParts := strings.Split(content, " ")
			amount, _ := strconv.Atoi(contentParts[0])
			innerBagKey := strings.Join(contentParts[1:], " ")
			if []rune(innerBagKey)[len([]rune(innerBagKey))-1] != 's' {
				innerBagKey = fmt.Sprintf("%ss", innerBagKey)
			}
			directory[bagKey][innerBagKey] = amount
		}
	}

	return directory
}

func solution1(directory bagDirectory, target string) int {
	total := 0
	for bag := range directory {
		if directory.leadsToTarget(bag, target) {
			total++
		}
	}

	return total
}

func solution2(directory bagDirectory, bag string) int {
	return directory.countInnerBags(bag)
}

func main() {
	var input []string
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	directory := parseBags(input)
	fmt.Println(solution1(directory, "shiny gold bags"))
	fmt.Println(solution2(directory, "shiny gold bags"))
}
