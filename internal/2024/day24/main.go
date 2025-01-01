package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strings"
)

var wires map[string]func(a, b bool)bool
var gates map[string][2]string
var cache map[string]bool

func main() {
	data := reader.FileToArray("data/2024/day24.txt")

	wires = map[string]func(a, b bool)bool{}
	gates = map[string][2]string{}
	cache = map[string]bool{}

	for i := range data {
		if data[i] == "" {
			parseWires(data[:i]) 
			parseGates(data[i+1:])
		}
	}

	zWires := getWires('z')
	fmt.Println("First Part Number:", parseNumber(zWires))
}

func parseNumber(keys []string) int {
	num := 0
	for i, key := range keys {
		if !isOn(key) {
			continue
		}
		if i == 0 {
			num += 1
			continue
		}
		num += 2 << (i-1)
	}
	return num
}

func isOn(key string) bool {
	if val, ok := cache[key]; ok {
		return val
	}
	if _, ok := gates[key]; !ok {
		cache[key] = wires[key](true, true)
		return cache[key]
	}

	w1 := gates[key][0]
	w2 := gates[key][1]

	cache[key] = wires[key](isOn(w1), isOn(w2))
	return cache[key]
}

func getWires(first byte) []string {
	keys := []string{}
	for wire := range wires {
		if wire[0] == first {
			keys = append(keys, wire)
		}
	}
	slices.Sort(keys)
	return keys
}

func parseWires(data []string) {
	for _, line := range data {
		wire, value, _ := strings.Cut(line, ": ")
		if value == "1" {
			wires[wire] = func(_, _ bool) bool { return true }
		} else {
			wires[wire] = func(_, _ bool) bool { return false }
		}
	}
}

func parseGates(data []string) {
	for _, line := range data {
		op, dest, _ := strings.Cut(line, " -> ") 
		parts := strings.Split(op, " ")
		switch parts[1] {
		case "AND":
			wires[dest] = andGate
		case "OR":
			wires[dest] = orGate
		case "XOR":
			wires[dest] = xorGate
		}
		gates[dest] = [2]string{parts[0], parts[2]}
	}
}

func andGate(a, b bool) bool {
	return a && b
}

func orGate(a, b bool) bool {
	return a || b
}

func xorGate(a, b bool) bool {
	return (a && !b) || (!a && b)
}
