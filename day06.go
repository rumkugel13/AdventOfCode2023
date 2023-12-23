package main

import (
	"fmt"
	"strings"
	"strconv"
)

func day06() {
	input := getLines("input/06.txt")
	times := strToNumArray(input[0][5:])
	distances := strToNumArray(input[1][9:])

	winners := 1
	for race := 0; race < len(times); race++ {
		winCount := 0
		for time := 1; time < times[race]; time++ {
			if time * (times[race] - time) > distances[race] {
				winCount++
			}
		}
		winners *= winCount
	}

	var result = winners
	fmt.Println("Day 06 Part 1 Result: ", result)

	actualTime,_ := strconv.Atoi(strings.ReplaceAll(input[0][5:], " ", ""))
	actualDist,_ := strconv.Atoi(strings.ReplaceAll(input[1][9:], " ", ""))

	start, end := 0,0
	for time := 1; time < actualTime; time++ {
		if time * (actualTime - time) > actualDist {
			start = time
			break
		}
	}
	for time := actualTime; time > 0; time-- {
		if time * (actualTime - time) > actualDist {
			end = time
			break
		}
	}

	var result2 = (end - start) + 1
	fmt.Println("Day 06 Part 2 Result: ", result2)
}