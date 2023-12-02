package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day02() {
	lines := getLines("02.txt")
	var sum, powersum int

	for _, line := range lines {
		const rmax, gmax, bmax = 12, 13, 14
		possible := true
		var rmin, gmin, bmin = 0, 0, 0

		split := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Fields(split[0])[1])

		split = strings.Split(split[1], ";")
		for _, set := range split {
			draw := strings.Split(set, ",")
			for _, colorDraw := range draw {
				colorAmount := strings.Fields(strings.TrimSpace(colorDraw))
				amount, _ := strconv.Atoi(colorAmount[0])
				color := colorAmount[1]
				switch color {
				case "red":
					if amount > rmax {
						possible = false
					}
					if amount > rmin {
						rmin = amount
					}
				case "green":
					if amount > gmax {
						possible = false
					}
					if amount > gmin {
						gmin = amount
					}
				case "blue":
					if amount > bmax {
						possible = false
					}
					if amount > bmin {
						bmin = amount
					}
				}
			}
		}

		if possible {
			sum += id
		}

		powersum += rmin * gmin * bmin
	}

	fmt.Println("Day 02 Part 1 Result: ", sum)
	fmt.Println("Day 02 Part 2 Result: ", powersum)
}
