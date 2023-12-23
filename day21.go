package main

import (
	"fmt"
)

func day21() {
	grid := getLines("input/21.txt")

	startingPoint := findInGrid(grid, 'S')
	visited := map[Point]bool{startingPoint: true}
	for i := 0; i < 64; i++ {
		visited = oneStep(grid, visited)
	}

	var result = len(visited)
	fmt.Println("Day 21 Part 1 Result: ", result)

	queue := []Point{startingPoint}
	visitedOdd, visitedEven := map[Point]bool{}, map[Point]bool{startingPoint: true}
	count := make([]int, (len(grid)/2)+2*len(grid)+1)
	for i := 1; i <= (len(grid)/2)+2*len(grid); i++ {
		var currentVisited *map[Point]bool
		if i%2 == 0 {
			currentVisited = &visitedEven
		} else {
			currentVisited = &visitedOdd
		}

		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			point := queue[i]
			for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nextLocation := Point{point.x + dir.x, point.y + dir.y}
				if grid[mod(nextLocation.y, len(grid))][mod(nextLocation.x, len(grid[0]))] != '#' {
					if _, found := (*currentVisited)[nextLocation]; !found {
						(*currentVisited)[nextLocation] = true
						queue = append(queue, nextLocation)
					}
				}
			}
		}
		queue = queue[queueLen:]
		count[i] = len(*currentVisited)
	}

	var result2 = simplifiedLagrange(count[len(grid)/2], count[(len(grid)/2)+len(grid)], count[(len(grid)/2)+2*len(grid)])
	fmt.Println("Day 21 Part 2 Result: ", result2)
}

func simplifiedLagrange(v0, v1, v2 int) int {
	a := v0/2 - v1 + v2/2
	b := -3*(v0/2) + 2*v1 - v2/2
	c := v0
	x := 26501365 / 131
	return a*x*x + b*x + c
}

func oneStep(grid []string, visitable map[Point]bool) map[Point]bool {
	newVisitable := map[Point]bool{}

	for point := range visitable {
		for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nextLocation := Point{point.x + dir.x, point.y + dir.y}
			if insideGrid(grid, nextLocation) && grid[nextLocation.y][nextLocation.x] != '#' {
				newVisitable[nextLocation] = true
			}
		}
	}
	return newVisitable
}
