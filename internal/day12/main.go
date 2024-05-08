package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type config struct {
	springs      string
	broken       []int
	arrangements []string
}

func main() {
	lines := reader.FileToArray("data/day12.txt")
	configs := parseConfig(lines)

	for _, c := range configs {
		fmt.Println("springs:", c.springs, " broken: ", c.broken)
	}

	fmt.Println(solveFirstProblem(configs))
}

func solveFirstProblem(configs []config) int {
	sum := 0
	for _, config := range configs {
		sum += config.GetArrangements()
	}
	return sum
}

func (c config) GetArrangements() int {
	arrs := []string{}
	pieces := strings.Split(c.springs, ".")
	for _, i := range c.broken {
		fmt.Println(i, "looking for pattern in", pieces)
	}
	return len(arrs)
}

func parseConfig(lines []string) []config {
	configs := []config{}
	re := regexp.MustCompile(`\.+`)
	for _, l := range lines {
		line := strings.Split(l, " ")
		springs := strings.TrimSpace(line[0])
		springs = re.ReplaceAllString(springs, ".")

		broken := strings.Split(line[1], ",")
		brokenInt := []int{}
		for _, b := range broken {
			bInt, err := strconv.Atoi(b)
			if err != nil {
				panic(err)
			}
			brokenInt = append(brokenInt, bInt)
		}
		configs = append(configs, config{springs: springs, broken: brokenInt})
	}
	return configs
}
