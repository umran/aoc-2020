package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solution1 ...
func Solution1(input []int, sum int) int {
	compliments := make(map[int]int)
	for _, val := range input {
		compliment := sum - val
		if _, ok := compliments[compliment]; ok {
			return val * compliment
		}
		compliments[val] = val
	}
	panic("nope!")
}

// Solution2 ...
func Solution2(input []int, sum int) int {
	for i, vali := range input {
		compliments := make(map[int]int)
		for j, valj := range input {
			if i == j {
				continue
			}
			compliment := sum - vali - valj
			if _, ok := compliments[compliment]; ok {
				return vali * valj * compliment
			}
			compliments[valj] = valj
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

	fmt.Println(Solution1(input, 2020))
	fmt.Println(Solution2(input, 2020))
}

func intFromString(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
