package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToString("data/2024/day9.txt")

	blocks := ParseBlocks(data)
	sorted := SortBlocks(blocks)
	checksum := CreateChecksum(sorted)

	fmt.Println("Checksum:", checksum)
}

func CreateChecksum(files []int) int {
	var total int = 0
	for i := 0; i < len(files); i++ {
		total += files[i] * i
	} 
	return total 
}

func SortBlocks(blocks []*int) []int {
	sorted := []int{}

	last := len(blocks)-1
	for blocks[last] == nil { 
		last--
	}

	for i := 0; i < len(blocks); i++ {
		if i > last {
			break
		}
		if blocks[i] != nil {
			sorted = append(sorted, *blocks[i])
		} else {
			sorted = append(sorted, *blocks[last])
			last--
			for blocks[last] == nil {
				last--
			}
		}
	}

	return sorted
}

func ParseBlocks(numstr string) []*int {
	blocks := []*int{}
	var idx int = 0
	for i, num := range numstr {
		even := i % 2 == 0
		for range int(num) - 48 {
			if even {
				entry := idx
				blocks = append(blocks, &entry)
			} else {
				blocks = append(blocks, nil) 
			}
		}
		if even {
			idx++
		}
	}
	return blocks
}
