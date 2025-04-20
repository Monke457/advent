package main

import (
	"fmt"
)

func main() {
	const ELVES = 3004953

	elves := make([]int, ELVES)
	for i  := range ELVES {
		elves[i] = i+1
	}

	flag := map[int]int{}

	length := ELVES
	start := length >> 1 
	i := 0
	for len(elves) > 1 {
		idx := i + length >> 1 
		fmt.Printf("\rCalculating elf with all the presents currently: %d     ", elves[0])

		if idx >= length {
			skip := 0
			for j := start; j+skip < len(elves); j++ {
				skip += flag[j]
				flag[j] = 0
				if j+skip >= len(elves) {
					break
				}
				elves[j] = elves[j+skip]
			}
			elves = append(elves[i:length], elves[:i]...)
			start = length >> 1
			i = 0
			continue
		}

		flag[idx]++
		length--
		i++
	}

	fmt.Printf("\rElf with all the presents: %d                       \n", elves[0])
}
