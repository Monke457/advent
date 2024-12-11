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

	temp := []int{}
	mul := map[int]int{}

	for _, stone := range stones {
		if stone == 0 {
			mul[1]++
			if mul[1] > 1 {
				continue
			}
			temp = append(temp, 1)
			continue
		}

		numstr := strconv.Itoa(stone)
		if len(numstr) % 2 != 0 {
			num := stone * 2024
			mul[num]++
			if mul[num] > 1 {
				continue
			}
			temp = append(temp, num)
			continue
		}

		half := len(numstr) / 2
		left, right := numstr[:half], numstr[half:]
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)

		mul[l]++
		if mul[l] <= 1 {
			temp = append(temp, l)
		}

		mul[r]++
		if mul[r] <= 1 {
			temp = append(temp, r)
		}
	}

	arr := []int{}
	l := 0
	for _, t := range temp {
		if _, ok := cache[t]; !ok {
			cache[t] = map[int]int{}
		}
		if _, ok := cache[t][n-1]; ok {
			l += max(mul[t], 1) * cache[t][n-1]
			continue
		}

		if v, ok := mul[t]; ok {
			res := transformStones([]int{t}, n-1)
			cache[t][n-1] = res
			l += v * cache[t][n-1]
		} else {
			arr = append(arr, t)
		}
	}
	return l + transformStones(arr, n-1)
}
