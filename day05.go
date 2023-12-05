package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Range struct {
	destination, source, length int
}

type RangeMap = []Range
type RangeMaps = []RangeMap

func day05() {
	lines := getLines("input/05.txt")

	var seeds = strToNumArray(lines[0][6:])
	var maps RangeMaps = parseRangeMaps(lines)

	var smallest = math.MaxInt32
	for _, seed := range seeds {
		var num = seed
		for i := 0; i < len(maps); i++ {
			var rangeMaps = maps[i]
			for _, m := range rangeMaps {
				if num >= m.source && num <= m.source+m.length {
					num = (num - m.source) + m.destination
					break
				}
			}
		}
		smallest = min(num, smallest)
	}

	var result = smallest
	var result2 = 0
	fmt.Println("Day 05 Part 1 Result: ", result)
	fmt.Println("Day 05 Part 2 Result: ", result2)
}

func parseRangeMaps(lines []string) RangeMaps {
	var maps RangeMaps = RangeMaps{RangeMap{}}
	var section = 0
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			section++
			maps = append(maps, RangeMap{})
			i += 2
			if i >= len(lines) {
				break
			}
		}
		maps[section] = append(maps[section], parseRange(lines[i]))
	}
	return maps
}

func parseRange(line string) Range {
	var r Range
	split := strings.Fields(line)
	r.destination, _ = strconv.Atoi(split[0])
	r.source, _ = strconv.Atoi(split[1])
	r.length, _ = strconv.Atoi(split[2])
	return r
}
