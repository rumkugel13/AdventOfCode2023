package main

import (
	"fmt"
	"strings"
)

func day24() {
	lines := getLines("input/24.txt")
	hailStones := parseHailstones(lines)

	areaMin, areaMax := float64(200000000000000), float64(400000000000000)
	// areaMin, areaMax := float64(7), float64(27)
	intersectCount := 0
	for i := 0; i < len(hailStones)-1; i++ {
		for j := i + 1; j < len(hailStones); j++ {
			a, b := hailStones[i], hailStones[j]
			if point, does := hailstonesIntersect(a, b); does {
				if point.x >= areaMin && point.x <= areaMax &&
					point.y >= areaMin && point.y <= areaMax {
					dx := point.x - a.pos.x
					dy := point.y - a.pos.y
					if (dx > 0) == (a.vel.x > 0) && (dy > 0) == (a.vel.y > 0) {
						dx = point.x - b.pos.x
						dy = point.y - b.pos.y
						if (dx > 0) == (b.vel.x > 0) && (dy > 0) == (b.vel.y > 0) {
							intersectCount++
						}
					}
				}
			}
		}
	}

	var result = intersectCount
	var result2 = 0
	fmt.Println("Day 24 Part 1 Result: ", result)
	fmt.Println("Day 24 Part 2 Result: ", result2)
}

type Vector3 struct {
	x, y, z float64
}

type Vector2 struct {
	x, y float64
}

type Hailstone struct {
	pos, vel Vector3
}

func hailstonesIntersect(a, b Hailstone) (Vector2, bool) {
	a2 := Vector2{a.vel.x, a.vel.y}
	b2 := Vector2{b.vel.x, b.vel.y}
	d2 := Vector2{b.pos.x - a.pos.x, b.pos.y - a.pos.y}

	det := vectorCross(a2, b2)
	// parallel
	if det == 0 {
		return Vector2{-1, -1}, false
	}

	u := vectorCross(d2, b2) / det
	return Vector2{a.pos.x + a.vel.x*u, a.pos.y + a.vel.y*u}, true
}

func vectorCross(a, b Vector2) float64 {
	return (a.x * b.y) - (a.y * b.x)
}

func parseHailstones(lines []string) []Hailstone {
	hailStones := make([]Hailstone, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, " @ ")
		coords := commaSepToIntArr(split[0])
		vels := commaSepToIntArr(split[1])
		hailStone := Hailstone{Vector3{float64(coords[0]), float64(coords[1]), float64(coords[2])}, Vector3{float64(vels[0]), float64(vels[1]), float64(vels[2])}}
		hailStones = append(hailStones, hailStone)
	}
	return hailStones
}
