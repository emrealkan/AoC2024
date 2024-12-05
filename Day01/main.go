package main

import (
	"fmt"
	"math"
	"os"
	"log"
	"sort"
	"strings"
	"strconv"
)

func part1(lines []string) int {
	var leftList []int
	var rightList []int

	for _, line := range lines {
		var left, right int
		parts := strings.Fields(line)
		left, _ = strconv.Atoi(parts[0])
		right, _ = strconv.Atoi(parts[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return totalDistance
}


func part2(lines []string) int {
	var leftList []int
	var rightList []int

	for _, line := range lines {
		var left, right int
		parts := strings.Fields(line)
		left, _ = strconv.Atoi(parts[0])
		right, _ = strconv.Atoi(parts[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	rightFrequency := make(map[int]int)

	for _, num := range rightList {
		rightFrequency[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		countInRight := rightFrequency[num]
		similarityScore += num * countInRight
	}

	return similarityScore
}

func main() {
	input, err := os.ReadFile("./input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	inputString := string(input)
	splitStr := strings.Split(inputString, "\n")
	
	fmt.Println("Part1:", part1(splitStr))
	fmt.Println("Part2:", part2(splitStr))
}