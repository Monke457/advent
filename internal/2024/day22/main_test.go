package main

import "testing"

func TestBitwiseMultiply(t *testing.T) {
	examples := map[[2]int]int{
		{0, 0}:0,
		{1, 1}:2,
		{1, 2}:4,
		{2, 3}:16,
	}

	for in, out := range examples {
		res := bitwiseMultiply(in[0], in[1]) 
		if res != out {
			t.Errorf("Expected %d got %d\n", out, res)
		}
	}
}

func TestBitwiseDivide(t *testing.T) {
	examples := map[[2]int]int{
		{1, 2}:0,
		{2, 1}:1,
		{10, 2}:2,
		{5, 2}:1,
	}

	for in, out := range examples {
		res := bitwiseDivide(in[0], in[1]) 
		if res != out {
			t.Errorf("Expected %d got %d\n", out, res)
		}
	}
}

func TestMix(t *testing.T) {
	examples := map[[2]int]int{
		{0, 0}:0,
		{42, 15}:37,
	}

	for in, out := range examples {
		res := mix(in[0], in[1]) 
		if res != out {
			t.Errorf("Expected %d got %d\n", out, res)
		}
	}
}

func TestPrune(t *testing.T) {
	examples := map[[2]int]int{
		{5, 2}:1,
		{100000000,16777216}:16113920,
		{54,25}:4,
	}

	for in, out := range examples {
		res := prune(in[0], in[1]) 
		if res != out {
			t.Errorf("Expected %d got %d\n", out, res)
		}
	}
}
