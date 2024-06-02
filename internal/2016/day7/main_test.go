package main

import (
	"strings"
	"testing"
)

func TestContainsABBA(t *testing.T) {
	contains := []string{"feihfwffw", "effesfijhe", "efflgighdkjggj"}
	none := []string{"fiehfwlkeenola", "efohqpeivnoflai", "evowfnwld"}

	for _, str := range contains {
		if !containsABBA(str) {
			t.Error("Pattern not found:", str)
		}
	}

	for _, str := range none {
		if containsABBA(str) {
			t.Error("Pattern found where there should be none:", str)
		}
	}
}

type testdata struct {
	in  string
	out string
	res bool
}

func TestFindABAS(t *testing.T) {
	data := []testdata{
		{in: "abaijfe", out: "babk", res: true},
		{in: "xyxxyx", out: "xyx", res: false},
		{in: "xyhhfidihe", out: "kdidglljil", res: true},
		{in: "xyefljeoe", out: "fjeojf", res: false},
	}

	for _, d := range data {
		abas := findABAS(d.in)

		if len(abas) == 0 {
			t.Error("no patterns found", d.in)

		} else {
			res := strings.Contains(d.out, abas[0])
			if res != d.res {
				t.Error("wrong output, want", d.res, "got", res)
			}
		}
	}
}
