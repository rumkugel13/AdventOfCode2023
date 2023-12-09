package main

import (
	"fmt"
)

func day09() {
	lines := getLines("input/09.txt")

	predictSum, predictSum2 := 0, 0
	for _, line := range lines {
		sequence := strToNumArray(line)
		predictSum += predict(sequence, false)
		predictSum2 += predict(sequence, true)
	}

	var result = predictSum
	var result2 = predictSum2
	fmt.Println("Day 09 Part 1 Result: ", result)
	fmt.Println("Day 09 Part 2 Result: ", result2)
}

func predict(sequence []int, part2 bool) int {
	differences := make([]int, len(sequence)-1)
	allZeroes := true
	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
		if differences[i] != 0 {
			allZeroes = false
		}
	}
	nextVal := 0
	if !allZeroes {
		nextVal = predict(differences, part2)
	}
	if part2 {
		return sequence[0] - nextVal
	}
	return sequence[len(sequence)-1] + nextVal
}
