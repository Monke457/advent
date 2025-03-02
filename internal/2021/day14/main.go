package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2021/day14.txt")

	monomers := map[string]int{}
	for i := 0; i < len(data[0])-1; i++ {
		monomer := data[0][i:i+2]
		monomers[monomer]++
	}
	first := data[0][0]

	rules := map[string]byte{}
	for _, line := range data[2:] {
		pair, insert, _ := strings.Cut(line, " -> ")
		rules[pair] = insert[0]
	}

	steps := 40
	for range steps {
		monomers = process(monomers, rules)
	}

	mcv, lcv := examine(monomers, first)

	fmt.Println("Difference in rate of most common and least common elements:", mcv - lcv)
}

func examine(mono map[string]int, first byte) (int, int) {
	m := map[byte]int{first: 1}
	for val, count := range mono {
		m[val[1]] += count
	}
	mcv := 0
	lcv := math.MaxInt
	for _, v := range m {
		mcv = max(v, mcv)
		lcv = min(v, lcv)
	}
	return mcv, lcv
}

func process(mono map[string]int, rules map[string]byte) map[string]int {
	newMonomers := map[string]int{}
	for pair, count := range mono {
		if insert, ok := rules[pair]; ok {
			a := fmt.Sprintf("%c%c", pair[0], insert)
			b := fmt.Sprintf("%c%c", insert, pair[1])
			newMonomers[a] += count
			newMonomers[b] += count
		} else {
			newMonomers[pair] += count
		}
	}
	return newMonomers
}
