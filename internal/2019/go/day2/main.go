package main

import (
	"advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day2.txt")
	first, err := intcode.RunDay2(data, 12, 2)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Part 1:", first)
	}

	var alarmcode int
	target := 19690720
	loop:
	for noun := 0; noun < 100; noun++ {	
		for verb := 0; verb < 100; verb++ {
			output, err := intcode.RunDay2(data, noun, verb) 
			if err != nil {
				panic(err)
			} else if output == target  {
				alarmcode = 100 * noun + verb
				break loop
			}
		}
	}
	fmt.Println("Part 2:", alarmcode)
}
