package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
	"time"
)

func main() {
	data := reader.FileToArray("data/2017/day4.txt")

	solveFirst(data)
	solveSecond(data)
}

func solveFirst(data []string) {
	start := time.Now()
	count := 0
	for _, line := range data {
		words := strings.Split(line, " ")
		if containsDuplicate(words) {
			continue
		}
		count++
	}
	fmt.Println("First:", count, "time:", time.Now().Sub(start)) 
}

func solveSecond(data []string) {
	start := time.Now()
	count := 0
	for _, line := range data {
		words := strings.Split(line, " ")
		if rearrangable(words) {
			continue
		}
		count++
	}
	fmt.Println("Second:", count, "time:", time.Now().Sub(start)) 
}

func rearrangable(words []string) bool {
	wordMaps := []map[rune]int{}
	for i := 0; i < len(words); i++ {
		wordMap := mapWord(words[i])
		for _, m := range wordMaps {
			if canRearrange(m, wordMap) {
				return true
			}
		}
		wordMaps = append(wordMaps, wordMap)
	}

	return false 
}

func canRearrange(src, dest map[rune]int) bool {
	if len(src) != len(dest) {
		return false
	}
	for k, v := range dest {
		if src[k] != v {
			return false
		}
	}
	return true
}

func mapWord(word string) map[rune]int {
	res := map[rune]int{}
	for _, r := range word {
		res[r]++
	}
	return res
}

func containsDuplicate(words []string) bool {
	for i := 0; i < len(words); i++ {
		for j := i+1; j < len(words); j++ {
			if words[i] == words[j] {
				return true
			}
		}
	}
	return false
}
