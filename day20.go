package main

import (
	"fmt"
	"strings"
)

const (
	Broadcaster int = 0
	FlipFlop    int = 1
	Conjunction int = 2
)

type Module struct {
	name          string
	moduleType    int
	destinations  []string
	prevStates    map[string]bool
	flipFlopState *bool
}

func day20() {
	lines := getLines("input/20.txt")
	modules := parseModules(lines)

	lowPulseCount, highPulseCount := 0, 0
	hasInputs := []Module{}
	for i := 0; i < 1000; i++ {
		inputs := sendPulse(modules, false, "button", []string{"broadcaster"}, &lowPulseCount, &highPulseCount)
		hasInputs = append(hasInputs, inputs...)
		for len(hasInputs) > 0 {
			currentModule := hasInputs[0]
			hasInputs = hasInputs[1:]

			switch currentModule.moduleType {
			case Broadcaster:
				pulse := currentModule.prevStates["button"]
				delete(currentModule.prevStates, "button")
				inputs := sendPulse(modules, pulse, currentModule.name, currentModule.destinations, &lowPulseCount, &highPulseCount)
				hasInputs = append(hasInputs, inputs...)
			case FlipFlop:
				for _, pulse := range currentModule.prevStates {
					if !pulse {
						*currentModule.flipFlopState = !*currentModule.flipFlopState
						inputs := sendPulse(modules, *currentModule.flipFlopState, currentModule.name, currentModule.destinations, &lowPulseCount, &highPulseCount)
						hasInputs = append(hasInputs, inputs...)
					}
				}
				clear(currentModule.prevStates)
			case Conjunction:
				totalPulse := true
				for _, pulse := range currentModule.prevStates {
					if !pulse {
						totalPulse = false
					}
				}
				inputs := sendPulse(modules, !totalPulse, currentModule.name, currentModule.destinations, &lowPulseCount, &highPulseCount)
				hasInputs = append(hasInputs, inputs...)
			}
		}
	}

	var result = lowPulseCount * highPulseCount
	var result2 = 0
	fmt.Println("Day 20 Part 1 Result: ", result)
	fmt.Println("Day 20 Part 2 Result: ", result2)
}

func sendPulse(modules map[string]Module, pulse bool, sender string, dest []string, low, high *int) (inputs []Module) {
	for _, name := range dest {
		_, found := modules[name]
		if found {
			modules[name].prevStates[sender] = pulse
			inputs = append(inputs, modules[name])
		}
	}
	if pulse {
		*high += len(dest)
	} else {
		*low += len(dest)
	}
	return
}

func parseModules(lines []string) map[string]Module {
	modules := map[string]Module{}
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		name := split[0]
		destinations := strings.Split(split[1], ", ")
		val := false
		if name[0] == 'b' {
			modules[name] = Module{name, Broadcaster, destinations, map[string]bool{}, &val}
		} else if name[0] == '%' {
			modules[name[1:]] = Module{name[1:], FlipFlop, destinations, map[string]bool{}, &val}
		} else if name[0] == '&' {
			modules[name[1:]] = Module{name[1:], Conjunction, destinations, map[string]bool{}, &val}
		}
	}

	for name, module := range modules {
		for _, val := range module.destinations {
			if dest, found := modules[val]; found && dest.moduleType == Conjunction {
				dest.prevStates[name] = false
			}
		}
	}

	return modules
}
