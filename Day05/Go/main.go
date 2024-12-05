package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	First  int
	Second int
}

type ParsedInput struct {
	Rules   []Rule
	Updates [][]int
}

func main() {
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sections := strings.Split(string(input), "\n\n")
	parsedInput := parseInput(sections)
	solvePuzzles(parsedInput.Rules, parsedInput.Updates)
}

func parseInput(sections []string) ParsedInput {
	ruleLines := strings.Split(sections[0], "\n")
	rules := make([]Rule, len(ruleLines))

	for i, line := range ruleLines {
		parts := strings.Split(line, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		rules[i] = Rule{First: first, Second: second}
	}

	updateLines := strings.Split(sections[1], "\n")
	updates := make([][]int, len(updateLines))

	for i, line := range updateLines {
		values := strings.Split(line, ",")
		update := make([]int, len(values))
		for j, v := range values {
			update[j], _ = strconv.Atoi(v)
		}
		updates[i] = update
	}

	return ParsedInput{Rules: rules, Updates: updates}
}

func solvePuzzles(rules []Rule, updates [][]int) {
	sumOfPart1 := 0
	sumOfPart2 := 0

	for _, update := range updates {
		if isLineOrderValid(update, rules) {
			middleIndex := len(update) / 2
			sumOfPart1 += update[middleIndex]
		} else {
			modified := fixLineOrder(update, rules)
			middleIndex := len(modified) / 2
			sumOfPart2 += modified[middleIndex]
		}
	}

	fmt.Printf("Part1: %d\n", sumOfPart1)
	fmt.Printf("Part2: %d\n", sumOfPart2)
}

func isLineOrderValid(updates []int, rules []Rule) bool {
	for _, rule := range rules {
		indexOfFirstRule := findIndex(updates, rule.First)
		indexOfSecondRule := findIndex(updates, rule.Second)
		if indexOfFirstRule != -1 && indexOfSecondRule != -1 {
			if indexOfFirstRule >= indexOfSecondRule {
				return false
			}
		}
	}
	return true
}

func findIndex(updates []int, item int) int {
	for index, value := range updates {
		if value == item {
			return index
		}
	}
	return -1
}

func fixLineOrder(originalUpdates []int, rules []Rule) []int {
	copyUpdates := append([]int(nil), originalUpdates...)

	wasUpdated := true
	for wasUpdated {
		wasUpdated = false
		for _, rule := range rules {
			indexOfFirstRule, indexOfSecondRule := -1, -1
			for i, copyUpdate := range copyUpdates {
				if copyUpdate == rule.First {
					indexOfFirstRule = i
				}
				if copyUpdate == rule.Second {
					indexOfSecondRule = i
				}
				if indexOfFirstRule != -1 && indexOfSecondRule != -1 {
					break
				}
			}

			if indexOfFirstRule != -1 && indexOfSecondRule != -1 && indexOfFirstRule >= indexOfSecondRule {
				copyUpdates[indexOfFirstRule], copyUpdates[indexOfSecondRule] = copyUpdates[indexOfSecondRule], copyUpdates[indexOfFirstRule]
				wasUpdated = true
			}
		}
	}
	return copyUpdates
}
