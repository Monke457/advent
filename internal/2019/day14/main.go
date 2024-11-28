package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var cache = map[string]int{}
var yields = map[string]int{}
var reactions = map[string]func(int)int{}

const (
	base = "ORE"
	target = "FUEL"
	trillion = 1000000000000
)

func main() {
	data := reader.FileToArray("data/2019/day14.txt")

	reactions[base] = func(n int)int{return n}

	for _, line := range data {
		fn, output, _ := strings.Cut(line, " => ")
		fnParts := mapFunction(fn)
		amount, name, _ := strings.Cut(output, " ")
		n, err := strconv.Atoi(amount)
		if err != nil {
			panic(err)
		}
		yields[name] = n

		reactions[name] = func(n int) int {
			ore := 0
			if cache[name] >= n {
				cache[name] -= n
				return ore 
			}
			total := int(math.Ceil(float64((n - cache[name])) / float64(yields[name])))
			for key, val := range fnParts {
				needed := total * val 
				if cache[key] < needed {
					ore += reactions[key](needed)
				} else {
					cache[key] -= needed
				}
			}
			cache[name] += total * yields[name] - n
			return ore 
		}
	}

	first := reactions[target](1)
	fmt.Println("First:", first)

	fuel := 1
	for {
		resetCache()
		result := reactions[target](fuel+1)
		if result > trillion {
			break
		}
		fuel = max(fuel + 1, int(math.Floor(float64((fuel + 1) * trillion / result))));
	}

	fmt.Println("Second:", fuel)
}

func resetCache() {
	for key := range cache {
		delete(cache, key)
	}
}

func mapFunction(fn string) map[string]int {
	fns := strings.Split(fn, ", ")
	result := map[string]int{}
	for _, fn := range fns {
		amount, name, _ := strings.Cut(fn, " ")
		n, err := strconv.Atoi(amount)
		if err != nil {
			panic(err)
		}
		result[name] = n 
	}
	return result
}
