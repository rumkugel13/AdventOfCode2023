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
	fmt.Println("Day 18 Part 1 Result: ", result)

	nodes := make([]Point, 0, len(lines))
	nodes = append(nodes, Point{0, 0})
	point = Point{0, 0}
	for _, line := range lines {
		data := strings.Fields(line)
		color := data[2][2 : len(data[2])-1]
		hexNum := color[:5]
		distance, _ := (strconv.ParseInt(hexNum, 16, 0))
		dir := color[5]
		switch dir {
		case '0':
			point.x += int(distance)
		case '1':
			point.y += int(distance)
		case '2':
			point.x -= int(distance)
		case '3':
			point.y -= int(distance)
		}
		nodes = append(nodes, point)
	}

	var result2 = findAreaShoelace(nodes) + 1
	fmt.Println("Day 18 Part 2 Result: ", result2)
}

func findAreaShoelace(nodes []Point) int {
	area := 0

	for i := 0; i < len(nodes); i++ {
		pointA, pointB := nodes[i], nodes[(i+1)%(len(nodes))]
		area += (pointA.x * pointB.y) - (pointB.x * pointA.y) + max(abs(pointA.x-pointB.x), abs(pointA.y-pointB.y))
	}

	return area / 2
}

func floodFill(trench map[Point]string, point Point) {
	trench[point] = ""
	for _, p := range [4]Point{{point.x, point.y - 1}, {point.x, point.y + 1}, {point.x - 1, point.y}, {point.x + 1, point.y}} {
		if _, contains := trench[p]; !contains {
			floodFill(trench, p)
		}
	}
}
