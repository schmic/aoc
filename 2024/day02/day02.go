package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

var (
	Up   = Direction(0)
	Down = Direction(1)
)

func main() {
	data, _ := os.ReadFile("./input.txt")

	safeLines := 0

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		numbers := parseNumbers(line)
		direction := getDirection(numbers, 0)

		fmt.Print(numbers)
		if checkLine(numbers, direction) {
			fmt.Println(" 👌")
			safeLines++
		} else {
			fmt.Println(" 👎")
		}
	}

	fmt.Println("Safe lines: ", safeLines)
}

func checkLine(numbers []int, direction Direction) bool {
	for idx := range len(numbers) - 1 {
		if !isSafe(numbers, idx) {
			return false
		}

		idxDirection := getDirection(numbers, idx)
		if idxDirection != direction {
			return false
		}

	}
	return true
}

func isSafe(numbers []int, idx int) bool {
	left := numbers[idx]
	right := numbers[idx+1]

	if abs(right-left) >= 1 && abs(right-left) <= 3 {
		return true
	}
	return false
}

func parseNumbers(line string) []int {
	numbers := make([]int, 0)
	for _, numstr := range strings.Split(line, " ") {
		num, _ := strconv.Atoi(numstr)
		numbers = append(numbers, num)
	}
	return numbers
}

func getDirection(numbers []int, idx int) Direction {
	if numbers[idx] > numbers[idx+1] {
		return Up
	}
	return Down
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
