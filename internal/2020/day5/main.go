package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2020/day5.txt")

	ids := make([]int, 0)
	for _, line := range data {
		res := getId(line)
		ids = insertId(ids, res)
	}
	fmt.Println("First:", ids[len(ids)-1])
	fmt.Println("First:", findMissing(ids))
}

func findMissing(ids []int) int {
	for i := 0; i < len(ids)-1; i++ {
		if ids[i] + 1 != ids[i+1] {
			return ids[i] + 1
		}
	}
	return 0
}

func insertId(ids []int, id int) []int {
	idx := 0
	for idx < len(ids) && ids[idx] < id {
		idx++
	}
	return append(ids[:idx], append([]int{id}, ids[idx:]...)...)
}

func getId(bsp string) int {
	col := binarySearch(bsp[:7], 'F', 0, 127)
	row := binarySearch(bsp[7:], 'L', 0, 7)
	return col * 8 + row
}

func binarySearch(bsp string, front byte, low, high int) int {
	for i := 0; i < len(bsp)-1; i++ {
		if bsp[i] == front {
			high = low + (high - low) >> 1 
		} else { 
			low = low + 1 + (high - low) >> 1
		}
	}
	if bsp[len(bsp)-1] == front {
		return low
	} else {
		return high 
	}
}

