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

		nextState := HeatState{Point{current.point.x + current.dir.x, current.point.y + current.dir.y}, current.dir, current.streak + 1}
		if insideGrid(grid, nextState.point) && current.streak < maxStreak {
			totalHeatLoss := visited[current] + int(grid[nextState.point.y][nextState.point.x]-'0')
			if val, found := visited[nextState]; !found || val > totalHeatLoss {
				visited[nextState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, nextState)
			}
		}

		leftDir := dirLeft(current.dir)
		nextState = HeatState{Point{current.point.x + leftDir.x, current.point.y + leftDir.y}, leftDir, 1}
		if insideGrid(grid, nextState.point) && current.streak >= minStreak {
			totalHeatLoss := visited[current] + int(grid[nextState.point.y][nextState.point.x]-'0')
			if val, found := visited[nextState]; !found || val > totalHeatLoss {
				visited[nextState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, nextState)
			}
		}

		rightDir := dirRight(current.dir)
		nextState = HeatState{Point{current.point.x + rightDir.x, current.point.y + rightDir.y}, rightDir, 1}
		if insideGrid(grid, nextState.point) && current.streak >= minStreak {
			totalHeatLoss := visited[current] + int(grid[nextState.point.y][nextState.point.x]-'0')
			if val, found := visited[nextState]; !found || val > totalHeatLoss {
				visited[nextState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, nextState)
			}
		}
	}

	return minHeatLoss
}
