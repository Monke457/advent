package main

import (
	"advent/pkg/reader"
	"advent/pkg/sorter"
	"fmt"
	"strconv"
	"strings"
)

// so sadge :(
func main() {
	data := reader.FileToArray("data/2015/day13.txt")

	prefs := parsePrefs(data)
	names := sorter.GetOrderByKey(prefs)

	matrix := make([][]*int, len(names))
	for i, n1 := range names {
		for j, n2 := range names {
			if j == i {
				matrix[i] = append(matrix[i], nil)
				continue
			}
			sum := prefs[n1][n2] + prefs[n2][n1]
			matrix[i] = append(matrix[i], &sum)
		}
	}
	fmt.Println(matrix)
}

func parsePrefs(data []string) map[string]map[string]int {
	res := map[string]map[string]int{}
	for _, l := range data {
		parts := strings.Split(l, " ")
		name := parts[0]
		neighbour := strings.Trim(parts[len(parts)-1], ".")
		val, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err)
		}
		if parts[2] == "lose" {
			val *= -1
		}
		if _, ok := res[name]; !ok {
			res[name] = map[string]int{}
		}
		res[name][neighbour] = val
	}
	return res
}
