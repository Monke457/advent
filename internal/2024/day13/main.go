package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type machine struct {
	a [2]int
	b [2]int
	prize [2]int
}

const adjustment = 10000000000000

func main() {
	data := reader.FileToArray("data/2024/day13.txt")

	machines := parseMachines(data)

	total := 0
	for _, machine := range machines {
		tokens := machine.calculate(0)
		if tokens < 0 {
			continue
		}
		total += tokens
	}

	fmt.Println("Min tokens needed to win all prizes:", total)

	total = 0
	for _, machine := range machines {
		tokens := machine.calculate(adjustment)
		if tokens < 0 {
			continue
		}
		total += tokens
	}
	fmt.Println("Min tokens needed to win all prizes (adjusted):", total)
}

func multiplyCoords(coords [2]int, m int) (int, int) {
	return coords[0] * m, coords[1] * m
}

func (m machine) calculate(adjust int) int {
	c, d := m.prize[0] + adjust, m.prize[1] + adjust
	x1, y1, x2, y2 := m.a[0], m.a[1], m.b[0], m.b[1] 
	a := float64((c*y2 - d*x2)) / float64((x1*y2 - y1*x2))
	b := float64((d*x1 - c*y1)) / float64((x1*y2 - y1*x2))
	if math.Round(a) == a && math.Round(b) == b {
		return int(3 * a + b)
	}
	return -1
}

func parseMachines(data []string) []machine {
	machines := []machine{{}}

	i := 0
	for _, line := range data {
		if line == "" {
			machines = append(machines, machine{})
			i++
			continue
		}

		pos := parsePosition(line)

		if strings.Contains(line, "Button A") {
			machines[i].a = pos
		}
		if strings.Contains(line, "Button B") {
			machines[i].b = pos
		}
		if strings.Contains(line, "Prize") {
			machines[i].prize = pos
		}
	}

	return machines
}

func parsePosition(line string) [2]int {
	_, coords, _ := strings.Cut(line, ": ")
	parts := strings.Split(coords, ", ")
	xstr, ystr := parts[0][2:], parts[1][2:]
	x, _ := strconv.Atoi(xstr)
	y, _ := strconv.Atoi(ystr)
	return [2]int{x, y}
}
