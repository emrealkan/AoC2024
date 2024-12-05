package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func calculateInstructionValue(input string) int {
	matched := regexp.MustCompile(`\d+`).FindAllString(input, -1)

	if len(matched) < 2 {
		return 0
	}

	a, b := matched[0], matched[1]
	left, _ := strconv.Atoi(a)
	right, _ := strconv.Atoi(b)

	return left * right
}

func part1(input string) int {
	instructionValue := 0
	instructionMatches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllString(input, -1)

	for _, value := range instructionMatches {
		instructionValue += calculateInstructionValue(value)
	}

	return instructionValue
}

func part2(input string) int {
	instructionValue := 0
	enableInstruction := true
	instructionMatches := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(don't\(\))|(do\(\))`).FindAllString(input, -1)

	for _, value := range instructionMatches {
		if value == "do()" {
			enableInstruction = true
		} else if value == "don't()" {
			enableInstruction = false
		} else if enableInstruction {
			instructionValue += calculateInstructionValue(value)
		}
	}

	return instructionValue
}

func main() {
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part1: %d\n", part1(string(input)))
	fmt.Printf("Part2: %d\n", part2(string(input)))
}
