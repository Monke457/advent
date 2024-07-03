package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

const mybag = "shiny gold"

var bags = make(map[string]map[string]int)

func main() {
	data := reader.FileToArray("data/2020/day7.txt")

	for _, line := range data {
		id, content := parseLine(line)
		bags[id] = content
	}

	count := 0
	for id := range bags {
		if checkValid(id) {
			count++
		}
	}
	fmt.Println("First:", count)

	fmt.Println("Second:", innerBagCount(mybag, 1))
}

func innerBagCount(id string, sum int) int {
	fmt.Println("checking bag", id)
	if len(bags[id]) == 0 {
		return 1
	}
	for k := range bags[id] {
		sum += bags[id][k] * innerBagCount(k, sum) 
		fmt.Println(k, bags[id][k], len(bags[k]), sum)
	}
	return sum
}

func checkValid(id string) bool {
	if id == mybag {
		return false
	}
	if _, ok := bags[id][mybag]; ok {
		return true
	}
	for k := range bags[id] {
		if checkValid(k) {
			return true
		}
	}
	return false
} 

func parseLine(line string) (string, map[string]int) {
	bagId, rest, _ := strings.Cut(line, " bags contain ")

	if rest[:3] == "no " {
		return bagId, nil
	}

	result := map[string]int{}
	bags := strings.Split(rest, ", ")
	for _, bag := range bags {
		count, err := strconv.Atoi(string(bag[0]))
		if err != nil {
			panic(err)
		}
		id := strings.Split(bag, " ")[1:3]
		result[strings.Join(id, " ")] = count
	}

	return bagId, result 
}
