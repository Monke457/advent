package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2024/day23.txt")

	connections := parseConnections(data)
	trios := map[[3]string]bool{}
	for id := range connections {
		for trio := range findTrios(connections, id) {
			trios[trio] = true
		}
	}

	subset := findTriosWithT(trios)
	fmt.Println("trios with 't' computer id found:", len(subset))

	computerSets := findComputerSets(connections, trios) 
	var size int 
	var set string 
	for key, val := range computerSets {
		if val > size {
			set = key 
			size = val
		}
	}
	fmt.Println("largest set with length", size, "found:", set)
}



func findComputerSets(conns map[string][]string, trios map[[3]string]bool) map[string]int {
	sets := [][]string{}
	seen := map[string]bool{}

	loop:
	for trio := range trios {
		for _, id := range trio {
			if seen[id] {
				continue loop
			}
			seen[id] = true
		}
		newSet := buildSet(conns, trio[:])
		sets = append(sets, newSet)
	}
	
	result := map[string]int{}
	for _, set := range sets {
		slices.Sort(set)
		key := strings.Join(set, ",")
		if _, ok := result[key]; ok {
			continue
		}
		result[key] = len(set)
	}

	return result
}

func buildSet(conns map[string][]string, ids []string) []string {
	remaining := map[string]bool{}
	for conn := range conns {
		if slices.Contains(ids, conn) {
			continue
		}
		remaining[conn] = true
	}

	for len(remaining) > 0 {
		loop:
		for conn := range remaining {
			for _, id := range ids {
				if !slices.Contains(conns[conn], id) {
					delete(remaining, conn)
					continue loop
				}
			}
			ids = append(ids, conn)
			delete(remaining, conn)
		}
	}
	return ids
}

func findTriosWithT(trios map[[3]string]bool) map[[3]string]bool {
	subset := map[[3]string]bool{}
	loop:
	for trio := range trios {
		for _, id := range trio {
			if id[0] == 't' {
				subset[trio] = true
				continue loop
			}
		}
	}
	return subset
}

func findTrios(connections map[string][]string, id string) map[[3]string]bool {
	trios := [][]string{}
	for _, conn := range connections[id] {
		for _, subconn := range connections[conn] {
			if subconn == id {
				continue
			} 
			if !slices.Contains(connections[subconn], id) {
				continue
			}
			trios = append(trios, []string{id, conn, subconn})
		}
	}
	result := map[[3]string]bool{}
	for _, trio := range trios {
		slices.Sort(trio)
		result[[3]string{trio[0], trio[1], trio[2]}] = true
	}
	return result
}

func parseConnections(data []string) map[string][]string {
	compIds := map[string][]string{}
	for _, line := range data {
		id1, id2, _ := strings.Cut(line, "-")
		compIds[id1] = append(compIds[id1], id2)
		compIds[id2] = append(compIds[id2], id1)
	}
	return compIds
}
