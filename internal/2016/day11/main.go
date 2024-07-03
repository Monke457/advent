package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

type column struct {
	floors []floor
	elevator int
}

type floor struct {
	id int
	content []string
}

func main() {
	data := reader.FileToArray("data/2016/day11.txt")

	column := initialLocations(data) 
	column.print()
}

func initialLocations(data []string) column {
	res := column{}
	for i, line := range data {
		_, content, _ := strings.Cut(line, " contains")
	
		if strings.Contains(content, "nothing") {
			res.floors = append(res.floors, floor{ id: i, content: []string{} })
		} else {
			res.floors = append(res.floors, floor{ id: i, content: getElements(content) })
		}
	}
	return res
}

func getElements(str string) []string {
	res := []string{}
	parts := strings.Split(str, "a ")
	for _, part := range parts {
		part = strings.Trim(part, "and ")
		els := strings.Split(part, " ")
		l := len(els)
		if l < 2 {
			continue
		}
		res = append(res, fmt.Sprintf("%c%c", els[l-2][0], els[l-1][0]))
	}
	return res
}

func (c column) print() {
	m := c.maxLen()
	for i := len(c.floors); i > 0; i-- {
		fmt.Printf("F%d ", i)
		if c.elevator == i-1 {
			fmt.Printf("E ")
		} else {
			fmt.Printf(". ")
		}
		j := 0
		for _, v := range c.floors[i-1].content {
			j++	
			fmt.Printf("%s ", v)
		}
		for j < m {
			j++
			fmt.Printf(" . ")
		}
		fmt.Println()
	}
}

func (c column) maxLen() int {
	m := 0
	for _, f := range c.floors {
		m = max(m, len(f.content))
	}
	return m
}
