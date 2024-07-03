package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type passcode struct {
	policy [2]int
	char byte 
	words map[byte][]int
}

func main() {
	data := reader.FileToArray("data/2020/day2.txt")

	count1 := 0
	count2 := 0
	for _, line := range data {
		pass := parseLine(line)
		if pass.isValid1() {
			count1++
		}
		if pass.isValid2() {
			count2++
		}
	}

	fmt.Println("First:", count1)
	fmt.Println("Second:", count2)
}

func (p *passcode) isValid1() bool {
	count := len(p.words[p.char])
	return count >= p.policy[0] && count <= p.policy[1]
}

func (p *passcode) isValid2() bool {
	var count int 
	pos := p.words[p.char]
	if slices.Contains(pos, p.policy[0]) {
		count++
	}
	if slices.Contains(pos, p.policy[1]) {
		count++
	}
	return count == 1
}

func parseLine(line string) passcode {
	low, rest, _ := strings.Cut(line, "-")
	lowNum, err := strconv.Atoi(low)
	if err != nil {
		panic(err)
	}

	high, rest, _ := strings.Cut(rest, " ")
	highNum, err := strconv.Atoi(high)
	if err != nil {
		panic(err)
	}

	char := rest[0]
	_, word, _ := strings.Cut(rest, " ")

	words := make(map[byte][]int)
	for i, b := range word {
		words[byte(b)] = append(words[byte(b)], i + 1) 
	}

	return passcode{ 
		policy: [2]int{ lowNum, highNum }, 
		char: char, 
		words: words,
	}
}
