package main

import (
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day5.txt")

	comp := ic.NewComputer(data, 0, 1, false)
	output := run(comp)
	fmt.Println("First:", output)

	comp = ic.NewComputer(data, 0, 5, false)
	output = run(comp)
	fmt.Println("Second:", output)
}

func run(c ic.Computer) int {
	var output int
	for {
		ch := make(chan int)
		done := make(chan bool)
		go c.Run(ch, done)
		select {
		case out := <-ch:
			if output != 0 {
				fmt.Println("There was an error in the diagnostic (multiple non-zero outputs)")
				return out
			}
			output = out
		case <-done:
			return output
		}
	}
}

