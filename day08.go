package main

import (
	"fmt"
)

type LRNode struct {
	left, right string
}

func day08() {
	input := getLines("input/08.txt")
	instructions := input[0]

	nodes := make(map[string]LRNode, len(input[2:]))
	for _, node := range input[2:] {
		nodes[node[0:3]] = LRNode{node[7:10], node[12:15]}
	}

	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		leftOrRight := instructions[steps%len(instructions)]
		if leftOrRight == 'L' {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}

		steps++
	}

	var result = steps
	fmt.Println("Day 08 Part 1 Result: ", result)

	stepCount := []int{}
	for key := range nodes {
		if key[2] == 'A' {
			currentNode := key
			steps := 0
			for currentNode[2] != 'Z' {
				leftOrRight := instructions[steps%len(instructions)]
				if leftOrRight == 'L' {
					currentNode = nodes[currentNode].left
				} else {
					currentNode = nodes[currentNode].right
				}

				steps++
			}
			stepCount = append(stepCount, steps)
		}
	}

	var result2 = LCM(stepCount[0], stepCount[1], stepCount[2:]...)
	fmt.Println("Day 08 Part 2 Result: ", result2)
}

// following code from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
