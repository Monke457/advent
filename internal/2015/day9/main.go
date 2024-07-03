package main

import (
	"advent/internal/pkg/reader"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2015/day9.txt")

	distances := parseDistances(data)

	fmt.Println(solveFirstProblem(distances))
	fmt.Println(solveSecondProblem(distances))
}

var cache = []string{}

func solveFirstProblem(ds map[string]map[string]int) int {
	res := math.MaxInt
	for k := range ds {
		cache = []string{}
		res = min(res, findShortestRoute(ds, k))
	}
	return res
}

func solveSecondProblem(ds map[string]map[string]int) int {
	res := 0
	for k := range ds {
		cache = []string{}
		res = max(res, findLongestRoute(ds, k))
	}
	return res
}

func findLongestRoute(ds map[string]map[string]int, loc string) (v int) {
	if len(cache) == len(ds) {
		return v
	}

	cache = append(cache, loc)
	keys := []string{}
	for k := range ds[loc] {
		if slices.Contains(cache, k) {
			continue
		}
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(ds[loc][b], ds[loc][a])
	})

	for _, k := range keys {
		v += ds[loc][k]
		return v + findLongestRoute(ds, k)
	}

	return v
}

func findShortestRoute(ds map[string]map[string]int, loc string) (v int) {
	if len(cache) == len(ds) {
		return v
	}

	cache = append(cache, loc)
	keys := []string{}
	for k := range ds[loc] {
		if slices.Contains(cache, k) {
			continue
		}
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(ds[loc][a], ds[loc][b])
	})

	for _, k := range keys {
		v += ds[loc][k]
		return v + findShortestRoute(ds, k)
	}

	return v
}

func parseDistances(data []string) map[string]map[string]int {
	res := make(map[string]map[string]int)
	for _, l := range data {
		parts := strings.Split(l, " ")
		d, err := strconv.Atoi(parts[4])
		if err != nil {
			panic(err)
		}
		if _, ok := res[parts[0]]; !ok {
			res[parts[0]] = make(map[string]int)
		}
		if _, ok := res[parts[2]]; !ok {
			res[parts[2]] = make(map[string]int)
		}
		res[parts[0]][parts[2]] = d
		res[parts[2]][parts[0]] = d
	}
	return res
}
