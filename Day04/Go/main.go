package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var DIRECTIONS = [8][2]int{
	{1, 0},   // Down
	{-1, 0},  // Up
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 1},   // Down-right diagonal
	{-1, -1}, // Up-left diagonal
	{1, -1},  // Down-left diagonal
	{-1, 1},  // Up-right diagonal
}

func calculateXMAS(matrix [][]string, x, y int) int {
	xmasCount := 0

	for _, direction := range DIRECTIONS {
		dx, dy := direction[0], direction[1]
		word := ""

		xx, yy := x, y
		for i := 0; i < 4; i++ {
			if xx < 0 || xx >= len(matrix) || yy < 0 || yy >= len(matrix[0]) {
				break
			}

			word += matrix[xx][yy]
			xx += dx
			yy += dy
		}

		if word == "XMAS" {
			xmasCount++
		}
	}

	return xmasCount
}

func part1(matrix [][]string) int {
	xmasCount := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			xmasCount += calculateXMAS(matrix, x, y)
		}
	}

	return xmasCount
}

func part2(matrix [][]string) int {
	fmt.Println("TODO")
	return 2
}

func main() {
	var matrix [][]string
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	// Convert each line into a row in the matrix
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}
	fmt.Printf("Part1: %d\n", part1(matrix))
	fmt.Printf("Part2: %d\n", part2(matrix))
}
