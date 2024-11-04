package main

import (
	"advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day5.txt")

	err := intcode.RunDay5(data, 1)
	if err != nil {
		fmt.Print(err)
	}
}
