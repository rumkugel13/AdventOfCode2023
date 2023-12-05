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
	fmt.Println("Day 05 Part 1 Result: ", result)

	var seedRanges = seedsToRangeArray(seeds)
	var smallestRange = math.MaxInt32
	for _, seedRange := range seedRanges {
		var sourceRanges = []Range{seedRange}
		
		for _, rangeMaps := range maps {
			var mappedRanges []Range
			for _, rangeMap := range rangeMaps {
				var length = len(sourceRanges)
				for s := 0; s < length; s++ {
					var r = sourceRanges[s]
					if o, overlap := overlaps(r, rangeMap); o {
						var split = splitRange(r, overlap)
						if split[0].length > 0 {
							sourceRanges = append(sourceRanges, split[0])
						}
						if split[2].length > 0 {
							sourceRanges = append(sourceRanges, split[2])
						}
						var mapped = split[1]
						mapped.source = (mapped.source - rangeMap.source) + rangeMap.destination
						mappedRanges = append(mappedRanges, mapped)
						sourceRanges[s] = Range{}
					} 
				}
			}
			mappedRanges = append(mappedRanges, sourceRanges...)
			clear(sourceRanges)
			sourceRanges = append(sourceRanges, mappedRanges...)
		}
		for _,sources := range sourceRanges {
			if sources.length > 0 {
				smallestRange = min(smallestRange, sources.source)
			}
		}
	}

	var result2 = smallestRange
	fmt.Println("Day 05 Part 2 Result: ", result2)
}

func splitRange(seedRange, overlap Range) []Range {
	var ranges []Range
	ranges = append(ranges, Range{0, min(seedRange.source, overlap.source), abs(overlap.source-seedRange.source)})
	ranges = append(ranges, overlap)
	ranges = append(ranges, Range{0, min(seedRange.source+seedRange.length, overlap.source+overlap.length), 
		abs((overlap.source + overlap.length)-(seedRange.source +seedRange.length))})
	return ranges
}

func overlaps(seedRange, mapRange Range) (bool, Range) {
	left1, right1 := seedRange.source, seedRange.source + seedRange.length
	left2, right2 := mapRange.source, mapRange.source + mapRange.length
	if left1 <= right2 && right1 >= left2 {
		return true, Range{0, max(left1, left2), min(right1, right2) - max(left1, left2)}
	}
	return false, Range{}
}

func seedsToRangeArray(seeds []int) RangeMap {
	var ranges RangeMap
	for i := 0; i < len(seeds); i += 2 {
		var r = Range{seeds[i], seeds[i], seeds[i+1]}
		ranges = append(ranges, r)
	}
	return ranges
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
