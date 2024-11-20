package main

import(
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day9.txt")

	comp := ic.NewComputer(data, 0, 1, false)
	fmt.Println("First:", run(comp))

	comp = ic.NewComputer(data, 0, 2, false)
	fmt.Println("Second:", run(comp))
}

func run(comp ic.Computer) int {
	result := 0
	loop:
	for {
		out := make(chan int)
		done := make(chan bool)
		go comp.Run(out, done)
		select {
			case yield := <- out: 
			result = yield
			fmt.Println("Yielded output", yield)
		case <-done:
			break loop
		}
	}
	return result
}
