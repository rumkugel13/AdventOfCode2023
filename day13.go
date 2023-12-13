package main

import (
	"fmt"
)

func day13() {
	lines := getLines("input/13.txt")

	lineCount, lineCount2 := 0, 0
	start := 0
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			pattern := lines[start:i]
			i++
			start = i
			v, vok := verticalReflection(pattern, false)
			h, hok := horizontalReflection(pattern, false)
			if vok {
				lineCount += v
			}
			if hok {
				lineCount += 100 * h
			}
			v2, vok2 := verticalReflection(pattern, true)
			h2, hok2 := horizontalReflection(pattern, true)
			if vok2 {
				lineCount2 += v2
			}
			if hok2 {
				lineCount2 += 100 * h2
			}
		}
	}

	var result = lineCount
	var result2 = lineCount2
	fmt.Println("Day 13 Part 1 Result: ", result)
	fmt.Println("Day 13 Part 2 Result: ", result2)
}

func horizontalReflection(pattern []string, part2 bool) (int, bool) {
	for row := 0; row < len(pattern)-1; row++ {
		found := true
		smudges := 0
		for col := 0; col < len(pattern[0]); col++ {
			if pattern[row][col] != pattern[row+1][col] {
				smudges++
				if smudges > 1 {
					found = false
					break
				}
			}
		}
		if found {
			ok, alt := checkHorizontal(pattern, row, 1-smudges)
			if (!part2 && ok && !alt) || (part2 && alt) {
				return row + 1, true
			}
		}
	}

	return 0, false
}

func checkHorizontal(pattern []string, horLine, maxSmudge int) (bool, bool) {
	smudges := 0
	for row := horLine - 1; row >= 0 && (horLine+(horLine-row)+1) < len(pattern); row-- {
		for col := 0; col < len(pattern[0]); col++ {
			if pattern[row][col] != pattern[horLine+(horLine-row)+1][col] {
				smudges++
				if smudges > maxSmudge {
					return false, false
				}
			}
		}
	}
	return true, smudges == maxSmudge
}

func verticalReflection(pattern []string, part2 bool) (int, bool) {
	for col := 0; col < len(pattern[0])-1; col++ {
		found := true
		smudges := 0
		for row := 0; row < len(pattern); row++ {
			if pattern[row][col] != pattern[row][col+1] {
				smudges++
				if smudges > 1 {
					found = false
					break
				}
			}
		}
		if found {
			ok, alt := checkVertical(pattern, col, 1-smudges)
			if (!part2 && ok && !alt) || (part2 && alt) {
				return col + 1, true
			}
		}
	}

	return 0, false
}

func checkVertical(pattern []string, verLine, maxSmudge int) (bool, bool) {
	smudges := 0
	for col := verLine - 1; col >= 0 && (verLine+(verLine-col)+1) < len(pattern[0]); col-- {
		for row := 0; row < len(pattern); row++ {
			if pattern[row][col] != pattern[row][verLine+(verLine-col)+1] {
				smudges++
				if smudges > maxSmudge {
					return false, false
				}
			}
		}
	}
	return true, smudges == maxSmudge
}
