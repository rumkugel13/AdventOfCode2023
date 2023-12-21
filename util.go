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

func strToNumArray(numStr string) []int {
	fields := strings.Fields(strings.TrimSpace(numStr))
	var nums = make([]int, len(fields))
	for i, field := range fields {
		num, _ := strconv.Atoi(field)
		nums[i] = num
	}
	return nums
}
