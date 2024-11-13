package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2019/day6.txt")

	orbits := parseOrbits(data)

	first := calculate(orbits)
	fmt.Printf("First: %d  \n", first)

	distance := getDistance(orbits, "YOU", "SAN")
	fmt.Printf("Second: %d  \n", distance)
}

func getDistance(orbits map[string]string, from, to string) int {
	youMap := getMap(orbits, "YOU", []string{})
	sanMap := getMap(orbits, "SAN", []string{})

	var lcd string
	var distance int
	for i, step := range youMap {
		if slices.Contains(sanMap, step) {
			distance = i
			lcd = step
			break
		}
	}
	
	distance += slices.Index(sanMap, lcd)

	return distance - 2
}

func getMap(orbits map[string]string, obj string, res []string) []string {
	res = append(res, obj)
	if orb, ok := orbits[obj]; ok {
		res = getMap(orbits, orb, res)
	}
	return res
}

func calculate(orbits map[string]string) int {
	count := 0
	cache := map[string]int{}
	for obj := range orbits {
		if sum, ok := cache[obj]; ok {
			count += sum
			continue
		}
		count += getOrbits(orbits, cache, obj)
	}
	return count
}

func getOrbits(orbits map[string]string, cache map[string]int, obj string) int {
	next, ok := orbits[obj]
	if !ok {
		return 0 
	}
	if sum, ok := cache[obj]; ok {
		return sum
	}
	sum := getOrbits(orbits, cache, next) + 1
	cache[obj] = sum
	return sum
}

func parseOrbits(data []string) map[string]string {
	var orbits = map[string]string{}
	for _, line := range data {
		obj, orb, ok := strings.Cut(line, ")")
		if !ok {
			panic("Something went wrong")
		}
		_, ok = orbits[orb]
		orbits[orb] = obj
	}
	return orbits
}
