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

		if current.point == end {
			if current.streak >= minStreak {
				minHeatLoss = min(minHeatLoss, visited[current])
			}
			continue
		}

		skipLeft, skipRight := false, false
		nextPoint := Point{current.point.x + current.dir.x, current.point.y + current.dir.y}
		if insideGrid(grid, nextPoint) {
			if current.streak < maxStreak {
				nextState := HeatState{nextPoint, current.dir, current.streak + 1}
				totalHeatLoss := visited[current] + int(grid[nextState.point.y][nextState.point.x]-'0')
				if val, found := visited[nextState]; !found || val > totalHeatLoss {
					visited[nextState] = totalHeatLoss
					pointsToCheck = append(pointsToCheck, nextState)
				}
			}
		} else {
			if current.dir.x == 1 || current.dir.y == -1 {
				skipLeft = true
			} else if current.dir.y == 1 || current.dir.x == -1 {
				skipRight = true
			}
		}

		leftDir := dirLeft(current.dir)
		nextPoint = Point{current.point.x + leftDir.x, current.point.y + leftDir.y}
		if !skipLeft && insideGrid(grid, nextPoint) && current.streak >= minStreak {
			nextState := HeatState{nextPoint, leftDir, 1}
			totalHeatLoss := visited[current] + int(grid[nextState.point.y][nextState.point.x]-'0')
			if val, found := visited[nextState]; !found || val > totalHeatLoss {
				visited[nextState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, nextState)
			}
		}

		rightDir := dirRight(current.dir)
		nextPoint = Point{current.point.x + rightDir.x, current.point.y + rightDir.y}
		if !skipRight && insideGrid(grid, nextPoint) && current.streak >= minStreak {
			nextState := HeatState{nextPoint, rightDir, 1}
			totalHeatLoss := visited[current] + int(grid[nextState.point.y][nextState.point.x]-'0')
			if val, found := visited[nextState]; !found || val > totalHeatLoss {
				visited[nextState] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, nextState)
			}
		}
	}

	return minHeatLoss
}
