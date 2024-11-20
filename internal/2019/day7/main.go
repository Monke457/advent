package main

import (
	ic "advent/internal/pkg/intcode"
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day7.txt")

	phaseCombos := getCombos([]int{0, 1, 2, 3, 4}) 

	signal := 0
	for _, combo := range phaseCombos {
		output := runFirst(data, combo)
		if output > signal {
			signal = output
		}
	}
	
	fmt.Println("First:", signal)

	signal = 0
	best := [5]int{}
	phaseCombos = getCombos([]int{5, 6, 7, 8, 9}) 
	for _, combo := range phaseCombos {
		comps := []ic.Computer{}
		for _, phase := range combo {
			comps = append(comps, ic.NewComputer(data, phase, 0, true))
		}
		output := runSecond(comps)
		if output > signal {
			best = combo
			signal = output
		}
	}
	fmt.Println("Second", signal, best)
}

func runFirst(data []int, combo [5]int) int {
	comps := []ic.Computer{}
	for _, phase := range combo {
		comps = append(comps, ic.NewComputer(data, phase, 0, true))
	}

	input := 0
	loop:
	for _, comp := range comps {
		comp.Input = input
		for { 
			output := make(chan int)
			done := make(chan bool)
			go run(&comp, output, done)
			select {
			case input = <-output:
			case <-done:
				continue loop
			}
		}
	}
	return input 
}

func runSecond(comps []ic.Computer) int {
	input := 0
	for {
		loop:
		for i, comp := range comps {
			if comp.Status == ic.Halted {
				continue
			}
			for { 
				output := make(chan int)
				done := make(chan bool)
				go run(&comp, output, done)
				select {
				case input = <-output:
					if i+1 >= len(comps) {
						comps[0].Input = input
					} else {
						comps[i+1].Input = input
					}
				case <-done:
				}
				comps[i] = comp
				continue loop
			}
		}
		if allHalted(comps) {
			break
		}
	}
	return input 
}

func allHalted(comps []ic.Computer) bool {
	for _, comp := range comps {
		if comp.Status != ic.Halted {
			return false
		}
	}
	return true
}

func run(c *ic.Computer, output chan int, done chan bool) {
	for {
		ch := make(chan int)
		halted := make(chan bool)
		go c.Run(ch, halted)
		select {
		case out := <-ch:
			output<-out
			return
		case shutdown := <-halted:
			done<-shutdown
			return
		}
	}
}

func getCombos(in []int) (out [][5]int) {
	n := len(in)

	indices := []int{0,1,2,3,4}

	count := 0
	for {
		count++
		temp := map[int]bool{}
		for i := 0; i < n; i++ {
			key := in[indices[i]]
			temp[key] = true
		}
		if len(temp) == n {
			arr := [5]int{}
			i := 0
			for key := range temp {
				arr[i] = key
				i++
			}
			out = append(out, arr)
		}

		next := n - 1
		for next >= 0 && indices[next]+1 >= len(in) {
			next--
		}

		if next < 0 {
			return
		}

		indices[next]++

		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}
	}
}

