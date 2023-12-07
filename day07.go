package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TypeAndBid struct {
	hand          string
	handType, bid int
}

func day07() {
	input := getLines("input/07.txt")
	entries := []TypeAndBid{}
	entries2 := []TypeAndBid{}

	for _, hand := range input {
		handAndBid := strings.Fields(hand)
		handType := getHandType(handAndBid[0], false)
		handType2 := getHandType(handAndBid[0], true)
		bid, _ := strconv.Atoi(handAndBid[1])
		entries = append(entries, TypeAndBid{handAndBid[0], handType, bid})
		entries2 = append(entries2, TypeAndBid{handAndBid[0], handType2, bid})
	}

	var result = getWinnings(entries, false)
	fmt.Println("Day 07 Part 1 Result: ", result)

	var result2 = getWinnings(entries2, true)
	fmt.Println("Day 07 Part 2 Result: ", result2)
}

func getWinnings(entries []TypeAndBid, part2 bool) int {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].handType == entries[j].handType {
			return lessHandType(entries[i].hand, entries[j].hand, part2)
		}
		return entries[i].handType < entries[j].handType
	})

	winnings := 0
	for i, entry := range entries {
		winnings += (i + 1) * entry.bid
	}
	return winnings
}

func lessHandType(hand1, hand2 string, part2 bool) bool {
	nums := map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	if part2 {
		nums['J'] = 1
	}
	for i, a := range hand1 {
		if byte(a) == hand2[i] {
			continue
		} else if nums[byte(a)] < nums[hand2[i]] {
			return true
		} else {
			return false
		}
	}
	return false
}

func getHandType(hand string, part2 bool) int {
	cards := map[byte]int{}
	for _, card := range hand {
		cards[byte(card)]++
	}

	jokers := cards['J']

	switch len(cards) {
	case 1:
		return 6
	case 2:
		if part2 && jokers > 0 {
			return 6
		}
		for _, count := range cards {
			if count == 4 {
				return 5
			}
		}
		return 4
	case 3:
		for _, count := range cards {
			if count == 3 {
				if part2 && jokers > 0 {
					return 5
				}
				return 3
			}
		}
		if part2 && jokers > 1 {
			return 5
		} else if part2 && jokers > 0 {
			return 4
		}
		return 2
	case 4:
		if part2 && jokers > 0 {
			return 3
		}
		return 1
	case 5:
		if part2 && jokers > 0 {
			return 1
		}
		return 0
	}
	return 0
}
