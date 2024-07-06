package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

type game struct {
	state [][][][]bool
	floor int
	instance int
}

var neighbours = initializeNeighbours()

func main() {
	data := reader.FileTo2DArray("data/2020/day17.txt")

	for _, n := range neighbours {
		fmt.Println(n)
	}

	state := [][][][]bool{{{}}}
	for i, line := range data {
		state[0][0] = append(state[0][0], []bool{})
		for _, b := range line {
			state[0][0][i] = append(state[0][0][i], b == '#')
		}
	}

	game := game{ state: state }

	for range 6 {
		game.tick()
	}

	fmt.Println("Second:", game.active())
}

func (g *game) tick() { 
	livingCells := [][4]int{}
	maxes := [4]int{0, 0, 0, 0}
	mins := [4]int{math.MaxInt8, math.MaxInt8, math.MaxInt8, math.MaxInt8}

	for x := -1; x <= len(g.state); x++ {
		for y := -1; y <= len(g.state[0]); y++ {
			for z := -1; z <= len(g.state[0][0]); z++ {
				for w := -1; w <= len(g.state[0][0][0]); w++ {
					if g.newCell(x, y, z, w) {
						x1, y1, z1, w1 := x+1, y+1, z+1, w+1
						livingCells = append(livingCells, [4]int{x1, y1, z1, w1})
						if x1 >= maxes[0] {
							maxes[0] = x1+1
						}
						if y1 >= maxes[1] {
							maxes[1] = y1+1
						}
						if z1 >= maxes[2] {
							maxes[2] = z1+1
						}
						if w1 >= maxes[3] {
							maxes[3] = w1+1
						}
						if x1 < mins[0] {
							mins[0] = x1
						}
						if y1 < mins[1] {
							mins[1] = y1
						}
						if z1 < mins[2] {
							mins[2] = z1
						}
						if w1 < mins[3] {
							mins[3] = w1
						}
					}
				}
			}
		}
	}
	g.state = newState(livingCells, maxes, mins) 
	g.instance += mins[0]-1
	g.floor += mins[1]-1
}

func newState(living [][4]int, maxes, mins [4]int) [][][][]bool{
	state := make([][][][]bool, maxes[0] - mins[0])

	for x := 0; x < maxes[0] - mins[0]; x++ {
		state[x] = make([][][]bool, maxes[1] - mins[1])
		for y := 0; y < maxes[1] - mins[1]; y++ {
			state[x][y] = make([][]bool, maxes[2] - mins[2])
			for z := 0; z < maxes[2] - mins[2]; z++ {
				state[x][y][z] = make([]bool, maxes[3] - mins[3])
			}
		}
	}

	for _, l := range living {
		state[l[0] - mins[0]][l[1] - mins[1]][l[2] - mins[2]][l[3] - mins[3]] = true
	}
	return state
}

func (g *game) newCell(x, y, z, w int) bool {
	c := g.countN(x, y, z, w)
	if g.alive(x, y, z, w) {
		if c == 2 || c == 3 {
			return true	
		}
	} else {
		if c == 3 {
			return true
		}
	}
	return false
}

func (g *game) alive(x, y, z, w int) bool {
	if g.oob(x, y, z, w) {
		return false
	}
	return g.state[x][y][z][w] 
}

func (g *game) countN(x, y, z, w int) int {
	count := 0
	for _, n := range neighbours {
		if 0 == n[0] && 0 == n[1] && 0 == n[2] && 0 == n[3] {
			continue
		}
		x1, y1, z1, w1 := x+n[0], y+n[1], z+n[2], w+n[3]
		if g.oob(x1, y1, z1, w1) {
			continue
		}
		if g.state[x1][y1][z1][w1] {
			count++
		}
	}
	return count 
}

func (g *game) oob(x, y, z, w int) bool {
	if x < 0 || y < 0 || z < 0 || w < 0 {
		return true
	}
	if x >= len(g.state) || y >= len(g.state[x]) || z >= len(g.state[x][y]) || w >= len(g.state[x][y][z]) {
		return true
	}
	return false
}

func (g *game) active() int {
	count := 0
	for _, x := range g.state {
		for _, y := range x {
			for _, z := range y {
				for _, cell := range z {
					if cell {
						count++
					}
				}
			}
		}
	}
	return count
}

func (g *game) printState() {
	i := g.instance
	j := g.floor
	for _, x := range g.state {
		for _, y := range x {
			fmt.Printf("z=%d, w=%d\n", j, i)
			for _, z := range y {
				for _, b := range z {
					if b {
						fmt.Printf("#")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Printf("\n")
			}
			j++
		}
		j = g.floor
		i++
	}
	fmt.Println()
}
