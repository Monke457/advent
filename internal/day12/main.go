package main

import(
	"fmt"
	"advent/internal/pkg/reader"
)
func main() {
	lines := reader.FileToArray("data/day12.txt")
	for _, l := range lines {
		fmt.Println(l)
	}
}
