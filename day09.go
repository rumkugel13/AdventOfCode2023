package main

import (
	"fmt"
)

func day09() {
	lines := getLines("input/09.txt")

	predictSum, predictSum2 := 0, 0
	for _, line := range lines {
		sequence := strToNumArray(line)
		part1, part2 := predict(sequence)
		predictSum += part1
		predictSum2 += part2
	}

	var result = predictSum
	var result2 = predictSum2
	fmt.Println("Day 09 Part 1 Result: ", result)
	fmt.Println("Day 09 Part 2 Result: ", result2)
}

func predict(sequence []int) (next, prev int) {
	differences := make([]int, len(sequence)-1)
	allZeroes := true
	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
		if differences[i] != 0 {
			allZeroes = false
		}
	}
	prevVal, nextVal := 0, 0
	if !allZeroes {
		nextVal, prevVal = predict(differences)
	}
	return sequence[len(sequence)-1] + nextVal, sequence[0] - prevVal
}
