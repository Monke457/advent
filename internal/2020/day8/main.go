package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2020/day8.txt")
	res, _ := run(data)
	fmt.Println("First:", res) 
	fmt.Println("Second:", second(data))
}

func run(data []string) (int, bool) {
	var acc int
	cache := []int{}

	i := 0
	for i < len(data) {
		if slices.Contains(cache, i) {
			cache = append(cache, i)
			break
		}

		cache = append(cache, i)
		op, val := parseOperation(data[i])

		switch op {
			case "acc":
				acc += val
				i++
			case "jmp":
				i += val
			case "nop":
				i++
		}
	}
	return acc, i != len(data) 
}

func second(data []string) int { 
	for i := len(data)-1; i >= 0; i-- {
		op, _ := parseOperation(data[i]); 
		newOp := strings.Builder{} 

		switch op {
			case "acc":
				continue
			case "jmp":
				newOp.WriteString("nop")
			case "nop":
				newOp.WriteString("jmp")
		}

		newOp.WriteString(data[i][3:])

		cp := make([]string, len(data))
		copy(cp, data)

		cp[i] = newOp.String()

		if res, broken := run(cp); !broken {
			return res
		}
	}
	return 0 
}

func parseOperation(line string) (string, int) {
	op, value, _ := strings.Cut(line, " ")
	val, err := strconv.Atoi(value[1:])
	if err != nil {
		panic(err)
	}
	if value[0] == '-' {
		val *= -1
	}
	return op, val 
}

