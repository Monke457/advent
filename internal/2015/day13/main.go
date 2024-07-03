package main

import (
	"advent/internal/pkg/reader"
	"advent/internal/pkg/sorter"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	name string
	firstChild *node
	nextSibling *node
}

// so sadge :(
func main() {
	data := reader.FileToArray("data/2015/day13.txt")

	prefs := parsePrefs(data)
	names := sorter.GetOrderByKey(prefs)
	
	root := node{ name: names[0] }
	createTree(root, names[1:], 0)

	fmt.Println("ligma")
}

func createTree(root node, names []string, start int) {
	if start == len(names) {
		return
	}

	var curr *node = &root
	for i := start; i < len(names); i++ {
		n := i % len(names)
		curr.name = names[n]
		curr.firstChild = &node{}
		curr = curr.firstChild
	}

	createTree(root, names, start+1) 
}

func parsePrefs(data []string) map[string]map[string]int {
	res := map[string]map[string]int{}
	for _, l := range data {
		parts := strings.Split(l, " ")
		name := parts[0]
		neighbour := strings.Trim(parts[len(parts)-1], ".")
		val, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err)
		}
		if parts[2] == "lose" {
			val *= -1
		}
		if _, ok := res[name]; !ok {
			res[name] = map[string]int{}
		}
		res[name][neighbour] = val
	}
	return res
}
