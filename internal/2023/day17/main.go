package main

import (
	"log"
	"math"

	tea "github.com/charmbracelet/bubbletea"

	"advent/internal/pkg/reader"
	"fmt"
	"time"
)

type model struct {
	orientation int8 
	dest [2]int 
	data [][]int
	turnCD int8
	heatLoss int
	seen map[[2]int]bool
}

type tickMsg time.Time

func main() {
	data := reader.FileTo2DIntArray("data/2023/day17.txt")

	model := model{ 
		data: data, 
		seen: map[[2]int]bool{ 
			{0, 0}: true,
		},
	}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		m.orient()
		err := m.move() 
		if err != nil {
			fmt.Println(err)
			return m, tea.Quit
		}
		m.seen[m.dest] = true
		m.turnCD++
		m.heatLoss += m.data[m.dest[0]][m.dest[1]]
		return m, tick
	}
	return m, nil
}

func tick() tea.Msg {
	time.Sleep(500*time.Millisecond)
	return tickMsg{} 
}

func (m *model) orient() {
	nh := m.getNeighbourHeat()
	o := m.orientation
	val := math.MaxInt8
	for i := int8(0); i < 4; i++ {
		if m.orientation != i {
			continue
		}
		if nh[i] != nil && m.turnCD < 3 {	
			val = *nh[i]
			o = i
		}
		j := (i + 1) % 4
		if nh[j] != nil && *nh[j] < val {	
			val = *nh[j]
			o = j 
		}
		j = (i + 3) % 4
		if nh[j] != nil && *nh[j] < val {	
			val = *nh[j]
			o = j
		}
		break
	}

	if o != m.orientation {
		m.turnCD = 0
		m.orientation = o
	}
}

func (m model) getNeighbourHeat() [4]*int {
	res := [4]*int{}
	if _, ok := m.seen[[2]int{m.dest[0], m.dest[1] + 1}]; !ok {
		res[0] = m.getHeat(m.dest[0] - 1, m.dest[1] + 1)
	}
	if _, ok := m.seen[[2]int{m.dest[0] - 1, m.dest[1]}]; !ok {
		res[1] = m.getHeat(m.dest[0] + 1, m.dest[1] - 1)
	}
	if _, ok := m.seen[[2]int{m.dest[0], m.dest[1] - 1}]; !ok {
		res[2] = m.getHeat(m.dest[0] - 1, m.dest[1] - 3)
	}
	if _, ok := m.seen[[2]int{m.dest[0] - 1, m.dest[1]}]; !ok {
		res[3] = m.getHeat(m.dest[0] - 3, m.dest[1] - 1)
	}
	return res
}

func (m model) getHeat(x, y int) *int {
	if x >= len(m.data[0]) - 1 || y >= len(m.data) - 1 {
		return nil
	}
	if x + 2 < 0 || y + 2 < 0 {
		return nil
	}

	var res, count int 
	
	for i := x; i < i + 3; i++ {
		if i < 0 {
			continue
		}
		if i >= len(m.data) {
		 	break	
		}
		for j := y; j < y + 3; j++ {
			if j < 0 {
				continue
			}
			if j >= len(m.data[0]) {
				break
			}
			res += m.data[i][j]
			count++
		}
	}
	val := res / count
	return &val
}

func (m *model) move() error {
	switch m.orientation {
	case 0:
		m.dest[1]++
	case 1: 
		m.dest[0]++
	case 2:
		m.dest[1]--
	case 3:
		m.dest[0]--
	}
	if m.dest[0] < 0 || m.dest[0] >= len(m.data) || m.dest[1] < 0 || m.dest[1] >= len(m.data[0]) {
		return fmt.Errorf("Out of bounds destination!! %d", m.dest)
	}
	return nil
} 

func (m model) View() string {
	var view string
	for i := range m.data {
		for j := range m.data {
			if m.dest[0] == i && m.dest[1] == j {
				view += fmt.Sprintf("\033[31m%d\033[0m", m.data[i][j])
			} else if _, ok := m.seen[[2]int{i, j}]; ok {
				view += fmt.Sprintf("\033[32m%d\033[0m", m.data[i][j])
			} else {
				view += fmt.Sprintf("%d", m.data[i][j])
			}
		}
		view += fmt.Sprintln()
	}
	view += fmt.Sprintf("\nHeat loss: %d\n", m.heatLoss)
	return view
}
