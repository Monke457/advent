package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArray("data/2021/day1.txt")

	var prev int
	count := 0
	for i := 0; i < len(data) - 2; i++ {
		if i == 0 {
			prev = data[i] + data[i+1] + data[i+2]
			continue
		}
		next := data[i] + data[i+1] + data[i+2]
		if next > prev {
			count++
		}
		prev = next
	}
	fmt.Println("Second:", count)
}
