package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
)

type guard struct {
	start [2]int
	pos [2]int
	dir rune
}

func main() {
	data := reader.FileTo2DArray("data/2024/day6.txt")

	route, _ := mapRoute(data, nil)
	positions := getPositions(route)
	fmt.Println("positions the guard could be in:", len(positions))

	ch := make(chan bool)
	for _, pos := range positions {
		go func(out chan bool) {
			_, looping := mapRoute(data, &pos)
			out<-looping
		}(ch)
	}

	obstacles := 0
	for range len(positions) {
		select {
		case found := <-ch:
			if found {
				obstacles++
			}
		}
	}

	fmt.Println("positions for the obstacle to be in:", obstacles)
}

func copyData(raw [][]rune) [][]rune {
	data := make([][]rune, len(raw))
	for i, row := range raw {
		rowcp := make([]rune, len(row))
		copy(rowcp, row)
		data[i] = rowcp 
	}
	return data
}

func mapRoute(raw [][]rune, obstacle* [2]int) ([][]rune, bool) {
	data := copyData(raw)

	g, err := getGuard(data)
	if err != nil {
		drawMap(data, [2]int{-1, -1})
		panic(err)
	}

	if obstacle != nil && *obstacle != g.start {
		data[obstacle[0]][obstacle[1]] = '0'	
	}

	cache := map[[2]int][]rune{}
	for {
		pos, dir, exits := g.getNewPosition(data)
		data = g.markMap(data)
		if exits {
			break
		}
		if slices.Contains(cache[pos], dir) {
			return data, true
		}
		cache[pos] = append(cache[pos], dir)
		g.pos = pos
		g.dir = dir
	}

	return data, false 
}

func (g guard) markMap(data [][]rune) [][]rune {
	var r rune
	y, x := g.pos[0], g.pos[1]
	switch g.dir {
	case '^', 'v':
		if data[y][x] == '-' {
			r = '+'
		} else {
			r = '|'
		}
	case '>', '<':
		if data[y][x] == '|' {
			r = '+'
		} else {
			r = '-'
		}
	}
	data[y][x] = r
	return data
}

func getPositions(data [][]rune) [][2]int {
	positions := [][2]int{}
	for i, row := range data {
		for j, cell := range row {
			if cell != '.' && cell != '#' && cell != '0' {
				positions = append(positions, [2]int{i, j})
			}
		}
	}
	return positions
}

func (g guard) getNewPosition(data [][]rune) ([2]int, rune, bool) {
	var y, x int
	dir := g.dir

	for range 4 {
		switch dir {
		case '^':
			y = g.pos[0]-1
			x = g.pos[1]
		case '>':
			y = g.pos[0]
			x = g.pos[1]+1
		case 'v':
			y = g.pos[0]+1
			x = g.pos[1]
		case '<':
			y = g.pos[0]
			x = g.pos[1]-1
		}
		if exitsMap(data, y, x) {
			return [2]int{y, x}, dir, true
		}
		if validMovement(data, y, x) {
			return [2]int{y, x}, dir, false
		}
		dir = turn(dir)
	}

	drawMap(data, g.start)
	panic(fmt.Errorf("Error: could not find a valid movement for the guard\n"))
}

func exitsMap(data [][]rune, y, x int) bool {
	if y < 0 || y >= len(data) {
		return true 
	}
	if x < 0 || x >= len(data[y]) {
		return true 
	}
	return false
}
func validMovement(data [][]rune, y, x int) bool {
	if data[y][x] == '#' {
		return false
	}
	if data[y][x] == '0' {
		return false
	}
	return true
}

func turn(dir rune) rune {
	switch dir {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}
	panic(fmt.Errorf("Error: not a valid direction  dir %c\n", dir))
}

func getGuard(data [][]rune) (guard, error) {
	g := guard{}
	for i, row := range data {
		for j, cell := range row {
			if cell == '^' || cell == '>' || cell == 'v' || cell == '<' {
				g.start = [2]int{i, j}
				g.pos = [2]int{i, j}
				g.dir = cell
				return g, nil
			}
		}
	}
	return g, fmt.Errorf("Error: could not find a guard on the map\n", )
}

func drawMap(data [][]rune, start [2]int) {
	for i := range data {
		for j := range data[i] {
			if start == [2]int{i, j} {
				fmt.Printf("^")
				continue
			}
			fmt.Printf("%c", data[i][j])
		}
		fmt.Printf("\n")
	}
}
