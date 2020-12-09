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
	for i, val := range input[preambleSize:] {
		prev25 := input[i : i+preambleSize]
		if summands := twoSum(prev25, val); len(summands) != 2 {
			answer = val
			return
		}
	}

	panic("couldn't find a solution")
}

func solution2(input []int) (answer int) {
	target := solution1(input)
	var summands []int

search:
	for i, vali := range input {
		total := vali
		for j, valj := range input[i+1:] {
			total += valj
			if total > target {
				continue search
			}
			if total == target {
				summands = input[i : i+j+1]
				break search
			}
		}
	}

	if len(summands) < 2 {
		panic("couldn't find a solution")
	}

	sort.Ints(summands)
	answer = summands[0] + summands[len(summands)-1]
	return
}

func twoSum(input []int, sum int) (summands []int) {
	compliments := make(map[int]struct{})
	for _, val := range input {
		compliment := sum - val
		if _, ok := compliments[compliment]; ok {
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
