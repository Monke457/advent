package main

import (
	"advent/internal/pkg/reader"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := strings.Split(reader.FileToString("data/2018/day9.txt"), " ",)

	players, _ := strconv.Atoi(data[0])
	maxpoints, _ := strconv.Atoi(data[6])
	
	maxpoints *= 100
	scores := make(map[int]int)

	marbles := list.New()
	current := marbles.PushBack(0)

	for next := 1; next < maxpoints; next++ {
		if next % 23 == 0 {
			cm := current
			for range 7 {
				cm = cm.Prev()
				if cm == nil {
					cm = marbles.Back()
				}
			}
			
			p := next % players
			scores[p] += next + cm.Value.(int)
			current = cm.Next()
			marbles.Remove(cm)

		} else {
			n := current.Next()
			if n == nil {
				n = marbles.Front()
			}
			current = marbles.InsertAfter(next, n)
		}
		fmt.Printf("\r%d/%d   ", current.Value, maxpoints)
	}
	fmt.Println("Highest score:", getHighest(scores))
}

func getHighest(scores map[int]int) int {
	highest := 0
	for _, v := range scores {
		if v > highest {
			highest = v
		}
	}
	return highest
}
