package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToString("data/2016/day9.txt")
	res1 := findCompLen(data)
	fmt.Println("First:", res1)
	res2 := findCompLenRec(data)
	fmt.Println("Second:", res2)
}

func findCompLenRec(data string) int {
	re := regexp.MustCompile(`\(\d*x\d*\)`)
	match := re.FindStringIndex(data)

	if match == nil {
		return len(data)
	}

	r := match[0]
	l, c := parseCompLen(data[match[0]:match[1]])
	r += c * findCompLenRec(data[match[1]:match[1]+l])
	return r + findCompLenRec(data[match[1]+l:])
}

func findCompLen(data string) int {
	re := regexp.MustCompile(`\(\d*x\d*\)`)
	match := re.FindStringIndex(data)

	r := 0
	for match != nil {
		r += match[0]
		l, c := parseCompLen(data[match[0]:match[1]])
		r += c * l
		data = data[match[1]+l:]
		match = re.FindStringIndex(data)
	}

	r += len(data)
	return r
}

func parseCompLen(str string) (int, int) {
	length, count, _ := strings.Cut(strings.Trim(str, "()"), "x")

	l, err := strconv.Atoi(length)
	if err != nil {
		panic(err)
	}

	c, err := strconv.Atoi(count)
	if err != nil {
		panic(err)
	}

	return l, c
}
