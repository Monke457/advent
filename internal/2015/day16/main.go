package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

var tickertape string = `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`

type info map[string]int

func main() {
	data := reader.FileToArray("data/2015/day16.txt")

	legend := info{}
	for _, line := range strings.Split(tickertape, "\n") {
		key, value := parseData(line)
		legend[key] = value
	}

	sues := map[int]info{}

	for _, line := range data {
		namestr, reststr, _ := strings.Cut(line, ": ")
		namestr = namestr[4:]

		name, _ := strconv.Atoi(namestr)
		sues[name] = info{}

		rest := strings.Split(reststr, ", ")
		for _, data := range rest {
			key, value := parseData(data)
			sues[name][key] = value
		}
	}

	correct := 0
	loop:
	for sue, data := range sues {
		for key, value := range legend {
			val, ok := data[key]
			if key == "cats" || key == "trees" {
				if ok && val <= value {
					continue loop
				}
			} else if key == "pomeranians" || key == "goldfish" {
				if ok && val >= value {
					continue loop
				}
			} else if ok && val != value {
				continue loop
			}
		}
		correct = sue
		break
	}

	fmt.Println(legend)
	fmt.Println("corrent sue:", correct, sues[correct])
}

func parseData(line string) (string, int) {
	key, vstr, _ := strings.Cut(line, ": ")
	value, _ := strconv.Atoi(vstr)
	return key, value
}
