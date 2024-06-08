package main

import (
	"advent/pkg/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToString("data/2016/day9.txt")
	res := findCompLen(data)
	fmt.Println("First:", res)
}

func findCompLen(data string) int {
	re := regexp.MustCompile(`\(\d*x\d*\)`)
	result := strings.Builder{}
	match := re.FindStringIndex(data)

	for match != nil {
		result.WriteString(data[:match[0]])
		
		l, c := parseCompLen(data[match[0]:match[1]])

		for range c {
			result.WriteString(data[match[1]:match[1]+l])
		}

		data = data[match[1]+l:]
		match = re.FindStringIndex(data)
	}

	result.WriteString(data)
	
	return result.Len() 
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
