package main

import (
	"fmt"
)

func day14() {
	lines := getLines("input/14.txt")
	grid := toByteGrid(lines)

	rollNorth(grid)
	part1 := countRocks(grid)

	rollWest(grid)
	rollSouth(grid)
	rollEast(grid)
	cycleLoad := countRocks(grid)
	previous := map[string]int{gridToString(grid): 1}
	targetCycle := 1000000000

	for cycle := 2; cycle <= targetCycle; cycle++ {
		rollNorth(grid)
		rollWest(grid)
		rollSouth(grid)
		rollEast(grid)
		cycleLoad = countRocks(grid)
		hash := gridToString(grid)
		if cycleStart, found := previous[hash]; found {
			cycleLength := cycle - cycleStart
			for ; cycle <= targetCycle && ((targetCycle-cycle)%cycleLength) != 0; cycle++ {
				rollNorth(grid)
				rollWest(grid)
				rollSouth(grid)
				rollEast(grid)
				cycleLoad = countRocks(grid)
			}
			break
		}
		previous[hash] = cycle
	}

	var result = part1
	var result2 = cycleLoad
	fmt.Println("Day 14 Part 1 Result: ", result)
	fmt.Println("Day 14 Part 2 Result: ", result2)
}

func gridToString(grid [][]byte) string {
	data := make([]byte, 0, len(grid)*len(grid[0]))
	for _, line := range grid {
		data = append(data, line...)
	}
	return string(data)
}

func countRocks(grid [][]byte) int {
	result := 0
	i := 1
	for row := len(grid) - 1; row >= 0; row-- {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'O' {
				result += i
			}
		}
		i++
	}
	return result
}

func rollNorth(grid [][]byte) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			val := grid[row][col]
			if val == 'O' {
				for i := row - 1; i >= 0; i-- {
					if grid[i][col] == '.' {
						grid[i+1][col] = '.'
						grid[i][col] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func rollWest(grid [][]byte) {
	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid); row++ {
			val := grid[row][col]
			if val == 'O' {
				for i := col - 1; i >= 0; i-- {
					if grid[row][i] == '.' {
						grid[row][i+1] = '.'
						grid[row][i] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func rollSouth(grid [][]byte) {
	for row := len(grid) - 1; row >= 0; row-- {
		for col := 0; col < len(grid[row]); col++ {
			val := grid[row][col]
			if val == 'O' {
				for i := row + 1; i < len(grid[row]); i++ {
					if grid[i][col] == '.' {
						grid[i-1][col] = '.'
						grid[i][col] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func rollEast(grid [][]byte) {
	for col := len(grid[0]) - 1; col >= 0; col-- {
		for row := len(grid) - 1; row >= 0; row-- {
			val := grid[row][col]
			if val == 'O' {
				for i := col + 1; i < len(grid[0]); i++ {
					if grid[row][i] == '.' {
						grid[row][i-1] = '.'
						grid[row][i] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
}

func toByteGrid(grid []string) [][]byte {
	bytes := make([][]byte, 0, len(grid))
	for _, line := range grid {
		bytes = append(bytes, []byte(line))
	}
	return bytes
}
