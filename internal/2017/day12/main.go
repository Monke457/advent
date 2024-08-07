package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

var groups = []map[int]bool{}

func main() {
	data := reader.FileToArray("data/2017/day12.txt")

	ids := []int{}
	conns := map[int][]int{}
	for _, line := range data {
		id, con := parseConns(line)
		ids = append(ids, id)
		conns[id] = con
	}
	
	groupId := 0
	for i := 0; i < len(ids); i++ {
		if groupContains(i) {
			continue
		}
		if len(groups) == groupId {
			groups = append(groups, map[int]bool{})
		}
		findGroups(conns, ids[i], groupId)
		groupId++
	}
	res := len(groups[0])
	fmt.Println("First:", res)
	fmt.Println("Second:", len(groups))
}

func groupContains(n int) bool {
	for _, group := range groups {
		_, ok := group[n]
		if ok {
			return true
		}
	}
	return false
}

func findGroups(conns map[int][]int, id, groupId int) {
	groups[groupId][id] = true
	for _, c := range conns[id] {
		if _, ok := groups[groupId][c]; ok {
			continue
		}
		findGroups(conns, c, groupId)
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
