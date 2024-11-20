package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//step 0: prepare data
	data := reader.FileToArray("data/2019/day3.txt")
	wires := parseWires(data)

	//step 1: plot coordinates of each wire
	coordinates1 := plotCoords(wires[0])
	coordinates2 := plotCoords(wires[1])

	//step 2: find common points (intersections)
	intersections := findIntersections(coordinates1, coordinates2)

	//step 3: find closest intersection to central point (0,0)
	closest := math.MaxInt
	shortest := math.MaxInt
	for steps, intersection := range intersections {
		if intersection == [2]int{0,0} {
			continue
		}
		dist := int(math.Abs(float64(intersection[0])) + math.Abs(float64(intersection[1])))
		if dist < closest {
			closest = dist
		}
		if steps < shortest {
			shortest = steps
		}
	}

	fmt.Println("Part 1:", closest)
	fmt.Println("Part 2:", shortest)
}

func parseWires(data []string) [2][]string {
	wires := [2][]string{}
	for i, line := range data {
		wire := []string{}
		for _, cmd := range strings.Split(line, ",") {
			wire = append(wire, cmd)
		}
		wires[i] = wire
	}
	return wires
}

func plotCoords(wire []string) map[int]map[int]int {
	coords := map[int]map[int]int{0: {0: 0}}
	
	steps := 0
	curr := [2]int{0,0}
	for _, cmd := range wire {
		dist, err := strconv.Atoi(cmd[1:])
		if err != nil {
			fmt.Println(fmt.Errorf("Error while converting string to int", "string", cmd[1:], "err", err))
			continue
		}

		switch cmd[0] {
		case 'R':
			for i := range dist {
				steps++
				key := curr[0] + i + 1
				if _, ok := coords[key]; !ok {
					coords[key] = map[int]int{}
				}
				coords[key][curr[1]] = steps
			}
			curr[0] = curr[0] + dist
			break
		case 'L':
			for i := range dist {
				steps++
				key := curr[0] - i - 1
				if _, ok := coords[key]; !ok {
					coords[key] = map[int]int{}
				}
				coords[key][curr[1]] = steps
			}
			curr[0] = curr[0] - dist
			break
		case 'U':
			for i := range dist {
				steps++
				coords[curr[0]][curr[1] + i + 1] = steps
			}
			curr[1] = curr[1] + dist
			break
		case 'D':
			for i := range dist {
				steps++
				coords[curr[0]][curr[1] - i - 1] = steps
			}
			curr[1] = curr[1] - dist
			break
		}
	}

	return coords
}

func findIntersections(coords1 map[int]map[int]int, coords2 map[int]map[int]int) map[int][2]int {
	results := map[int][2]int{}

	for x, ys1 := range coords1 {
		if ys2, ok := coords2[x]; ok {
			if len(ys1) < len(ys2) {
				for y, steps1 := range ys1 {
					if steps2, ok := ys2[y]; ok {
						results[steps1 + steps2] = [2]int{x, y}
					}
				}
			} else {
				for y, steps2 := range ys2 {
					if steps1, ok := ys1[y]; ok {
						results[steps1 + steps2] = [2]int{x, y}
					}
				}
			}
		} else {
			continue
		}
	}
	return results
}



