package main

import (
	"fmt"
)

func day16() {
	grid := getLines("input/16.txt")
	pos := Point{-1, 0}
	dir := Point{1, 0}
	energized := map[Point]int{}

	walkBeam(grid, pos, dir, energized)
	var result = len(energized)
	fmt.Println("Day 16 Part 1 Result: ", result)

	maxEnergized := len(energized)
	clear(energized)
	for row := range grid {
		dir = Point{1, 0}
		pos = Point{-1, row}
		walkBeam(grid, pos, dir, energized)
		maxEnergized = max(maxEnergized, len(energized))
		clear(energized)

		dir = Point{-1, 0}
		pos = Point{len(grid[0]), row}
		walkBeam(grid, pos, dir, energized)
		maxEnergized = max(maxEnergized, len(energized))
		clear(energized)
	}

	for col := range grid[0] {
		dir = Point{0, 1}
		pos = Point{col, -1}
		walkBeam(grid, pos, dir, energized)
		maxEnergized = max(maxEnergized, len(energized))
		clear(energized)

		dir = Point{0, -1}
		pos = Point{col, len(grid)}
		walkBeam(grid, pos, dir, energized)
		maxEnergized = max(maxEnergized, len(energized))
		clear(energized)
	}

	var result2 = maxEnergized
	fmt.Println("Day 16 Part 2 Result: ", result2)
}

type Whatever struct {
	point, dir Point
}

func walkBeam(grid []string, pos, dir Point, energized map[Point]int) {
	testPoint := Point{pos.x + dir.x, pos.y + dir.y}
	checkNext := []Whatever{{testPoint, dir}}

	for len(checkNext) > 0 {
		whatever := checkNext[0]
		checkNext = checkNext[1:]
		testPoint = whatever.point
		dir = whatever.dir

	mainLoop:
		for insideGrid(grid, testPoint) && energized[testPoint] < 4 {
			energized[testPoint]++
			switch grid[testPoint.y][testPoint.x] {
			case '/':
				switch dir {
				case Point{1, 0}:
					dir = Point{0, -1}
				case Point{-1, 0}:
					dir = Point{0, 1}
				case Point{0, 1}:
					dir = Point{-1, 0}
				case Point{0, -1}:
					dir = Point{1, 0}
				}
			case '\\':
				switch dir {
				case Point{1, 0}:
					dir = Point{0, 1}
				case Point{-1, 0}:
					dir = Point{0, -1}
				case Point{0, 1}:
					dir = Point{1, 0}
				case Point{0, -1}:
					dir = Point{-1, 0}
				}
			case '-':
				switch dir {
				case Point{0, 1}:
					fallthrough
				case Point{0, -1}:
					checkNext = append(checkNext, Whatever{testPoint, Point{1, 0}})
					checkNext = append(checkNext, Whatever{testPoint, Point{-1, 0}})
					break mainLoop
				}
			case '|':
				switch dir {
				case Point{1, 0}:
					fallthrough
				case Point{-1, 0}:
					checkNext = append(checkNext, Whatever{testPoint, Point{0, 1}})
					checkNext = append(checkNext, Whatever{testPoint, Point{0, -1}})
					break mainLoop
				}
			}
			testPoint = Point{testPoint.x + dir.x, testPoint.y + dir.y}
		}
	}
}

func insideGrid(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) &&
		pos.y >= 0 && pos.y < len(grid)
}
