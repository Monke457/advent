package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
)

func main() {
	data := reader.FileToArray("data/2020/day12.txt")

	//directions 0=N, 1=E, 2=S, 3=W
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
			dir += (val / 90) % 3
		case 'L':
			dir -= (val / 90) % 3
			dir = int(math.Abs(float64(dir)))
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
	fmt.Println("First:", math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}

func parseLine(line string) (byte, int) {
	op, value := line[0], line[1:]
	val, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return op, val
}
