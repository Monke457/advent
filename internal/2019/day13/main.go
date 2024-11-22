package main

import (
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
	"os"
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

type result struct {
	score int
	done bool
}

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day13.txt")

	ch := make(chan byte)
	go readInput(ch)

	data[0] = 2
	comp := ic.NewComputer(data, 0, 0, false)
	res := make(chan result)
	go run(comp, ch, res)
	for {
		select {
		case r :=<-res:
			if r.done == true {
				fmt.Printf("Final Score: %s\n\r", r.score)
				return
			}
			fmt.Printf("New Score: %s\n\r", r.score)
		}
	}
}

func readInput(ch chan byte) {
	b := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		ch<-b[0]
		if b[0] == 113 {
			return
		}
	}
}

func run(comp ic.Computer, ch chan byte, res chan result) int {
	idx := 0
	tiles := []tile{}
	result := result{}
	var current tile
	for {
		out := make(chan int)
		done := make(chan bool)
		go comp.Run(out, done)
		select {
		case input :=<-ch:
			if input == 104 {
				comp.Input = -1
			} else if input == 108 {
				comp.Input = 1
			}
		case output := <-out:
			switch idx % 3 {
			case 0:
				current = tile{x:output}
			case 1: 
				current.y = output
			case 2:
				if current.x == -1 && current.y == 0 {
					result.score = output
					res<-result
					continue
				}
				current.id = tileType(output)
				tiles = append(tiles, current) 
			}
			idx++	
		case <-done:
			drawGame(tiles)
			fmt.Printf("finished, but now what?\n\r")
		}
		comp.Input = 0
	}
}

func drawGame(tiles []tile) {
	if len(tiles) == 0 {
		return
	}
	w, h := getSize(tiles)
	grid := make([][]tileType, h)
	for i := range h {
		grid[i] = make([]tileType, w)
	}

	for _, tile := range tiles {
		grid[tile.y][tile.x] = tile.id
	}

	fmt.Print("\rGAME\n\r")
	for _, row := range grid {
		for _, cell := range row {
			switch cell {
			case empty:
				fmt.Printf(" ")
			case wall:
				fmt.Printf("|")
			case block:
				fmt.Printf("0")
			case paddle:
				fmt.Printf("_")
			case ball:
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n\r")
	}
	fmt.Printf("\n\n\r")
}

func getSize(tiles []tile) (int, int) {
	w := 0
	h := 0
	for _, tile := range tiles {
		if tile.x > w {
			w = tile.x
		}
		if tile.y > h {
			h = tile.y
		}
	}
	return w+1, h+1
}
