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
	startIdx := 0
	for startIdx < len(bricks) {
		for i := startIdx; i < len(bricks); i++ {
			if bricks[i].begin.z == 1 {
				startIdx++
			} else if canMove, newBrick, bricksTouchingBelow := canMoveDown(bricks[i], bricks[:i]); canMove {
				bricks[i] = newBrick
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
outer:
	for brick := 0; brick < len(bricks); brick++ {
		if list, found := bricksAbove[brick]; found {
			// for each brick above current brick
			for _, idx := range list {
				// if that bricks only has one brick below it (i.e. the current brick), it would fall
				if len(bricksBelow[idx]) <= 1 {
					continue outer
				}
			}
			// otherwise it can be disintegrated
		}
		// also if brick doesn't have bricks above
		canDisintegrate[brick] = true
	}

	var result = len(canDisintegrate)
	fmt.Println("Day 22 Part 1 Result: ", result)

	part2Sum := 0
	for i := 0; i < len(bricks); i++ {
		if _, found := canDisintegrate[i]; !found {
			// count how many blocks would fall
			disintegrated := map[int]bool{i: true}
			checkList := []int{i}
			for len(checkList) > 0 {
				// take first from checklist
				check := checkList[0]
				checkList = checkList[1:]
				// for all bricks above, check ...
				for _, above := range bricksAbove[check] {
					bricksRemoved := 0
					// whether all bricks below it are disintegrated ...
					for _, below := range bricksBelow[above] {
						if _, found := disintegrated[below]; found {
							bricksRemoved++
						}
					}
					if len(bricksBelow[above]) == bricksRemoved {
						// then it would fall
						checkList = append(checkList, above)
						disintegrated[above] = true
					}
				}
			}
			// do not include the brick we are currently checking
			part2Sum += len(disintegrated) - 1
		}
	}

	var result2 = part2Sum
	fmt.Println("Day 22 Part 2 Result: ", result2)
}

func canMoveDown(brick Brick, bricksBelow []Brick) (bool, Brick, []int) {
	brick.begin.z--
	brick.end.z--
	bricksTouching := []int{}
	for i, other := range bricksBelow {
		if bricksOverlap(brick, other) {
			bricksTouching = append(bricksTouching, i)
		}
	}
	return len(bricksTouching) == 0, brick, bricksTouching
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
