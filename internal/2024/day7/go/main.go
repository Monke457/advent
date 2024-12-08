package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

var operators = map[rune]func(int, int)int{
	'+': func(a, b int) int {return a+b},
	'*': func(a, b int) int {return a*b},
	'|': concat,
}

func main() {
	data := reader.FileToArray("data/2024/day7.txt")

	sum := 0

	for _, line := range data {
		total, vals, err := parseEquation(line)
		if err != nil {
			panic(err)
		}
		if ValidEquation(total, vals) {
			sum += total
			fmt.Printf("\rSum of valid equations: %d     ", sum)
		}
	}
	fmt.Println("\rSum of valid equations:", sum)
}

func ValidEquation(total int, vals []int) bool {
	if len(vals) == 0 {
		return false
	}
	if len(vals) == 1 {
		return total == vals[0]
	}
	if checkCalc(total, vals) {
		return true
	}
	for _, op := range operators {
		newVals := make([]int, len(vals)-1)
		copy(newVals, vals[1:])
		newVals[0] = op(vals[0], vals[1])
		if ValidEquation(total, newVals) {
			return true
		}
	}
	return false
}

func checkCalc(total int, vals []int) bool {
	for key := range operators {
		if total == applyCalculations(vals, key) {
			return true 
		} 
	}
	return false
}

func applyCalculations(vals []int, op rune) int {
	result := vals[0]
	for i := 1; i < len(vals); i++ {
		result = operators[op](result, vals[i]) 
	}
	return result
}

func parseEquation(equation string) (int, []int, error) {
	var total int
	vals := []int{}

	a, b, ok := strings.Cut(equation, ": ")
	if !ok {
		return total, vals, fmt.Errorf("Error: cut failed - ': ' not found in string '%s'", equation) 
	}

	total, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	strVals := strings.Split(b, " ")
	for _, strVal := range strVals {
		val, err := strconv.Atoi(strVal)
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}

	return total, vals, nil 
}

func concat(a, b int) int {
	comb := fmt.Sprintf("%d%d", a, b)
	val, err := strconv.Atoi(comb)
	if err != nil {
		panic(err)
	}
	return val
}
