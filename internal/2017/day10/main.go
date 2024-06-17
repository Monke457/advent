package main

import (
	"advent/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	lengths := reader.FileToIntArrayByComma("data/2017/day10.txt")
	solveFirst(lengths)

	binLengths := reader.FileToBinaryValue("data/2017/day10.txt")
	binLengths = append(binLengths, 17, 31, 73, 47, 23)
	solveSecond(binLengths)
}

func solveFirst(lengths []int) {
	values := makeRange(0, 256)
	hashRound(values, lengths, 0, 0)
	fmt.Println("First:", values[0] * values[1])
}

func solveSecond(lengths []int) {
	values := makeRange(0, 256)
	var pos, skip int

	for range 64 {
		pos, skip = hashRound(values, lengths, pos, skip)
	}

	dh := denseHash(values)
	res := strings.Builder{}
	for _, val := range dh {
		hx := fmt.Sprintf("%x", val)
		if len(hx) == 1 {
			res.WriteString("0")
		}
		res.WriteString(hx)
	}
	fmt.Println("Second:", res.String())
}

func denseHash(values []int) []int {
	res := []int{}
	for i := 0; i < len(values); i += 16 {
		var hash int
		for _, v := range values[i:i+16] {
			hash ^= v
		}
		res = append(res, hash)
	}
	return res
}

func hashRound(values, lengths []int, pos, skip int) (int, int) {
	for _, l := range lengths {
		if l > 1 {
			rev := []int{} 

			for i := 0; i < l; i++ {
				n := (pos + i) % len(values)
				rev = append(rev, values[n])
			}
			revPos := 0
			for i := l-1; i >= 0; i-- {
				n := (pos + i) % len(values)
				values[n] = rev[revPos]
				revPos++
			}
		}
		pos += (skip + l) % len(values)
		skip++
	}
	return pos, skip
}

func makeRange(a, b int) []int {
	res := []int{}
	for i := a; i < b; i++ {
		res = append(res, i)
	}
	return res
}
