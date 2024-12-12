package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
)

func main() {
	stones := reader.FileToIntArrayByDivider("data/2024/day11.txt", " ")

	blinks := 75
	length := transformStones(stones, blinks)
	fmt.Printf("Stones after %d blinks: %d\n", blinks, length)
}

var cache = map[int]map[int]int{}

func transformStones(stones []int, n int) int {
	if n == 0 {
		return len(stones)
	}

	counts := map[int]int{}

	for _, stone := range stones {
		if stone == 0 {
			counts[1]++
			continue
		}

		numstr := strconv.Itoa(stone)
		if len(numstr) % 2 != 0 {
			num := stone * 2024
			counts[num]++
			continue
		}

		half := len(numstr) / 2
		left, right := numstr[:half], numstr[half:]
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)

		counts[l]++
		counts[r]++
	}

	l := 0
	for s, v := range counts{
		if _, ok := cache[s]; !ok {
			cache[s] = map[int]int{}
		}
		if _, ok := cache[s][n-1]; ok {
			l += v * cache[s][n-1]
			continue
		}

		res := transformStones([]int{s}, n-1)
		cache[s][n-1] = res
		l += v * res
	}
	return l
}
