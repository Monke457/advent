package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToString("data/2017/day9.txt")
	solveFirst(data)
	solveSecond(data)
}

func solveFirst(data string) {
	level := 0
	score := 0
	garbage := false
	for i := 0; i < len(data); i++ {
		if data[i] == '!' {
			i++
			continue
		}
		if data[i] == '>' {
			garbage = false
		}
		if garbage {
			continue
		}
		if data[i] == '<' {
			garbage = true
		}
		if data[i] == '{' {
			level++
			score += level
		}
		if data[i] == '}' {
			level--
		}
	}
	fmt.Println("First:", score)
}

func solveSecond(data string) {
	chars := 0
	garbage := false
	for i := 0; i < len(data); i++ {
		if data[i] == '!' {
			i++
			continue
		}
		if data[i] == '>' {
			garbage = false
		}
		if garbage {
			chars++
		}
		if data[i] == '<' {
			garbage = true
		}
	}
	fmt.Println("Second:", chars)
}
