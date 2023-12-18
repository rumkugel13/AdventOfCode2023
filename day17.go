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
	point, dir   Point
	sameDirCount int
}

func countHeatLoss(grid []string, start, end Point, minStreak, maxStreak int) int {
	pointsToCheck := []HeatState{{start, Point{1, 0}, 0}, {start, Point{0, 1}, 0}}
	visited := map[HeatState]int{{start, Point{0, 0}, 0}: 0}
	minHeatLoss := 999999999

	for len(pointsToCheck) > 0 {
		current := pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]
		// fmt.Println("Visiting ", current, getHeatLoss(grid, current.point), " Total ", visited[current])

		if current.point == end && current.sameDirCount >= minStreak {
			minHeatLoss = min(minHeatLoss, visited[current])
			// return visited[current]
		}

		for _, dir := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			// opposite direction
			if dir.x == -current.dir.x && dir.y == -current.dir.y {
				continue
			}

			nextPoint := Point{current.point.x + dir.x, current.point.y + dir.y}
			if !insideGrid(grid, nextPoint) {
				continue
			}

			totalHeatLoss := visited[current] + getHeatLoss(grid, nextPoint)
			nextDirCount := 1
			if dir == current.dir {
				nextDirCount = current.sameDirCount + 1
				if nextDirCount <= maxStreak {
					nextState := HeatState{nextPoint, dir, nextDirCount}

					val, found := visited[nextState]
					if found && val <= totalHeatLoss {
						continue
					}
					visited[nextState] = totalHeatLoss
					// pointsToCheck = queueInsert(grid, pointsToCheck, nextState)
					pointsToCheck = append(pointsToCheck, nextState)
					continue
				}
			}
			if current.sameDirCount < minStreak || nextDirCount > maxStreak {
				continue
			}
			nextState := HeatState{nextPoint, dir, nextDirCount}

			val, found := visited[nextState]
			if found && val <= totalHeatLoss {
				continue
			}
			visited[nextState] = totalHeatLoss
			// pointsToCheck = queueInsert(grid, pointsToCheck, nextState)
			pointsToCheck = append(pointsToCheck, nextState)
		}
		// fmt.Println("queue", pointsToCheck)
	}

	return minHeatLoss
	// return 0
}

func getHeatLoss(grid []string, point Point) int {
	return int(grid[point.y][point.x] - '0')
}
