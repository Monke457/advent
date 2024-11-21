package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2019/day12.txt")

	for _, line := range data {
		fmt.Println(line)
	}
}
