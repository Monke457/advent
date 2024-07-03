package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

const (
	BLUE  int = 14
	GREEN     = 13
	RED       = 12
)

func main() {
	fmt.Println(solveFirstProblem())
	fmt.Println(solveSecondProblem())
}

func solveFirstProblem() int {
	lines := reader.FileToArray("data/2023/day2.txt")

	sum := 0
	for _, l := range lines {
		game := strings.Split(l, ":")

		if len(game) != 2 {
			continue
		}

		results := strings.Split(game[1], ";")

		b, g, r := getCubeCounts(results)

		if b > BLUE || g > GREEN || r > RED {
			continue
		}

		id, err := strconv.Atoi(strings.Trim(game[0], "Game "))
		if err != nil {
			panic(err)
		}
		sum += id
	}

	return sum
}

func solveSecondProblem() int {
	lines := reader.FileToArray("data/2023/day2.txt")

	sum := 0
	for _, l := range lines {
		game := strings.Split(l, ":")

		if len(game) != 2 {
			continue
		}

		results := strings.Split(game[1], ";")

		b, g, r := getCubeCounts(results)
		sum += b * g * r
	}

	return sum
}

func getCubeCounts(games []string) (int, int, int) {
	var b, g, r int

	for _, pick := range games {
		colours := strings.Split(pick, ",")

		for _, colour := range colours {
			vals := strings.Split(strings.Trim(colour, " "), " ")
			count, err := strconv.Atoi(vals[0])

			if err != nil {
				panic(err)
			}

			switch vals[1] {
			case "blue":
				if count > b {
					b = count
				}
			case "green":
				if count > g {
					g = count
				}
			case "red":
				if count > r {
					r = count
				}
			}
		}
	}

	return b, g, r
}
