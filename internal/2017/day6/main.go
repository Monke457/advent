package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
)

type cache struct {
	saved [][]int
}

func main() {
	banks := reader.FileToIntArrayByTab("data/2017/day6.txt")

	cache := cache{ saved: [][]int{} }
	cycle := 0
	for !cache.isCached(banks) {
		newBanks := make([]int, len(banks))
		copy(newBanks, banks)
		cache.saved = append(cache.saved, newBanks)

		m := getMaxIndex(banks)
		banks = redistribute(banks, m)
		cycle++
	}
	fmt.Println("First:", cycle)

	target := make([]int, len(banks)) 
	copy(target, banks)
	cycle = 0
	for {
		m := getMaxIndex(banks)
		banks = redistribute(banks, m)
		cycle++
		if slices.Equal(target, banks) {
			break
		}
	}
	fmt.Println("Second:", cycle)
}

func getMaxIndex(banks []int) int {
	res := 0
	for i := range banks {
		if banks[i] > banks[res] {
			res = i
		}
	}
	return res
}

func redistribute(banks []int, start int) []int {
	val := banks[start]
	banks[start] = 0
	next := (start + 1) % len(banks)
	for val > 0 {
		banks[next]++
		val--
		next = (next + 1) % len(banks)
	}
	return banks
}

func (c cache) isCached(state []int) bool {
	for _, s := range c.saved {
 		if slices.Equal(s, state) {
			return true 
		}
	}
	return false 
} 
