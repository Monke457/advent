package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2016/day11.txt")

	floors := [][]string{}

	for _, line := range data {
		floor := []string{}
		parts := strings.Split(line, " ")
		for j, part := range parts {
			if strings.Contains(part, "compatible") {
				floor = append(floor, strings.ToUpper(part[0:1]) + "M")
			}
			if strings.Contains(part, "generator") {
				floor = append(floor, strings.ToUpper(parts[j-1][0:1]) + "G")
			}
		}
		floors = append(floors, floor)
	}

	el := 0
	for range 7 {
		fmt.Println(floors, el)
		if allAtTop(floors) {
			break
		}
		items, dest := getNextMove(floors, el)
		floors = makeMove(floors, items, el, dest)
	}
}

func makeMove(floors [][]string, items []string, pos, dest int) [][]string {
	return floors
}

func getNextMove(floors [][]string, el int) ([]string, int) {
	dest := getDestination(floors, el)
	items := getItems(floors, dest)

	return items, dest
} 

func getDestination(floors [][]string, el int) int {
	if el == 3 { return 2 }
	if el == 0 { return 1 }

	return el+1
}

func getItems(floors [][]string, dest int) []string {
	chips := map[byte]bool{}
	jennies := map[byte]bool{}
	for _, item := range floors[dest] {
		if item[1] == 'M' {
			chips[item[0]] = true
		} else {
			jennies[item[0]] = true
		}
	}
	items := []string{}
	return items
}

func allAtTop(floors [][]string) bool {
	for i := 0; i < len(floors)-1; i++ {
		if len(floors[i]) > 0 {
			return false
		}
	}
	return true
}

