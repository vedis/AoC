package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	reader := strings.NewReader(string(data))
	scanner := bufio.NewScanner(reader)

	left := []int{}
	right := []int{}

	for scanner.Scan() {

		line := scanner.Text()
		values := strings.Fields(line)

		first, _ := strconv.Atoi(values[0])
		second, _ := strconv.Atoi(values[1])

		left = append(left, first)
		right = append(right, second)
	}

	slices.Sort(left)
	slices.Sort(right)

	part1(left, right)
	part2(left, right)
}

func part1(left []int, right []int) {

	distance := 0

	for i, v := range left {
		if v > right[i] {
			distance += v - right[i]
		} else {
			distance += right[i] - v
		}
	}

	fmt.Printf("distance: %v\n", distance)
}

func part2(left []int, right []int) {

	similarity := 0

	for _, leftValue := range left {
		score := 0
		for _, rightValue := range right {
			if leftValue == rightValue {
				score += 1
			}
		}
		similarity += leftValue * score
	}

	fmt.Printf("similarity: %v\n", similarity)
}
