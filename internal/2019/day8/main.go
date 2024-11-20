package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

func main() {
	data := reader.FileToIntArrayByRune("data/2019/day8.txt")

	width := 25
	height := 6

	layers := parseData(data, width, height)
	leastZeros := getLayerWithLeastZeros(layers)
	countOnes := countDigits(layers[leastZeros], 1)
	countTwos := countDigits(layers[leastZeros], 2)

	fmt.Println("First:", countOnes * countTwos)

	decoded := decode(layers, width, height)

	for _, row := range decoded {
		for _, val := range row {
			if val == 1 {
				fmt.Print(val)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func makeFilledArray(w, h, val int) [][]int {
	result := make([][]int, h)
	for i := range result {
		result[i] = make([]int, w)
		for j := range result[i] {
			result[i][j] = val
		}
	}
	return result
}

func decode(layers [][][]int, w, h int) [][]int {
	result := makeFilledArray(w, h, 2)

	for _, layer := range layers {
		for i, row := range layer {
			for j, val := range row {
				if result[i][j] != 2 {
					continue
				}
				result[i][j] = val
			}
		}
	}

	return result
}

func parseData(data []int, w, h int) [][][]int {
	result := [][][]int{}
	l := 0
	lsize := w*h
	for { 
		if len(data) < (1+l)*lsize {
			break	
		}
		result = append(result, [][]int{})
		for i := 0; i < h; i++ {
			row := make([]int, w)
			start := l * lsize + i * w
			copy(row, data[start:start+w])
			result[l] = append(result[l], row)
		}
		l++
	}
	return result
}

func getLayerWithLeastZeros(layers [][][]int) int {
	result := 0
	count := math.MaxInt
	loop:
	for i, layer := range layers {
		temp := 0
		for _, row := range layer {
			for _, val := range row {
				if val == 0 {
					temp++
				}
				if temp > count {
					continue loop
				}
			}
		}
		if temp < count {
			count = temp
			result = i 
		}
	}
	return result
}

func countDigits(layer [][]int, digit int) int {
	count := 0
	for _, row := range layer {
		for _, val := range row {
			if val == digit {
				count++
			}
		}
	}
	return count
}
