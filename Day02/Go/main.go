package main

import (
	"fmt"
	"math"
	"os"
	"log"
	"strconv"
	"strings"
)

func checkLevelSafety(levels []int) bool {
	var direction int

	for i := 0; i < len(levels)-1; i++ {
		difference := levels[i] - levels[i+1]
		sign := 0

		if difference > 0 {
			sign = 1
		} else if difference < 0 {
			sign = -1
		}

		absDifference := int(math.Abs(float64(difference)))

		if direction != 0 && direction != sign {
			return false
		}
		if absDifference < 1 || absDifference > 3 {
			return false
		}

		direction = sign
	}

	return true
}

func part1(lines []string) int {
	count := 0
	for _, report := range lines {
		levels := make([]int, 0)
		for _, str := range strings.Split(report, " ") {
			var level int
			level, _ = strconv.Atoi(str)
			levels = append(levels, level)
		}

		if checkLevelSafety(levels) {
			count++
		}
	}
	return count
}


func part2(lines []string) int {
	count := 0
	for _, report := range lines {
		levels := make([]int, 0)
		for _, str := range strings.Split(report, " ") {
			var level int
			level, _ = strconv.Atoi(str)
			levels = append(levels, level)
		}

		if checkLevelSafety(levels) {
			count++
		} else {
			isSafe := false
			for i := 0; i < len(levels); i++ {
				// Create a new slice without the i-th element
				subLevels := append([]int(nil), levels[:i]...) // Copy slice before i
				subLevels = append(subLevels, levels[i+1:]...)  // Append slice after i
				if checkLevelSafety(subLevels) {
					isSafe = true
					break
				}
			}
			if isSafe {
				count++
			}
		}
	}
	return count
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