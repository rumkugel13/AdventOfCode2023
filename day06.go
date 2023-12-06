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

	winningCounts := make([]int, len(times))
	for race := 0; race < len(times); race++ {
		totalTime := times[race]
		totalDistance := distances[race]
		for time := 1; time < totalTime; time++ {
			if getDistance(time, totalTime) > totalDistance {
				winningCounts[race]++
			}
		}
	}

	var result = 1
	for _,count := range winningCounts {
		result *= count
	}
	fmt.Println("Day 06 Part 1 Result: ", result)

	actualTime,_ := strconv.Atoi(strings.ReplaceAll(input[0][5:], " ", ""))
	actualDist,_ := strconv.Atoi(strings.ReplaceAll(input[1][9:], " ", ""))

	var winningCount = 0
	for time := 1; time < actualTime; time++ {
		if getDistance(time, actualTime) > actualDist {
			winningCount++
		}
	}

	var result2 = winningCount
	fmt.Println("Day 06 Part 2 Result: ", result2)
}

func getDistance(holdTime, totalTime int) int {
	remainingTime := totalTime - holdTime
	return holdTime * remainingTime
}