package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2017/day13.txt")

	layers := parseLayers(data) 

	fmt.Println("First:", calculateSeverity(layers, 0))

	r := 0 
	for _, layer := range layers {
		if layer[1] > r {
			r = layer[1]
		}
	}

	delay := 0
	for !isSafe(layers, delay) {
		delay++
	}
	fmt.Println("Second:", delay)
}

func isSafe(layers [][2]int, delay int) bool {
	for _, layer := range layers {
		sec := layer[0] + delay
		scanner := sec % ((layer[1] << 1) - 2)
		if scanner == 0 {
			return false 
		}
	}
	return true 
}

func calculateSeverity(layers [][2]int, delay int) int {
	severity := 0
	for _, layer := range layers {
		sec := layer[0] + delay
		scanner := sec % ((layer[1] << 1) - 2)
		if scanner == 0 {
			s := layer[0] * layer[1]
			severity += s 
		}
	}
	return severity
}

func parseLayers(data []string) [][2]int {
	res := [][2]int{}
	for _, line := range data {
		depth, rng := parseLine(line)
		res = append(res, [2]int{depth, rng}) 
	}
	return res
}

func parseLine(line string) (int, int) {
	d, r, _ := strings.Cut(line, ": ")
	depth, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}
	rng, err := strconv.Atoi(r)
	if err != nil {
		panic(err)
	}
	return depth, rng
}
