package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2024/day2.txt")

	reports := parseReports(data)

	count := 0
	withDampener := 0
	for _, report := range reports {
		if reportSafe(report) {
			count++
			continue
		}
		if safeWithDampener(report) {
			withDampener++
		}
	}
	fmt.Println("Safe reports:", count)
	fmt.Println("Safe reports with problem dampener:", count + withDampener)
}

func parseReports(data []string) [][]int {
	result := [][]int{}
	for _, line := range data {
		report := []int{}
		for _, level := range strings.Split(line, " ") {
			num, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			report = append(report, num)
		}
		result = append(result, report)
	}
	return result
}

func safeWithDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		dampened := make([]int, len(report))
		copy(dampened, report)
		dampened = append(dampened[:i], dampened[i+1:]...)
		if reportSafe(dampened) {
			return true
		}
	}
	return false
}

func reportSafe(report []int) bool {
	l := len(report)
	if l < 2 {
		return true
	}

	diff := report[0] - report[1]
	if diff == 0 || math.Abs(float64(diff)) > 3 {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		next := report[i] - report[i+1]
		if intAsBool(diff) != intAsBool(next) { 
			return false
		}
		if next == 0 {
			return false
		}
		if math.Abs(float64(next)) > 3 {
			return false
		}
		diff = next
	}

	return true
}

//true if positive otherwise false
func intAsBool(i int) bool {
	if i > 0 {
		return true
	}
	return false
}
