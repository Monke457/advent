package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2024/day5.txt")

	rules := map[int][]int{}
	sumOfMiddle := 0
	sumOfMiddleCorrected := 0
	endOfRules := false

	for _, line := range data {
		if line == "" {
			endOfRules = true
			continue
		}

		if !endOfRules {
			a, b, err := parseRule(line)
			if err != nil {
				fmt.Errorf("Failed to parse rules:", line, err)
				continue
			}
			if _, ok := rules[a]; !ok {
				rules[a] = []int{}
			}
			rules[a] = append(rules[a], b)
			continue
		}

		pages, err := parsePages(line)
		if err != nil {
			fmt.Errorf("Failed to parse pages:", line, err)
			continue
		}
		if _, _, ok := isCorrect(rules, pages); ok {
			sumOfMiddle += pages[len(pages)/2]
			continue
		}

		pages = corrected(rules, pages)
		sumOfMiddleCorrected += pages[len(pages)/2]
	}

	fmt.Println("Sum of middle pages:", sumOfMiddle)
	fmt.Println("Sum of middle pages corrected:", sumOfMiddleCorrected)
}

func parseRule(line string) (int, int, error) {
	aStr, bStr, _ := strings.Cut(line, "|")
	a, err := strconv.Atoi(aStr)
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(bStr)
	if err != nil {
		return 0, 0, err
	}
	return a, b, nil
}

func corrected(rules map[int][]int, pages []int) []int {
	corrected := make([]int, len(pages))
	copy(corrected, pages)
	for {
		idxA, idxB, corr := isCorrect(rules, corrected) 
		if corr {
			break
		}
		temp := []int{}
		temp = append(temp, corrected[:idxB]...)
		temp = append(temp, corrected[idxA])
		temp = append(temp, corrected[idxB:idxA]...)
		temp = append(temp, corrected[idxA+1:]...)
		corrected = append([]int{}, temp...) 
	}
	return corrected 
}

func parsePages(line string) ([]int, error) {
	result := []int{}
	pagesStr := strings.Split(line, ",")
	for _, p := range pagesStr {
		page, err := strconv.Atoi(p)
		if err != nil {
			return result, err
		}
		result = append(result, page)
	}
	return result, nil
}

func isCorrect(rules map[int][]int, pages []int) (int, int, bool) {
	for i := len(pages)-1; i > -1; i-- {
		page := pages[i]
		for j := 0; j < i; j++ {
			for _, r := range rules[page] {
				if pages[j] == r {
					return i, j, false
				}
			}
		}
	}
	return -1, -1, true 
}
