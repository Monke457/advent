package main

import (
	"advent/internal/pkg/reader"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	data := reader.FileToString("data/2015/day4.txt")
	fmt.Println(data)

	i := 1
	var first int
	var second int
	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", data, i)))
		res := hex.EncodeToString(hash[:])
		if res[:5] == "00000" {
			if res[:6] == "000000" {
				second = i
				break
			}
			if first == 0 {
				first = i
			}
		}
		i++
	}
	fmt.Println("First problem:", first)
	fmt.Println("Second problem:", second)
}
