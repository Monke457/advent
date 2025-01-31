package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := strings.Split(reader.FileToString("data/2018/day9.txt"), " ",)

	players, _ := strconv.Atoi(data[0])
	points, _ := strconv.Atoi(data[6])
	
	//points *= 100

	player := 0
	scores := map[int]int{}
	rounds := []int{0}
	current := 0

	for marble := 1; marble <= points; marble++ {
		if marble % 23 == 0 {
			scores[player] += marble 
			current = current-6
			if current < 0 {
				current = len(rounds)+current
			}
			
			scores[player] += rounds[current]
			rounds = append(rounds[:current], rounds[current+1:]...)
			current--

		} else {
			current = (current+2) % len(rounds)
			rounds = insert(rounds, current, marble)
		}
		player = (player+1) % players
		fmt.Printf("\r%d/%d   ", marble, points)
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

func insert(original []int, index, value int) []int {
	if index >= len(original) {
		return append(original, value)
	}
	arr := make([]int, len(original)+1)
	i := 0
	for _, val := range original {
		arr[i] = val
		if i == index {
			i++
			arr[i] = value
		}
		i++
	}
	return arr
}
