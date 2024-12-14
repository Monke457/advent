package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2018/day6.txt")

	coords := map[[2]int]byte{}
	for i, line := range data {
		coords[parseCoords(line)] = byte(i+65)
	}
	
	remotestAreas := getRemoteAreas(coords)
	largest := getLargestArea(remotestAreas)

	popAreaSize := measurePopulousArea(coords, 10000)

	fmt.Println("Largest finite area:", largest)
	fmt.Println("Size of well populated area:", popAreaSize)
}

func measurePopulousArea(coords map[[2]int]byte, maxDist int) int {
	w, h := getSize(coords)
	areas := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			score := getDistScore(coords, x, y)
			if score < maxDist {
				areas++
			}
		}
	}
	return areas
}

func getDistScore(coords map[[2]int]byte, x, y int) int {
	score := 0.0
	for k := range coords {
		score += math.Abs(float64(k[0] - x)) + math.Abs(float64(k[1] - y))
	}
	return int(score)
}

func getLargestArea(areas map[byte]int) int {
	largest := 0
	for _, v := range areas {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func getRemoteAreas(coords map[[2]int]byte) map[byte]int {
	w, h := getSize(coords)
	areas := map[byte]int{}
	infinite := []byte{}
	for y := 0; y < h; y++ {
		loop:
		for x := 0; x < w; x++ {
			b := getClosestPoint(coords, x, y)
			if b == nil {
				continue
			}
			for _, inf := range infinite {
				if inf == *b {
					continue loop
				}
			}
			if y == 0 || x == 0 || y == h-1 || x == w-1 {
				infinite = append(infinite, *b)
				delete(areas, *b)
			} else {
				areas[*b]++
			}
		}
	}
	return areas
}

func getClosestPoint(coords map[[2]int]byte, x, y int) *byte {
	var closest *byte
	dist := math.MaxFloat64
	for k, v := range coords {
		newDist := math.Abs(float64(k[0] - x)) + math.Abs(float64(k[1] - y))
		if newDist < dist {
			closest = &v
			dist = newDist
		} else if newDist == dist {
			closest = nil
		}
	}
	return closest
}

func parseCoords(str string) [2]int {
	parts := strings.Split(str, ", ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return [2]int{x, y}
}

func getSize(coords map[[2]int]byte) (int, int) {
	w, h := 0, 0
	for k := range coords {
		if k[0] > w {
			w = k[0]
		}
		if k[1] > h {
			h = k[1]
		}
	}
	return w+1, h+1
}
