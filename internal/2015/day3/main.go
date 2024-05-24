package main

import (
	"advent/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToString("data/2015/day3.txt")
	fmt.Println("First problem:", solveFirstProblem(data))
	fmt.Println("Second problem:", solveSecondProblem(data))
}

func solveFirstProblem(data string) int {
	curr := [2]int{0, 0}
	houses := map[[2]int]int{
		curr: 1,
	}
	for _, r := range data {
		switch r {
		case '^':
			curr[0] += 1
		case 'v':
			curr[0] -= 1
		case '>':
			curr[1] += 1
		case '<':
			curr[1] -= 1
		}
		houses[curr] += 1
	}
	return len(houses)
}

func solveSecondProblem(data string) int {
	santas := map[bool][2]int{
		true:  {0, 0},
		false: {0, 0},
	}
	houses := map[[2]int]int{
		santas[true]: 1,
	}
	for i := 0; i < len(data); i++ {
		s := i&1 == 0
		switch data[i] {
		case '^':
			n := santas[s][0]
			santas[s] = [2]int{n + 1, santas[s][1]}
		case 'v':
			n := santas[s][0]
			santas[s] = [2]int{n - 1, santas[s][1]}
		case '>':
			n := santas[s][1]
			santas[s] = [2]int{santas[s][0], n + 1}
		case '<':
			n := santas[s][1]
			santas[s] = [2]int{santas[s][0], n - 1}
		}
		houses[santas[s]] += 1
	}
	return len(houses)
}
