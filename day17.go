package main

import (
	"fmt"
)

func day17() {
	grid := getLines("input/17.txt")
	start, end := Point{0, 0}, Point{len(grid[0]) - 1, len(grid) - 1}

	heatLoss := countHeatLoss(grid, start, end, 1, 3)
	var result = heatLoss
	fmt.Println("Day 17 Part 1 Result: ", result)

	heatLoss2 := countHeatLoss(grid, start, end, 4, 10)
	var result2 = heatLoss2
	fmt.Println("Day 17 Part 2 Result: ", result2)
}

type HeatState struct {
	point, dir Point
	streak     int
}

func countHeatLoss(grid []string, start, end Point, minStreak, maxStreak int) int {
	pointsToCheck := []HeatState{{start, Point{1, 0}, 0}, {start, Point{0, 1}, 0}}
	visited := map[HeatState]int{{start, Point{0, 0}, 0}: 0}
	minHeatLoss := 999999999

	for len(pointsToCheck) > 0 {
		current := pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]

		if current.point == end && current.streak >= minStreak {
			minHeatLoss = min(minHeatLoss, visited[current])
		}

		for _, dir := range [3]Point{current.dir, dirLeft(current.dir), dirRight(current.dir)} {
			nextPoint := Point{current.point.x + dir.x, current.point.y + dir.y}
			if !insideGrid(grid, nextPoint) {
				continue
			}

			totalHeatLoss := visited[current] + int(grid[nextPoint.y][nextPoint.x] - '0')
			nextStreak := 1
			if dir == current.dir {
				nextStreak = current.streak + 1
			}
			if (dir == current.dir && current.streak < maxStreak) ||
				(dir != current.dir && current.streak >= minStreak) {
				nextState := HeatState{nextPoint, dir, nextStreak}
				if val, found := visited[nextState]; !found || val > totalHeatLoss {
					visited[nextState] = totalHeatLoss
					pointsToCheck = append(pointsToCheck, nextState)
				}
			}
		}
	}

	return minHeatLoss
}

func dirLeft(p Point) Point {
	return Point{p.y, -p.x}
}

func dirRight(p Point) Point {
	return Point{-p.y, p.x}
}
