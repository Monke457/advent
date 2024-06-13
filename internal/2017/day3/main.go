package main

import (
	"fmt"
	"math"
)

func main() {
	data := 312051
	rt, sq := findSquare(data)
	result := findDistance(data, rt, sq)
	fmt.Println("First:", result)
}

func findDistance(n, rt, sq int) int {
	l := max(rt-1, 1)
	c := (rt-2) * (rt-2) + 1
	low, high := sq, 1
	for low > n {
		low = max(low-l, c) 
	}
	for high < n {
		high += l 
	}
	mid := high - (l >> 1)
	diff := int(math.Abs(float64(mid - n)))
	return (l >> 1) + diff 
}

func findSquare(n int) (int, int) {
	res := 1
	sq := 0
	for sq < n {
		sq = res * res 
		res += 2
	}
	res -= 2
	return res, sq
}
