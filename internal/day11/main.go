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

	fmt.Println(solve(data, galaxies, 2))
	fmt.Println(solve(data, galaxies, 1000000))
}

func solve(data []string, galaxies map[int][2]int, offset int) int {
	newGalaxies := map[int][2]int{}
	for i := range galaxies {
		newGalaxies[i] = offsetGalaxy(data, galaxies[i], offset)
	}
	return calculateDistance(newGalaxies)
}

func calculateDistance(galaxies map[int][2]int) int {
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
	return sum
}

func offsetGalaxy(data []string, g [2]int, offset int) [2]int {
	result := [2]int{g[0], g[1]}

	for i := 0; i < g[0]; i++ {
		if data[i][0] == 42 {
			result[0] -= 1
			result[0] += offset
		}
	}
	stars := strings.Count(data[0][:g[1]], "*")
	result[1] -= stars
	result[1] += stars * offset

	return result
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
			column++
			if col[j] {
				sb.WriteString(s)
				continue
			}
			s = "*"
			sb.WriteString(s)
		}
		row++
		if strings.Contains(l, "#") {
			data = append(data, sb.String())
			continue
		}
		data = append(data, strings.Repeat("*", len(sb.String())))
	}
	return data, galaxies
}
