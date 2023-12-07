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

	for _, hand := range input {
		handAndBid := strings.Fields(hand)
		handType := getHandType(handAndBid[0])
		bid, _ := strconv.Atoi(handAndBid[1])
		entries = append(entries, TypeAndBid{handAndBid[0], handType, bid})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].handType == entries[j].handType {
			return lessHandType(entries[i].hand, entries[j].hand)
		}
		return entries[i].handType < entries[j].handType
	})

	count := 0
	for i, entry := range entries {
		count += (i + 1) * entry.bid
	}

	var result = count
	var result2 = 0
	fmt.Println("Day 07 Part 1 Result: ", result)
	fmt.Println("Day 07 Part 2 Result: ", result2)
}

func lessHandType(hand1, hand2 string) bool {
	nums := map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
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

func getHandType(hand string) int {
	cards := map[byte]int{}
	for _, card := range hand {
		cards[byte(card)]++
	}

	switch len(cards) {
	case 1:
		return 6
	case 2:
		for _, count := range cards {
			if count == 4 {
				return 5
			}
		}
		return 4
	case 3:
		for _, count := range cards {
			if count == 3 {
				return 3
			}
		}
		return 2
	case 4:
		return 1
	case 5:
		return 0
	}
	return 0
}
