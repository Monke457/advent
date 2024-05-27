package main

import (
	"advent/pkg/reader"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2015/day13.txt")

	prefs := parsePrefs(data)

	optimal := 0
	for name := range prefs {
		cache = []string{}
		optimal = max(optimal, findOptimal(prefs, name))
	}

	fmt.Println("First problem:", optimal)
}

var cache = []string{}

func findOptimal(p map[string]map[string]int, s string) (res int) {
	if len(cache) == len(p) {
		return res
	}
	cache = append(cache, s)
	keys := []string{}
	for name := range p[s] {
		if slices.Contains(cache, name) {
			continue
		}
		keys = append(keys, name)
	}
	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(p[s][b], p[s][a])
	})
	for _, key := range keys {
		return p[s][key] + p[key][s] + findOptimal(p, key)
	}

	return res + p[s][cache[0]] + p[cache[0]][s]
}

func parsePrefs(data []string) map[string]map[string]int {
	res := map[string]map[string]int{}
	for _, l := range data {
		parts := strings.Split(l, " ")
		name := parts[0]
		neighbour := strings.Trim(parts[len(parts)-1], ".")
		val, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err)
		}
		if parts[2] == "lose" {
			val *= -1
		}
		if _, ok := res[name]; !ok {
			res[name] = map[string]int{}
		}
		res[name][neighbour] = val
	}
	return res
}
