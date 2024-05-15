package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type hashMap struct {
	m map[int][]*lense
}

type lense struct {
	label string
	focus int
}

func main() {
	content := reader.FileToString("data/day15.txt")

	data := strings.Split(content, ",")

	fmt.Println("First:", solveFirstProblem(data))
	fmt.Println("Second:", solveSecondProblem(data))
}

func solveFirstProblem(data []string) int {
	sum := 0
	for _, d := range data {
		sum += hashValue(d)
	}
	return sum
}

func solveSecondProblem(data []string) int {
	hm := hashMap{m: map[int][]*lense{}}
	for _, d := range data {
		box := hashLabel(d)
		if strings.Contains(d, "=") {
			hm.Add(box, d)
		} else if strings.Contains(d, "-") {
			hm.Remove(box, d)
		}
	}
	return hm.CalculatePower()
}

func (hm hashMap) Remove(box int, label string) hashMap {
	label = label[:len(label)-1]
	lnew := []*lense{}
	for _, l := range hm.m[box] {
		if l.label == label {
			continue
		}
		lnew = append(lnew, l)
	}
	hm.m[box] = lnew
	return hm
}

func (hm hashMap) Add(box int, label string) hashMap {
	l, f, _ := strings.Cut(label, "=")
	focus, err := strconv.Atoi(f)
	if err != nil {
		panic(err)
	}
	for _, lense := range hm.m[box] {
		if lense.label == l {
			lense.focus = focus
			return hm
		}
	}
	hm.m[box] = append(hm.m[box], &lense{l, focus})
	return hm
}

func (hm hashMap) CalculatePower() (res int) {
	for box, lenses := range hm.m {
		for slot, lense := range lenses {
			res += (box + 1) * (slot + 1) * lense.focus
		}
	}
	return
}

func hashValue(str string) (res int) {
	for _, s := range str {
		res += int(s)
		res *= 17
		res %= 256
	}
	return
}

func hashLabel(l string) (res int) {
	for _, s := range l {
		if s == '=' || s == '-' {
			return
		}
		res += int(s)
		res *= 17
		res %= 256
	}
	return
}
