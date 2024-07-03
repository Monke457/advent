package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

var ops = map[string]func(int, int)bool{
	"==": func(key int, n int)bool{ return key == n },
	"!=": func(key int, n int)bool{ return key != n },
	">=": func(key int, n int)bool{ return key >= n },
	"<=": func(key int, n int)bool{ return key <= n },
	">": func(key int, n int)bool{ return key > n },
	"<": func(key int, n int)bool{ return key < n },
}

var registers = map[string]int{}

func main() {
	data := reader.FileToArray("data/2017/day8.txt")

	mx := 0
	for _, line := range data {
		inst, cond, _ := strings.Cut(line, " if ")
		if parseCondition(cond) {
			newVal := parseInstruction(inst)
			if newVal > mx {
				mx = newVal
			}
		}
	}
	
	res := 0
	for _, v := range registers {
		if v > res {
			res = v
		}
	}

	fmt.Println("First:", res)
	fmt.Println("Second:", mx)
}

func parseCondition(str string) bool {
	parts := strings.Split(str, " ")
	f := ops[parts[1]]
	val, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	if f(registers[parts[0]], val) {
		return true
	}
	return false 
}

func parseInstruction(str string) int {
	parts := strings.Split(str, " ")
	val, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	switch parts[1] {
	case "dec":
		registers[parts[0]] -= val
	case "inc":
		registers[parts[0]] += val
	}
	return registers[parts[0]]
}


