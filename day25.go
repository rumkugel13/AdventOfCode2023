package main

import (
	"fmt"
	"strings"
)

type Edge struct {
	from, to string
}

func day25() {
	lines := getLines("input/25.txt")
	nodes := getNodes(lines)

	edgeCount := countEdges(nodes)
	first := findAndRemoveMax(edgeCount)
	removeEdge(nodes, first)
	edgeCount = countEdges(nodes)
	second := findAndRemoveMax(edgeCount)
	removeEdge(nodes, second)
	edgeCount = countEdges(nodes)
	third := findAndRemoveMax(edgeCount)
	removeEdge(nodes, third)

	countA := countNodes(nodes, first.from)
	countB := len(nodes) - countA

	var result = countA * countB
	var result2 = 0
	fmt.Println("Day 25 Part 1 Result: ", result)
	fmt.Println("Day 25 Part 2 Result: ", result2)
}

func removeEdge(nodes map[string][]string, edge Edge) {
	new := []string{}
	for _, val := range nodes[edge.from] {
		if val != edge.to {
			new = append(new, val)
		}
	}
	nodes[edge.from] = new
	new = []string{}
	for _, val := range nodes[edge.to] {
		if val != edge.from {
			new = append(new, val)
		}
	}
	nodes[edge.to] = new
}

func findAndRemoveMax(edges map[Edge]int) Edge {
	max := 0
	var maxEdge Edge
	for key, val := range edges {
		if val > max {
			max = val
			maxEdge = key
		}
	}
	delete(edges, maxEdge)
	return maxEdge
}

func countEdges(nodes map[string][]string) map[Edge]int {
	encountered := map[Edge]int{}
	i := 0
	for from := range nodes {
		walkNodes(nodes, from, encountered)
		i++
		// if this doesn't work, increase value
		if i > 32 {
			break
		}
	}
	return encountered
}

func countNodes(nodes map[string][]string, start string) int {
	visited := map[string]bool{}
	queue := []string{start}

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]

		for _, to := range nodes[from] {
			if _, found := visited[to]; found {
				continue
			}
			queue = append(queue, to)
			visited[to] = true
		}
	}
	return len(visited)
}

func walkNodes(nodes map[string][]string, start string, encountered map[Edge]int) {
	visited := map[string]bool{}
	queue := []string{start}

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]

		for _, to := range nodes[from] {
			if _, found := visited[to]; found {
				continue
			}
			queue = append(queue, to)
			visited[to] = true
			var edge Edge
			if from < to {
				edge = Edge{from, to}
			} else {
				edge = Edge{to, from}
			}
			encountered[edge]++
		}
	}
}

func getNodes(lines []string) map[string][]string {
	nodes := map[string][]string{}
	for _, line := range lines {
		split := strings.Split(line, ": ")
		from := split[0]
		if _, inside := nodes[from]; !inside {
			nodes[from] = []string{}
		}

		to := strings.Fields(split[1])
		for _, target := range to {
			nodes[from] = append(nodes[from], target)
			if _, inside := nodes[target]; !inside {
				nodes[target] = []string{}
			}
			nodes[target] = append(nodes[target], from)
		}
	}
	return nodes
}
