package main

import (
	"advent/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2016/day11.txt")
	for _, line := range data {
		fmt.Println(line)
	}
}
