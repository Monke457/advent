package main

import (
	"advent/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArray("data/2017/day1.txt")

	fmt.Println("First:", solve(data, plusOne))
	fmt.Println("Second:", solve(data, halfWay))
}

func solve(nums []int, fn func(int, int) int) int {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		n := fn(i, l)
		if nums[i] == nums[n] {
			sum += nums[i]
		}
	}
	return sum
}

func plusOne(n, l int) int {
	return (n + 1) % l
}

func halfWay(n, l int) int {
	return (n + (l >> 1)) % l
}
