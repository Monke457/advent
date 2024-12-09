package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
)

func main() {
	data := reader.FileToString("data/2024/day9.txt")

	blocks := ParseBlocks(data)
	sorted := SortBlocks(blocks)
	checksum := CreateChecksum(sorted)

	fmt.Println("Checksum:", checksum)
}

func CreateChecksum(files []string) int {
	var total int = 0
	for i := 0; i < len(files); i++ {
		num, _ := strconv.Atoi(files[i]) 
		total += num * i
	} 
	return total 
}

func SortBlocks(blocks []string) []string {
	sorted := []string{}

	last := len(blocks)-1
	for blocks[last] == "." {
		last--
	}

	for i := 0; i < len(blocks); i++ {
		if i > last {
			break
		}
		if blocks[i] != "." {
			sorted = append(sorted, blocks[i])
		} else {
			sorted = append(sorted, blocks[last])
			last--
			for blocks[last] == "." {
				last--
			}
		}
	}

	return sorted
}

func getFill(blocks []int, length int) []int {
	res := []int{}
	last := len(blocks)-1
	val := blocks[last-1]
	count := blocks[last]
	for {
		if length <= 0 {
			break
		}
		if count >= length {
			res = append(res, val, length)
			break
		}
		res = append(res, val, count)
		length -= count
		last -= 3
		val = blocks[last-1]
		count = blocks[last]
	}
	return res
}

func ParseBlocks(numstr string) []string {
	blocks := []string{}
	var idx int = 0
	for i, num := range numstr {
		even := i % 2 == 0
		for range int(num) - 48 {
			if even {
				blocks = append(blocks, strconv.Itoa(idx))
			} else {
				blocks = append(blocks, ".") 
			}
		}
		if even {
			idx++
		}
	}
	return blocks
}
