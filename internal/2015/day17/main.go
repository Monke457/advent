package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

const (
	TARGET int = 150
	TARGET_TEST = 25
)

func main() {
	containers := reader.FileToIntArrayByDivider("data/2015/day17.txt", "\n")
	//containers := reader.FileToIntArrayByDivider("data/2015/day17_test.txt", "\n")

	cMap :=	mapValues(containers)
	combos := getCombinations(containers, cMap, TARGET, 0)

	fmt.Println(cache)
	minValue := getMin(cache)

	fmt.Println("Total ways to fill containers", TARGET, "=", combos)
	fmt.Println("Way to fill min no. of containers", minValue)
}

func getMin(cache map[int]int) int {
	var val int
	key := math.MaxInt
	for k, v := range cache {
		if k < key {		
			key = k
			val = v
		}
	} 
	return val
}

var cache = map[int]int{}

func getCombinations(values []int, cMap map[int]int, target, n int) int {
	if target == 0 {
		cache[n]++
		return  1
	}
	if target < 0 || len(values) == 0 {
		return 0
	}
	combos := 0
	for i, v := range values {
		c := getCombinations(values[i+1:], cMap, target - v, n+1)
		if cMap[v] > 1 {
			c *= cMap[v]
			cMap[v]--
		}
		combos += c
	}
	return combos
}

func mapValues(values []int) map[int]int {
	res := map[int]int{}
	for _, v := range values {
		res[v]++
	}
	return res
}
