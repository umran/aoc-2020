package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution1(input []string) (highest int) {
	highest = sortedIds(input)[len(input)-1]
	return
}

func solution2(input []string) int {
	ids := sortedIds(input)
	missingIds := make([]int, 0)
	for i, id := range ids {
		nextIndex := i + 1
		if nextIndex > len(ids)-1 {
			break
		}

		if ids[nextIndex]-id == 2 {
			missingIds = append(missingIds, id+1)
		}
	}

	return missingIds[0]
}

func sortedIds(input []string) []int {
	ids := make([]int, len(input))
	for i, bp := range input {
		rowInfo := string([]rune(bp)[:7])
		colInfo := string([]rune(bp)[7:])

		row := findRow(rowInfo)
		col := findCol(colInfo)

		id := row*8 + col
		ids[i] = id
	}

	sort.Ints(ids)
	return ids
}

func findRow(input string) (location int) {
	location = locate(input, 128, 'F')
	return
}

func findCol(input string) (location int) {
	location = locate(input, 8, 'L')
	return
}

func locate(input string, space int, lower rune) (location int) {
	for _, c := range input {
		if c == lower {
			space = location + (space-location)/2
		} else {
			location = location + (space-location)/2
		}
	}

	return
}

func main() {
	var input []string
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("the solution to part 1 is: %d\n", solution1(input))
	fmt.Printf("the solution to part 2 is: %d\n", solution2(input))
}
