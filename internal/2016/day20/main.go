package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2016/day20.txt")

	ipRanges := parseData(data)

	slices.SortFunc(ipRanges, rangeSort)

	count := ipRanges[0][0]
	curr := ipRanges[0]

	var valid int
	for i := 1; i < len(ipRanges); i++ {
		if ipRanges[i][0] <= curr[1] + 1 {
			curr[1] = max(ipRanges[i][1], curr[1])
		} else {
			if valid == 0 {
				valid = curr[1] + 1
			}
			count += ipRanges[i][0] - (curr[1] + 1)
			curr = ipRanges[i]
		}
	}

 	count += 4294967295 - curr[1]

	fmt.Println("smallest valid ip:", valid)
	fmt.Println("ammount allowed:", count)
}

func rangeSort(a, b [2]int) int {
	if a[0] > b[0] { return 1 }
	if a[0] < b[0] { return -1 }
	if a[1] > b[1] { return 1 }
	return -1
}

func parseData(data []string) [][2]int {
	ipRanges := [][2]int{}
	for _, line := range data {
		aStr, bStr, _ := strings.Cut(line, "-")
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)

		ipRanges = append(ipRanges, [2]int{a, b})
	}
	return ipRanges
}
