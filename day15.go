package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Lens struct {
	label       string
	focalLength int
}

func day15() {
	input := getLines("input/15.txt")
	sequence := strings.Split(input[0], ",")

	hashSum := 0
	for _, step := range sequence {
		hashSum += calculateHash(step)
	}

	var result = hashSum
	fmt.Println("Day 15 Part 1 Result: ", result)

	boxes := make([][]Lens, 256)
	for _, step := range sequence {
		if strings.Contains(step, "-") {
			data := strings.Split(step, "-")
			box := calculateHash(data[0])
			removeFromBox(boxes, box, data[0])
		} else {
			data := strings.Split(step, "=")
			box := calculateHash(data[0])
			focal, _ := strconv.Atoi(data[1])
			addOrUpdateBox(boxes, box, data[0], focal)
		}
	}

	focusingPower := 0
	for boxNum, box := range boxes {
		for i, lens := range box {
			power := (boxNum + 1) * (i + 1) * lens.focalLength
			focusingPower += power
		}
	}

	var result2 = focusingPower
	fmt.Println("Day 15 Part 2 Result: ", result2)
}

func addOrUpdateBox(boxes [][]Lens, box int, label string, focal int) {
	for i, lens := range boxes[box] {
		if lens.label == label {
			boxes[box][i].focalLength = focal
			return
		}
	}
	boxes[box] = append(boxes[box], Lens{label, focal})
}

func removeFromBox(boxes [][]Lens, box int, label string) {
	for i, lens := range boxes[box] {
		if lens.label == label {
			boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
			break
		}
	}
}

func calculateHash(data string) int {
	result := 0
	for _, char := range data {
		result += int(char)
		result *= 17
		result %= 256
	}
	return result
}
