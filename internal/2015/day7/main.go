package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2015/day7.txt")

	muhmap := parseInstructions(data)

	first := solveFirstProblem(muhmap)
	fmt.Println("First problem:", first)
	fmt.Println("Second problem:", solveSecondProblem(muhmap, first))
}

var cache map[string]int

func solveFirstProblem(mm map[string]string) int {
	cache = map[string]int{}
	res := getSignal(mm, "a")
	return res
}

func solveSecondProblem(mm map[string]string, override int) int {
	cache = map[string]int{}
	cache["b"] = override
	res := getSignal(mm, "a")
	return res
}

func getSignal(mm map[string]string, val string) int {
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	if v, ok := cache[val]; ok {
		return v
	}
	if strings.Contains(mm[val], "AND") {
		a, b, _ := strings.Cut(mm[val], " AND ")
		cache[a] = getSignal(mm, a)
		cache[b] = getSignal(mm, b)
		return cache[a] & cache[b]

	} else if strings.Contains(mm[val], "OR") {
		a, b, _ := strings.Cut(mm[val], " OR ")
		cache[a] = getSignal(mm, a)
		cache[b] = getSignal(mm, b)
		return cache[a] | cache[b]

	} else if strings.Contains(mm[val], "NOT") {
		a := strings.Split(mm[val], " ")[1]
		cache[a] = getSignal(mm, a)
		return ^cache[a]

	} else if strings.Contains(mm[val], "LSHIFT") {
		a, b, _ := strings.Cut(mm[val], " LSHIFT ")
		bint, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		cache[a] = getSignal(mm, a)
		return cache[a] << bint

	} else if strings.Contains(mm[val], "RSHIFT") {
		a, b, _ := strings.Cut(mm[val], " RSHIFT ")
		bint, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		cache[a] = getSignal(mm, a)
		return cache[a] >> bint
	}

	i, err := strconv.Atoi(mm[val])
	if err != nil {
		cache[mm[val]] = getSignal(mm, mm[val])
		return cache[mm[val]]
	}
	cache[val] = i
	return i
}

func parseInstructions(data []string) map[string]string {
	res := map[string]string{}
	for _, l := range data {
		f, key, _ := strings.Cut(l, " -> ")
		res[key] = f
	}
	return res
}
