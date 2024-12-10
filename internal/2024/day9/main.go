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
	fmt.Println("Checksum first:", checksum)

	sorted = sortWholeBlocks(blocks)
	checksum = CreateChecksum(sorted)
	fmt.Println("Checksum second:", checksum)
}

func CreateChecksum(files []*int) int {
	var total int = 0
	for i := 0; i < len(files); i++ {
		if files[i] == nil {
			continue
		}
		total += *files[i] * i
	} 
	return total 
}

func sortWholeBlocks(blocks []*int) []*int {
	sorted := make([]*int, len(blocks))
	copy(sorted, blocks)

	idxEnd := len(blocks)-1
	idxStart := idxEnd
	for {
		idxStart, idxEnd = getNextIndexes(sorted, idxEnd)
		if idxStart == idxEnd {
			break
		}

		length := idxEnd - idxStart
		gapStart := findGap(sorted, length)

		if gapStart < idxStart && gapStart > 0 {
			for i := 0; i < length; i++ {
				sorted[gapStart+i] = sorted[idxStart+i]
				sorted[idxStart+i] = nil
			}
		}

		idxEnd = idxStart-1
	}

	return sorted
}

func getNextIndexes(blocks []*int, e int) (int, int){
	for e >= 0 && blocks[e] == nil {
		e--
	}

	s := e
	for s > 0 && blocks[s] != nil && *blocks[s] == *blocks[e] {
		s--
	}

	s++
	e++
	return s, e
}

func findGap(blocks []*int, length int) int {
	for i := 0; i < len(blocks); i++ {
		idx := i
		gapLength := 0
		for blocks[idx] == nil {
			gapLength++
			idx++
			if idx >= len(blocks) {
				break
			}
		}
		if gapLength >= length {
			return i
		}
	}
	return -1 
}

func SortBlocks(blocks []*int) []*int {
	sorted := []*int{}

	last := len(blocks)-1
	for blocks[last] == nil {
		last--
	}

	for i := 0; i < len(blocks); i++ {
		if i > last {
			break
		}
		if blocks[i] != nil {
			entry := *blocks[i]
			sorted = append(sorted, &entry)
		} else {
			entry := *blocks[last]
			sorted = append(sorted, &entry)
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
