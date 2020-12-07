package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution1(input []string, targetKey string) int {
	bagMap := createBagMap(input)
	total := 0
	for bagKey := range bagMap {
		if existsInBag(bagMap, bagKey, targetKey) {
			total++
		}
	}

	return total
}

func solution2(input []string, bagKey string) int {
	bagMap := createBagMap(input)
	return countTotalBags(bagMap, bagKey)
}

func existsInBag(bagMap map[string]map[string]int, bagKey, targetKey string) bool {
	for candidateKey := range bagMap[bagKey] {
		if candidateKey == targetKey {
			return true
		}
		if existsInBag(bagMap, candidateKey, targetKey) {
			return true
		}
	}

	return false
}

func countTotalBags(bagMap map[string]map[string]int, bagKey string) int {
	total := 0
	for innerKey, amount := range bagMap[bagKey] {
		total += amount
		innerAmount := countTotalBags(bagMap, innerKey)
		if innerAmount > 0 {
			total += amount * innerAmount
		}
	}
	return total
}

func createBagMap(input []string) map[string]map[string]int {
	bagMap := make(map[string]map[string]int)
outer:
	for _, ruleDot := range input {
		rule := strings.TrimSuffix(ruleDot, ".")
		parts := strings.Split(rule, "contain")
		bagKey := strings.TrimSuffix(parts[0], " ")
		spaceContents := strings.Split(parts[1], ",")
		bagMap[bagKey] = make(map[string]int)
		for _, spaceContent := range spaceContents {
			content := strings.TrimPrefix(spaceContent, " ")
			if content == "no other bags" {
				continue outer
			}
			contentParts := strings.Split(content, " ")
			contentAmount, _ := strconv.Atoi(contentParts[0])
			contentKey := strings.Join(contentParts[1:], " ")
			if []rune(contentKey)[len([]rune(contentKey))-1] != 's' {
				contentKey = fmt.Sprintf("%ss", contentKey)
			}
			bagMap[bagKey][contentKey] = contentAmount
		}
	}

	return bagMap
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
