package main

import (
	"advent/internal/pkg/reader"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToString("data/2016/day5.txt")

	//solveFirstProblem(data)
	solveSecondProblem(data)
}

func solveFirstProblem(data string) {
	fmt.Println("Processing first problem...")
	result := strings.Builder{}
	i := 0
	for result.Len() < 8 {
		plain := fmt.Sprintf("%s%d", data, i)
		hash := md5.Sum([]byte(plain))
		hex := fmt.Sprintf("%x", hash)
		if hex[:5] == "00000" {
			result.WriteByte(hex[5])
		}
		i++
	}
	fmt.Println(result.String())
}

func solveSecondProblem(data string) {
	fmt.Println("Processing second problem...")

	result := [8]*byte{}
	count := 0
	i := 0
	for count < 8 {
		plain := fmt.Sprintf("%s%d", data, i)
		hash := md5.Sum([]byte(plain))
		hex := fmt.Sprintf("%x", hash)
		if hex[:5] == "00000" {
			if hex[5] > '7' {
				i++
				continue
			}
			idx, err := strconv.Atoi(string(hex[5]))
			if err != nil {
				panic(err)
			}
			if result[idx] != nil {
				i++
				continue
			}
			fmt.Printf("found %s, %s\n", plain, hex)
			b := hex[6]
			result[idx] = &b
			count++
		}
		i++
	}
	for _, c := range result {
		fmt.Printf("%c", *c)
	}
	fmt.Println()
}
