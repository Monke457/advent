package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

func main() {
	data := reader.FileTo2DIntArray("data/2025/day3.txt")

	var joltage int
	b_len := 12
	for _, bank := range data {
		joltage += getJoltage(bank, b_len)
	}
	fmt.Println(joltage)
}

func getJoltage(bank []int, l int) int {
	var sum, prev, temp int
	batteries := make([]int, l)
	copy(batteries, bank[len(bank)-l:])

	for i := len(bank)-l-1; i >= 0; i-- {
		prev = bank[i]
		for j := 0; j < l; j++ {
			if batteries[j] > prev { break }
			temp = batteries[j]
			batteries[j] = prev
			prev = temp
		}
	}
	fmt.Println(batteries)
	for i, b := range batteries {
		sum += b * int(math.Pow10(l-i-1))
	}
	return sum
}
