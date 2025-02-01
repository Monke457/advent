package main_test

import (
	"advent/internal/2018/day9"
	"slices"
	"testing"
)

type data struct {
	input []int
	expected []int
	index int
	value int
}

func TestInsert(t *testing.T) {
	case1 := data {
		input: []int{0, 2, 3, 4, 5, 6},
		expected: []int{0, 2, 3, 99, 4, 5, 6},
		index: 3,
		value: 99,
	}
	case2 := data {
		input: []int{0, 2, 3, 4, 5, 6},
		expected: []int{0, 2, 3, 4, 5, 6, 2},
		index: 6,
		value: 2,
	}
	case3 := data {
		input: []int{0, 2, 3, 4, 5, 6},
		expected: []int{0, 5, 2, 3, 4, 5, 6},
		index: 0,
		value: 5,
	}

	cases := []data{case1, case2, case3}
	for _, c := range cases {
		output := main.Insert(c.input, c.index, c.value)
		if !slices.Equal(output, c.expected) {
			t.Fatalf("Expected: %d Actual: %d", c.expected, output)
		}
	}
}


