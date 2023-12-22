package main

import (
	"fmt"
	"sort"
	"strings"
)

type Cube struct {
	x, y, z int
}

type Brick struct {
	begin, end Cube
}

func day22() {
	lines := getLines("input/22.txt")
	bricks := getBricks(lines)
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].begin.z < bricks[j].begin.z
	})

	bricksBelow := make(map[int][]int, len(lines))
	bricksAbove := make(map[int][]int, len(lines))
	brickMoved := true
	startIdx := 0
	for brickMoved {
		brickMoved = false
		for i := startIdx; i < len(bricks); i++ {
			if canMove, newBrick, bricksTouchingBelow := canMoveDown(bricks[i], bricks[:i]); canMove {
				bricks[i] = newBrick
				brickMoved = true
				break
			} else {
				bricksBelow[i] = bricksTouchingBelow
				for _, brick := range bricksTouchingBelow {
					bricksAbove[brick] = append(bricksAbove[brick], i)
				}
				startIdx++
			}
		}
	}

	canDisintegrate := map[int]bool{}
	for brick := 0; brick < len(bricks); brick++ {
		list, found := bricksAbove[brick]
		if found {
			unique := make(map[int]bool, len(list))
			for _, idx := range list {
				unique[idx] = true
			}
			mightDisintegrate := true
			for idx := range unique {
				if len(bricksBelow[idx]) <= 1 {
					mightDisintegrate = false
				}
			}
			if mightDisintegrate {
				canDisintegrate[brick] = true
			}
		} else {
			canDisintegrate[brick] = true
		}
	}

	var result = len(canDisintegrate)
	var result2 = 0
	fmt.Println("Day 22 Part 1 Result: ", result)
	fmt.Println("Day 22 Part 2 Result: ", result2)
}

func canMoveDown(brick Brick, bricksBelow []Brick) (bool, Brick, []int) {
	if brick.begin.z == 1 {
		return false, brick, []int{}
	}
	brick.begin.z--
	brick.end.z--
	bricksTouching := []int{}
	canMove := true
	for i, other := range bricksBelow {
		if bricksOverlap(brick, other) {
			canMove = false
			bricksTouching = append(bricksTouching, i)
		}
	}
	return canMove, brick, bricksTouching
}

func bricksOverlap(a, b Brick) bool {
	return (a.begin.x <= b.end.x && a.end.x >= b.begin.x) &&
		(a.begin.y <= b.end.y && a.end.y >= b.begin.y) &&
		(a.begin.z <= b.end.z && a.end.z >= b.begin.z)
}

func getBricks(lines []string) []Brick {
	bricks := make([]Brick, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, "~")
		begin, end := commaSepToIntArr(strings.Split(split[0], ",")), commaSepToIntArr(strings.Split(split[1], ","))
		bricks = append(bricks, Brick{Cube{begin[0], begin[1], begin[2]}, Cube{end[0], end[1], end[2]}})
	}
	return bricks
}
