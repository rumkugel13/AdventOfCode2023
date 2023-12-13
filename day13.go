package main

import (
	"fmt"
)

func day13() {
	lines := getLines("input/13.txt")

	lineCount := 0
	start := 0
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			pattern := lines[start:i]
			i++
			start = i
			v, vok := verticalReflection(pattern)
			h, hok := horizontalReflection(pattern)
			if vok {
				lineCount += v
			}
			if hok {
				lineCount += 100 * h
			}
		}
	}

	var result = lineCount
	var result2 = 0
	fmt.Println("Day 13 Part 1 Result: ", result)
	fmt.Println("Day 13 Part 2 Result: ", result2)
}

func horizontalReflection(pattern []string) (int, bool) {
	for row := 0; row < len(pattern)-1; row++ {
		if pattern[row] == pattern[row+1] && checkHorizontal(pattern, row) {
			return row + 1, true
		}
	}

	return 0, false
}

func checkHorizontal(pattern []string, horLine int) bool {
	for row := horLine - 1; row >= 0 && (horLine+(horLine-row)+1) < len(pattern); row-- {
		if pattern[row] != pattern[horLine+(horLine-row)+1] {
			return false
		}
	}
	return true
}

func verticalReflection(pattern []string) (int, bool) {
	for col := 0; col < len(pattern[0])-1; col++ {
		found := true
		for row := 0; row < len(pattern); row++ {
			if pattern[row][col] != pattern[row][col+1] {
				found = false
				break
			}
		}
		if found && checkVertical(pattern, col) {
			return col + 1, true
		}
	}

	return 0, false
}

func checkVertical(pattern []string, verLine int) bool {
	for col := verLine - 1; col >= 0 && (verLine+(verLine-col)+1) < len(pattern[0]); col-- {
		for row := 0; row < len(pattern); row++ {
			if pattern[row][col] != pattern[row][verLine+(verLine-col)+1] {
				return false
			}
		}
	}
	return true
}
