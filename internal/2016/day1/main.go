package main

import (
	"advent/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	north int = 0
	east      = 1
	south     = 2
	west      = 3
)

func main() {
	data := reader.FileToString("data/2016/day1.txt")

	solveFirstProblem(data)
	solveSecondProblem(data)
}

func solveFirstProblem(data string) {
	var x, y, dir int
	for _, v := range strings.Split(data, ",") {
		str := strings.Trim(v, " ")
		d, val := str[0], str[1:]
		numVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		if d == 'L' {
			dir = (dir + 3) % 4
		}
		if d == 'R' {
			dir = (dir + 1) % 4
		}
		switch dir {
		case north:
			y += numVal
		case east:
			x += numVal
		case south:
			y -= numVal
		case west:
			x -= numVal
		}
	}
	result := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Println(result)
}

func solveSecondProblem(data string) {
	var seen = map[[2]int]bool{}
	var x, y int
	var dir int
loop:
	for _, v := range strings.Split(data, ",") {
		str := strings.Trim(v, " ")
		d, val := str[0], str[1:]
		numVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		if d == 'L' {
			dir = (dir + 3) % 4
		}
		if d == 'R' {
			dir = (dir + 1) % 4
		}
		switch dir {
		case north:
			for range numVal {
				y++
				if seen[[2]int{y, x}] {
					break
				}
				seen[[2]int{y, x}] = true
			}
		case east:
			for range numVal {
				x++
				if seen[[2]int{y, x}] {
					break loop
				}
				seen[[2]int{y, x}] = true
			}
		case south:
			for range numVal {
				y--
				if seen[[2]int{y, x}] {
					break loop
				}
				seen[[2]int{y, x}] = true
			}
		case west:
			for range numVal {
				x--
				if seen[[2]int{y, x}] {
					break loop
				}
				seen[[2]int{y, x}] = true
			}
		}
	}
	result := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Println(result)
}
