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
		variationSum += countVariations([]byte(springs), groups, 0)
	}

	var result = variationSum
	var result2 = 0
	fmt.Println("Day 12 Part 1 Result: ", result)
	fmt.Println("Day 12 Part 2 Result: ", result2)
}

func countVariations(springs []byte, groups []int, start int) int {
	for i := start; i < len(springs); i++ {
		if springs[i] == '?' {
			springs[i] = '.'
			count := countVariations(springs, groups, i+1)
			springs[i] = '#'
			count += countVariations(springs, groups, i+1)
			springs[i] = '?'
			return count
		}
	}
	if checkVariation(springs, groups) {
		return 1
	}
	return 0
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
