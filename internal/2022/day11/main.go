package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type monkey struct {
	items []int
	op func(int)int
	test func(int)int
}

func main() {
	data := reader.FileToArray("data/2022/day11.txt")

	worrymod := 1
	monkeys := []monkey{}
	for i := 0; i < len(data); i += 7 {
		monkey, mod := parseMonkey(data[i:i+6])
		monkeys = append(monkeys, monkey)
		worrymod *= mod
	}

	inspections := make([]int, len(monkeys))
	rounds := 10000
	for range rounds {
		for idx := range monkeys {
			for _, item := range monkeys[idx].items {
				inspections[idx]++
				res := monkeys[idx].op(item) % worrymod
				throw := monkeys[idx].test(res)
				monkeys[throw].items = append(monkeys[throw].items, res)
			}
			monkeys[idx].items = []int{}
		}
	}

	slices.SortFunc(inspections, func(a, b int)int {
		if a < b {
			return 1
		}
		return -1
	})
	fmt.Println(inspections)

	product := inspections[0] * inspections[1]
	fmt.Printf("Monkey business: %d\n", product)
}

func (m monkey) print(i int) {
	fmt.Printf("Monkey %d: %d\n", i, m.items)
}

func parseMonkey(data []string) (monkey, int) {
	items := parseItems(data[1])
	op := parseOperation(data[2])
	test, mod := parseTest(data[3:])

	return monkey{
		items:items, 
		op:op, 
		test:test,
	}, mod
}

func parseItems(data string) []int {
	items := []int{}
	_, itemsstr, _ := strings.Cut(data, ": ")
	for _, itemstr := range strings.Split(itemsstr, ", ") {
		item, _ := strconv.Atoi(itemstr)
		items = append(items, item)
	}
	return items
}

func parseOperation(data string) func(int)int {
	_, opstr, _ := strings.Cut(data, "= old ")
	op := opstr[0]
	valstr := opstr[2:]
	val, err := strconv.Atoi(valstr)
	if op == '*' {
		return func(old int) int {
			if err != nil {
				return old * old
			}
			return old * val
		}
	} else {
		return func(old int) int {
			if err != nil {
				return old + old
			}
			return old + val
		}
	}
}

func parseTest(data []string) (func(int)int, int) {
	divstr := data[0][21:]
	div, _ := strconv.Atoi(divstr)
	tstr := data[1][29:]
	tMonkey, _ := strconv.Atoi(tstr)
	fstr := data[2][30:]
	fMonkey, _ := strconv.Atoi(fstr)

	return func(val int) int {
		if val % div == 0 {
			return tMonkey
		}
		return fMonkey
	}, div
}
