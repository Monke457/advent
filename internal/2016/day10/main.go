package main

import (
	"advent/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2016/day10.txt")

	bots, rest := initialBots(data)

	res1 := swapperoo(rest, bots)
	fmt.Println("First:", res1)

	res2 := swapperoo2(rest, bots, map[int]int{})
	fmt.Println("Second:", res2)
}

func swapperoo2(data []string, bots map[int][]int, outputs map[int]int) int {
	rest := []string{}
	for i, line := range data {
		bot, lo, hi, out :=  parseCmd(line)
		if len(bots[bot]) == 2 {
			l := slices.Min(bots[bot])
			h := slices.Max(bots[bot])
			if out[0] != nil {
				outputs[*out[0]] = l
			}
			if out[1] != nil {
				outputs[*out[1]] = h
			}
			if lo != nil {
				bots[*lo] = append(bots[*lo], l)
			}
			if hi != nil {
				bots[*hi] = append(bots[*hi], h)
			}
			bots[bot] = []int{} 
		} else {
			rest = append(rest, data[i])
		}
		one, ok1 := outputs[0];
		two, ok2 := outputs[1];
		three, ok3 := outputs[2];
		if ok1 && ok2 && ok3 {
			return one * two * three
		}
	}
	return swapperoo2(rest, bots, outputs)
}

func swapperoo(data []string, bots map[int][]int) int {
	rest := []string{}
	for i, line := range data {
		bot, lo, hi, _ :=  parseCmd(line)
		if len(bots[bot]) == 2 {
			l := slices.Min(bots[bot])
			h := slices.Max(bots[bot])
			if l == 17 && h == 61 {
				return bot
			}
			if lo != nil {
				bots[*lo] = append(bots[*lo], l)
			}
			if hi != nil {
				bots[*hi] = append(bots[*hi], h)
			}
			bots[bot] = []int{} 
		} else {
			rest = append(rest, data[i])
		}
	}
	return swapperoo(rest, bots)
}

func parseCmd(cmd string) (bot int, lo, hi *int, out [2]*int) {
	parts := strings.Split(cmd, " ")
	bot, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	l, err := strconv.Atoi(parts[6])
	if err != nil {
		panic(err)
	}
	low := &l
	if parts[5] == "output" {
		low = nil
		out[0] = &l
	}
	h, err := strconv.Atoi(parts[11])
	if err != nil {
		panic(err)
	}
	high := &h
	if parts[10] == "output" {
		high = nil
		out[1] = &h
	}
	return bot, low, high, out
}

func initialBots(data []string) (map[int][]int, []string) {
	bots := map[int][]int{}
	rest := []string{}

	for _, line := range data {
		if strings.Contains(line, "goes") {
			parts := strings.Split(line, " ")
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			bot, err := strconv.Atoi(parts[5])
			if err != nil {
				panic(err)
			}
			bots[bot] = append(bots[bot], val)
		} else {
			rest = append(rest, line)
		}
	}

	return bots, rest
}
