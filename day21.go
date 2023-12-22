package main

import (
	"fmt"
)

func day21() {
	grid := getLines("input/21.txt")

	startingPoint := findInGrid(grid, 'S')
	visited := map[Point]bool{startingPoint: true}
	for i := 0; i < 64; i++ {
		visited = oneStep(grid, visited, false)
	}

	var result = len(visited)
	fmt.Println("Day 21 Part 1 Result: ", result)

	visited2 := map[Point]bool{startingPoint: true}
	yValues := [3]int{}
	for i := 1; i <= (len(grid)/2)+2*len(grid); i++ {
		visited2 = oneStep(grid, visited2, true)
		if i == (len(grid) / 2) {
			yValues[0] = len(visited2)
		} else if i == (len(grid)/2)+len(grid) {
			yValues[1] = len(visited2)
		} else if i == (len(grid)/2)+2*len(grid) {
			yValues[2] = len(visited2)
		}
	}

	var result2 = simplifiedLagrange(yValues[:])
	fmt.Println("Day 21 Part 2 Result: ", result2)
}

func simplifiedLagrange(values []int) int {
	a := values[0]/2 - values[1] + values[2]/2
	b := -3*(values[0]/2) + 2*values[1] - values[2]/2
	c := values[0]
	x := 26501365 / 131
	return a*x*x + b*x + c
}

func oneStep(grid []string, visitable map[Point]bool, part2 bool) map[Point]bool {
	newVisitable := map[Point]bool{}

	for point := range visitable {
		for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nextLocation := Point{point.x + dir.x, point.y + dir.y}
			if (part2 && grid[mod(nextLocation.y, len(grid))][mod(nextLocation.x, len(grid[0]))] != '#') ||
				(!part2 && insideGrid(grid, nextLocation) && grid[nextLocation.y][nextLocation.x] != '#') {
				newVisitable[nextLocation] = true
			}
		}
	}
	return newVisitable
}
