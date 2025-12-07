package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArrayByDivider("data/2025/day2.txt", ",")

	var sum int = 0
	var vals [2]int

	for _, line := range data {
		vals = parseLine(line)
		for val := vals[0]; val <= vals[1]; val++ {
			if isRepeated(val) {
				sum += val
			}
		}
	}
	fmt.Println("invalid:", sum)
}

func parseLine(line string) [2]int {
	a, b, _ := strings.Cut(line, "-")
	val1, _ := strconv.Atoi(a)
	val2, _ := strconv.Atoi(b)
	return [2]int{val1, val2}
}

func isRepeated(val int) bool {
	str := strconv.Itoa(val)
	for i := 1; i <= len(str)>>1; i++ {
		if strings.Repeat(str[:i], len(str) / i) == str { 
			return true
		}
	}
	return false
}
