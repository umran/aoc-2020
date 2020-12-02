package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution1(input []string) int {
	valid := 0
	for _, entry := range input {
		pwpol := strings.Split(entry, ":")
		pwpolr := strings.Split(pwpol[0], " ")
		r := strings.Split(pwpolr[0], "-")
		min := intFromString(r[0])
		max := intFromString(r[1])

		ch := pwpolr[1]

		count := 0
		for _, c := range pwpol[1] {
			if string(c) == ch {
				count++
			}
		}

		if count >= min && count <= max {
			valid++
		}
	}

	return valid
}

func solution2(input []string) int {
	valid := 0
	for _, entry := range input {
		pwpol := strings.Split(entry, ":")
		pwpolr := strings.Split(pwpol[0], " ")
		r := strings.Split(pwpolr[0], "-")
		posA := intFromString(r[0])
		posB := intFromString(r[1])

		ch := pwpolr[1]

		pw := strings.Split(pwpol[1], " ")[1]
		if string(pw[posA-1]) == ch && string(pw[posB-1]) == ch {
			continue
		} else if string(pw[posA-1]) == ch || string(pw[posB-1]) == ch {
			valid++
		}
	}

	return valid
}

func main() {
	var input []string
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(solution1(input))
	fmt.Println(solution2(input))
}

func intFromString(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
