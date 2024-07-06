package main

import (
	"advent/internal/pkg/reader"
	"cmp"
	"fmt"
	"slices"
)

func main() {
	data := reader.FileToIntArray("data/2020/day10.txt")
	slices.SortFunc(data, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	data = append([]int{0}, data[:]...)
	data = append(data, data[len(data)-1] + 3)

	fmt.Println("First:", first(data)) 
	fmt.Println("Second:", second(data, 0))
}

func first(data []int) int {
	jolts := map[int]int{}
	curr := 0 
	for _, line := range data {
		jolts[line - curr]++
		curr = line
	}
	return jolts[1] * jolts[3]
}

var cache = map[int]int{}

func second(data []int, idx int) int {
	if val, ok := cache[idx]; ok {
		return val
	}
	res := 1
	curr := data[idx]
	for i := idx+1; i < len(data); i++ {
		if i < len(data) - 1 && data[i+1] - curr < 4 {
			res += second(data, i+1)
		}
		if i < len(data) - 2 && data[i+2] - curr < 4 {
			res += second(data, i+2)
		}
		curr = data[i]
	}
	cache[idx] = res
	return res 
}
