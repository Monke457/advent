package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2018/day7.txt")

	locked := map[byte][]byte{}
	all := map[byte]bool{}

	for _, line := range data {
		req := line[5]
		step := line[36]

		all[req] = true
		all[step] = true

		if _, ok := locked[step]; !ok {
			locked[step] = []byte{}
		}
		locked[step] = append(locked[step], req)
	}

	order := getOrder(locked, all)
	fmt.Printf("Order: %s\n", order)

	workers := 5
	t := completeTasks(locked, all, workers)
	fmt.Println("seconds taken to complete the tasks with", workers, "workers:", t)
}

func completeTasks(locked map[byte][]byte, all map[byte]bool, w int) int {
	complete := map[byte]bool{}
	inprogress := map[byte]int{}
	seconds := 0
	var i byte
	for len(locked) > 0 || len(inprogress) > 0 {
		seconds++
		if len(locked) > 0 {
			loop:
			for i = 65; i < 65 + 26; i++ {
				if len(inprogress) >= w {
					break loop
				}
				if complete[i] || inprogress[i] > 0 || !all[i] {
					continue
				}
				if reqs, ok := locked[i]; ok {
					for _, req := range reqs {
						if !complete[req] {
							continue loop
						}
					}
				} 
				inprogress[i]--
				delete(locked, i)
			}
		}
		for k, v := range inprogress {
			if v < 0 {
				inprogress[k] = 1
			} else {
				inprogress[k]++
			}
			if inprogress[k] >= int(k) - 4 {
				complete[k] = true
				delete(inprogress, k)
			}
		}
	}
	return seconds
}

func getOrder(original map[byte][]byte, all map[byte]bool) string {
	order := []byte{}
	open := map[byte]bool{}
	locked := copyMap(original)
	var i byte
	loop:
	for i = 65; i < 65 + 26; i++ {
		if open[i] || !all[i] {
			continue
		}
		if reqs, ok := locked[i]; ok {
			for _, req := range reqs {
				if !open[req] {
					continue loop
				}
			}
		} 
		order = append(order, i)
		open[i] = true
		delete(locked, i)
		i = 64
	}
	return string(order)
}

func copyMap(o map[byte][]byte) map[byte][]byte {
	m := map[byte][]byte{}
	for k, v := range o {
		newV := make([]byte, len(v))
		copy(newV, v)
		m[k] = newV
	}
	return m
}
