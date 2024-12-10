package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2018/day3.txt")


	claims := map[int][2][2]int{}

	for _, line := range data {
		id, claim := parseClaim(line)
		claims[id] = claim
	}

	claimed := mapClaims(claims)

	overlap := 0
	for _, ids := range claimed {
		if len(ids) > 1 {
			overlap++
		}
	}

	noOverlap := []int{}
	loop:
	for id, val := range claims {
		for y := val[0][0]; y < val[1][0]; y++ {
			for x := val[0][1]; x < val[1][1]; x++ {
				spot := [2]int{y, x}
				if len(claimed[spot]) > 1 {
					continue loop
				}
			}
		}
		noOverlap = append(noOverlap, id)
	}

	fmt.Println("Overlapping claims:", overlap)
	fmt.Println("No overlap:", noOverlap)

}

func mapClaims(claims map[int][2][2]int) map[[2]int][]int {
	claimed := map[[2]int][]int{}
	for key, val := range claims {
		for y := val[0][0]; y < val[1][0]; y++ {
			for x := val[0][1]; x < val[1][1]; x++ {
				spot := [2]int{y, x}
				if _, ok := claimed[spot]; !ok {
					claimed[spot] = []int{}
				}
				claimed[spot] = append(claimed[spot], key)
			}
		}
	}
	return claimed
}

func parseClaim(claim string) (int, [2][2]int) {
	parts := strings.Split(claim, " ")

	idStr := strings.Trim(parts[0], "#")
	id, _:= strconv.Atoi(idStr)

	xstr, ystr, _ := strings.Cut(strings.Trim(parts[2], ":"), ",") 
	y, _ := strconv.Atoi(ystr)
	x, _ := strconv.Atoi(xstr)

	wstr, hstr, _ := strings.Cut(parts[3], "x")
	w, _ := strconv.Atoi(wstr)
	h, _ := strconv.Atoi(hstr)

	return id, [2][2]int{{y, x}, {y+h, x+w}} 
}
