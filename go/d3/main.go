package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution1(input []string, slope []int) int {
	return treesForSlope(input, slope[0], slope[1])
}

func solution2(input []string, slopes [][]int) int {
	answer := 1
	for _, slope := range slopes {
		answer = answer * treesForSlope(input, slope[0], slope[1])
	}

	return answer
}

func treesForSlope(input []string, dx, dy int) int {
	trees := 0
	x := 0
	for y := 0; y < len(input); y += dy {
		row := []rune(input[y])
		if row[x] == []rune("#")[0] {
			trees++
		}
		x = (x + dx) % len(row)
	}
	return trees
}

func main() {
	var input []string
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(solution1(input, []int{3, 1}))
	fmt.Println(solution2(input, [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}))
}
