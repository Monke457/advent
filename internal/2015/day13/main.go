package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

// so sadge :(
func main() {
	data := reader.FileToArray("data/2015/day13.txt")

	for _, line := range data {
		fmt.Println(line)
	}
}
