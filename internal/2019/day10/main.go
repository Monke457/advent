package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
)

func main() {
	data := reader.FileTo2DArray("data/2019/day10.txt")

	seenAsteroids := [][2]int{}
	seesMost := make([]int, 2)

	for y := range data {
		for x := range data[y] {
			if data[y][x] != '#' {
				continue
			}
			seen := getSeenAsteroids(data, [2]int{y, x})

			if len(seen) > len(seenAsteroids) {
				seesMost[0] = y
				seesMost[1] = x
				seenAsteroids = seen[:]
			} 
		} 
	}

	printMap(data, seesMost, seenAsteroids)

}

func printMap(data [][]rune, center []int, seen [][2]int) {
	fmt.Println("MAP OF POSITION", center, "SEES:", len(seen), "ASTEROIDS")
	for y, line := range data {
		for x, val := range line {
			if center[0] == y && center[1] == x {
				fmt.Printf("\033[31m%c\033[0m", val)
				continue
			}
			coord := [2]int{y, x}
			if slices.Contains(seen, coord) {
				fmt.Printf("\033[32m%c\033[0m", val)
				continue
			}
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}

func getSeenAsteroids(asteroids [][]rune, pos [2]int) [][2]int {
	seen := map[float64]float64{}
	coords := map[float64][2]int{}
	for y := range asteroids {
		for x := range asteroids[y] {
			if asteroids[y][x] != '#' {
				continue
			}
			fy := float64(y - pos[0])
			fx := float64(x - pos[1])
			if fx == 0 && fy == 0 {
				continue
			}
			angle := calculateAngle(fy, fx)
			if dist, ok := seen[angle]; ok { 
				if dist < math.Abs(fx) + math.Abs(fy) {
					continue
				}
			}
			seen[angle] = math.Abs(fx) + math.Abs(fy)
			coords[angle] = [2]int{y, x}
		}
	} 
	result := [][2]int{}
	for _, coord := range coords {
		result = append(result, coord)
	}

	return result 
}


func calculateAngle(y, x float64) float64 {
	if x == 0 && y == 0 {
		panic("there should not be two 0 values here :(")
	}
	if x == 0 {
		if y > 0 {
			return 90
		}
		return 270
	}
	if y == 0 {
		if x > 0 {
			return 0
		}
		return 180
	}
	degrees := math.Atan(math.Abs(y)/math.Abs(x)) * 180 / math.Pi
	if x > 0 {
		if y > 0 {
			return degrees
		}
		return 270 + degrees
	}

	if y > 0 {
		return 90 + degrees
	}
	return 180 + degrees
}
