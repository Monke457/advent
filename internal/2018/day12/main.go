package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2018/day12.txt")

	var sum, shift, offset int

	generations := 50000000000
	plants := data[0][15:]
	prev := plants
	rules := []string{}

	for _, row := range data[2:] {
		if row[9] == '#' {
			rules = append(rules, row[:5])
		}
	}

	for i := range generations {
		plants, shift = spread(plants, rules)
		if plants == prev {
			offset += shift * (generations - i)
			break
		}
		offset += shift
		prev = plants
	}

	sum = sumPots(plants, offset)
	fmt.Println("\nSum of pot numbers after", generations, "generations:", sum)
}

func sumPots(plants string, offset int) int {
	sum := 0
	for i, plant := range plants {
		if plant == '#' {
			sum += i-offset
		}
	}
	return sum
}

func spread(plants string, rules []string) (string, int) {
	var result string
	var shift int

	plants = "..." + plants + "..."

	loop:
	for i := 2; i < len(plants)-2; i++ {
		for _, rule := range rules {
			if plants[i-2:i+3] == rule {
				if i == 2 {
					shift++
				}
				result += "#"
				continue loop
			}
		}
		if i == 2 {
			continue
		}
		result += "."
	}
	for result[0] == '.' {
		result = result[1:]
		shift--
	}
	result = strings.TrimRight(result, ".")
	return result, shift 
}


