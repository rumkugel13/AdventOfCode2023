package main

import (
	"os"
	"strconv"
	"strings"
)

func getLines(file string) []string {
	data, _ := os.ReadFile(file)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func isDigit(a byte) bool {
	return a >= '0' && a <= '9'
}

type Point struct {
	x, y int
}

func insideGrid(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}

func findInGrid(grid []string, char rune) Point {
	for y, row := range grid {
		for x, col := range row {
			if col == char {
				return Point{x, y}
			}
		}
	}
	return Point{-1, -1}
}

func strToNumArray(numStr string) []int {
	fields := strings.Fields(strings.TrimSpace(numStr))
	var nums = make([]int, len(fields))
	for i, field := range fields {
		num, _ := strconv.Atoi(field)
		nums[i] = num
	}
	return nums
}
