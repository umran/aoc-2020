package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagDirectory map[string]map[string]int

func (directory bagDirectory) hasPathToBag(bagKey, targetKey string) bool {
	for candidateKey := range directory[bagKey] {
		if candidateKey == targetKey {
			return true
		}
		if directory.hasPathToBag(candidateKey, targetKey) {
			return true
		}
	}

	return false
}

func (directory bagDirectory) countInnerBags(bagKey string) int {
	total := 0
	for innerKey, amount := range directory[bagKey] {
		total += amount
		innerAmount := directory.countInnerBags(innerKey)
		if innerAmount > 0 {
			total += amount * innerAmount
		}
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
				continue
			}
			contentParts := strings.Split(content, " ")
			contentAmount, _ := strconv.Atoi(contentParts[0])
			contentKey := strings.Join(contentParts[1:], " ")
			if []rune(contentKey)[len([]rune(contentKey))-1] != 's' {
				contentKey = fmt.Sprintf("%ss", contentKey)
			}
			directory[bagKey][contentKey] = contentAmount
		}
	}

	return directory
}

func solution1(input []string, targetKey string) int {
	directory := parseBags(input)
	total := 0
	for bagKey := range directory {
		if directory.hasPathToBag(bagKey, targetKey) {
			total++
		}
	}

	return total
}

func solution2(input []string, bagKey string) int {
	directory := parseBags(input)
	return directory.countInnerBags(bagKey)
}

func main() {
	var input []string
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(solution1(input, "shiny gold bags"))
	fmt.Println(solution2(input, "shiny gold bags"))
}
