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
	minHeatLoss := 999999999
	queue := BucketHeap{map[int][]HeatState{}, minHeatLoss}
	queue.insert(HeatState{start, Point{1, 0}, 0}, 0)
	queue.insert(HeatState{start, Point{0, 1}, 0}, 0)
	visited := map[HeatState]int{{start, Point{0, 0}, 0}: 0}

	for len(queue.buckets) > 0 {
		current := queue.popMin()

		if current.point == end {
			if current.streak >= minStreak {
				minHeatLoss = min(minHeatLoss, visited[current])
				break
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
					queue.insert(nextState, totalHeatLoss)
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
				queue.insert(nextState, totalHeatLoss)
			}
		}

		rightDir := dirRight(current.dir)
		nextPoint = Point{current.point.x + rightDir.x, current.point.y + rightDir.y}
		if !skipRight && insideGrid(grid, nextPoint) && current.streak >= minStreak {
			nextState := HeatState{nextPoint, rightDir, 1}
			totalHeatLoss := visited[current] + int(grid[nextState.point.y][nextState.point.x]-'0')
			if val, found := visited[nextState]; !found || val > totalHeatLoss {
				visited[nextState] = totalHeatLoss
				queue.insert(nextState, totalHeatLoss)
			}
		}
	}

	return minHeatLoss
}

type BucketHeap struct {
	buckets   map[int][]HeatState
	minBucket int
}

func (bh *BucketHeap) insert(state HeatState, heatloss int) {
	bh.buckets[heatloss] = append(bh.buckets[heatloss], state)
	if heatloss < bh.minBucket {
		bh.minBucket = heatloss
	}
}

func (bh *BucketHeap) popMin() HeatState {
	minHeatstate := bh.buckets[bh.minBucket][0]
	bh.buckets[bh.minBucket] = bh.buckets[bh.minBucket][1:]

	if len(bh.buckets[bh.minBucket]) == 0 {
		delete(bh.buckets, bh.minBucket)
		bh.minBucket = 999999999
		for bucket := range bh.buckets {
			if bucket < bh.minBucket {
				bh.minBucket = bucket
			}
		}
	}

	return minHeatstate
}
