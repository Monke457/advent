package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type race struct {
	time int
	dist int
}

func main() {
	content := reader.FileToArray("data/day6.txt")

	fmt.Println(solveFirstProblem(content))
	fmt.Println(solveSecondProblem(content))
}

func solveFirstProblem(c []string) int {
	races := parseRaces(c)
	final := 0
	for _, race := range races {
		waysToBeat := countWaysToBeat(race)
		if final == 0 {
			final = waysToBeat
		} else {
			final *= waysToBeat
		}
	}

	return final
}

func solveSecondProblem(c []string) int {
	time := parseSingle(strings.SplitAfter(c[0], ":")[1])
	dist := parseSingle(strings.SplitAfter(c[1], ":")[1])

	return countWaysToBeat(race{time: time, dist: dist})
}

func countWaysToBeat(race race) int {
	count := 0
	for i := 0; i < race.time; i++ {
		t := i * (race.time - i)
		if t > race.dist {
			count++
		} else if count > 0 {
			break
		}
	}
	return count
}

func parseRaces(c []string) []race {
	races := []race{}

	times := parseLine(strings.Split(c[0], ":")[1])
	dists := parseLine(strings.Split(c[1], ":")[1])

	for i := 0; i < len(times); i++ {
		time := parseValue(times[i])
		dist := parseValue(dists[i])
		races = append(races, race{time: time, dist: dist})
	}

	return races
}

func parseValue(c string) int {
	conv, err := strconv.Atoi(c)
	if err != nil {
		panic(err)
	}
	return conv
}

func parseLine(line string) []string {
	vals := strings.Split(line, " ")

	for i := 0; i < len(vals); {
		vals[i] = strings.Trim(vals[i], " ")
		if vals[i] == "" {
			vals = append(vals[:i], vals[i+1:]...)
		} else {
			i++
		}
	}
	return vals
}

func parseSingle(line string) int {
	v := ""
	v = strings.ReplaceAll(line, " ", v)
	return parseValue(v)
}
