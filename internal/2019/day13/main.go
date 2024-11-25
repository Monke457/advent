package main

import (
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
	"time"
)

type tileType int 
const (
	empty tileType = iota
	wall
	block
	paddle
	ball
)

type tile struct {
	x int
	y int
	id tileType
}

type model struct {
	comp ic.Computer
	tiles [24][45]tile
	score int 
	ch chan bool
}

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day13.txt")

	mod := initialModel(data)

	// hack paddle blocks along bottom
	for pos := 1585; pos <= 1627; pos++ {
		mod.comp.Data[pos] = 3; 
	}

	// I am so confused. Why does it behave like this? :(
	go mod.runGame()

	for {
		select {
		case <-time.After(50 * time.Millisecond):
			mod.drawGame()
		case <-mod.ch:
			fmt.Println("Game over. Final score:", mod.score)
			return
		}
	}
}

func initialModel(data []int) model {
	data[0] = 2
	comp := ic.NewComputer(data, 0, 0, false)
	tiles := [24][45]tile{}
	ch := make(chan bool)
	return model{comp: comp, tiles: tiles, ch: ch}
}

func (m *model) runGame() {
	idx := 0
	var current tile

	for {
		out := make(chan int)
		done := make(chan bool)
		go m.comp.Run(out, done)
		select {
		case output := <-out:
			switch idx {
			case 0:
				current = tile{x:output}
			case 1: 
				current.y = output
			case 2:
				if current.x == -1 && current.y == 0 {	
					m.score = output
					continue
				}
				current.id = tileType(output)
				m.tiles[current.y][current.x] = current
			}
			idx = (idx+1)%3
		case <-done:
			m.ch<-true
			return 
		}
	}
}

func (m *model) drawGame() {
	if len(m.tiles) == 0 {
		return
	}

	result := ""
	for i, row := range m.tiles {
		for j, cell := range row {
			switch cell.id {
			case empty:
				result += fmt.Sprintf(" ")
			case wall:
				if i == 0 {
					if j == 0 || j == len(row) -1 {
						result += fmt.Sprintf(" ")
						break
					}
					result += fmt.Sprintf("_")
					break
				}
				result += fmt.Sprintf("|")
			case block:
				result += fmt.Sprintf("=")
			case paddle:
				result += fmt.Sprintf("_")
			case ball:
				result += fmt.Sprintf(".")
			}
		}
		result += fmt.Sprintf("\n\r")
	}
	result += fmt.Sprintf("Score: %d\n\n\r", m.score)
	fmt.Print(result)
}
