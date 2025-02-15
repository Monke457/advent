package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2022/day4.txt")

	full := 0
	overlap := 0

	for _, line := range data {
		elfa, elfb, _ := strings.Cut(line, ",")

		a1s, a2s, _ := strings.Cut(elfa, "-")
		b1s, b2s, _ := strings.Cut(elfb, "-")

		a1, _ := strconv.Atoi(a1s)
		a2, _ := strconv.Atoi(a2s)
		b1, _ := strconv.Atoi(b1s)
		b2, _ := strconv.Atoi(b2s)

		if (a1 < b2 && a2 < b1) || (a1 > b2 && a2 > b1) {
			continue
		}
		if (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2) {
			full++
		}
		overlap++
	}

	fmt.Println("number of fully nested section pairs:", full)
	fmt.Println("number of overlapping section pairs:", overlap)
}
