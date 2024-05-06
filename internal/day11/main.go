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

	fmt.Println(solveFirstProblem(galaxies))
	fmt.Println(solveSecondProblem(data, galaxies))
}

func solveFirstProblem(galaxies map[int][2]int) int {
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

func solveSecondProblem(data []string, galaxies map[int][2]int) int {
	sum := 0
	for i := 1; i <= len(galaxies); i++ {
		for j := i; j <= len(galaxies); j++ {
			if i == j {
				continue
			}
			a := offsetGalaxy(data, galaxies[i])
			b := offsetGalaxy(data, galaxies[j])
			fmt.Println(i, ":", a, j, ":", b)
			sum += math.AbsDiff(a[0], b[0])
			sum += math.AbsDiff(a[1], b[1])
		}
	}
	return sum
}

func offsetGalaxy(data []string, g [2]int) [2]int {
	result := [2]int{g[0], g[1]}

	for i := 0; i < g[0]; i++ {
		if data[i][0] == 42 {
			result[0] -= 1
			result[0] += 10
		}
	}
	stars := strings.Count(data[0][:g[1]], "*")
	result[1] -= stars
	result[1] += stars * 10

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
			sb.WriteString(s)
			column++
			if col[j] {
				continue
			}
			sb.WriteString("*")
			column++
		}
		data = append(data, sb.String())
		row++
		if strings.Contains(l, "#") {
			continue
		}
		data = append(data, strings.Repeat("*", len(sb.String())))
		row++
	}
	return data, galaxies
}
