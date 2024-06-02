package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToString("data/2015/day12.txt")
	fmt.Println("First problem:", solveFirstProblem(data))
	fmt.Println("Second problem:", solveSecondProblem(data))
}

func solveFirstProblem(data string) int {
	return sumInts(data)
}

func solveSecondProblem(data string) int {
	return sumWithoutRed(data)
}

func sumWithoutRed(data string) int {
	var total int
	status := map[int]bool{}
	idx := 0
	sb := strings.Builder{}
	rb := map[int]*strings.Builder{}
	nums := []int{}
	for _, b := range data {
		if b == '{' {
			status[idx] = true
			idx++
		}
		if b == '[' {
			status[idx] = false
			idx++
		}
		if b == ']' {
			idx--
		}
		if _, ok := rb[idx]; !ok {
			rb[idx] = &strings.Builder{}
		}
		if b == '}' {
			if !strings.Contains(rb[idx].String(), "red") {
				for _, i := range nums {
					total += i
				}
			}
			nums = []int{}
			idx--
		}
		if b == '-' {
			sb.WriteRune(b)
		} else if b > 47 && b < 58 {
			sb.WriteRune(b)
		} else {
			num := 0
			if sb.Len() > 0 {
				val, err := strconv.Atoi(sb.String())
				if err != nil {
					panic(err)
				}
				num = val
				sb.Reset()
			}
			if status[idx] {
				rb[idx].WriteRune(b)
				nums = append(nums, num)
			} else {
				total += num
			}
		}
	}
	return total
}

func sumInts(data string) int {
	var total int
	sb := strings.Builder{}
	for _, b := range data {
		if b == '-' {
			sb.WriteRune(b)
		} else if b > 47 && b < 58 {
			sb.WriteRune(b)
		} else if sb.Len() > 0 {
			val, err := strconv.Atoi(sb.String())
			if err != nil {
				panic(err)
			}
			total += val
			sb.Reset()
		}
	}
	return total
}
