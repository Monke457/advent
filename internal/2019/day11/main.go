package main

import (
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

type direction rune 
const (
	left direction = '<'
	up = '^'
	right = '>'
	down ='v'
)

type color rune 
const (
	white color = '#' 
	black = '.'	
)

type robot struct {
	dir direction
	pos [2]int
}

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day11.txt")

	start := [2]int{0, 0}
	robot := robot{dir: up, pos: start}
	grid := map[[2]int]color{start:white}

	run(data, robot, grid)
}

func run(data []int, robot robot, grid map[[2]int]color) {
	comp := ic.NewComputer(data, 0, 0, false)
	paintPhase := true 
	painted := map[[2]int]bool{}
	loop:
	for {
		if grid[robot.pos] == white {
			comp.Input = 1
		} else {
			comp.Input = 0 
		}
		out := make(chan int)
		done := make(chan bool)
		go comp.Run(out, done)
		select {
		case output := <-out:
			if paintPhase {
				grid = robot.paint(grid, output)
				painted[robot.pos] = true
			} else {
				robot.turn(output)
				robot.move()
			}
			paintPhase = !paintPhase
		case <-done:
			break loop	
		}
	}
	drawGrid(grid, robot)
	fmt.Println("Painted", len(painted), "panels")
}

func (r robot) paint(grid map[[2]int]color, i int) map[[2]int]color {
	if i == 0 {
		grid[r.pos] = black 
	}
	if i == 1 {
		grid[r.pos] = white 
	}
	return grid
}

func (r *robot) turn(i int) {
	if i == 0 {
		switch r.dir {
		case left:
			r.dir = down
		case up:
			r.dir = left 
		case right:
			r.dir = up 
		case down:
			r.dir = right 
		}
	} else if i == 1 {
		switch r.dir {
		case left:
			r.dir = up 
		case up:
			r.dir = right 
		case right:
			r.dir = down 
		case down:
			r.dir = left 
		}
	}
}

func (r *robot) move() {
	switch r.dir {
	case left:
		r.pos[0] -= 1
	case up:
		r.pos[1] -= 1
	case right:
		r.pos[0] += 1
	case down:
		r.pos[1] += 1
	}
}

func drawGrid(grid map[[2]int]color, robot robot) {
	yAxis := getYAxis(grid, robot)
	xAxis := getXAxis(grid, robot)
	for y := yAxis[0]; y <= yAxis[1]; y++ {
		for x := xAxis[0]; x <= xAxis[1]; x++ {
			cell := [2]int{x, y}
			if robot.pos == cell {
				fmt.Printf("%c", robot.dir)
				continue
			}
			if color, ok := grid[cell]; ok {
				fmt.Printf("%c", color)
			} else {
				fmt.Printf("%c", black)
			}
		}
		fmt.Println()
	} 
}

func getXAxis(grid map[[2]int]color, robot robot) [2]int {
	lowest := robot.pos[0] 
	highest := robot.pos[0] 
	for key := range grid {
		if key[0] < lowest {
			lowest = key[0]
		}
		if key[0] > highest {
			highest = key[0]
		}
	}
	return [2]int{lowest, highest}
}

func getYAxis(grid map[[2]int]color, robot robot) [2]int {
	lowest := robot.pos[1] 
	highest := robot.pos[1] 
	for key := range grid {
		if key[1] < lowest {
			lowest = key[1]
		}
		if key[1] > highest {
			highest = key[1]
		}
	}
	return [2]int{lowest, highest}
}
