package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"slices"
)

func main() {
	data := reader.FileTo2DArray("data/2019/day10.txt")

	center := findCenter(data)
	visible := getSortedVisible(data, center)

	fmt.Println("First:", len(visible))

	printMap(data, center, visible)

	pos := vaporizeAndGet200thPos(data, center, visible)
	fmt.Println("Second:", pos[1] * 100 + pos[0])
}

func vaporizeAndGet200thPos(raw [][]rune, center [2]int, visible [][2]int) [2]int {
	data := copyMap(raw)
	i := 0
	var result [2]int
	for {
		if i + len(visible) >= 200 {
			coords := visible[199-i]
			result = [2]int{coords[0], coords[1]}
		} else if i < 200 {
			i += len(visible)
		}
		data = vaporizeWave(data, center, visible)
		visible = getSortedVisible(data, center)
		if len(visible) == 0 {
			fmt.Println("ALL CLEAR!")
			printMap(data, center, visible)
			return result 
		}
	}
}

func copyMap(raw [][]rune) [][]rune {
	data := make([][]rune, len(raw))
	for i := range data {
		row := make([]rune, len(raw[i]))
		copy(row, raw[i])
		data[i] = append(data[i], row...)
	}
	return data
}

func vaporizeWave(raw [][]rune, center [2]int, visible [][2]int) [][]rune {
	fmt.Println("VAPORIZING", len(visible), "VISIBLE ASTEROIDS FROM POSITION:", center)
	data := copyMap(raw)
	for _, coords := range visible {
		data[coords[0]][coords[1]] = '.'
	}
	return data
}

func findCenter(data [][]rune) [2]int {
	seen := 0
	center := [2]int{0, 0}

	for y := range data {
		for x := range data[y] {
			if data[y][x] != '#' {
				continue
			}
			temp := len(getVisible(data, [2]int{y, x}))

			if temp > seen {
				center[0] = y
				center[1] = x
				seen = temp
			} 
		} 
	}
	return center
}

func getSortedVisible(data [][]rune, center [2]int) [][2]int {
	visible := getVisible(data, center)
	keys := []float64{}
	for key := range visible {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	result := [][2]int{}
	for _, key := range keys {
		result = append(result, visible[key])
	}
	return result 
}

func printMap(data [][]rune, center [2]int, visible[][2]int) {
	fmt.Println("MAP OF POSITION", center, "SEES:", len(visible), "ASTEROIDS")
	for y, line := range data {
		for x, val := range line {
			if center[0] == y && center[1] == x {
				fmt.Printf("\033[31m%c\033[0m", val)
				continue
			}
			coord := [2]int{y, x}
			if slices.Contains(visible, coord) {
				fmt.Printf("\033[32m%c\033[0m", val)
				continue
			}
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}

func getVisible(asteroids [][]rune, pos [2]int) map[float64][2]int {
	seen := map[float64][2]int{}

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
			if coords, ok := seen[angle]; ok { 
				dist := math.Abs(float64(pos[0] - coords[0])) + math.Abs(float64(pos[1] - coords[1]))
				if dist < math.Abs(fx) + math.Abs(fy) {
					continue
				}
			}
			seen[angle] = [2]int{y, x}
		}
	} 

	return seen 
}


func calculateAngle(y, x float64) float64 {
	if x == 0 && y == 0 {
		panic("there should not be two 0 values here :(")
	}
	if x == 0 {
		if y > 0 {
			return 180 
		}
		return 0 
	}
	if y == 0 {
		if x > 0 {
			return 90
		}
		return 270 
	}
	degrees := math.Atan(math.Abs(y)/math.Abs(x)) * 180 / math.Pi
	if x > 0 {
		if y > 0 {
			return 90 +degrees
		}
		return degrees
	}

	if y > 0 {
		return 180 + degrees
	}
	return 270 + degrees
}
