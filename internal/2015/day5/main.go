package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strings"
)

var illegal = [4]string{"ab", "cd", "pq", "xy"}
var vowels = [5]rune{'a', 'e', 'i', 'o', 'u'}

func main() {
	data := reader.FileToArray("data/2015/day5.txt")

	niceCount := 0
	nicerCount := 0
	for _, l := range data {
		if nice(l) {
			niceCount++
		}
		if nicer(l) {
			nicerCount++
		}
	}
	fmt.Println("First problem:", niceCount)
	fmt.Println("Second problem:", nicerCount)
}

// rule 1: at least 3 vowels
// rule 2: at least 1 double letter
// rule 3: none of these pairs: ab, cd, pq, xy
func nice(str string) bool {
	for _, s := range illegal {
		if strings.Contains(str, s) {
			return false
		}
	}

	var r1, r2 bool
	var last rune
	v := 0
	for _, r := range str {
		if !r1 {
			if slices.Contains(vowels[:], r) {
				v++
			}
			if v >= 3 {
				r1 = true
			}
		}
		if !r2 {
			if last == r {
				r2 = true
			}
			last = r
		}
		if r1 && r2 {
			break
		}
	}
	return r1 && r2
}

// rule 1: at least two appearances of a pair (non overlapping)
// rule 2: at least one repeat letter with a letter between
func nicer(str string) bool {
	var r1, r2 bool
	var last, laster rune
	for _, r := range str {
		if !r1 {
			pair := fmt.Sprintf("%c%c", last, r)
			if strings.Count(str, pair) > 1 {
				r1 = true
			}
		}
		if !r2 {
			if laster == r {
				r2 = true
			}
			laster = last
		}
		last = r
		if r1 && r2 {
			break
		}
	}
	return r1 && r2
}
