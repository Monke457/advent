package main_test

import (
	d "advent/internal/day17"
	"testing"
)

func TestStartWalk(t *testing.T) {
	cells := [][]int{
		{1, 5, 4},
		{1, 5, 1},
		{1, 2, 5},
		{1, 1, 1},
	}
	blocks := d.NewBlocks(cells)
	blocks.Print()

	res := blocks.FindRoute()

	if res != 6 {
		t.Errorf("StartWalk = %d; want 6", res)
	}
}
