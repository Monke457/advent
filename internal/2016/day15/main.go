package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type disc struct {
	positions int
	start int
}

func main() {
	data := reader.FileToArray("data/2016/day15.txt")

	discs := []disc{}
	for _, line := range data {
		disc := parseDisc(line)
		discs = append(discs, disc)
	}

	t := 0
	for {
		if dropCapsule(discs, t) {
			break
		}
		t++
	}

	fmt.Println("Capsule dropped after", t, "seconds")
}

func dropCapsule(discs []disc, t int) bool {
	for i, disc := range discs {
		if (disc.start + i+1 + t) % disc.positions > 0 {
			return false
		}
	}
	return true
}

func parseDisc(str string) disc {
	parts := strings.Split(str, " ")
	positions, _ := strconv.Atoi(parts[3])
	start, _ := strconv.Atoi(parts[11][:len(parts[11])-1])
	return disc {
		positions: positions,
		start: start,
	} 
}
