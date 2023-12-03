package main

import (
	"fmt"
	"strconv"
)

func day03() {
	schematics := getLines("03.txt")
	var partNumberSum int

	for y := 0; y < len(schematics); y++ {
		for x := 0; x < len(schematics[y]); x++ {
			char := schematics[y][x]
			if isDigit(char) {
				var dx = x + 1
				for ; dx < len(schematics[y]); dx++ {
					if !isDigit(schematics[y][dx]) {
						break
					}
				}

				var num, _ = strconv.Atoi(schematics[y][x:dx])

				if isPartNumber(schematics, x, dx, y) {
					partNumberSum += num
				}

				x = dx
			}
		}
	}

	var result = partNumberSum
	var result2 = 0
	fmt.Println("Day 03 Part 1 Result: ", result)
	fmt.Println("Day 03 Part 2 Result: ", result2)
}

func isPartNumber(schematics []string, x1, x2, y int) bool {
	for dy := max(y-1, 0); dy < min(y+2, len(schematics)); dy++ {
		for dx := max(x1-1, 0); dx < min(x2+1, len(schematics[y])); dx++ {
			char := schematics[dy][dx]
			if !(isDigit(char) || char == '.') {
				return true
			}
		}
	}
	return false
}
