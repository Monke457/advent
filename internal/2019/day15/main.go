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
	oxy = 'O'
)

var area = map[[2]int]tile{}

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day15.txt")

	d := [2]int{0, 0}
	area[d] = droid
	comp := ic.NewComputer(data, 0, int(north), false) 

	steps := run(d, comp)

	drawArea()
	fmt.Println("Steps to find system:", steps)

	mins := 0
	for {
		if flooded() {
			break
		}
		floodArea()
		mins++
	}

	drawArea()
	fmt.Println("Mins to flood (incorrect)", mins) 
}

func run(d [2]int, comp ic.Computer) int {
	steps := []cmd{}
	retracing := false
	var foundIn int 
	for {
		out := make(chan int)
		done := make(chan bool)
		go comp.Run(out, done)
		select {
		case output := <-out:
			switch output {
			case 0:
				markTile(d, cmd(comp.Input), wall)
				steps, comp.Input, retracing = getNextStep(d, steps)
			case 1:
				if area[d] != oxy {
					markTile(d, 0, floor)
				}
				d = markTile(d, cmd(comp.Input), droid)
				if retracing {
					steps, comp.Input, retracing = getNextStep(d, steps)
				} else {
					steps = append(steps, cmd(comp.Input))
				}
			case 2:
				markTile(d, 0, oxy)
				steps = append(steps, cmd(comp.Input))
				foundIn = len(steps) 
			}
		case <-done:
			fmt.Println("Computer shutdown")
			return foundIn 
		}
	}
}

func flooded() bool {
	for _, val := range area {
		if val == floor || val == droid {
			return false
		}
	}
	return true
}

func getNeighbours(pos [2]int) [][2]int {
	result := [][2]int{}
	n := [2]int{pos[0]-1, pos[1]}
	s := [2]int{pos[0]+1, pos[1]}
	e := [2]int{pos[0], pos[1]+1}
	w := [2]int{pos[0], pos[1]-1}
	if area[n] == floor || area[n] == droid {
		result = append(result, n)
	}
	if area[s] == floor || area[s] == droid {
		result = append(result, s)
	}
	if area[e] == floor || area[e] == droid {
		result = append(result, e)
	}
	if area[w] == floor || area[w] == droid {
		result = append(result, w)
	}
	return result
}

func floodArea() {
	marked := [][2]int{}
	for key, val := range area {
		if val == oxy {
			marked = append(marked, getNeighbours(key)...)
		}
	}
	for _, m := range marked {
		area[m] = oxy
	}
}

func getNextStep(pos [2]int, steps []cmd) ([]cmd, int, bool){
	dir, found := findUnexplored(pos)
	if found {
		return steps, dir, false
	}
	if len(steps) == 0 {
		return steps, 0, false
	}
	switch steps[len(steps)-1] {
	case north:
		return steps[:len(steps)-1], int(south), true 
	case south:
		return steps[:len(steps)-1], int(north), true 
	case west:
		return steps[:len(steps)-1], int(east), true 
	}
	return steps[:len(steps)-1], int(west), true 
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
