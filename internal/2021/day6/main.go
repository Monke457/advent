package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	fishies := reader.FileToIntArrayByComma("data/2021/day6.txt")

	fishmap := mapFishies(fishies)
	queue := [2]int{0, 0}

	days := 256
	for i := range days {
		cycle := i % 7
		queuedFish := fishmap[cycle]

		newFish := queue[0]
		queue[0] = queue[1]
		queue[1] = queuedFish 

		if newFish > 0 {
			fishmap[cycle] += newFish
		}
	}

	fmt.Printf("No. of fishies after %d days: %d\n", days, countfish(fishmap, queue))
}

func countfish(fish map[int]int, queue [2]int) int {
	sum := 0
	for _, no := range fish {
		sum += no
	}
	return sum + queue[0] + queue[1]
}

func mapFishies(fish []int) map[int]int {
	fishmap := map[int]int{}
	for i := range fish {
		fishmap[fish[i]]++
	}
	return fishmap
}
