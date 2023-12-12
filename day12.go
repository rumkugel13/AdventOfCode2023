package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day12() {
	lines := getLines("input/12.txt")

	variationSum := 0
	for _, line := range lines {
		split := strings.Fields(line)
		springs := split[0]
		groups := commaSepToIntArr(strings.Split(split[1], ","))
		variationSum += countVariations(springs, groups)
	}

	var result = variationSum
	var result2 = 0
	fmt.Println("Day 12 Part 1 Result: ", result)
	fmt.Println("Day 12 Part 2 Result: ", result2)
}

func countVariations(springs string, groups []int) int {
	count := 0
	data := []byte(springs)
	variations := makeVariations(data, 0)
	for _, variation := range variations {
		if checkVariation(variation, groups) {
			count++
		}
	}
	return count
}

func makeVariations(springs []byte, start int) [][]byte {
	variations := [][]byte{}
	complete := true
	for i := start; i < len(springs); i++ {
		if springs[i] == '?' {
			var springCopy = make([]byte, len(springs))
			copy(springCopy, springs)
			springCopy[i] = '.'
			variations = append(variations, makeVariations(springCopy, i+1)...)
			var springCopy2 = make([]byte, len(springs))
			copy(springCopy2, springs)
			springCopy2[i] = '#'
			variations = append(variations, makeVariations(springCopy2, i+1)...)
			complete = false
			break
		}
	}
	if complete {
		variations = append(variations, springs)
	}
	return variations
}

func checkVariation(springs []byte, groups []int) bool {
	group := 0
	for i := 0; i < len(springs); i++ {
		if springs[i] == '#' {
			start := i
			for ; i < len(springs) && springs[i] == '#'; i++ {
			}
			if group < len(groups) && i-start != groups[group] {
				return false
			} else {
				group++
			}
		}
	}
	return group == len(groups)
}

func commaSepToIntArr(data []string) []int {
	result := make([]int, len(data))
	for i, val := range data {
		num, _ := strconv.Atoi(val)
		result[i] = num
	}
	return result
}
