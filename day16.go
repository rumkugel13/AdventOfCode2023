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
	dirToNum := map[Point]int{{1, 0}: 1, {0, 1}: 2, {-1, 0}: 4, {0, -1}: 8}
	energized := make(map[Point]int, len(grid)*len(grid[0]))
	testTile := Point{pos.x + dir.x, pos.y + dir.y}
	stillToCheck := []TileAndDir{{testTile, dir}}

	for len(stillToCheck) > 0 {
		testTile, dir = stillToCheck[0].tile, stillToCheck[0].dir
		stillToCheck = stillToCheck[1:]

	mainLoop:
		for insideGrid(grid, testTile) && (energized[testTile]&dirToNum[dir]) == 0 {
			energized[testTile] += dirToNum[dir]
			switch grid[testTile.y][testTile.x] {
			case '/':
				if dir.x == 0 {
					dir = dirRight(dir)
				} else {
					dir = dirLeft(dir)
				}
			case '\\':
				if dir.y == 0 {
					dir = dirRight(dir)
				} else {
					dir = dirLeft(dir)
				}
			case '-':
				if dir.x == 0 {
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{1, 0}})
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{-1, 0}})
					break mainLoop
				}
			case '|':
				if dir.y == 0 {
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{0, 1}})
					stillToCheck = append(stillToCheck, TileAndDir{testTile, Point{0, -1}})
					break mainLoop
				}
			}
			testTile = Point{testTile.x + dir.x, testTile.y + dir.y}
		}
	}

	return len(energized)
}
