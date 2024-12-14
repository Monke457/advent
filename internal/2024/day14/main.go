package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type robot struct {
	pos [2]int
	vel [2]int
}

const (
	WIDTH int = 101
	HEIGHT = 103
)

type model struct {
	robots map[[2]int][]robot
	tickCount int
	tickSize int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "+":
			m.tickSize++
		case "-":
			m.tickSize--
		case "l":
			m.tickCount += m.tickSize
			m.robots = tick(m.robots, m.tickSize)
		case "h":
			m.tickCount -= m.tickSize
			m.robots = tick(m.robots, m.tickSize * -1)
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return buildmap(m.robots) + fmt.Sprintf("\nTick size: %d\nTick count: %d\n", m.tickSize, m.tickCount)
}

func main() {
	data := reader.FileToArray("data/2024/day14.txt")

	//solved with trial and error because I'm dumb lol
	m := model{
		robots: parseRobots(data),
		tickSize: 103,
	}
	offset := 51
	m.robots = tick(m.robots, offset)
	m.tickCount = offset

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func countRobots(robots map[[2]int][]robot, quad int) int {
	w := (WIDTH -1) / 2
	h := (HEIGHT -1) / 2
	x, y := getStartingCoords(w, h, quad)

	count := 0
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			p := [2]int{i, j}
			if list, ok := robots[p]; ok {
				count += len(list)
			}
		}
	}

	return count
}

func getStartingCoords (w, h, quad int) (int, int) {
	switch quad {
	case 1:
		return w+1, 0
	case 2:
		return 0, h+1
	case 3:
		return w+1, h+1
	}
	return 0, 0
}

func tick(robots map[[2]int][]robot, seconds int) map[[2]int][]robot {
	results := map[[2]int][]robot{}
	for _, robotList := range robots {
		for _, r := range robotList {
			r.move(seconds)
			if _, ok := results[r.pos]; !ok {
				results[r.pos] = []robot{}
			}
			results[r.pos] = append(results[r.pos], r)
		}
	}
	return results 
}

func (r *robot) move(n int) {
	newPos := [2]int{
		wrap(r.pos[0] + n * r.vel[0], WIDTH),
		wrap(r.pos[1] + n * r.vel[1], HEIGHT),
	}
	r.pos = newPos
}

func wrap(i, n int) int {
	return ((i % n) + n) % n
}

func parseRobots(data []string) map[[2]int][]robot {
	robots := map[[2]int][]robot{}
	for _, line := range data {
		r := parseRobot(line)
		if _, ok := robots[r.pos]; !ok {
			robots[r.pos] = []robot{}
		}
		robots[r.pos] = append(robots[r.pos], r)
	}
	return robots
}

func parseRobot(str string) robot {
	parts := strings.Split(str, " ")
	pos := parseValues(parts[0][2:])
	vel := parseValues(parts[1][2:])
	return robot{pos: pos, vel: vel}
}

func parseValues(str string) [2]int {
	parts := strings.Split(str, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return [2]int{x, y}
}

func drawmap(robots map[[2]int][]robot) {
	fmt.Print(buildmap(robots))
}

func buildmap(robots map[[2]int][]robot) string {
	s := strings.Builder{}
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if rs, ok := robots[[2]int{j, i}]; ok {
				s.WriteString(fmt.Sprintf("%d", len(rs)))
			} else {
				s.WriteRune(' ')
			}
		}
		s.WriteRune('\n')
	}
	return s.String()
}
