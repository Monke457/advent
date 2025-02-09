package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type digit string

const (
	ONE digit = "cf"
	TWO = "acdeg"
	THREE = "acdfg"
	FOUR = "bcdf"
	FIVE = "abdfg"
	SIX = "abdefg"
	SEVEN = "acf"
	EIGHT = "abcdefg"
	NINE = "abcdfg"
	ZERO = "abcefg"
)

func main() {
	data := reader.FileToArray("data/2021/day8.txt")

	sum := 0
	for _, line := range data {
		sum += decodeOutput(line)
	}

	fmt.Println("sum of all outputs:", sum)
}

func decodeOutput(raw string) int {
	digitmap := map[digit]string{}
	values, output := parseValues(raw)

	for _, value := range values {
		switch len(value) {
		case 2:
			digitmap[ONE] = value

		case 3:
			digitmap[SEVEN] = value

		case 4:
			digitmap[FOUR] = value

		case 5:
			if containsAll(value, digitmap[SEVEN]) {
				digitmap[THREE] = value
			}
		case 6:
			if containsAll(value, digitmap[FOUR]) {
				digitmap[NINE] = value
			} else if containsAll(value, digitmap[SEVEN]) {
				digitmap[ZERO] = value
			} else {
				digitmap[SIX] = value
			}
		case 7:
			digitmap[EIGHT] = value
		}
	}
	runemap := remapRunes(digitmap)
	digitmap[TWO] = digit(TWO).convertDigit(runemap)
	digitmap[FIVE] = digit(FIVE).convertDigit(runemap)

	return decode(digitmap, output)
}

func decode(digitmap map[digit]string, output []string) int {
	resstr := ""
	for _, o := range output {
		resstr += findDigitAsString(digitmap, o)
	}
	res, _ := strconv.Atoi(resstr)
	return res
}

func findDigitAsString(digitmap map[digit]string, original string) string {
	if equalRunes(digitmap[ONE], original) {
		return "1"
	}
	if equalRunes(digitmap[TWO], original) {
		return "2"
	}
	if equalRunes(digitmap[THREE], original) {
		return "3"
	}
	if equalRunes(digitmap[FOUR], original) {
		return "4"
	}
	if equalRunes(digitmap[FIVE], original) {
		return "5"
	}
	if equalRunes(digitmap[SIX], original) {
		return "6"
	}
	if equalRunes(digitmap[SEVEN], original) {
		return "7"
	}
	if equalRunes(digitmap[EIGHT], original) {
		return "8"
	}
	if equalRunes(digitmap[NINE], original) {
		return "9"
	}
	if equalRunes(digitmap[ZERO], original) {
		return "0"
	}
	return ""
}

func equalRunes(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !strings.ContainsRune(b, rune(a[i])) {
			return false
		}
	}
	return true
}

func (d digit) convertDigit(runemap map[rune]rune) string {
	res := ""
	for _, r := range d {
		res += string(runemap[r])
	}
	return res
}

func remapRunes(digits map[digit]string) map[rune]rune {
	result := map[rune]rune{}
	if r, l := mask(digits[SEVEN], digits[ONE]); l == 1 {
		result['a'] = rune(r[0])
	}
	if r, l := mask(digits[FOUR], digits[THREE]); l == 1 {
		result['b'] = rune(r[0])
	}
	if r, l := mask(digits[ONE], digits[SIX]); l == 1 {
		result['c'] = rune(r[0])
	}
	if r, l := mask(digits[EIGHT], digits[ZERO]); l == 1 {
		result['d'] = rune(r[0])
	}
	if r, l := mask(digits[EIGHT], digits[NINE]); l == 1 {
		result['e'] = rune(r[0])
	}
	if r, l := mask(digits[ONE], string(result['c'])); l == 1 {
		result['f'] = rune(r[0])
	}
	if r, l := mask(digits[NINE], digits[FOUR], digits[SEVEN]); l == 1 {
		result['g'] = rune(r[0])
	}
	return result
}

func mask(val string, masks... string) (string, int) {
	r := val
	for _, mask := range masks {
		r = removeAll(r, mask)
	}
	return r, len(r)
}

func removeAll(a, b string) string {
	rest := ""
	for _, r := range a {
		if !strings.ContainsRune(b, r) {
			rest += string(r)
		}
	}
	return rest 
}

func containsAll(a, b string) bool {
	for _, r := range b {
		if !strings.ContainsRune(a, r) {
			return false
		}
	}
	return true
}

func parseValues(line string) ([]string, []string) {
	s, o, _ := strings.Cut(line, " | ")
	values := strings.Split(s, " ")
	output := strings.Split(o, " ")
	values = append(values, output...)
	slices.SortFunc(values, func(a, b string) int {
		if len(a) < len(b) {
			return -1
		}
		return 1
	})
	return values, output
}
