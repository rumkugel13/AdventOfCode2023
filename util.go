package main

import ( 
	"os" 
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

func isDigit(a byte) bool {
	return a >= '0' && a <='9'
}

type Point struct {
	x, y int
}