package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
)

var CLOSURE = map[rune]rune{'(':')', '[':']', '{':'}', '<':'>'}
var SCORES_A = map[rune]int{')':3, ']':57, '}':1197, '>':25137}
var SCORES_B = map[rune]int{')':1, ']':2, '}':3, '>':4}

func main() {
	data := reader.FileToArray("data/2021/day10.txt")

	sum := 0
	scores := []int{}

	for _, line := range data {
		stack, err := findFirstError(line) 
		if err != nil {
			sum += SCORES_A[*err]
			continue
		} 
		if len(stack) == 0 {
			continue
		}
		score := 0
		for len(stack) > 0 {
			score *= 5
			score += SCORES_B[pop(&stack)]
		}
		scores = append(scores, score)
	}

	slices.Sort(scores)
	middle := scores[len(scores)>>1]

	fmt.Println("sum of first error in each line:", sum)
	fmt.Println("middle completion score:", middle)
}

func findFirstError(line string) ([]rune, *rune) {
	stack := []rune{}
	for _, r := range line {
		if closure, ok := CLOSURE[r]; ok {
			push(&stack, closure)
			continue
		}
		expected := pop(&stack)
		if expected != r {
			return stack, &r
		}
	}
	return stack, nil
}

func push(stack *[]rune, item rune) {
	(*stack) = append(*stack, item)
}

func pop(stack *[]rune) rune {
	idx := len(*stack)-1
	item := (*stack)[idx]
	(*stack) = (*stack)[:idx]
	return item
}
