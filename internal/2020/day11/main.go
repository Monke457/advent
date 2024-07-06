package main

import (
	"advent/internal/pkg/grid"
	"advent/internal/pkg/reader"
	"fmt"
)

type plan struct {
	seats [][]rune
}

var neighbours = grid.Neighbours2D()

func main() {
	first()
	second()
}

func first() {
	plan := plan{ seats: reader.FileTo2DArray("data/2020/day11.txt") }

	var done bool
	for !done {
		done = plan.round1();
	}
	fmt.Println("First:", plan.occupied())
}

func second() {
	plan := plan{ seats: reader.FileTo2DArray("data/2020/day11.txt") }
	var done bool
	for !done {
		done = plan.round2();
	}
	fmt.Println("Second:", plan.occupied())
}

func (p *plan) round1() bool {
	changes := map[[2]int]rune{}
	for i, row := range p.seats {
		for j, seat := range row {
			if seat == '.' {
				continue
			}
			n := p.countNeighbours(i, j)
			if seat == 'L' {
				if n == 0 {
					changes[[2]int{i, j}] = '#'
				}
			}
			if seat == '#' {
				if n >= 4 {
					changes[[2]int{i, j}] = 'L'
				}
			}
		}
	}

	for k, v := range changes {
		p.seats[k[0]][k[1]] = v 
	}
	return len(changes) == 0 
}

func (p *plan) round2() bool {
	changes := map[[2]int]rune{}
	for i, row := range p.seats {
		for j, seat := range row {
			if seat == '.' {
				continue
			}
			n := p.countNeighbours2(i, j)
			if seat == 'L' {
				if n == 0 {
					changes[[2]int{i, j}] = '#'
				}
			}
			if seat == '#' {
				if n >= 5 {
					changes[[2]int{i, j}] = 'L'
				}
			}
		}
	}

	for k, v := range changes {
		p.seats[k[0]][k[1]] = v 
	}
	return len(changes) == 0 
}

func (p *plan) countNeighbours(i, j int) int {
	count := 0
	for _, n := range neighbours {
		i1, j1 := i+n[0], j+n[1]
		if p.oob(i1, j1) {
			continue
		}
		if p.seats[i1][j1] == '#' {
			count++
		}
	}
	return count
}

func (p *plan) countNeighbours2(i, j int) int {
	count := 0
	for _, n := range neighbours {
		i1, j1 := i, j
		for {
			i1 += n[0]
			j1 += n[1]  
			if p.oob(i1, j1) {
				break	
			}
			if p.seats[i1][j1] == '.' {
				continue
			}
			if p.seats[i1][j1] == '#' {
				count++
			}
			break
		}
	}
	return count
}

func (p *plan) oob(i, j int) bool {
	if i < 0 || j < 0 {
		return true
	}
	return i >= len(p.seats) || j >= len(p.seats[i])
}

func (p *plan) occupied() int {
	count := 0
	for _, row := range p.seats {
		for _, seat := range row {
			if seat == '#' {
				count++
			}
		}
	}
	return count
}
