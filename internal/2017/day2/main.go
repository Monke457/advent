package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

func main() {
	data := reader.FileTo2DIntArray("data/2017/day2.txt")
	solveFirst(data)
	solveSecond(data)
}

func solveFirst(data [][]int) {
	sum := 0
	for _, line := range data {
		mn, mx := extremeValues(line)
		sum += mx - mn
	}
	fmt.Println("First:", sum)
}

func solveSecond(data [][]int) {
	sum := 0
	for _, line := range data {
		mn, mx := divisibleValues(line)
		sum += mx / mn 
	}
	fmt.Println("Second:", sum)
}

func extremeValues(vals []int) (int, int) {
	mn, mx := math.MaxInt, 0
	for _, val :=  range vals {
		if val < mn {
			mn = val
		}
		if val > mx {
			mx = val
		}
	}
	return mn, mx
}

func divisibleValues(vals []int) (int, int) {
	for i, a := range vals {
		for j, b := range vals {
			if i == j {
				continue
			}
			if a % b == 0 {
				return b, a
			}
			if b % a == 0 {
				return a, b
			}
		}
	}
	panic("No divisible numbers found in array")
}
