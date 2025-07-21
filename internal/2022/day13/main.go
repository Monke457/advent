package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2022/day13.txt")

	pkg := 1
	sum := 0
	for i := 0; i < len(data); i+=3 {
		if (compare(data[i], data[i+1]) == 1) {
			fmt.Println("package", pkg) 
			fmt.Println(data[i]) 
			fmt.Println(data[i+1]) 
			sum	+= pkg
		}
		pkg++
	}
	fmt.Println("Sum of indices of packages in correct order:", sum)
}

func compare(list_a, list_b string) int {
	ret_val, level_a, level_b := 0, 0, 0
	var next_a, next_b *rune

	for i_a, i_b := 0, 0; i_a < len(list_a) || i_b < len(list_b); {

		loop_a:
		for {
			if (i_a >= len(list_a)) { next_a = nil; break loop_a }
			item := rune(list_a[i_a])
			switch item {
			case ',': i_a++;
			case '[':
				level_a++; i_a++
				if (list_a[i_a] == ']') { next_a = nil; break loop_a }
			case ']': level_a--; i_a++
			default: next_a = &item; break loop_a
			}
		}

		loop_b:
		for {
			if (i_b >= len(list_b)) { next_b = nil; break loop_b }
			item := rune(list_b[i_b])
			switch item {
			case ',': i_b++;
			case '[': 
				level_b++; i_b++
				if (list_b[i_b] == ']') { next_b = nil; break loop_b }
			case ']': level_b--; i_b++
			default: next_b = &item; break loop_b
			}
		}

		if (next_a == nil && next_b == nil) { continue }
		if (next_a == nil) {ret_val = 1; break }
		if (next_b == nil) {ret_val = -1; break }

		if (*next_a < *next_b) { ret_val = 1; break }
		if (*next_a > *next_b) { ret_val = -1; break }

		if (level_a < level_b) { ret_val = 1; break }
		if (level_a > level_b) { ret_val = -1; break }

		i_a++; i_b++
	} 
	return ret_val
}

