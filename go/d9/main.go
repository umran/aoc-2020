package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const preambleSize = 25

func solution1(input []int) (answer int) {
	for i, val := range input {
		if i < preambleSize {
			continue
		}

		prev25 := input[i-preambleSize : i]
		if exists, _ := twoSum(prev25, val); !exists {
			answer = val
			break
		}
	}

	return
}

func solution2(input []int) (answer int) {
	target := solution1(input)
	var summands []int

search:
	for i, vali := range input {
		total := vali
		for j, valj := range input {
			if j <= i {
				continue
			}
			total += valj
			if total > target {
				continue search
			}
			if total == target {
				summands = input[i : j+1]
				break search
			}
		}
	}

	sort.Ints(summands)
	answer = summands[0] + summands[len(summands)-1]
	return
}

func twoSum(input []int, sum int) (exists bool, summands []int) {
	compliments := make(map[int]struct{})
	for _, val := range input {
		compliment := sum - val
		if _, ok := compliments[compliment]; ok {
			exists = true
			summands = []int{val, compliment}
			break
		}
		compliments[val] = struct{}{}
	}
	return
}

func main() {
	var input []int
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed, _ := strconv.Atoi(scanner.Text())
		input = append(input, parsed)
	}

	fmt.Println(solution1(input))
	fmt.Println(solution2(input))
}
