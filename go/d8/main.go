package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type computer struct {
	accumulator  int
	visitedCount map[int]int
	instructions []string
}

func (c *computer) runProgramme(instructions []string) (status bool) {
	c.accumulator = 0
	c.visitedCount = make(map[int]int)
	c.instructions = make([]string, len(instructions))
	copy(c.instructions, instructions)

	pointer := 0
	for {
		if pointer >= len(c.instructions) {
			status = true
			break
		}

		ins := c.instructions[pointer]
		c.visitedCount[pointer]++
		if count := c.visitedCount[pointer]; count > 1 {
			break
		}

		opVal := strings.Split(ins, " ")
		op := opVal[0]
		val, _ := strconv.Atoi(opVal[1])

		switch op {
		case "acc":
			c.accumulator += val
			pointer++
		case "jmp":
			pointer += val
		case "nop":
			pointer++
		}
	}

	return
}

func solution1(instructions []string) int {
	c := new(computer)
	c.runProgramme(instructions)
	return c.accumulator
}

func solution2(instructions []string) (fixed bool, answer int) {
	c := new(computer)
	for i, in := range instructions {
		if !strings.Contains(in, "nop") && !strings.Contains(in, "jmp") {
			continue
		}

		candidateInstructions := make([]string, len(instructions))
		copy(candidateInstructions, instructions)
		if strings.Contains(in, "nop") {
			candidateInstructions[i] = strings.Replace(in, "nop", "jmp", 1)
		} else {
			candidateInstructions[i] = strings.Replace(in, "jmp", "nop", 1)
		}

		fixed = c.runProgramme(candidateInstructions)
		if fixed {
			answer = c.accumulator
			break
		}
	}

	return
}

func main() {
	instructions := make([]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	fmt.Println(solution1(instructions))
	fmt.Println(solution2(instructions))
}
