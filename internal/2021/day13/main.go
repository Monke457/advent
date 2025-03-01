package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"

	"github.com/Monke457/printge"
)

func main() {
	data := reader.FileToArray("data/2021/day13.txt")

	dots := [][2]int{}
	h, w := 0, 0
	var folds [][2]int

	for i := range data {
		if data[i] == "" {
			folds = makeFolds(data[i+1:])
			break
		}
		xStr, yStr, _ := strings.Cut(data[i], ",")
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)

		dots = append(dots, [2]int{y, x})
		if y > h { h = y }
		if x > w { w = x }
	}

	paper := makePaper(dots, h+1, w+1)

	for i := 0; i < len(folds); i++ {
		paper = foldPaper(paper, folds[i])
		if i == 0 {
			count := countDots(paper)
			fmt.Println("Number of dots after 1 fold:", count)
		}
	}

	for _, row := range paper {
		for _, cell := range row {
			if cell == '#' {
				s := string(cell)
				printge.Print(s+s, printge.RED)
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
}

func countDots(paper [][]rune) int {
	count := 0
	for _, row := range paper {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}

func foldPaper(paper [][]rune, fold [2]int) [][]rune {
	folded, temp := [][]rune{}, [][]rune{}
	if fold[0] == 0 {
		for i := range paper {
			folded = append(folded, []rune{})
			for j := 0; j < fold[1] && j < len(paper[i]); j++ {
				folded[i] = append(folded[i], paper[i][j])
			}
			temp = append(temp, []rune{})
			for j := fold[1]+1; j < len(paper[i]); j++ {
				temp[i] = append([]rune{paper[i][j]}, temp[i]...)
			}
		}
	} else {
		for i := 0; i < fold[0] && i < len(paper); i++ {
			folded = append(folded, paper[i]) 
		}
		for i := fold[0]+1; i < len(paper); i++ {
			temp = append([][]rune{paper[i]}, temp...) 
		}
	}

	for i := len(temp)-1; i >= 0; i-- {
		for j := len(temp[i])-1; j >= 0; j-- {
			if temp[i][j] == '#' {
				folded[i][j] = '#'
			}
		}
	}

	return folded
}

func makeFolds(data []string) [][2]int {
	folds := [][2]int{}
	for _, line := range data {
		info, valStr, _ := strings.Cut(line, "=")
		val, _ := strconv.Atoi(valStr)
		fold := [2]int{0,0}
		if info[len(info)-1] == 'y' {
			fold[0] = val
		} else {
			fold[1] = val
		}
		folds = append(folds, fold)
	}
	return folds
}

func makePaper(dots [][2]int, h, w int) [][]rune {
	paper := make([][]rune, h)
	for i := range paper {
		val := strings.Repeat(".", w)
		paper[i] = []rune(val)
	}
	for _, dot := range dots {
		paper[dot[0]][dot[1]] = '#'
	}
	return paper
}
