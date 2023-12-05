package main

import (
	"fmt"
	"math"
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
		for _, rangeMaps := range maps {
			for _, rangeMap := range rangeMaps {
				if num >= rangeMap.source && num <= rangeMap.source+rangeMap.length {
					num = (num - rangeMap.source) + rangeMap.destination
					break
				}
			}
		}
		smallest = min(num, smallest)
	}

	var result = smallest
	fmt.Println("Day 05 Part 1 Result: ", result)

	var smallestRange = math.MaxInt32
	for i := 0; i < len(seeds); i += 2 {
		var seedRange = Range{0, seeds[i], seeds[i+1]}
		var sourceRanges = []Range{seedRange}
		for _, rangeMaps := range maps {
			for _, rangeMap := range rangeMaps {
				var length = len(sourceRanges)
				for s := 0; s < length; s++ {
					var sourceRange = sourceRanges[s]
					if itDoes, overlap := overlapsRange(sourceRange, rangeMap); itDoes {
						var split = splitRange(sourceRange, overlap)
						var mapped = split[0]
						mapped.source = (mapped.source - rangeMap.source) + rangeMap.destination
						sourceRanges[s] = mapped
						for j := 1; j <= 2; j++ {
							if split[j].length > 0 {
								sourceRanges = append(sourceRanges, split[j])
							}
						}
					}
				}
			}
		}
		for _, sources := range sourceRanges {
			smallestRange = min(smallestRange, sources.source)
		}
	}

	var result2 = smallestRange
	fmt.Println("Day 05 Part 2 Result: ", result2)
}

func splitRange(seedRange, overlap Range) []Range {
	var ranges []Range
	ranges = append(ranges, overlap)
	ranges = append(ranges, Range{0, min(seedRange.source, overlap.source), abs(overlap.source - seedRange.source)})
	ranges = append(ranges, Range{0, min(seedRange.source+seedRange.length, overlap.source+overlap.length),
		abs((overlap.source + overlap.length) - (seedRange.source + seedRange.length))})
	return ranges
}

func overlapsRange(seedRange, mapRange Range) (bool, Range) {
	left1, right1 := seedRange.source, seedRange.source+seedRange.length
	left2, right2 := mapRange.source, mapRange.source+mapRange.length
	if left1 <= right2 && right1 >= left2 {
		return true, Range{0, max(left1, left2), min(right1, right2) - max(left1, left2)}
	}
	return false, Range{}
}

func parseRangeMaps(lines []string) RangeMaps {
	var maps RangeMaps = RangeMaps{RangeMap{}}
	var section = 0
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			section++
			maps = append(maps, RangeMap{})
			i++
		} else {
			nums := strToNumArray(lines[i])
			maps[section] = append(maps[section], Range{nums[0], nums[1], nums[2]})
		}
	}
	return maps
}
