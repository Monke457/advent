package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"time"
)

func main() {
	lines := reader.FileTo2DArray("data/day14.txt")
	//fmt.Println("first:", solveFirstProblem(lines))
	fmt.Println("second:", solveSecondProblem(lines))
}

func solveFirstProblem(lines [][]rune) int {
	north := tiltNorth(lines, 0)
	load := countLoad(north)
	return load
}

func solveSecondProblem(rocks [][]rune) int {
	arr := []int{}

	c := make(chan int)
	go func(c chan int) {
		start := time.Now()
		for range 100 {
			rocks = cycle(rocks)
			c <- countLoad(rocks)
		}
		fmt.Println("100 cycles", time.Now().Sub(start))
		return
	}(c)

	c2 := make(chan [2]int)
	for {
		select {
		case load := <-c:
			arr = append(arr, load)
			go findPattern(arr, 3, c2)
		case p := <-c2:
			fmt.Printf("pattern found: %d\n", arr[p[0]:p[0]+p[1]])
			return p[1]
		}
	}
}

func findPattern(loads []int, l int, c chan [2]int) {
	if l+l+1 >= len(loads) {
		return
	}

	for i := 0; i+l+l < len(loads); i++ {
		if slices.Equal(loads[i:i+l], loads[i+l:i+l+l]) {

			fmt.Printf("pattern found %d, %d\n", loads[i:i+l], loads[i+l:i+l+l])
			c <- [2]int{i, l}
			return
		}
	}
	findPattern(loads, l+1, c)
}

func printRocks(rocks [][]rune) {
	fmt.Println()
	for _, r := range rocks {
		fmt.Printf("%c\n", r)
	}
}

func tiltNorth(rocks [][]rune, pos int) [][]rune {
	if pos >= len(rocks)-1 {
		return rocks
	}
	for i := pos; i < len(rocks)-1; i++ {
		for j := 0; j < len(rocks[i]); j++ {
			if rocks[i][j] == '.' && rocks[i+1][j] == 'O' {
				rocks[i][j] = 'O'
				rocks[i+1][j] = '.'
				if pos > 0 {
					tiltNorth(rocks, pos-1)
				}
			}
		}
	}
	return tiltNorth(rocks, pos+1)
}

func cycle(rocks [][]rune) [][]rune {
	res := [][]rune{}
	for i := len(rocks[0]) - 1; i >= 0; i-- {
		res = append(res, cycleCol(rocks, i, true))
	}
	rocks = res
	res = [][]rune{}
	for i := len(rocks[0]) - 1; i >= 0; i-- {
		res = append(res, cycleCol(rocks, i, false))
	}
	rocks = res
	res = [][]rune{}
	for i := len(rocks[0]) - 1; i >= 0; i-- {
		res = append(res, cycleCol(rocks, i, true))
	}
	rocks = res
	res = [][]rune{}
	for i := len(rocks[0]) - 1; i >= 0; i-- {
		res = append(res, cycleCol(rocks, i, false))
	}
	return res
}

func cycleCol(rocks [][]rune, col int, up bool) []rune {
	r, e := 0, 0

	res := []rune{}
	for i := 0; i < len(rocks); i++ {
		if rocks[i][col] == '.' {
			e++
		} else if rocks[i][col] == 'O' {
			r++
		} else if rocks[i][col] == '#' {
			res = addRunes(r, e, true, res, up)
			r, e = 0, 0
		}
	}
	res = addRunes(r, e, false, res, up)
	return res
}

func addRunes(r, e int, h bool, res []rune, up bool) []rune {
	if up {
		for range r {
			res = append(res, 'O')
		}
		for range e {
			res = append(res, '.')
		}
	} else {
		for range e {
			res = append(res, '.')
		}
		for range r {
			res = append(res, 'O')
		}
	}
	if h {
		res = append(res, '#')
	}
	return res
}

func countLoad(rocks [][]rune) int {
	count := 0

	for i, row := range rocks {
		for _, r := range row {
			if r == 'O' {
				count += len(rocks) - i
			}
		}
	}

	return count
}
