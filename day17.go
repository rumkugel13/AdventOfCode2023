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

		straightState := HeatState{Point{current.point.x + current.dir.x, current.point.y + current.dir.y}, current.dir, current.streak + 1}
		if insideGrid(grid, straightState.point) && current.streak < maxStreak {
			totalHeatLoss := visited[current] + int(grid[straightState.point.y][straightState.point.x] - '0')
			if val, found := visited[straightState]; !found || val > totalHeatLoss {
				visited[straightState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, straightState)
			}
		}

		leftDir := dirLeft(current.dir)
		leftState := HeatState{Point{current.point.x + leftDir.x, current.point.y + leftDir.y}, leftDir, 1}
		if insideGrid(grid, leftState.point) && current.streak >= minStreak {
			totalHeatLoss := visited[current] + int(grid[leftState.point.y][leftState.point.x] - '0')
			if val, found := visited[leftState]; !found || val > totalHeatLoss {
				visited[leftState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, leftState)
			}
		}

		rightDir := dirRight(current.dir)
		rightState := HeatState{Point{current.point.x + rightDir.x, current.point.y + rightDir.y}, rightDir, 1}
		if insideGrid(grid, rightState.point) && current.streak >= minStreak {
			totalHeatLoss := visited[current] + int(grid[rightState.point.y][rightState.point.x] - '0')
			if val, found := visited[rightState]; !found || val > totalHeatLoss {
				visited[rightState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, rightState)
			}
		}
	}

	return minHeatLoss
}
