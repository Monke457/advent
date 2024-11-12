package main

import (
	"advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day5.txt")

	var err error
	err = intcode.RunDay5(data, 5)
	if err != nil {
		fmt.Print(err)
	}
}
