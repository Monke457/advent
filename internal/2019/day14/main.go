package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
	"time"
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
			for key, val := range fnParts {
				if cache[key] < val {
					ore += reactions[key](val)
				}
				cache[key] -= val 
			}
			cache[name] += yields[name]
			if cache[name] < n {
				return ore + reactions[name](n)
			}
			return ore 
		}
	}

	total := reactions[target](1)
	fmt.Println("First:", total)

	start := time.Now()
	for {
		total += reactions[target](1)
		fmt.Printf("%s: %d       %s: %d       %v\t\t\r", target, cache[target], base, cache[base]*-1, time.Since(start))
		if zeroCache() {
			fmt.Println("Zero cache found", cache)
			break
		}
		if total >= trillion {
			fmt.Println()
			break
		}
	}

	fmt.Println("Second:", cache[target] * trillion / total)
}

func zeroCache() bool {
	for key, val := range cache {
		if key == target || key == base {
			continue
		}
		if val > 0 {
			return false
		}
	}
	return true
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
