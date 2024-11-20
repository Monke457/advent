package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToString("data/2019/day4.txt")

	startStr, endStr, _ := strings.Cut(data, "-")

	start, err := strconv.Atoi(startStr)
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(endStr)
	if err != nil {
		panic(err)
	}

	found := []int{}
	for i := start; i <= end; i++ {
		if !hasDouble(i) {
			continue
		}
		if !onlyInc(i) {
			continue
		}
		found = append(found, i)
	}

	fmt.Println("Part 1:", len(found))

	valid := 0
	for _, val := range found {
		if !validDouble(val) {
			continue
		}
		valid++
	}
	fmt.Println("Part 2:", valid)
}

func validDouble(value int) bool {
	for {
		last := value % 10
		value = value / 10
		if value < 1 {
			break
		}
		if value % 10 == last {
			if value / 10 % 10 != last {
				return true
			}
			for {
				value = value / 10
				if value % 10 != last {
					break
				}
			}
		}
	}
	return false
} 

func hasDouble(value int) bool {
	for {
		last := value % 10
		value = value / 10
		if value < 1 {
			break
		}
		if value % 10 == last {
			return true
		}
	}
	return false
}

func onlyInc(value int) bool {
	last := 0
	for {
		last = value % 10
		value = value / 10
		if value < 1 {
			break
		}
		if value % 10 > last {
			return false
		}
	}
	return value < last 
}
