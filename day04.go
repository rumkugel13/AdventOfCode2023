package main

import (
	"fmt"
	"strings"
)

func day04() {
	lines := getLines("input/04.txt")
	var points = 0
	var winnings = make([]int, len(lines))

	for i, line := range lines {
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		haveNums := strToNumArray(numbers[0])
		winningNums := strToNumArray(numbers[1])
		winnings[i] = countWinnings(haveNums, winningNums)
		points += pow2(winnings[i])
	}

	var cards = make([]int, len(lines))
	var cardCount = 0
	for i := range lines {
		cards[i] += 1
		winCount := winnings[i]
		for j := i + 1; j <= i+winCount; j++ {
			cards[j] += cards[i]
		}
		cardCount += cards[i]
	}

	var result = points
	var result2 = cardCount
	fmt.Println("Day 04 Part 1 Result: ", result)
	fmt.Println("Day 04 Part 2 Result: ", result2)
}

func countWinnings(have, winners []int) int {
	var count = 0
	for _, num := range have {
		for _, winNum := range winners {
			if num == winNum {
				count++
			}
		}
	}
	return count
}

func pow2(n int) int {
	if n > 0 {
		return 1 << (n - 1)
	}
	return 0
}
