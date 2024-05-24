package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type present struct {
	dims [3]int
}

func main() {
	data := reader.FileToArray("data/2015/day2.txt")
	presents := parsePresents(data)

	fmt.Println("First problem:", solveFirstProblem(presents))
	fmt.Println("Second problem:", solveSecondProblem(presents))
}

func solveFirstProblem(presents []present) int {
	pape := 0
	for _, pres := range presents {
		a := pres.dims[0] * pres.dims[1]
		b := pres.dims[0] * pres.dims[2]
		c := pres.dims[1] * pres.dims[2]
		sum := 2 * (a + b + c)
		sum += min(a, b, c)
		pape += sum
	}
	return pape
}

func solveSecondProblem(presents []present) int {
	ribbon := 0
	for _, pres := range presents {
		a, b := getMins(pres.dims)
		sum := 2 * (a + b)
		sum += pres.dims[0] * pres.dims[1] * pres.dims[2]
		ribbon += sum
	}
	return ribbon
}

func getMins(dim [3]int) (int, int) {
	a := 0
	b := 0
	if dim[0] < dim[1] {
		a = dim[0]
		b = min(dim[1], dim[2])
	} else {
		a = dim[1]
		b = min(dim[0], dim[2])
	}
	return a, b
}

func parsePresents(data []string) []present {
	res := []present{}
	for _, l := range data {
		dims := strings.Split(l, "x")
		iDims := [3]int{}
		for i, d := range dims {
			if i > len(iDims) {
				panic("Too many dimensions!!")
			}
			iDim, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}
			iDims[i] = iDim
		}
		res = append(res, present{dims: iDims})
	}
	return res
}
