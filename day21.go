package main

import (
	"fmt"
)

func day21() {
	grid := getLines("input/21.txt")

	startingPoint := findStart(grid)
	visitable := map[Point]bool{startingPoint: true}
	for i := 0; i < 64; i++ {
		oneStep(grid, visitable)
	}

	var result = len(visitable)
	var result2 = 0
	fmt.Println("Day 21 Part 1 Result: ", result)
	fmt.Println("Day 21 Part 2 Result: ", result2)
}

func oneStep(grid []string, visitable map[Point]bool) {
	visitCopy := []Point{}
	for point := range visitable {
		visitCopy = append(visitCopy, point)
	}

	for _, point := range visitCopy {
		delete(visitable, point)
		for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nextLocation := Point{point.x + dir.x, point.y + dir.y}
			if insideGrid(grid, nextLocation) && grid[nextLocation.y][nextLocation.x] != '#' {
				visitable[nextLocation] = true
			}
		}
	}
}
