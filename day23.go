package main

import (
	"fmt"
)

func day23() {
	grid := getLines("input/23.txt")
	start, end := Point{1, 0}, Point{len(grid[0]) - 2, len(grid) - 1}

	visited := map[Point]int{start: 0}
	currentDir := Point{0, 1}
	walkTrail(grid, start, currentDir, visited)

	var result = visited[end]
	var result2 = 0
	fmt.Println("Day 23 Part 1 Result: ", result)
	fmt.Println("Day 23 Part 2 Result: ", result2)
}

func walkTrail(grid []string, start, currentDir Point, visited map[Point]int) {
	current := start
	currentStep := visited[current]

	for _, dir := range [3]Point{currentDir, dirLeft(currentDir), dirRight(currentDir)} {
		next := Point{current.x + dir.x, current.y + dir.y}
		if insideGrid(grid, next) && grid[next.y][next.x] != '#' {
			char := grid[next.y][next.x]
			switch dir {
			case Point{1, 0}:
				if char == '<' {
					continue
				}
			case Point{-1, 0}:
				if char == '>' {
					continue
				}
			case Point{0, 1}:
				if char == '^' {
					continue
				}
			case Point{0, -1}:
				if char == 'v' {
					continue
				}
			}

			if val, found := visited[next]; !found || val < currentStep+1 {
				visited[next] = currentStep + 1
				walkTrail(grid, next, dir, visited)
			}
		}
	}
}
