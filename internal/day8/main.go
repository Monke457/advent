package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

const (
	L rune = 76
	R      = 82
)

func main() {
	lines := reader.FileToArray("data/day8.txt")

	dirs := parseDirections(lines[0])
	m := parseMap(lines[1:])

	fmt.Println(solveFirstProblem(dirs, m))
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
