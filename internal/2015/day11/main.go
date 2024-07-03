package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	raw := reader.FileToString("data/2015/day11.txt")

	data := make([]*rune, len(raw))
	for i, r := range raw {
		data[i] = &r
	}

	for i := range 2 {
		for !validPassword(data) {
			increment(data, len(data)-1)
			shiftBaddies(data)
		}
		fmt.Printf("Problem %d:", i+1)
		printPtrs(data)

		increment(data, len(data)-1)
	}
}

func increment(data []*rune, pos int) {
	if pos < 0 {
		return
	}
	if *data[pos] == 'z' {
		*data[pos] = 'a'
		increment(data, pos-1)
	} else {
		*data[pos]++
	}
	return
}

func shiftBaddies(data []*rune) {
	for i, r := range data {
		if *r == 'o' || *r == 'i' || *r == 'l' {
			increment(data, i)
		}
	}
}

// rule 1: must contain a string of 3 consecutive letters
// rule 2: no i o or l
// rule 3: 2 non-overlapping pairs of letters
func validPassword(data []*rune) bool {
	var r1, r3 bool
	c1, c3 := [3]rune{}, 0
	var prev rune
	for _, r := range data {
		if *r == 'i' || *r == 'o' || *r == 'l' {
			return false
		}
		if !r1 {
			c1 = shiftRunes(c1, *r)
			if checkRunes(c1[0], c1[1], c1[2]) {
				r1 = true
			}
		}
		if !r3 {
			if prev == *r {
				c3++
				if c3 == 2 {
					r3 = true
				}
				prev = ' '
				continue
			}
			prev = *r
		}
	}
	return r1 && r3
}

func checkRunes(a, b, c rune) bool {
	if a-1 == b && b-1 == c {
		return true
	}
	return false
}

func shiftRunes(arr [3]rune, r rune) [3]rune {
	arr[2] = arr[1]
	arr[1] = arr[0]
	arr[0] = r
	return arr
}

func printPtrs(data []*rune) {
	for _, r := range data {
		if r == nil {
			continue
		}
		fmt.Printf("%c", *r)
	}
	fmt.Printf("\n")
}
