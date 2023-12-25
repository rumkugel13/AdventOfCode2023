package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("AoC 2023")

	funcs := []func(){day01, day02, day03, day04, day05, day06, day07, day08, day09, day10,
		day11, day12, day13, day14, day15, day16, day17, day18, day19, day20, day21, day22, day23, day24, day25}

	for _, f := range funcs {
		then := time.Now()
		f()
		fmt.Println("Took", time.Since(then))
	}
}
