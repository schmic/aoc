package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("./input.txt")

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	muls := re.FindAllStringSubmatch(string(data), -1)
	sumOfMuls := 0

	for _, mul := range muls {
		mulNum1, _ := strconv.Atoi(mul[1])
		mulNum2, _ := strconv.Atoi(mul[2])
		sumOfMuls += mulNum1 * mulNum2
	}

	fmt.Printf("Sum of matches: %d\n", sumOfMuls)
}
