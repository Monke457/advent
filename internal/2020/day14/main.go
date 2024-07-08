package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2020/day14.txt")
	fmt.Println("First:", first(data))
}

func first(data []string) int {
	var mask string
	memory := map[int][]bool{}

	for _, line := range data {
		f, b, _ := strings.Cut(line, " = ")
		if f == "mask" {
			mask = b
			continue
		}

		f = strings.Trim(f, "mem[]")
		addr, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		val, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		memory[addr] = apply(val, mask) 
	}

	final := 0
	for _, mem :=  range memory {
		for i, val := range mem {
			if !val {
				continue
			}
			final += 1 << (35-i)
		}
	}
	return final
}

func apply(val int, mask string) []bool{
	bin := []rune(fmt.Sprintf("%b", val))
	slices.Reverse(bin)
	res := []bool{} 
	for i, m := range mask {
		pos := 35 - i
		if m == '0' {
			res = append(res, false)
		} else if m == '1' {
			res = append(res, true)
		} else if pos < len(bin) && bin[pos] == '1' {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res  
}
