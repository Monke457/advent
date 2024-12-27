package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

var cache = map[string]int{}

func main() {
	data := reader.FileToArray("data/2024/day19.txt")
	patterns := strings.Split(data[0], ", ")
	designs := make([]string, len(data[2:]))

	for i, line := range data[2:] {
		designs[i] = line
	}

	possible := 0
	variants := 0
	for _, design := range designs {
		pos := possibleVariants(patterns, design)
		if pos > 0 {
			possible++
		}
		variants += pos
	}

	fmt.Println("Possible designs:", possible)
	fmt.Println("Possible variants:", variants)
}

func possibleVariants(patterns []string, design string) int {
	if p, ok := cache[design]; ok {
		return p
	}
	if len(design) == 0 {
		return 1
	}
	count := 0
	for _, pattern := range patterns {
		if len(pattern) > len(design) {
			continue
		}
		if pattern == design[:len(pattern)] {
			count += possibleVariants(patterns, design[len(pattern):])
		}
	}
	cache[design] = count
	return count
}
