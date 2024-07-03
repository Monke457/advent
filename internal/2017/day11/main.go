package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

var move = map[string]func() int {
	"n": func() int { return max(inc('q'), dec('r')) },
	"ne": func() int { return max(inc('q'), dec('s')) },
	"se": func() int { return max(inc('r'), dec('s')) }, 
	"s": func() int { return max(inc('r'), dec('q')) },
	"sw": func() int { return max(inc('s'), dec('q')) },
	"nw": func() int { return max(inc('s'), dec('r')) },
}

var position = map[rune]int{
	's': 0,
	'r': 0,
	'q': 0,
}

func main() {
	data := reader.FileToString("data/2017/day11.txt")

	dirs := strings.Split(data, ",")

	res2 := 0
	for _, dir := range dirs {
		distance := move[dir]()
		if distance > res2 {
			res2 = distance
		}
	}
	
	convertToAbs()
	res := 0
	for _, v := range position {
		if v > res {
			res = v
		}
	}
	fmt.Println("First:", res)
	fmt.Println("Second:", res2)
}

func inc(key rune) int {
	position[key]++
	return abs(position[key])
}

func dec(key rune) int {
	position[key]--
	return abs(position[key])
}

func convertToAbs() {
	for k, v := range position {
		if v < 0 {
			position[k] = -1 * v
		}
	}
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}
