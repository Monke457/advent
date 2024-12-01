package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var cache = map[int]int{}

func main() {
	data := reader.FileToArray("data/2024/day1.txt")

	lists := parseLists(data)

	diff := 0
	for i := range lists[0] {
		diff += int(math.Abs(float64(lists[0][i]) - float64(lists[1][i])))
	}

	fmt.Println("Difference:", diff)

	similarity := 0
	for _, val1 := range lists[0] {
		if c, ok := cache[val1]; ok {
			similarity += c
			continue
		}
		count := 0
		for _, val2 := range lists[1] {
			if val1 == val2 {
				count++
			}
		}
		cache[val1] = val1 * count
		similarity += cache[val1]
	}

	fmt.Println("Similarity:", similarity)
}

func parseLists(data []string) [2][]int {
	result := [2][]int{}
	for _, line := range data {
		first, second, _ := strings.Cut(line, "   ")

		iFirst, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}
		iSecond, err := strconv.Atoi(second)
		if err != nil {
			panic(err)
		}

		result[0] = append(result[0], iFirst)
		result[1] = append(result[1], iSecond)
	}

	slices.Sort(result[0])
	slices.Sort(result[1])

	return result
}
