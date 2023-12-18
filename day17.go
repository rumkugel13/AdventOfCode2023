package main

import (
	"fmt"
	"sort"
)

func day17() {
	grid := getLines("input/17.txt")
	start, end := Point{0, 0}, Point{len(grid[0]) - 1, len(grid) - 1}
	heatLoss := countHeatLoss(grid, start, end)

	var result = heatLoss
	var result2 = 0
	fmt.Println("Day 17 Part 1 Result: ", result)
	fmt.Println("Day 17 Part 2 Result: ", result2)
}

type HeatState struct {
	point, dir   Point
	sameDirCount int
}

func countHeatLoss(grid []string, start, end Point) int {
	pointsToCheck := []HeatState{{start, Point{0, 0}, 0}}
	visited := map[HeatState]int{{start, Point{0, 0}, 0}: 0}
	minHeatLoss := 999999999

	for len(pointsToCheck) > 0 {
		current := pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]
		// fmt.Println("Visiting ", current, getHeatLoss(grid, current.point), " Total ", visited[current])

		if current.point == end {
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

			nextDirCount := 1
			if dir == current.dir {
				nextDirCount = current.sameDirCount + 1
			}
			if nextDirCount > 3 {
				continue
			}
			nextState := HeatState{nextPoint, dir, nextDirCount}
			nextHeatLoss := getHeatLoss(grid, nextPoint)

			val, found := visited[nextState]
			if found && val <= visited[current]+nextHeatLoss {
				continue
			}
			visited[nextState] = visited[current] + nextHeatLoss
			// pointsToCheck = queueInsert(grid, pointsToCheck, nextState)
			pointsToCheck = append(pointsToCheck, nextState)
		}
		// fmt.Println("queue", pointsToCheck)
	}

	return minHeatLoss
	// return 0
}

func queueInsert(grid []string, queue []HeatState, state HeatState) []HeatState {
	i := sort.Search(len(queue), func(i int) bool {
		return getHeatLoss(grid, queue[i].point) > getHeatLoss(grid, state.point)
	})
	queue = append(queue, HeatState{})
	copy(queue[i+1:], queue[i:])
	queue[i] = state
	return queue
}

func getHeatLoss(grid []string, point Point) int {
	return int(grid[point.y][point.x] - '0')
}
