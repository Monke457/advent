package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2020/day6.txt")

	var sum1, sum2, groupSize int
	answers := make(map[rune]int) 

	for _, line := range data {
		if len(line) == 0 {
			sum1 += len(answers)
			for _, v := range answers {
				if groupSize == v {
					sum2++
				}
			}
			clear(answers)
			groupSize = 0
			continue
		}
		groupSize++
		for _, r := range line {
			answers[r]++ 
		}
	}

	//don't forget the last answers, stupid.
	sum1 += len(answers)
	for _, v := range answers {
		if groupSize == v {
			sum2++
		}
	}

	fmt.Println("First:", sum1)
	fmt.Println("Second:", sum2)
}
