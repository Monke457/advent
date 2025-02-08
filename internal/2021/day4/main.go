package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	id int
	cells [][]*cell
}

type cell struct {
	val int
	hit bool 
}

func main() {
	data := reader.FileToArray("data/2021/day4.txt")

	rounds := parseNums(strings.Split(data[0], ",")) 
	boards := parseBoards(data[2:])

	var firstScore int
	var finalScore int

	loop:
	for _, round := range rounds {
		for i, board := range boards {
			if len(boards) < 1 {
				panic("Big problem!")
			}
			board.playRound(round)
			if board.win() {
				if len(boards) == 1 {
					finalScore = board.sumNonHits() * round
					break loop
				}
				if firstScore == 0 {
					firstScore = board.sumNonHits() * round
				}
				delete(boards, i)
			}
		}
	}

	fmt.Println("First:", firstScore)
	fmt.Println("Last:", finalScore)
}

func (b board) sumNonHits() int {
	sum := 0
	for _, row := range b.cells {
		for _, cell := range row {
			if cell.hit {
				continue
			}
			sum += cell.val
		}
	}
	return sum
}

func (b board) win() bool {
	return checkHorizontal(b.cells) || checkVertical(b.cells)
}

func checkHorizontal(c [][]*cell) bool {
	loop:
	for _, row := range c {
		for _, cell := range row {
			if !cell.hit {
				continue loop
			}
		}
		return true
	}
	return false 
}

func checkVertical(c [][]*cell) bool {
	loop:
	for i := 0; i < len(c[0]); i++ {
		for _, row := range c {
			if !row[i].hit {
				continue loop
			}
		}
		return true
	}
	return false 
}

func (b board) playRound(val int) bool {
	for _, row := range b.cells {
		for _, cell := range row {
			if cell.val == val {
				cell.hit = true
				return true
			}
		}
	}
	return false
}

func parseBoards(data []string) map[int]board {
	boards := map[int]board{}
	cells := [][]*cell{}
	i := 1
	for _, line := range data {
		if len(line) == 0 {
			boards[i] = board{ id:i, cells: cells }
			cells = [][]*cell{}
			i++
			continue
		}
		cells = append(cells, parseCells(strings.Split(line, " ")))
	}
	boards[i] = board{ id: i, cells: cells }
	return boards
}

func parseCells(nums []string) []*cell {
	res := []*cell{}
	for _, num := range nums {
		if num == "" {
			continue
		}
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error while converting number", num)
			continue
		}
		res = append(res, &cell{ val: val })
	}
	return res 
}

func parseNums(nums []string) []int {
	res := []int{}
	for _, num := range nums {
		if num == "" {
			continue
		}
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error while converting number", num)
			continue
		}
		res = append(res, val)
	}
	return res 
}

const color = "\033[1;32m"
const colorless = "\033[0m"

func (b board) print() {
	fmt.Println("Board", b.id)
	for _, row := range b.cells {
		for _, cell := range row {
			if cell.hit {
				fmt.Fprintf(os.Stdout, "%s%d %s", color, cell.val, colorless)
			} else {
				fmt.Fprintf(os.Stdout, "%d ", cell.val)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
