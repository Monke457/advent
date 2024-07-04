package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
)

type game struct {
	state [][][]bool
	floor int
}


var neighbours3D [][]int
var neighbours4D [][]int

func main() {
	data := reader.FileTo2DArray("data/2020/day9.txt")

	neighbours3D = initializeNeighbours(3, [][]int{{}})
	neighbours4D = initializeNeighbours(4, [][]int{{}})

	fmt.Println(len(neighbours3D))
	for _, n := range neighbours4D {
		fmt.Println(n)
	}
	
	state := [][][]bool{{}}
	for i, line := range data {
		state[0] = append(state[0], []bool{})
		for _, b := range line {
			state[0][i] = append(state[0][i], b == '#')
		}
	}

	game := game{ state: state, floor: 0 }

	for range 6 {
		game.tick()
	}

	fmt.Println("First:", game.active())
}

func (g *game) tick() { 
	livingCells := [][3]int{}
	maxes := [3]int{0, 0, 0}
	mins := [3]int{math.MaxInt8, math.MaxInt8, math.MaxInt8}

	for x := -1; x <= len(g.state); x++ {
		for y := -1; y <= len(g.state[0]); y++ {
			for z := -1; z <= len(g.state[0][0]); z++ {
				if g.newCell(x, y, z) {
					x1, y1, z1 := x+1, y+1, z+1
					livingCells = append(livingCells, [3]int{x1, y1, z1})
					if x1 >= maxes[0] {
						maxes[0] = x1+1
					}
					if y1 >= maxes[1] {
						maxes[1] = y1+1
					}
					if z1 >= maxes[2] {
						maxes[2] = z1+1
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
				}
			}
		}
	}
	g.state = newState(livingCells, maxes, mins) 
	g.floor += mins[0]-1
}

func newState(living [][3]int, maxes, mins [3]int) [][][]bool{
	state := make([][][]bool, maxes[0] - mins[0])

	for x := 0; x < maxes[0] - mins[0]; x++ {
		state[x] = make([][]bool, maxes[1] - mins[1])
		for y := 0; y < maxes[1] - mins[1]; y++ {
			state[x][y] = make([]bool, maxes[2] - mins[2])
		}
	}

	for _, l := range living {
		state[l[0] - mins[0]][l[1] - mins[1]][l[2] - mins[2]] = true
	}
	return state
}

func (g *game) newCell(x, y, z int) bool {
	c := g.countN3D(x, y, z)
	if g.alive(x, y, z) {
		if c == 2 || c == 3 {
			return true		}
	} else {
		if c == 3 {
			return true
		}
	}
	return false
}

func (g *game) alive(x, y, z int) bool {
	if g.oob(x, y, z) {
		return false
	}
	return g.state[x][y][z] 
}

func (g *game) countN3D(x, y, z int) int {
	count := 0
	for _, n := range neighbours3D {
		if 0 == n[0] && 0 == n[1] && 0 == n[2] {
			continue
		}
		x1, y1, z1 := x+n[0], y+n[1], z+n[2]
		if g.oob(x1, y1, z1) {
			continue
		}
		if g.state[x1][y1][z1] {
			count++
		}
	}
	return count 
}

func (g *game) oob(x, y, z int) bool {
	if x < 0 || y < 0 || z < 0 {
		return true
	}
	if x >= len(g.state) || y >= len(g.state[x]) || z >= len(g.state[x][y]) {
		return true
	}
	return false
}

func (g *game) active() int {
	count := 0
	for _, x := range g.state {
		for _, y := range x {
			for _, cell := range y {
				if cell {
					count++
				}
			}
		}
	}
	return count
}

func (g *game) printState() {
	i := g.floor
	for _, layer := range g.state {
		fmt.Println("z=", i)
		for _, line := range layer {
			for _, b := range line {
				if b {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
		i++
	}
	fmt.Println()
}

func initializeNeighbours(d int, arr [][]int) [][]int {
	if d == 0 {
		return arr
	} 
	result := [][]int{}
	for i := -1; i < 2; i++ {
		old := make([][]int, len(arr))
		copy(old, arr)
		for r := range old {
			old[r] = append(old[r], i) 
		}
		result = append(result, old[:]...) 
	}
	arr = append([][]int{}, result[:]...)
	return initializeNeighbours(d-1, arr)
}
