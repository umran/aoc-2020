package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func solution1(groups []string) int {
	total := 0
	for _, group := range groups {
		seenQuestions := make(map[rune]struct{})
		individuals := strings.Split(group, "\n")
		for _, individual := range individuals {
			for _, q := range individual {
				if _, seen := seenQuestions[q]; seen {
					continue
				}
				seenQuestions[q] = struct{}{}
				total++
			}
		}
	}

	return total
}

func solution2(groups []string) int {
	total := 0
	for _, group := range groups {
		answerCounts := make(map[rune]int)
		individuals := strings.Split(group, "\n")
		for _, individual := range individuals {
			for _, q := range individual {
				answerCounts[q]++
				if answerCounts[q] == len(individuals) {
					total++
				}
			}
		}
	}

	return total
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	text := string(bytes)
	groups := strings.Split(text, "\n\n")

	fmt.Println(solution1(groups))
	fmt.Println(solution2(groups))
}
