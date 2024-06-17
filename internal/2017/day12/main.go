package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

var groups = map[int]bool{}

func main() {
	data := reader.FileToArray("data/2017/day12.txt")

	conns := map[int][]int{}
	for _, line := range data {
		id, con := parseConns(line)
		conns[id] = con
	}
	
	findGroups(conns, 0)
	res := len(groups)
	fmt.Println("First:", res)
}

func findGroups(conns map[int][]int, id int) {
	groups[id] = true
	for _, c := range conns[id] {
		if _, ok := groups[c]; ok {
			continue
		}
		findGroups(conns, c)
	}
}

func parseConns(raw string) (int, []int) {
	idStr, connsStr, _ := strings.Cut(raw, " <-> ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	conns := []int{}
	for _, connStr := range strings.Split(connsStr, ", ") {
		conn, err := strconv.Atoi(connStr)
		if err != nil {
			panic(err)
		}
		conns = append(conns, conn)
	}

	return id, conns
}
