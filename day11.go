package main

import (
	"fmt"
)

func day11() {
	universe := getLines("input/11.txt")
	galaxies := findGalaxies(universe)
	emptyRows, emptyCols := emptySpace(universe, galaxies)
	distances, dist2 := getDistances(galaxies, emptyRows, emptyCols)

	var result = distances
	var result2 = dist2
	fmt.Println("Day 11 Part 1 Result: ", result)
	fmt.Println("Day 11 Part 2 Result: ", result2)
}

func findGalaxies(universe []string) map[Point]bool {
	galaxies := map[Point]bool{}
	for y, row := range universe {
		for x, char := range row {
			if char == '#' {
				galaxies[Point{x, y}] = true
			}
		}
	}
	return galaxies
}

func getDistances(galaxies map[Point]bool, emptyRows, emptyCols []int) (int, int) {
	galaxyList := make([]Point, 0, len(galaxies))
	for galaxy := range galaxies {
		galaxyList = append(galaxyList, galaxy)
	}

	distances, distances2 := 0, 0
	for i := 0; i < len(galaxyList); i++ {
		galaxyA := galaxyList[i]
		for j := i + 1; j < len(galaxyList); j++ {
			galaxyB := galaxyList[j]
			minx, miny := min(galaxyA.x, galaxyB.x), min(galaxyA.y, galaxyB.y)
			maxx, maxy := max(galaxyA.x, galaxyB.x), max(galaxyA.y, galaxyB.y)

			expansionX, expansionY := 0, 0
			expansionX2, expansionY2 := 0, 0
			for _, val := range emptyCols {
				if minx < val && val < maxx {
					expansionX++
					expansionX2 += 999_999
				}
			}

			for _, val := range emptyRows {
				if miny < val && val < maxy {
					expansionY++
					expansionY2 += 999_999
				}
			}

			dist := (maxx - minx) + (maxy - miny) + expansionX + expansionY
			dist2 := (maxx - minx) + (maxy - miny) + expansionX2 + expansionY2
			distances += dist
			distances2 += dist2
		}
	}

	return distances, distances2
}

func emptySpace(universe []string, galaxies map[Point]bool) ([]int, []int) {
	emptyRows, emptyCols := []int{}, []int{}
	for i := 0; i < len(universe); i++ {
		if rowEmpty(universe, galaxies, i) {
			emptyRows = append(emptyRows, i)
		}
		if colEmpty(universe, galaxies, i) {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyRows, emptyCols
}

func rowEmpty(universe []string, galaxies map[Point]bool, row int) bool {
	for x := 0; x < len(universe[0]); x++ {
		if _, found := galaxies[Point{x, row}]; found {
			return false
		}
	}
	return true
}

func colEmpty(universe []string, galaxies map[Point]bool, col int) bool {
	for y := 0; y < len(universe); y++ {
		if _, found := galaxies[Point{col, y}]; found {
			return false
		}
	}
	return true
}
