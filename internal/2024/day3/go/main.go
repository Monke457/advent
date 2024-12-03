package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToString("data/2024/day3.txt")

	first := multiplyMuls(data)

	fmt.Println("Sum of multiples:", first)

	second := 0
	for {
		idx := strings.Index(data, "don't()")
		if idx == -1 {
			break
		}
		second += multiplyMuls(data[:idx])
		data = data[idx:]
		idx = strings.Index(data, "do()")
		if idx == -1 {
			break
		}
		data = data[idx:]
	}
	fmt.Println("Sum of multiples with do and dont:", second)
}

func multiplyMuls(data string) int {
	re := regexp.MustCompile(`mul\(\d*,\d*\)`)
	sum := 0 
	for _, match := range re.FindAll([]byte(data), -1) {
		nums := match[4:len(match)-1]
		x, y, _ := strings.Cut(string(nums), ",")
		xi, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		yi, err := strconv.Atoi(y)
		if err != nil {
			panic(err)
		}
		sum += xi * yi
	}
	return sum
}
