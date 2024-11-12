package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)


func main() {
	data := reader.FileToArray("data/2019/day6.txt")

	orbits := parseOrbits(data)

	for obj, orb := range orbits {
		fmt.Printf("object: %s, in orbit: %v\n", obj, orb)
	}
}

func parseOrbits(data []string) map[string][]string {
	var orbits = map[string][]string{}
	for _, line := range data {
		obj, orb, ok := strings.Cut(line, ")")
		if !ok {
			panic("Something went wrong")
		}
		_, ok = orbits[obj]
		if !ok {
			orbits[obj] = []string{}
		}
		orbits[obj] = append(orbits[obj], orb)
	}
	return orbits
}
