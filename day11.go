package main

import (
	"fmt"
)

func day11() {
	universe := getLines("input/11.txt")
	galaxies := findGalaxies(universe)
	emptyRows, emptyCols := getEmptySpace(universe)
	distances, dist2 := getGalaxyDistances(galaxies, emptyRows, emptyCols)

	var result = distances
	var result2 = dist2
	fmt.Println("Day 11 Part 1 Result: ", result)
	fmt.Println("Day 11 Part 2 Result: ", result2)
}

func findGalaxies(universe []string) []Point {
	galaxies := []Point{}
	for y, row := range universe {
		for x, char := range row {
			if char == '#' {
				galaxies = append(galaxies, Point{x, y})
			}
		}
	}
	return galaxies
}

func getGalaxyDistances(galaxies []Point, emptyRows, emptyCols []int) (int, int) {
	distances, distances2 := 0, 0
	for i := 0; i < len(galaxies); i++ {
		galaxyA := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			galaxyB := galaxies[j]
			dist, dist2 := getGalaxyDistance(galaxyA, galaxyB, emptyRows, emptyCols)
			distances += dist
			distances2 += dist2
		}
	}

	return distances, distances2
}

func getGalaxyDistance(galaxyA, galaxyB Point, emptyRows, emptyCols []int) (int, int) {
	minx, miny := min(galaxyA.x, galaxyB.x), min(galaxyA.y, galaxyB.y)
	maxx, maxy := max(galaxyA.x, galaxyB.x), max(galaxyA.y, galaxyB.y)

	expansionX, expansionX2 := expandSpaceBetween(emptyCols, minx, maxx)
	expansionY, expansionY2 := expandSpaceBetween(emptyRows, miny, maxy)

	dist := (maxx - minx) + (maxy - miny) + expansionX + expansionY
	dist2 := (maxx - minx) + (maxy - miny) + expansionX2 + expansionY2
	return dist, dist2
}

func expandSpaceBetween(emptySpace []int, min, max int) (a, b int) {
	for _, val := range emptySpace {
		if min < val && val < max {
			a++
			b += 999_999
		}
	}
	return
}

func getEmptySpace(universe []string) ([]int, []int) {
	emptyRows, emptyCols := []int{}, []int{}
	for i := 0; i < len(universe); i++ {
		if rowEmpty(universe, i) {
			emptyRows = append(emptyRows, i)
		}
		if colEmpty(universe, i) {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyRows, emptyCols
}

func rowEmpty(universe []string, row int) bool {
	for x := 0; x < len(universe[0]); x++ {
		if universe[row][x] == '#' {
			return false
		}
	}
	return true
}

func colEmpty(universe []string, col int) bool {
	for y := 0; y < len(universe); y++ {
		if universe[y][col] == '#' {
			return false
		}
	}
	return true
}
