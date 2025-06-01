package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type model struct {
	Lights map[[2]int][]Light 
	yRange [2]int
	xRange [2]int
	Ticks int
	TickSize int
}

type Light struct {
	yVel int
	xVel int
}

func (m *model) Tick(n int) {
	newLights := map[[2]int][]Light{}

	first := true
	for coords, lights := range m.Lights {
		for _, light := range lights {
			newCoord := [2]int{
				coords[0] + n * light.yVel, 
				coords[1] + n * light.xVel,
			}
			if _, ok := newLights[newCoord]; !ok {
				newLights[newCoord] = []Light{}
			}
			newLights[newCoord] = append(newLights[newCoord], light)
			if first {
				m.yRange[0] = newCoord[0]
				m.yRange[1] = newCoord[0]
				m.xRange[0] = newCoord[1]
				m.xRange[1] = newCoord[1]
				first = false
			} else {
				m.updateRanges(newCoord)
			}
		}
	} 
	m.Lights = newLights
}

func (m model) View() string {
	sb := strings.Builder{}
	for y := m.yRange[0]; y <= m.yRange[1]; y++ {
		for x := m.xRange[0]; x <= m.xRange[1]; x++ {
			if _, ok := m.Lights[[2]int{y, x}]; ok {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func newModel() model {
	return model{
		Lights: map[[2]int][]Light{},
		yRange: [2]int{}, 
		xRange: [2]int{}, 
		TickSize:1,
	}
}

func initModel(data []string) model {
	m := newModel()
	for _, line := range data {
		_, position, _ := strings.Cut(line, "<")
		position, rest, _ := strings.Cut(position, ">")
		_, velocity, _ := strings.Cut(rest, "<")
		velocity, _, _ = strings.Cut(velocity, ">")

		pos_a_str, pos_b_str, _ := strings.Cut(position, ", ")
		vel_a_str, vel_b_str, _ := strings.Cut(velocity, ", ")

		pos_a, _ := strconv.Atoi(strings.TrimSpace(pos_a_str))
		pos_b, _ := strconv.Atoi(strings.TrimSpace(pos_b_str))

		vel_a, _ := strconv.Atoi(strings.TrimSpace(vel_a_str))
		vel_b, _ := strconv.Atoi(strings.TrimSpace(vel_b_str))

		coord := [2]int{pos_b, pos_a} 
		if _, ok := m.Lights[coord]; !ok {
			m.Lights[coord] = []Light{}
		}
		m.Lights[coord] = append(m.Lights[coord], Light{yVel:vel_b, xVel: vel_a})
		m.updateRanges(coord)
	}

	return m
}

func (m *model) updateRanges(coord [2]int) {
	if coord[0] < m.yRange[0] { m.yRange[0] = coord[0] }
	if coord[0] > m.yRange[1] { m.yRange[1] = coord[0] }
	if coord[1] < m.xRange[0] { m.xRange[0] = coord[1] }
	if coord[1] > m.xRange[1] { m.xRange[1] = coord[1] }
}

func main() {
	data := reader.FileToArray("data/2018/day10.txt")

	model := initModel(data)
	model.Tick(10243)
	fmt.Print(model.View())
	fmt.Println("10243 seconds")
}
