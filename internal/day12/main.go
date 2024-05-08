package main

import(
	"fmt"
	"strings"
	"strconv"
	"advent/internal/pkg/reader"
)

type config struct {
	springs string
	broken []int
	arrangements []string
}

func main() {
	lines := reader.FileToArray("data/day12.txt")
	configs := parseConfig(lines)

	fmt.Println(solveFirstProblem(configs))

	for _, c := range configs {
		fmt.Println("springs:", c.springs, " broken: ", c.broken)

	}
}

func solveFirstProblem(configs []config) int {
	sum := 0
	for _, config := range configs {
		sum += config.GetArrangements()
	}
	return sum
}

func (c config) GetArrangements() int {
	return 1
}

func parseConfig(lines []string) []config {
	configs := []config{}

	for _, l := range lines {
		line := strings.Split(l, " ")
		springs := strings.Trim(line[0], " ")
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
