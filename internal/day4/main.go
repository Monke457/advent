package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solveFirstProblem())
	fmt.Println(solveSecondProblem())
}

func solveFirstProblem() int {
	lines := reader.FileToArray("data/day4.txt")

	points := 0
	for _, l := range lines {
		nums := strings.Split(strings.Split(l, ":")[1], "|")
		tar := strings.Split(nums[0], " ")
		vals := strings.Split(nums[1], " ")

		p := 0
		for _, v := range vals {
			if contains(tar, v) {
				if p == 0 {
					p++
				} else {
					p = p << 1
				}
			}
		}

		points += p
	}

	return points
}

func solveSecondProblem() int {
	lines := reader.FileToArray("data/day4.txt")

	m := map[int]int{}
	cards := 0
	for i, l := range lines {
		m[i] += 1
		nums := strings.Split(strings.Split(l, ":")[1], "|")
		tar := strings.Split(nums[0], " ")
		vals := strings.Split(nums[1], " ")

		p := 0
		for _, v := range vals {
			if contains(tar, v) {
				p++
			}
		}

		for j := 1; j <= p; j++ {
			m[i+j] += m[i] * 1
		}

		cards += m[i]
	}

	return cards
}

func contains(tar []string, val string) bool {
	for _, i := range tar {
		if i == "" {
			continue
		}
		if i == val {
			return true
		}
	}
	return false
}
