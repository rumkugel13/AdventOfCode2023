package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day18() {
	lines := getLines("input/18.txt")

	trench := map[Point]string{}
	point := Point{0, 0}

	for _, line := range lines {
		data := strings.Fields(line)
		dir := data[0]
		length, _ := strconv.Atoi(data[1])
		color := data[2][1 : len(data[2])-2]

		for i := 0; i < length; i++ {
			switch dir {
			case "U":
				point.y--
			case "D":
				point.y++
			case "L":
				point.x--
			case "R":
				point.x++
			}
			trench[point] = color
		}
	}

	floodFill(trench, Point{1, 1})

	var result = len(trench)
	var result2 = 0
	fmt.Println("Day 18 Part 1 Result: ", result)
	fmt.Println("Day 18 Part 2 Result: ", result2)
}

func floodFill(trench map[Point]string, point Point) {
	trench[point] = ""
	next := []Point{{point.x, point.y - 1}, {point.x, point.y + 1}, {point.x - 1, point.y}, {point.x + 1, point.y}}
	for _, p := range next {
		if _, contains := trench[p]; !contains {
			floodFill(trench, p)
		}
	}
}
