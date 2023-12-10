package main

import (
	"fmt"
)

func day10() {
	grid := getLines("input/10.txt")

	startingPoint := findStart(grid)
	visited := map[Point]int{startingPoint: 0}
	notChecked := []Point{startingPoint}

	maxDist := 0
	for len(notChecked) > 0 {
		current := notChecked[0]
		notChecked = notChecked[1:]
		next := nextPoints(grid, current)
		for _, point := range next {
			if _, found := visited[point]; !found {
				visited[point] = visited[current] + 1
				maxDist = max(maxDist, visited[current]+1)
				notChecked = append(notChecked, point)
			}
		}
	}

	var result = maxDist
	fmt.Println("Day 10 Part 1 Result: ", result)

	countInside := 0
	for y, row := range grid {
		count := 0
		for x := 0; x < len(row); x++ {
			tile := row[x]
			point := Point{x, y}
			if tile == 'S' {
				tile = findStartTile(point, grid)
			}
			if _, part := visited[point]; part {
				if tile == '|' {
					count++
				} else if tile != '-' && tile != '.' {
					if tile == 'L' {
						for x++; x < len(row) && row[x] == '-'; x++ {
						}
						if x < len(row) && row[x] == '7' {
							count++
						}
					}
					if tile == 'F' {
						for x++; x < len(row) && row[x] == '-'; x++ {
						}
						if x < len(row) && row[x] == 'J' {
							count++
						}
					}
				}
			} else if count%2 == 1 {
				countInside++
			}
		}
	}

	var result2 = countInside
	fmt.Println("Day 10 Part 2 Result: ", result2)
}

func findStart(grid []string) Point {
	for y, row := range grid {
		for x, col := range row {
			if byte(col) == 'S' {
				return Point{x, y}
			}
		}
	}
	return Point{}
}

func findStartTile(start Point, grid []string) byte {
	points := nextPoints(grid, start)
	minx, maxx, miny, maxy := min(points[0].x, points[1].x), max(points[0].x, points[1].x), min(points[0].y, points[1].y), max(points[0].y, points[1].y)
	if points[0].x == points[1].x {
		return '|'
	} else if points[0].y == points[1].y {
		return '-'
	} else if minx < start.x && miny < start.y {
		return 'J'
	} else if maxx > start.x && maxy > start.y {
		return 'F'
	} else if maxx > start.x && miny < start.y {
		return 'L'
	} else if minx < start.x && maxy > start.y {
		return '7'
	}
	return '.'
}

func nextPoints(grid []string, p Point) []Point {
	points := []Point{}
	switch grid[p.y][p.x] {
	case '|':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x, p.y - 1})
	case '-':
		points = append(points, Point{p.x + 1, p.y})
		points = append(points, Point{p.x - 1, p.y})
	case 'L':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x + 1, p.y})
	case 'J':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x - 1, p.y})
	case '7':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x - 1, p.y})
	case 'F':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x + 1, p.y})
	case '.':
	case 'S':
		down, right, up, left := grid[p.y+1][p.x], grid[p.y][p.x+1], grid[p.y-1][p.x], grid[p.y][p.x-1]
		if down == '|' || down == 'L' || down == 'J' {
			points = append(points, Point{p.x, p.y + 1})
		}
		if right == '-' || right == '7' || right == 'J' {
			points = append(points, Point{p.x + 1, p.y})
		}
		if up == '|' || up == '7' || up == 'F' {
			points = append(points, Point{p.x, p.y - 1})
		}
		if left == '-' || left == 'L' || left == 'F' {
			points = append(points, Point{p.x - 1, p.y})
		}
	}
	return points
}
