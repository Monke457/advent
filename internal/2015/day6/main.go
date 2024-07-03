package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

const (
	on     int = 0
	off        = 1
	toggle     = 2
)

type instruction struct {
	op    int
	start [2]int
	end   [2]int
}

func main() {
	data := reader.FileToArray("data/2015/day6.txt")

	instructions := parseInstructions(data)

	fmt.Println("First problem:", solveFirstProblem(instructions))
	fmt.Println("Second problem:", solveSecondProblem(instructions))
}

func solveFirstProblem(instructions []instruction) int {
	lights := map[[2]int]bool{}
	for _, inst := range instructions {
		for row := inst.start[0]; row <= inst.end[0]; row++ {
			for col := inst.start[1]; col <= inst.end[1]; col++ {
				val := lights[[2]int{row, col}]
				switch inst.op {
				case 0:
					//turn on
					val = true
				case 1:
					//turn off
					val = false
				case 2:
					//toggle
					val = !val
				}
				if !val {
					delete(lights, [2]int{row, col})
				} else {
					lights[[2]int{row, col}] = true
				}
			}
		}
	}
	return len(lights)
}

func solveSecondProblem(instructions []instruction) int {
	lights := map[[2]int]int{}
	for _, inst := range instructions {
		for row := inst.start[0]; row <= inst.end[0]; row++ {
			for col := inst.start[1]; col <= inst.end[1]; col++ {
				val := lights[[2]int{row, col}]
				switch inst.op {
				case 0:
					//turn up 1
					val++
				case 1:
					//turn down 1
					val = max(0, val-1)
				case 2:
					//turn up 2
					val += 2
				}
				lights[[2]int{row, col}] = val
			}
		}
	}

	brightness := 0
	for _, v := range lights {
		brightness += v
	}
	return brightness
}

func parseInstructions(data []string) []instruction {
	res := []instruction{}

	for _, l := range data {
		parts := strings.Split(l, " ")
		endstr := strings.Split(parts[len(parts)-1], ",")
		startstr := strings.Split(parts[len(parts)-3], ",")

		op := getOperation(parts[:2])
		end := convertToInt(endstr)
		start := convertToInt(startstr)

		res = append(res, instruction{op: op, start: start, end: end})
	}

	return res
}

func getOperation(arr []string) int {
	if len(arr) != 2 {
		panic("Expecting slice of size 2")
	}
	if arr[0] == "toggle" {
		return toggle
	}
	if arr[1] == "on" {
		return on
	}
	return off
}

func convertToInt(arr []string) [2]int {
	if len(arr) != 2 {
		panic("Expecting array of size 2")
	}
	res := [2]int{}
	for i, s := range arr {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res[i] = n
	}
	return res
}
