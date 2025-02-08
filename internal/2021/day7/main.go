package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

func main() {
	crabs := reader.FileToIntArrayByComma("data/2021/day7.txt")

	avg := getAverage(crabs)
	margin := len(crabs) >> 2
	lowestCost := math.MaxInt
	lowestCostF := math.MaxInt
	best := 0
	bestF := 0

	for i := avg - margin; i < avg + margin; i++ {
		cost := sumDiffs(crabs, i)
		if cost < lowestCost {
			lowestCost = cost
			best = i
		}
		costF := sumDiffFactorials(crabs, i)
		if costF < lowestCostF {
			lowestCostF = costF
			bestF = i
		}
	}

	fmt.Printf("best alignment: %d, cost: %d\n", best, lowestCost)
	fmt.Printf("best alignment factorial: %d, cost: %d\n", bestF, lowestCostF)
}

func sumDiffFactorials(nums []int, pos int) int {
	result := 0
	for _, num := range nums {
		diff := num - pos
		if diff < 0 {
			diff = -diff
		}
		result += factorial(diff)
	}
	return result
}

func factorial(num int) int {
	sum := 0
	for num > 0 {
		sum += num
		num--
	}
	return sum
}

func sumDiffs(nums []int, pos int) int {
	sum := 0
	for _, num := range nums {
		diff := num - pos
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func getAverage(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum / len(nums)
}
