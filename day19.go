package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	LessThan    byte = 0
	GreaterThan byte = 1
)

type WorkflowRule struct {
	category     byte
	operation    byte
	value        int
	nextWorkflow string
}

type Workflows = map[string][]WorkflowRule

type PartRange struct {
	start, end int
}

func day19() {
	lines := getLines("input/19.txt")
	workflows, parts := parseInput(lines)

	acceptedPartsSum := 0
	for _, part := range parts {
		current := "in"
		for current != "A" && current != "R" {
			currentWorkflow := workflows[current]
			for _, rule := range currentWorkflow {
				if rule.category == 0 || (rule.operation == LessThan && part[rule.category] < rule.value) ||
					(rule.operation == GreaterThan && part[rule.category] > rule.value) {
					current = rule.nextWorkflow
					break
				}
			}
		}
		if current == "A" {
			acceptedPartsSum += partRatingsSum(part)
		}
	}

	var result = acceptedPartsSum
	fmt.Println("Day 19 Part 1 Result: ", result)

	var result2 = countCombinations(workflows, "in", []PartRange{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}})
	fmt.Println("Day 19 Part 2 Result: ", result2)
}

func countCombinations(workflows Workflows, workflow string, ranges []PartRange) int {
	if workflow == "R" {
		return 0
	}
	if workflow == "A" {
		return partRangesProduct(ranges)
	}

	result := 0
	currentWorkflow := workflows[workflow]
	for _, rule := range currentWorkflow {
		newRanges := make([]PartRange, len(ranges))
		copy(newRanges, ranges[:])
		rangeIndex := map[byte]int{'x': 0, 'm': 1, 'a': 2, 's': 3}[rule.category]

		if rule.category == 0 {
			result += countCombinations(workflows, rule.nextWorkflow, ranges)
		} else if rule.operation == LessThan {
			newRanges[rangeIndex].end = rule.value - 1
			ranges[rangeIndex].start = rule.value
			result += countCombinations(workflows, rule.nextWorkflow, newRanges)
		} else {
			newRanges[rangeIndex].start = rule.value + 1
			ranges[rangeIndex].end = rule.value
			result += countCombinations(workflows, rule.nextWorkflow, newRanges)
		}
	}
	return result
}

func partRangesProduct(ranges []PartRange) int {
	result := 1
	for _, r := range ranges {
		result *= r.end - r.start + 1
	}
	return result
}

func partRatingsSum(part map[byte]int) int {
	result := 0
	for _, val := range part {
		result += val
	}
	return result
}

func parseInput(lines []string) (Workflows, []map[byte]int) {
	workflows := Workflows{}
	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			i++
			break
		}

		split := strings.Split(line, "{")
		workflow := split[0]
		split = strings.Split(split[1], ",")
		for j := 0; j < len(split); j++ {
			rule := WorkflowRule{}

			data := strings.Split(split[j], ":")
			if len(data) == 1 {
				rule.nextWorkflow = data[0][:len(data[0])-1]
				workflows[workflow] = append(workflows[workflow], rule)
				break
			}

			rule.nextWorkflow = data[1]
			rule.category = data[0][0]
			operation := data[0][1]
			if operation == '<' {
				rule.operation = LessThan
			} else {
				rule.operation = GreaterThan
			}
			value, _ := strconv.Atoi(data[0][2:])
			rule.value = value
			workflows[workflow] = append(workflows[workflow], rule)
		}
	}

	parts := []map[byte]int{}
	start := i
	for ; i < len(lines); i++ {
		parts = append(parts, map[byte]int{})
		line := lines[i]
		split := strings.Split(line[1:len(line)-1], ",")
		for j := 0; j < len(split); j++ {
			category := split[j][0]
			value, _ := strconv.Atoi(split[j][2:])
			parts[i-start][category] = value
		}
	}
	return workflows, parts
}
