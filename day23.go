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
	fmt.Println("Day 23 Part 1 Result: ", result)

	junctions := getJunctions(grid)
	junctions[start] = true
	junctions[end] = true

	paths := []Path{}
	for junction := range junctions {
		paths = append(paths, getPaths(grid, junction, junctions)...)
	}

	mappedPaths := map[Point][]Path{}
	for _, path := range paths {
		mappedPaths[path.start] = append(mappedPaths[path.start], path)
	}

	var result2 = findLongestPath(grid, mappedPaths, start, end, 0, map[Point]bool{start: true})
	fmt.Println("Day 23 Part 2 Result: ", result2)
}

type Path struct {
	start, end Point
	length     int
}

func getJunctions(grid []string) map[Point]bool {
	junctions := map[Point]bool{}
	for row, line := range grid {
		for col, char := range line {
			if char == '#' {
				continue
			}
			point := Point{col, row}
			neighbours := 0
			for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				next := Point{point.x + dir.x, point.y + dir.y}
				if insideGrid(grid, next) && grid[next.y][next.x] != '#' {
					neighbours++
				}
			}
			if neighbours > 2 {
				junctions[point] = true
			}
		}
	}
	return junctions
}

func getPaths(grid []string, junctionPoint Point, junctions map[Point]bool) []Path {
	paths := []Path{}
	for _, startDir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		currentPoint := Point{junctionPoint.x + startDir.x, junctionPoint.y + startDir.y}
		if insideGrid(grid, currentPoint) && grid[currentPoint.y][currentPoint.x] != '#' {
			paths = append(paths, getPath(grid, junctionPoint, currentPoint, startDir, 1, junctions))
		}
	}

	return paths
}

func getPath(grid []string, pathStart, currentPoint, currentDir Point, pathLength int, junctions map[Point]bool) Path {
	for _, dir := range [3]Point{currentDir, dirLeft(currentDir), dirRight(currentDir)} {
		next := Point{currentPoint.x + dir.x, currentPoint.y + dir.y}
		if grid[next.y][next.x] != '#' {
			if _, found := junctions[next]; found {
				return Path{pathStart, next, pathLength + 1}
			} else {
				return getPath(grid, pathStart, next, dir, pathLength+1, junctions)
			}
		}
	}
	return Path{pathStart, Point{-1, -1}, 0}
}

func findLongestPath(grid []string, paths map[Point][]Path, start, end Point, step int, visited map[Point]bool) int {
	if start == end {
		return step
	}
	
	maxStep := 0
	for _, nextJunction := range paths[start] {
		if _, found := visited[nextJunction.end]; !found {
			visited[nextJunction.end] = true
			maxStep = max(maxStep, findLongestPath(grid, paths, nextJunction.end, end, step+nextJunction.length, visited))
			delete(visited, nextJunction.end)
		}
	}
	return maxStep
}

func walkTrail(grid []string, start, currentDir Point, visited map[Point]int) {
	current := start
	currentStep := visited[current]

	for _, dir := range [3]Point{currentDir, dirLeft(currentDir), dirRight(currentDir)} {
		next := Point{current.x + dir.x, current.y + dir.y}
		if insideGrid(grid, next) && grid[next.y][next.x] != '#' {
			char := grid[next.y][next.x]
			oppositeChar := map[Point]byte{{1, 0}: '<', {-1, 0}: '>', {0, 1}: '^', {0, -1}: 'v'}
			if oppositeChar[dir] == char {
				continue
			}

			if val, found := visited[next]; !found || val < currentStep+1 {
				visited[next] = currentStep + 1
				walkTrail(grid, next, dir, visited)
			}
		}
	}
}
