package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2021/day2.txt")

	fmt.Println("First:", first(data))
	fmt.Println("Second:", second(data))
}

func first(data []string) int {
	hor, dep := 0, 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
			case "forward":
				hor += val
			case "up":
				dep -= val
			case "down":
				dep += val
		}
	}
	return hor * dep
}

func second(data []string) int {
	hor, dep, aim := 0, 0, 0
	for _, line := range data {
		parts := strings.Split(line, " ")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
			case "forward":
				hor += val
				dep += aim * val 
			case "up":
				aim -= val
			case "down":
				aim += val
		}
	}
	return hor * dep
}

