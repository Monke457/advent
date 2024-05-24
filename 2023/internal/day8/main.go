package main

import (
	"advent/pkg/math"
	"advent/pkg/reader"
	"fmt"
	"strings"
)

const (
	L rune = 76
	R      = 82
)

func main() {
	lines := reader.FileToArray("data/2023/day8.txt")

	dirs := parseDirections(lines[0])
	m := parseMap(lines[1:])

	fmt.Println(solveFirstProblem(dirs, m))
	fmt.Println(solveSecondProblem(dirs, m))
}

func solveFirstProblem(dirs []rune, m map[string][2]string) int {
	steps := 0
	dir := 0
	pos := "AAA"
	for i := 0; pos != "ZZZ"; {
		if i >= len(dirs) {
			i = 0
		}
		if dirs[i] == L {
			dir = 0
		} else {
			dir = 1
		}
		pos = m[pos][dir]
		steps++
		i++
	}
	return steps
}

func solveSecondProblem(dirs []rune, m map[string][2]string) int {
	start := startingPositions(m)
	c := make(chan int)

	for _, s := range start {
		go walk(dirs, s, m, c)
	}

	result := []int{}
	count := 0
	for {
		f := <-c
		count++
		result = append(result, f)
		if count == len(start) {
			return math.LCD(result[:]...)
		}
	}
}

func walk(dirs []rune, s string, m map[string][2]string, c chan<- int) {
	steps := 0
	dir := 0
	i := 0
	for {
		steps++
		if i >= len(dirs) {
			i = 0
		}
		if dirs[i] == L {
			dir = 0
		} else {
			dir = 1
		}
		s = m[s][dir]
		if s[2] == 90 {
			c <- steps
			return
		}
		i++
	}
}

func parseDirections(line string) []rune {
	dirs := []rune{}
	for _, r := range line {
		dirs = append(dirs, r)
	}
	return dirs
}

func parseMap(lines []string) map[string][2]string {
	m := map[string][2]string{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		l := strings.Split(line, "=")
		name := strings.Trim(l[0], " ")
		dests := strings.Trim(l[1], " (")
		dests = strings.Trim(dests, ")")
		destArr := strings.Split(dests, ",")
		left := destArr[0]
		right := strings.Trim(destArr[1], " ")

		m[name] = [2]string{left, right}
	}

	return m
}

func startingPositions(m map[string][2]string) []string {
	start := []string{}
	for key := range m {
		if key[2] == 65 {
			start = append(start, key)
		}
	}
	return start
}
