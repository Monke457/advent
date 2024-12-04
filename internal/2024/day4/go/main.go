package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileTo2DArray("data/2024/day4.txt")

	count := 0
	for y := range data {
		for x := range data[y] {
			count += xmasCount(data, y, x)
		}
	}
	fmt.Println("XMAS count:", count)

	count = 0
	for y := 0; y < len(data)-2; y++ {
		for x := 0; x < len(data[y])-2; x++ {
			temp := make([][]rune, 3)
			for i := 0; i < 3; i++ {
				temp[i] = data[y+i][x:x+3]
			}
			if hasXMas(temp) {
				count++
			}
		}
	}
	fmt.Println("X-MAS count:", count)
}

const (
	NORTH = "north"
	NORTH_EAST = "north east"
	EAST = "east"
	SOUTH_EAST = "south east"
	SOUTH = "south"
	SOUTH_WEST = "south west"
	WEST = "west"
	NORTH_WEST = "north west"
)

var cardinalities = map[string][2]int{
	NORTH: [2]int{-1, 0},
	NORTH_EAST: [2]int{-1, 1},
	EAST: [2]int{0, 1},
	SOUTH_EAST: [2]int{1, 1},
	SOUTH: [2]int{1, 0},
	SOUTH_WEST: [2]int{1, -1},
	WEST: [2]int{0, -1},
	NORTH_WEST: [2]int{-1, -1},
}

func xmasCount(data [][]rune, y, x int) int {
	result := 0
	target := []rune("XMAS")
	loop:
	for _, card := range cardinalities {
		for i := 0; i < len(target); i++ {
			posY := y + card[0] * i 
			posX := x + card[1] * i
			if posY < 0 || posY >= len(data) {
				continue loop
			}
			if posX < 0 || posX >= len(data[posY]) {
				continue loop
			}
			if data[posY][posX] != target[i] {
				continue loop
			}
		}
		result++
	}
	return result
}

func hasXMas(data [][]rune) bool {
	if data[1][1] != 'A' {
		return false 
	}

	target := "MS"
	tr := string(data[0][0]) + string(data[2][2])
	rt := string(data[2][2]) + string(data[0][0])
	if target != tr && target != rt {
		return false
	}

	br := string(data[2][0]) + string(data[0][2])
	rb := string(data[0][2]) + string(data[2][0])
	return target == br || target == rb
}
