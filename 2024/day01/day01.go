package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var lefts []int
var rights []int
var distanceSum = 0

func main() {
	data, _ := os.ReadFile("./input.txt")
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Fields(fmt.Sprintf(" %s ", line))

		leftValue, _ := strconv.Atoi(parts[0])
		lefts = append(lefts, leftValue)
		rightValue, _ := strconv.Atoi(parts[1])
		rights = append(rights, rightValue)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	for idx, left := range lefts {
		right := rights[idx]

		distance := right - left
		if distance < 0 {
			distance = distance * -1
		}
		distanceSum += distance
	}

	fmt.Println("Sum: ", distanceSum)
}
