package main

import (
	iter "advent/pkg/iterator"
	"advent/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByRune("data/2015/day10.txt")

	l := 0
	for i := range 50 {
		col := iter.NewCollection(data)
		iter := col.CreateIterator()
		data = []int{}
		count := 1
		for iter.HasNext() {
			n := iter.LookAhead(1)
			c := *iter.Peek()
			if n != nil && c == *n {
				count++
			} else {
				data = append(data, count, c)
				count = 1
			}
			iter.GetNext()
		}
		l = len(data)
		if i == 39 {
			fmt.Println("First problem:", l)
		}
	}
	fmt.Println("Second problem:", l)
}
