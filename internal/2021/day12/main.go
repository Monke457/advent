package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2021/day12.txt")

	smols := map[string]bool{}
	cavemap := map[string]map[string]bool{}
	for _, line := range data {
		from, to, _ := strings.Cut(line, "-")
		if _, ok := cavemap[from]; !ok {
			cavemap[from] = map[string]bool{}
		}
		if _, ok := cavemap[to]; !ok {
			cavemap[to] = map[string]bool{}
		}
		cavemap[from][to] = true
		cavemap[to][from] = true
		if isLowerCase(from) {
			smols[from] = true
		}
		if isLowerCase(to) {
			smols[to] = true
		}
	}

	start, end := "start", "end"
	count := countPaths(cavemap, map[string]bool{start:true}, start, end)
	fmt.Println("Number of possible routes:", count)

	for smol := range smols {
		if smol == start || smol == end {
			continue
		}
		count += countPathsExtra(cavemap, map[string]int{start:1}, start, end, smol)
	}

	fmt.Println("Number of possible routes with smol boy:", count)
}

func countPaths(cavemap map[string]map[string]bool, visited map[string]bool, pos, end string) int {
	if pos == end {
		return 1
	}
	stack := cavemap[pos]
	count := 0
	for next := range stack {
		if visited[next] && isLowerCase(next) {
			continue
		}
		visited[next] = true
		count += countPaths(cavemap, visited, next, end)
		delete(visited, next)
	}
	return count
}

func countPathsExtra(cavemap map[string]map[string]bool, visited map[string]int, pos, end, smol string) int {
	if pos == end {
		if visited[smol] < 2 {
			return 0
		}
		return 1
	}
	count := 0
	for next := range cavemap[pos] {
		if next == smol {
			if visited[next] > 1 {
				continue
			}
		} else if visited[next] > 0 && isLowerCase(next) {
			continue
		}
		visited[next]++
		count += countPathsExtra(cavemap, visited, next, end, smol)
		visited[next]--
	}
	return count
}

func isLowerCase(val string) bool {
	return int(val[0]) - int('Z') > 0
}
