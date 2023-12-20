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
	modules, modulesForPart2 := parseModules(lines)

	cyclesPart2 := map[string]int{}
	for _, name := range modulesForPart2 {
		cyclesPart2[name] = 999999999
	}

	finished := true
	lowPulseCount, highPulseCount := 0, 0
	i := 1
	for ; i <= 1000; i++ {
		buttonPress(modules, &lowPulseCount, &highPulseCount, cyclesPart2, i)
		for _, val := range cyclesPart2 {
			if val == 999999999 {
				finished = false
			}
		}
	}

	var result = lowPulseCount * highPulseCount
	fmt.Println("Day 20 Part 1 Result: ", result)

	for ; !finished; i++ {
		finished = true
		buttonPress(modules, &lowPulseCount, &highPulseCount, cyclesPart2, i)
		for _, val := range cyclesPart2 {
			if val == 999999999 {
				finished = false
			}
		}
	}

	valuesPart2 := []int{}
	for _, val := range cyclesPart2 {
		valuesPart2 = append(valuesPart2, val)
	}

	var result2 = LCM(valuesPart2[0], valuesPart2[1], valuesPart2...)
	fmt.Println("Day 20 Part 2 Result: ", result2)
}

func buttonPress(modules map[string]Module, lowPulseCount, highPulseCount *int, iterations map[string]int, iteration int) {
	hasInputs := []Module{}
	sendPulse(modules, &hasInputs, false, "button", []string{"broadcaster"}, lowPulseCount, highPulseCount, iterations, iteration)
	for len(hasInputs) > 0 {
		currentModule := hasInputs[0]
		hasInputs = hasInputs[1:]

		switch currentModule.moduleType {
		case Broadcaster:
			pulse := currentModule.prevStates["button"]
			delete(currentModule.prevStates, "button")
			sendPulse(modules, &hasInputs, pulse, currentModule.name, currentModule.destinations, lowPulseCount, highPulseCount, iterations, iteration)
		case FlipFlop:
			// note: the range might not be in correct order, but it still works if only one of the modules sends a low signal
			for _, pulse := range currentModule.prevStates {
				if !pulse {
					*currentModule.flipFlopState = !*currentModule.flipFlopState
					sendPulse(modules, &hasInputs, *currentModule.flipFlopState, currentModule.name, currentModule.destinations, lowPulseCount, highPulseCount, iterations, iteration)
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
			sendPulse(modules, &hasInputs, !totalPulse, currentModule.name, currentModule.destinations, lowPulseCount, highPulseCount, iterations, iteration)
		}
	}
}

func sendPulse(modules map[string]Module, inputs *[]Module, pulse bool, sender string, dest []string, low, high *int, iterations map[string]int, iteration int) {
	for _, name := range dest {
		_, found := modules[name]
		if found {
			modules[name].prevStates[sender] = pulse
			*inputs = append(*inputs, modules[name])
		}
	}
	if pulse {
		*high += len(dest)
	} else {
		*low += len(dest)
	}
	for p2 := range iterations {
		if sender == p2 && !pulse {
			iterations[sender] = min(iterations[sender], iteration)
		}
	}
}

func parseModules(lines []string) (map[string]Module, []string) {
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

	rxPredecessor := ""
	for name, module := range modules {
		for _, val := range module.destinations {
			if dest, found := modules[val]; found && dest.moduleType == Conjunction {
				dest.prevStates[name] = false
			} else if !found && val == "rx" {
				rxPredecessor = name
			}
		}
	}

	modulesNeededToSendHighPulse := []string{}
	for name, module := range modules {
		for _, val := range module.destinations {
			if val == rxPredecessor {
				modulesNeededToSendHighPulse = append(modulesNeededToSendHighPulse, name)
			}
		}
	}

	modulesNeededToSendLowPulse := []string{}
	for name, module := range modules {
		for _, val := range module.destinations {
			for _, m := range modulesNeededToSendHighPulse {
				if m == val {
					modulesNeededToSendLowPulse = append(modulesNeededToSendLowPulse, name)
				}
			}
		}
	}

	return modules, modulesNeededToSendLowPulse
}
