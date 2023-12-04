package main

import (
	"fmt"
	"strings"
)

var numbers = map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func day01() {
	lines := getLines("input/01.txt")

	var result = getSum(lines, false)
	var result2 = getSum(lines, true)
	fmt.Println("Day 01 Part 1 Result: ", result)
	fmt.Println("Day 01 Part 2 Result: ", result2)
}

func getSum(lines []string, part2 bool) int {
	var sum int

	for _, line := range lines {
		var first, last int
	out:
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = int(line[i] - '0')
				break
			} else if part2 {
				for k := range numbers {
					if strings.HasPrefix(line[i:], k) {
						first = numbers[k]
						break out
					}
				}
			}
		}
	out2:
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = int(line[i] - '0')
				break
			} else if part2 {
				for k := range numbers {
					if strings.HasSuffix(line[:i], k) {
						last = numbers[k]
						break out2
					}
				}
			}
		}
		sum += 10*first + last
	}
	return sum
}
