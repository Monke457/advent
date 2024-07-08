package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
)

func main() {
	data := reader.FileToArray("data/2020/day12.txt")

	fmt.Println("First:", first(data))
	fmt.Println("Second:", second(data))
}

func first(data []string) float64 {
	dir := 1 
	pos := [2]int{0, 0} 

	for _, line := range data {
		op, val := parseLine(line)
		switch op {
		case 'N':
			pos[0] += val
		case 'E':
			pos[1] += val
		case 'S':
			pos[0] -= val
		case 'W':
			pos[1] -= val
		case 'R':
			dir = wrap(dir + val / 90)
		case 'L':
			dir = wrap(dir - val / 90)
		case 'F':
			switch dir {
			case 0:
				pos[0] += val
			case 1:
				pos[1] += val
			case 2:
				pos[0] -= val
			case 3:
				pos[1] -= val
			}
		}
	}
	return math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
}

func second(data []string) float64 {
	wp := [4]int{ 1, 10, 0, 0 } 
	pos := [2]int{ 0, 0 } 

	for _, line := range data {
		op, val := parseLine(line)
		switch op {
		case 'N':
			wp[0] += val	
		case 'E':
			wp[1] += val 
		case 'S':
			wp[2] += val
		case 'W':
			wp[3] += val
		case 'R':
			wp = wrapWP(wp, val / 90)
		case 'L':
			wp = wrapWP(wp, -1 * val / 90)
		case 'F':
			pos[0] += val * wp[0]
			pos[1] += val * wp[1]
			pos[0] -= val * wp[2]
			pos[1] -= val * wp[3]
		}
	}
	return math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
}

func wrap(val int) int {
	if val > 3 {
		return val % 4
	}
	for val < 0 {
		val = 4 + val 
	}
	return val
}

func wrapWP(wp [4]int, val int) [4]int {
	return [4]int{ wp[wrap(0 - val)], wp[wrap(1 - val)], wp[wrap(2 - val)], wp[wrap(3 - val)] }
}

func parseLine(line string) (byte, int) {
	op, value := line[0], line[1:]
	val, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return op, val
}
