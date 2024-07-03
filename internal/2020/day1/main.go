package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArray("data/2020/day1.txt")

	var f, s = make(chan int), make(chan int)

	go first(data, f)
	go second(data, s)

	fmt.Println("First:", <-f)
	fmt.Println("Second:", <-s)
}

func first(data []int, c chan int) {
	for i := 0; i < len(data); i++ {
		for j := i+1; j < len(data); j++ {
			if data[i] + data[j] == 2020 {
				c<-data[i] * data[j]
				return	
			}
		}
	}
}

func second(data []int, c chan int) {
	for i := 0; i < len(data); i++ {
		for j := i+1; j < len(data); j++ {
			for k := j+1; k < len(data); k++ {
				if data[i] + data[j] + data[k] == 2020 {
					c <-data[i] * data[j] * data[k]
					return
				}
			}
		}
	}
}
