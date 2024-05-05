package main

import (
	"advent/internal/pkg/math"
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	lines := reader.FileToArray("data/day11.txt")

	data, galaxies := parseData(lines)

	fmt.Println(galaxies)
	for _, d := range data {
		fmt.Println(d)
	}

	sum := 0
	for i := 1; i <= len(galaxies); i++ {
		for j := i; j <= len(galaxies); j++ {
			if i == j {
				continue
			}
			sum += math.AbsDiff(galaxies[j][0], galaxies[i][0])
			sum += math.AbsDiff(galaxies[j][1], galaxies[i][1])
		}
	}
	fmt.Println("Sum:", sum)
}

func parseData(lines []string) ([]string, map[int][2]int) {
	data := []string{}
	col := map[int]bool{}
	for _, l := range lines {
		for i, r := range l {
			if r == 35 {
				col[i] = true
			}
		}
	}
	count := 0
	row := 0
	galaxies := map[int][2]int{}
	for _, l := range lines {
		column := 0
		sb := strings.Builder{}
		for j, r := range l {
			s := "."
			if r == 35 {
				count++
				s = fmt.Sprintf("%d", count)
				galaxies[count] = [2]int{row, column}
			}
			sb.WriteString(s)
			column++
			if col[j] {
				continue
			}
			sb.WriteString(s)
			column++
		}
		data = append(data, sb.String())
		row++
		if strings.Contains(l, "#") {
			continue
		}
		data = append(data, sb.String())
		row++
	}
	return data, galaxies
}
