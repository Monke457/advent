package main

import (
	"advent/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type config struct {
	springs []byte
	broken  []int
}

func (c config) Extend() config {
	nc := config{}
	for range 4 {
		nc.springs = append(nc.springs, c.springs[:]...)
		nc.springs = append(nc.springs, '?')
		nc.broken = append(nc.broken, c.broken[:]...)
	}
	nc.springs = append(nc.springs, c.springs[:]...)
	nc.broken = append(nc.broken, c.broken[:]...)
	return nc
}

func slicesMap[T, U any](t []T, f func(T) U) []U {
	us := make([]U, len(t))
	for i := range t {
		us[i] = f(t[i])
	}
	return us
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func mapsClear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

func main() {
	lines := reader.FileToArray("data/2023/day12.txt")
	configs := parseConfig(lines)

	fmt.Println(solveProblem(configs))
	fmt.Println(solveSecondProblem(configs))
}

func solveProblem(configs []config) int {
	sum := 0
	for _, c := range configs {
		sum += countPossible(c.springs, c.broken)
	}
	return sum
}

func solveSecondProblem(configs []config) int {
	nc := []config{}
	for _, c := range configs {
		nc = append(nc, c.Extend())
	}
	return solveProblem(nc)
}

func countPossible(s []byte, b []int) int {
	poss := 0

	cstates := map[[3]int]int{{0, 0, 0}: 1}
	nstates := map[[3]int]int{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		for state, num := range cstates {
			ci, cc, expdot := state[0], state[1], state[2]
			switch {
			case (c == '#' || c == '?') && ci < len(b) && expdot == 0:
				if c == '?' && cc == 0 {
					nstates[[3]int{ci, cc, expdot}] += num
				}
				cc++
				if cc == b[ci] {
					ci, cc, expdot = ci+1, 0, 1
				}
				nstates[[3]int{ci, cc, expdot}] += num
			case (c == '.' || c == '?') && cc == 0:
				expdot = 0
				nstates[[3]int{ci, cc, expdot}] += num
			}
		}
		cstates, nstates = nstates, cstates
		mapsClear(nstates)
	}

	for s, v := range cstates {
		if s[0] == len(b) {
			poss += v
		}
	}

	return poss
}

func parseConfig(lines []string) []config {
	configs := []config{}
	for _, l := range lines {
		springs, broken, _ := strings.Cut(l, " ")
		s := []byte(springs)

		b := slicesMap(strings.Split(broken, ","), atoi)
		configs = append(configs, config{springs: s, broken: b})
	}
	return configs
}
