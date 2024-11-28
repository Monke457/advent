package main

import (
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

type cmd int

const (
	current cmd = iota
	north
	south 
	west
	east
)

type tile rune 

const (
	wall tile = '#'
	droid = 'D'
	floor = '.'
	system= 'O'
)

var area = map[[2]int]tile{}

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day15.txt")

	d := [2]int{0, 0}
	area[d] = droid
	comp := ic.NewComputer(data, 0, int(north), false) 

	steps, found := run(d, comp)

	drawArea()
	if found {
		fmt.Println("Steps to find system:", steps)
	} else {
		fmt.Println("System not found")
	}
}

func run(d [2]int, comp ic.Computer) (int, bool) {
	steps := []cmd{}
	retracing := false
	for {
		out := make(chan int)
		done := make(chan bool)
		go comp.Run(out, done)
		select {
		case output := <-out:
			switch output {
			case 0:
				markTile(d, cmd(comp.Input), wall)
				dir, found := findUnexplored(d)
				if !found {
					retracing = true
					switch steps[len(steps)-1] {
					case north:
						comp.Input = int(south)
					case south:
						comp.Input = int(north)
					case west:
						comp.Input = int(east)
					case east:
						comp.Input = int(west)
					}
					steps = steps[:len(steps)-1]
				} else {
					retracing = false
					comp.Input = dir
				}
			case 1:
				markTile(d, 0, floor)
				d = markTile(d, cmd(comp.Input), droid)
				if retracing {
					dir, found := findUnexplored(d)
					if !found {
						switch steps[len(steps)-1] {
						case north:
							comp.Input = int(south)
						case south:
							comp.Input = int(north)
						case west:
							comp.Input = int(east)
						case east:
							comp.Input = int(west)
						}
						steps = steps[:len(steps)-1]
					} else {
						retracing = false
						comp.Input = dir
					}
				} else {
					steps = append(steps, cmd(comp.Input))
				}
			case 2:
				markTile(d, 0, system)
				return len(steps)+1, true
			}
		case <-done:
			fmt.Println("Computer shutdown")
			return len(steps)+1, false
		}
	}
}

func findUnexplored(d [2]int) (int, bool) {
	if _, ok := area[[2]int{d[0]-1, d[1]}]; !ok {
		return int(north), true
	}
	if _, ok := area[[2]int{d[0]+1, d[1]}]; !ok {
		return int(south), true
	}
	if _, ok := area[[2]int{d[0], d[1]+1}]; !ok {
		return int(east), true
	}
	if _, ok := area[[2]int{d[0], d[1]-1}]; !ok {
		return int(west), true
	}
	return 0, false
}

func markTile(d [2]int, i cmd, t tile) [2]int {
	coord := [2]int{d[0], d[1]}
	switch i {
	case north:
		coord[0]-=1
	case south:
		coord[0]+=1
	case west:
		coord[1]-=1
	case east:
		coord[1]+=1
	}
	area[coord] = t
	return coord
}

func drawArea() {
	y1, y2, x1, x2 := 0, 0, 0, 0
	for key := range area {
		if key[0] < y1 {
			y1 = key[0]
		}
		if key[0] > y2 {
			y2 = key[0]
		}
		if key[1] < x1 {
			x1 = key[1]
		}
		if key[1] > x2 {
			x2 = key[1]
		}
	}
	for i := y1; i <= y2; i++ {
		for j := x1; j <= x2; j++ {
			if t, ok := area[[2]int{i,j}]; ok {
				fmt.Printf("%c", t) 
			} else {
				fmt.Printf(" ") 
			}
		} 
		fmt.Println()
	}
}
