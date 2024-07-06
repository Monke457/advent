package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
)

func main() {
	data := reader.FileToIntArray("data/2020/day9.txt")

	var first int
	for i := 25; i < len(data); i++ { 
		if sumExists(data, data[i], i, 25) {
			continue
		}
		first = data[i]
		break
	}

	fmt.Println("First:", first)

	s, e, err := findRange(data, first)
	if err != nil {
		panic(err)
	}
	fmt.Println("Second:", sumExtremes(data[s:e]))
}

func sumExtremes(r []int) int {
	mn := slices.Min(r)
	mx := slices.Max(r)
	return mn + mx

}

func findRange(data []int, tar int) (int, int, error) {
	loop:
	for i := 0; i < len(data); i++ {
		for j := i+1; j < len(data); j++ {
			sum := sumRange(data[i:j])
			if sum > tar {
				continue loop
			}
			if sum == tar {
				return i, j, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("No valid sum range found") 
}

func sumRange(data []int) int {
	sum := 0
	for _, i := range data {
		sum += i
	}
	return sum 
}

func sumExists(data []int, tar, pos, length int) bool {
	for i := 1; i <= length; i++ {
		for j := i+1; j <= length; j++ {
			if data[pos-i] + data[pos-j] == tar {
				return true
			}
		}
	}
	return false
}
