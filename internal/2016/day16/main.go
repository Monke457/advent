package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
)

const LENGTH = 35651584

func main() {
	data := []byte{}

	for _, bit := range reader.FileToString("data/2016/day16.txt") {
		data = append(data, byte(bit - '0'))
	}

	for len(data) < LENGTH {
		data = applyDragonCurve(data)
	}
	
	checksum := generateChecksum(data[:LENGTH])
	fmt.Println("checksum", sprintCS(checksum)) 
}

func sprintCS(cs []byte) string {
	res := ""
	for _, bit := range cs {
		res += fmt.Sprintf("%b", bit)
	}
	return res
}

func generateChecksum(data []byte) []byte {
	cs := []byte{}
	for i := 0; i < len(data)-1; i+=2 {
		if data[i] == data[i+1] {
			cs = append(cs, 1)
		} else {
			cs = append(cs, 0)
		}
	}
	if len(cs) % 2 != 0 {
		return cs
	}
	return generateChecksum(cs)
}

func applyDragonCurve(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)

	slices.Reverse(b)
	for i := range b {
		if b[i] == 0 {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}

	a = append(a, 0)
	a = append(a, b...)
	return a
}
