package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution1(input []int, sum int) int {
	compliments := make(map[int]struct{})
	for _, val := range input {
		compliment := sum - val
		if _, ok := compliments[compliment]; ok {
			return val * compliment
		}
		compliments[val] = struct{}{}
	}
	panic("nope!")
}

func solution2(input []int, sum int) int {
	for i, vali := range input {
		compliments := make(map[int]struct{})
		for j, valj := range input {
			if i == j {
				continue
			}
			compliment := sum - vali - valj
			if _, ok := compliments[compliment]; ok {
				return vali * valj * compliment
			}
			compliments[valj] = struct{}{}
		}
	}
	panic("nope!")
}

func main() {
	var input []int
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, intFromString(scanner.Text()))
	}

	fmt.Println(solution1(input, 2020))
	fmt.Println(solution2(input, 2020))
}

func intFromString(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
