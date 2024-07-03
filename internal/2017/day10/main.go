package main

import (
	hash "advent/internal/pkg/hashing"
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	lengths := reader.FileToIntArrayByComma("data/2017/day10.txt")
	solveFirst(lengths)

	binLengths := reader.FileToASCIIValues("data/2017/day10.txt")
	binLengths = append(binLengths, 17, 31, 73, 47, 23)
	solveSecond(binLengths)
}

func solveFirst(lengths []int) {
	values := hash.MakeRange(0, 256)
	hashed := hash.KnotHash(values, lengths, 1)
	fmt.Println("First:", hashed[0] * hashed[1])
}

func solveSecond(lengths []int) {
	values := hash.MakeRange(0, 256)

	rounds := 64
	hashed := hash.KnotHash(values, lengths, rounds)

	dh := hash.DenseHash(hashed)
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
