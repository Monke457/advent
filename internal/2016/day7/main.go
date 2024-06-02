package main

import (
	"advent/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2016/day7.txt")
	solveFirstProblem(data)
	solveSecondProblem(data)
}

func solveFirstProblem(data []string) {
	count := 0
loop:
	for _, row := range data {
		ins, outs := splitHypernet(row)
		for _, out := range outs {
			if containsABBA(out) {
				continue loop
			}
		}
		for _, in := range ins {
			if containsABBA(in) {
				count++
				continue loop
			}
		}
	}
	fmt.Println(count)
}

func solveSecondProblem(data []string) {
	count := 0
loop:
	for _, row := range data {
		ins, outs := splitHypernet(row)
		abas := []string{}
		for _, in := range ins {
			abas = append(abas, findABAS(in)[:]...)
		}
		for _, aba := range abas {
			for _, out := range outs {
				if strings.Contains(out, aba) {
					count++
					continue loop
				}
			}

		}
	}
	fmt.Println(count)
}

func containsABBA(str string) bool {
	for i := 0; i < len(str)-3; i++ {
		if str[i] != str[i+3] {
			continue
		}
		if str[i+1] != str[i+2] {
			continue
		}
		if str[i] == str[i+1] {
			continue
		}
		return true
	}
	return false
}

func findABAS(str string) []string {
	res := []string{}
	for i := 0; i < len(str)-2; i++ {
		if str[i] != str[i+2] {
			continue
		}
		if str[i] == str[i+1] {
			continue
		}
		res = append(res, fmt.Sprintf("%c%c%c", str[i+1], str[i], str[i+1]))
	}
	return res
}

func splitHypernet(str string) ([]string, []string) {
	ins, outs := []string{}, []string{}
	sb := strings.Builder{}
	for _, c := range str {
		if c == '[' {
			ins = append(ins, sb.String())
			sb.Reset()
		} else if c == ']' {
			outs = append(outs, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(c)
		}
	}
	if sb.Len() > 0 {
		ins = append(ins, sb.String())
	}
	return ins, outs
}
