package main

import (
	"advent/pkg/reader"
	"cmp"
	"fmt"
	"slices"
)

func main() {
	data := reader.FileTo2DIntArray("data/2016/day3.txt")
	solveFirstProblem(data)
	solveSecondProblem(data)
}

func solveFirstProblem(data [][]int) {
	poss := 0
	for _, row := range data {
		tri := make([]int, len(row))
		copy(tri, row)
		if validTriangle(tri) {
			poss++
		}
	}
	fmt.Println(poss)
}

func solveSecondProblem(data [][]int) {
	poss := 0
	for i := 0; i < len(data); i += 3 {
		for j := 0; j < len(data[i]); j++ {
			tri := []int{data[i][j], data[i+1][j], data[i+2][j]}
			if validTriangle(tri) {
				poss++
			}
		}
	}
	fmt.Println(poss)
}

func validTriangle(t []int) bool {
	slices.SortFunc(t, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	if t[0]+t[1] > t[2] {
		return true
	}
	return false
}
