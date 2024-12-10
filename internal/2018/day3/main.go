package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2018/day3.txt")

	for _, line := range data {
		fmt.Println(line)
	}
}
