package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strings"
)

func main() {
	data := reader.FileToString("data/2018/day5.txt")

	result := collapsePolymer(data)

	smallest := math.MaxInt
	for i := 0; i < 25; i++ {
		poly := removeUnits(data, i)
		l := collapsePolymer(poly)
		if l < smallest {
			smallest = l
		}
	}

	fmt.Println("length after reactions:", result)
	fmt.Println("results after removing units:", smallest)
}

func removeUnits(str string, r int) string {
	r1 := byte(r+65)
	r2 := byte(r+97)
	sb := strings.Builder{}
	for i := 0; i < len(str); i++ {
		if str[i] != r1 && str[i] != r2 {
			sb.WriteByte(str[i])
		}
	}
	return sb.String()
}

func collapsePolymer(poly string) int {
	sb := strings.Builder{}
	for {
		done := true 
		for i := 0; i < len(poly)-1; i++ {
			if math.Abs(float64(poly[i]) - float64(poly[i+1])) == 32 {
				done = false
				i++
				continue
			}
			sb.WriteByte(poly[i])
		}
		sb.WriteByte(poly[len(poly)-1])
		if done {
			break
		}
		poly = sb.String()
		sb.Reset() 
	}
	return sb.Len()
}
