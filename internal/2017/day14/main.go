package main

import (
	hash "advent/pkg/hashing"
	"advent/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	data := reader.FileToString("data/2017/day14.txt")

	count := 0
	values := hash.MakeRange(0, 256)
	for i := range 128 {
		input := hash.ToASCIIValues(fmt.Sprintf("%s-%d", data, i))
		hashed := hash.KnotHash(values, input, 64)
		dense := hash.DenseHash(hashed)

		for _, h := range dense {
			val := fmt.Sprintf("%b", h)
			count += strings.Count(val, "1")
		}
	}

	fmt.Println("First:", count) 
}
