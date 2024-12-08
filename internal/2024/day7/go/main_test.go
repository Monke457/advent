package main_test

import (
	"testing"
	m "advent/internal/2024/day7/go"
)

func TestValidEquation(t *testing.T) {
	cases := [6][]int{}
	results := [6]int{}
	expected := [6]bool{}

	cases[0] = []int{5, 6, 2, 10, 23}
	results[0] = (5 * 6 + 2) * 10 + 23
	expected[0] = true 

	cases[1] = []int{6, 9, 1, 222, 345}
	results[1] = (6 + 9 + 1) * 222 * 345
	expected[1] = true 

	cases[2] = []int{234, 5, 23, 77, 94}
	results[2] = 234 * 5 + 23 + 709 * 14
	expected[2] = false 

	cases[3] = []int{9, 2, 23, 45}
	results[3] = 2349 * 14 * 347
	expected[3] = false 

	cases[4] = []int{72, 1, 6, 2, 149, 2, 854, 36}
	results[4] = 72 * 1 * 6 * 2 * 149 * 2 * 854 * 36
	expected[4] = true 

	cases[5] = []int{1, 0, 2}
	results[5] = 2
	expected[5] = true

	for i := range 6 {
		res := m.ValidEquation(results[i], cases[i]) 
		if res != expected[i] {
			t.Errorf("validEquation(%d, %d) = %v; want %v", results[i], cases[i], res, expected[i])
		}
	}
}
