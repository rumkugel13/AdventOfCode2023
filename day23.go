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

	paths := getPaths(grid, junctions, start, end)
	visited2 := make([]bool, len(junctions))
	visited2[paths[start][0].index] = true
	var result2 = findLongestPath(grid, paths, start, end, 0, visited2)
	fmt.Println("Day 23 Part 2 Result: ", result2)
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

type PathTo struct {
	end           Point
	length, index int
}

func getPaths(grid []string, junctions map[Point]bool, start, end Point) map[Point][]PathTo {
	paths := map[Point][]PathTo{}
	junctionIndex := 0
	for junctionPoint := range junctions {
		blocked := -1
		for i, startDir := range [4]Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
			currentPoint := Point{junctionPoint.x + startDir.x, junctionPoint.y + startDir.y}
			if insideGrid(grid, currentPoint) && grid[currentPoint.y][currentPoint.x] != '#' {
				path := getPath(grid, junctionPoint, currentPoint, startDir, 1, junctions)
				path.index = junctionIndex
				paths[junctionPoint] = append(paths[junctionPoint], path)
			} else {
				blocked = i
			}
		}
		if blocked != -1 && junctionPoint != start && junctionPoint != end {
			removeIndex := 0
			if blocked == 2 {
				// note: if up is blocked, then left is at index 0, otherwise at 1
				removeIndex = 1
			} 
			paths[junctionPoint] = append(paths[junctionPoint][:removeIndex], paths[junctionPoint][removeIndex+1:]...)
		}
		junctionIndex++
	}
	return paths
}

func getPath(grid []string, pathStart, currentPoint, currentDir Point, pathLength int, junctions map[Point]bool) PathTo {
	for _, dir := range [3]Point{currentDir, dirLeft(currentDir), dirRight(currentDir)} {
		next := Point{currentPoint.x + dir.x, currentPoint.y + dir.y}
		if grid[next.y][next.x] != '#' {
			if _, found := junctions[next]; found {
				return PathTo{next, pathLength + 1, 0}
			} else {
				return getPath(grid, pathStart, next, dir, pathLength+1, junctions)
			}
		}
	}
	return PathTo{Point{-1, -1}, 0, 0}
}

func findLongestPath(grid []string, paths map[Point][]PathTo, start, end Point, step int, visited []bool) int {
	maxStep := 0
	for _, path := range paths[start] {
		if path.end == end {
			return step + path.length
		}
		index := paths[path.end][0].index
		if !visited[index] {
			visited[index] = true
			maxStep = max(maxStep, findLongestPath(grid, paths, path.end, end, step+path.length, visited))
			visited[index] = false
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
