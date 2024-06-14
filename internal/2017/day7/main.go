package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type tower struct {
	name string
	weight int
	towers []*tower
}

var towers = []tower{}

func main() {
	data := reader.FileToArray("data/2017/day7.txt")

	for _, line := range data {
		towers = append(towers, parseTower(line))
	}
	
	root := &towers[0]
	for {
		next := root.getRoot()
		if next == nil {
			break
		}
		root = next 
	}
	fmt.Println("First:", root.name)

	imba, _ := root.findImbalance()
	fmt.Println("Second:", imba)
}

func (t tower) findImbalance() (int, bool) {
	res := t.weight
	if len(t.towers) == 0 {
		return res, false 
	}
	weights := map[int][]string{}
	for _, child := range t.towers { 
		c := findTower(child.name)
		weight, done := c.findImbalance()
		weights[weight] = append(weights[weight], child.name)
		if done {
			return weight, true
		}
	}

	tar := 0
	m := 0 
	for k, v := range weights {
		res += len(v) * k
		if tar == 0 && m == 0 {
			tar = k
			m = k
		}
		if len(v) < len(weights[m]) {
			m = k
		} else if len(v) > len(weights[m]) {
			tar = k
		}
	}

	if tar != m {
		bad := findTower(weights[m][0])
		return bad.weight + (tar - m), true 
	}	
	return res, false
} 

func (t tower) getRoot() *tower {
	for _, tower := range towers {
		if tower.name == t.name {
			continue
		}
		if len(tower.towers) == 0 {
			continue
		}
		for _, c := range tower.towers {
			if c.name == t.name {
				return &tower
			}
		}
	}
	return nil 
}

func parseTower(data string) tower {
	var name string

	res := findTower(name)
	if res == nil {
		res = &tower{}
	}

	if strings.Contains(data, " -> ") {
		_, towerParts, _ := strings.Cut(data, " -> ")
		tps := strings.Split(towerParts, ", ")
		for _, tp := range tps {
			t := findTower(tp)
			if t == nil {
				parsed := parseTower(tp)
				t = &parsed	
			}
			res.towers = append(res.towers, t)
		}
	}

	parts := strings.Split(data, " ")
	if res.name == "" {
		res.name = parts[0]
	}
	if res.weight == 0 {
		if len(parts) > 1 {
			weightStr := strings.Trim(parts[1], "()")
			w, err := strconv.Atoi(weightStr)
			if err != nil {
				panic(err)
			}
			res.weight = w
		}
	}
	return *res
}

func findTower(name string) *tower {
	for _, tower := range towers {
		if tower.name == name {
			return &tower
		}
	}
	return nil
}
