package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solution1(adapters []int) (answer int) {
	sort.Ints(adapters)
	lastAdapter := 0

	oneDiffs := 0
	threeDiffs := 0
	for _, adapter := range adapters {
		diff := adapter - lastAdapter
		if diff > 3 {
			continue
		}
		if diff == 1 {
			oneDiffs++
		} else if diff == 3 {
			threeDiffs++
		}
		lastAdapter = adapter
	}

	answer = oneDiffs * (threeDiffs + 1)
	return
}

func solution2(adapters []int) int {
	sort.Ints(adapters)
	equivalentJoltsCount := map[int]int{0: 1}
	for _, a := range adapters {
		equivalentJoltsCount[a] += equivalentJoltsCount[a-1] + equivalentJoltsCount[a-2] + equivalentJoltsCount[a-3]
	}
	target := adapters[len(adapters)-1] + 3
	return equivalentJoltsCount[target-1] + equivalentJoltsCount[target-2] + equivalentJoltsCount[target-3]
}

func solve2(ipt []int) {
	memo := make([]int, len(ipt))

	for i, v := range ipt[:3] {
		if v <= 3 {
			memo[i]++
		}
	}

	for i, v := range ipt {
		var sliceEnd int
		if i+4 < len(ipt) {
			sliceEnd = i + 4
		} else {
			sliceEnd = len(ipt)
		}

		for j, w := range ipt[i+1 : sliceEnd] {
			if j+i >= len(memo) {
				break
			}

			if w <= v+3 {
				memo[j+i] += memo[i]
			}
		}
	}

	fmt.Println(memo[len(ipt)-1])
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
	solve2(input)
}
