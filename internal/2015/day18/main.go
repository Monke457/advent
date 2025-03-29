package main

import (
	"advent/internal/pkg/grid"
	"advent/internal/pkg/reader"
	"fmt"
)

type model struct {
	lights [][]rune
}

func (m model) isCorner(y, x int) bool {
	if y == 0 {
		return x == 0 || x == len(m.lights[y])-1
	}
	if y == len(m.lights)-1 {
		return x == 0 || x == len(m.lights[y])-1
	}
	return false
}

func (m *model) tickLights() {
	newLights := make([][]rune, len(m.lights))

	for i := 0; i < len(m.lights); i++ {
		newLights[i] = make([]rune, len(m.lights[i]))

		for j := 0; j < len(m.lights[i]); j++ {
			if m.isCorner(i, j) {
				newLights[i][j] = '#'
				continue
			}
			count := 0
			for _, pos := range grid.GetNeighbours(m.lights, [2]int{i, j}) {
				if m.lights[pos[0]][pos[1]] == '#' {
					count++
				}
			}
			r := m.lights[i][j]
			if r == '#' && (count < 2 || count > 3) {
				r = '.'
			} else if r == '.' && count == 3 {
				r = '#'
			}
			newLights[i][j] = r
		}
	}
	m.lights = newLights
}

func (m model) countLights() int {
	count := 0
	for _, row := range m.lights {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}

const TICKS = 100

func main() {
	model := model{
		lights: reader.FileTo2DArray("data/2015/day18.txt"),
	}

	model.lights[0][0] = '#'
	model.lights[len(model.lights)-1][0] = '#'
	model.lights[0][len(model.lights)-1] = '#'
	model.lights[len(model.lights)-1][len(model.lights)-1] = '#'

	for range TICKS {
		model.tickLights()
	}

	fmt.Println("Lights after", TICKS, "ticks:", model.countLights())
}
