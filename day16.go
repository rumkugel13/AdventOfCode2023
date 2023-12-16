package main

import (
	"fmt"
)

func day16() {
	grid := getLines("input/16.txt")
	pos := Point{-1, 0}
	dir := Point{1, 0}

	var result = walkBeam(grid, pos, dir)
	fmt.Println("Day 16 Part 1 Result: ", result)

	maxEnergized := 0
	for row := range grid {
		maxEnergized = max(maxEnergized, walkBeam(grid, Point{-1, row}, Point{1, 0}))
		maxEnergized = max(maxEnergized, walkBeam(grid, Point{len(grid[0]), row}, Point{-1, 0}))
	}

	for col := range grid[0] {
		maxEnergized = max(maxEnergized, walkBeam(grid, Point{col, -1}, Point{0, 1}))
		maxEnergized = max(maxEnergized, walkBeam(grid, Point{col, len(grid)}, Point{0, -1}))
	}

	var result2 = maxEnergized
	fmt.Println("Day 16 Part 2 Result: ", result2)
}

type TileAndDir struct {
	tile, dir Point
}

func walkBeam(grid []string, pos, dir Point) int {
	energized := make(map[Point]int, len(grid)*len(grid[0]))
	visited := make(map[TileAndDir]bool, len(grid)*len(grid[0]))
	testTile := Point{pos.x + dir.x, pos.y + dir.y}
	stillToCheck := []TileAndDir{{testTile, dir}}

	for len(stillToCheck) > 0 {
		checkNext := stillToCheck[0]
		stillToCheck = stillToCheck[1:]
		testTile, dir = checkNext.tile, checkNext.dir

	mainLoop:
		for insideGrid(grid, testTile) && !visited[checkNext] {
			energized[testTile]++
			visited[checkNext] = true
			switch grid[testTile.y][testTile.x] {
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
				case Point{0, 1}, Point{0, -1}:
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{1, 0}})
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{-1, 0}})
					break mainLoop
				}
			case '|':
				switch dir {
				case Point{1, 0}, Point{-1, 0}:
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{0, 1}})
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{0, -1}})
					break mainLoop
				}
			}
			testTile = Point{testTile.x + dir.x, testTile.y + dir.y}
			checkNext.tile, checkNext.dir = testTile, dir
		}
	}

	return len(energized)
}

func insideGrid(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}
