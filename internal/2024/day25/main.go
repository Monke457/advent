package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
)

func main() {
	data := reader.FileToArray("data/2024/day25.txt")

	const LENGTH int = 6

	locks := [][5]int{}
	keys := [][5]int{}

	for i := 0; i < len(data)-LENGTH; i+=(LENGTH+2) {
		section := data[i:i+LENGTH+1]
		if section[0][0] == '#' {
			locks = append(locks, parseData(section[1:]))
		} else {
			slices.Reverse(section)
			keys = append(keys, parseData(section[1:]))
		}
	}

	fitCount := 0
	for _, lock := range locks {
		for _, key := range keys {
			if keyFits(lock, key, LENGTH) {
				fitCount++
			}
		}

	}

	fmt.Println("found", fitCount, "possible combinations")
}

func keyFits(lock, key [5]int, l int) bool {
	for i := range lock {
		if lock[i] + key[i] >= l {
			return false
		}
	}
	return true
}

func parseData(data []string) [5]int {
	res := [5]int{}
	for _, line := range data {
		for i, r := range line {
			if r == '#' {
				res[i]++
			}
		}
	}
	return res
}
