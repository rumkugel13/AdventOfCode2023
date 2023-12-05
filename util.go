package main

import ( 
	"os" 
	"strings"
	"strconv"
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

func isDigit(a byte) bool {
	return a >= '0' && a <='9'
}

type Point struct {
	x, y int
}

func strToNumArray(numStr string) []int {
	var nums []int
	for _, n := range strings.Fields(strings.TrimSpace(numStr)) {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}
	return nums
}