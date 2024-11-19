package main

import (
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day2.txt")

	comp := ic.NewComputer(data, 0, 0, false)

	first := run(comp, 12, 2)
	fmt.Println("Part 1:", first)

	var alarmcode int
	target := 19690720
	loop:
	for noun := 0; noun < 100; noun++ {	
		for verb := 0; verb < 100; verb++ {
			comp = ic.NewComputer(data, 0, 0, false)
			output := run(comp, noun, verb) 
			if output == target  {
				alarmcode = 100 * noun + verb
				break loop
			}
		}
	}
	fmt.Println("Part 2:", alarmcode)
}

func run(c ic.Computer, noun, verb int) int {
	c.Data[1] = noun 
	c.Data[2] = verb

	for {
		ch := make(chan int)
		done := make(chan bool)
		go c.Run(ch, done)
		select {
		case <-done:
			return c.Data[0] 
		}
	}
}
