package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day12() {
	lines := getLines("input/12.txt")

	variationSum, varSum2 := 0, 0
	for _, line := range lines {
		split := strings.Fields(line)
		springs := []byte(split[0])
		groups := commaSepToIntArr(strings.Split(split[1], ","))
		variationSum += countVariations(springs, groups, 0)

		p2 := springs
		p2g := groups
		for i := 0; i < 4; i++ {
			p2 = append(p2, '?')
			p2 = append(p2, springs...)
			p2g = append(p2g, groups...)
		}
		varSum2 += countVariations2(p2, p2g, map[string]int{})
	}

	var result = variationSum
	var result2 = varSum2
	fmt.Println("Day 12 Part 1 Result: ", result)
	fmt.Println("Day 12 Part 2 Result: ", result2)
}

func countVariations2(springs []byte, groups []int, cache map[string]int) int {
	key := string(springs)
	for _, num := range groups {
		key += string(num + '0')
	}

	if _, found := cache[key]; found {
		return cache[key]
	}

	// no more springs to check
	if len(springs) == 0 {
		// all groups found
		if len(groups) == 0 {
			cache[key] = 1
			return 1
		}
		// still groups remaining
		cache[key] = 0
		return 0
	}

	count := 0
	// check both variations
	if springs[0] == '?' {
		dup1 := dup(springs)
		dup1[0] = '.'
		count += countVariations2(dup1, groups, cache)
		dup2 := dup(springs)
		dup2[0] = '#'
		count += countVariations2(dup2, groups, cache)

		cache[key] = count
		return count
	}

	// skip character (todo: refactor)
	if springs[0] == '.' {
		return countVariations2(springs[1:], groups, cache)
	}

	// no groups left for #-sequences
	if len(groups) == 0 {
		cache[key] = 0
		return 0
	}

	// count sequence and check with next group
	if springs[0] == '#' {
		i := 0
		for ; i < len(springs) && springs[i] != '.' && !(springs[i] == '?' && i == groups[0]); i++ {
		}
		if i == groups[0] {
			if i < len(springs) {
				// still springs left, skip a gap in between groups as well
				count += countVariations2(springs[i+1:], groups[1:], cache)
				cache[key] = count
				return count
			}
			count += countVariations2(springs[i:], groups[1:], cache)
			cache[key] = count
			return count
		} else {
			// sequence length doesn't match next group
			cache[key] = 0
			return 0
		}
	}
	cache[key] = 0
	return 0
}

func countVariations(springs []byte, groups []int, start int) int {
	for i := start; i < len(springs); i++ {
		if springs[i] == '?' {
			dup1 := dup(springs)
			dup1[i] = '.'
			count := countVariations(dup1, groups, i+1)
			dup2 := dup(springs)
			dup2[i] = '#'
			count += countVariations(dup2, groups, i+1)
			return count
		}
	}
	if checkVariation(springs, groups) {
		return 1
	}
	return 0
}

func dup(data []byte) []byte {
	result := make([]byte, len(data))
	copy(result, data)
	return result
}

func checkVariation(springs []byte, groups []int) bool {
	for i := 0; i < len(springs); i++ {
		if springs[i] == '#' {
			start := i
			for ; i < len(springs) && springs[i] == '#'; i++ {
			}
			if 0 < len(groups) && i-start == groups[0] {
				groups = groups[1:]
			} else {
				return false
			}
		}
	}
	return len(groups) == 0
}

func commaSepToIntArr(data []string) []int {
	result := make([]int, len(data))
	for i, val := range data {
		num, _ := strconv.Atoi(val)
		result[i] = num
	}
	return result
}
