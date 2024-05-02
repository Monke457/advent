package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

var seeds []int
var seedRanges map[int][]int
var soil, fert, water, light, temp, hum, loc map[int][]int
var maps []*map[int][]int = []*map[int][]int{
	&soil, &fert, &water, &light, &temp, &hum, &loc,
}

func main() {
	//fmt.Println(solveFirstProblem())
	fmt.Println(solveSecondProblem())
}

func solveFirstProblem() int {
	lines := reader.FileToArray("data/day5.txt")

	seeds, soil, fert, water, light, temp, hum, loc = normalize(lines)

	location := -1
	for _, seed := range seeds {
		l := convert(seed, 0)
		if location == -1 {
			location = l
		} else {
			location = min(location, l)
		}
	}
	return location
}

func solveSecondProblem() int {
	lines := reader.FileToArray("data/day5.txt")

	seeds, soil, fert, water, light, temp, hum, loc = normalize(lines)

	seedRanges = getRanges()
	fmt.Println(seedRanges)
	fmt.Println(soil)

	location := -1
	sp := 0
	for _, r := range seedRanges {
		fmt.Println("range", r)
		for i := r[0]; i < r[1]; i++ {
			l := convert(i, 0)
			if location == -1 {
				location = l
			} else {
				location = min(location, l)
			}
			i += soil[sp][2]
		}
	}
	return location
}

func convert(n, m int) int {
	if m >= len(maps) {
		fmt.Println("converted", n)
		return n
	}
	fmt.Println("converting", n)
	for _, vals := range *maps[m] {
		if vals[1] <= n && vals[1]+vals[2] >= n {
			n += vals[0] - vals[1]
			m++
			return convert(n, m)
		}
	}
	m++
	return convert(n, m)
}

func normalize(lines []string) (
	[]int,
	map[int][]int,
	map[int][]int,
	map[int][]int,
	map[int][]int,
	map[int][]int,
	map[int][]int,
	map[int][]int,
) {

	var toSoil map[int][]int = map[int][]int{}
	var toFertilizer map[int][]int = map[int][]int{}
	var toWater map[int][]int = map[int][]int{}
	var toLight map[int][]int = map[int][]int{}
	var toTemp map[int][]int = map[int][]int{}
	var toHumidity map[int][]int = map[int][]int{}
	var toLocation map[int][]int = map[int][]int{}

	var s []int = []int{}
	flag := ""
	n := 0
	for i, l := range lines {
		if i == 0 {
			s = getSeedValues(l)
			continue
		}
		if l == "" {
			flag = ""
			n = 0
			continue
		}

		switch flag {
		case "":
			flag = getFlag(l)
		case "to-soil":
			toSoil[n] = addToMap(toSoil[n], l)
			n++
		case "to-fertilizer":
			toFertilizer[n] = addToMap(toFertilizer[n], l)
			n++
		case "to-water":
			toWater[n] = addToMap(toWater[n], l)
			n++
		case "to-light":
			toLight[n] = addToMap(toLight[n], l)
			n++
		case "to-temp":
			toTemp[n] = addToMap(toTemp[n], l)
			n++
		case "to-humidity":
			toHumidity[n] = addToMap(toHumidity[n], l)
			n++
		case "to-location":
			toLocation[n] = addToMap(toLocation[n], l)
			n++
		}
	}

	return s, toSoil, toFertilizer, toWater,
		toLight, toTemp, toHumidity, toLocation
}

func getSeedValues(l string) []int {
	s := []int{}
	l = strings.Split(l, ":")[1]
	for _, val := range strings.Split(l, " ") {
		if val == "" {
			continue
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		s = append(s, num)
	}
	return s
}

func getRanges() map[int][]int {
	ranges := map[int][]int{}

	n := 0
	for i := 0; i < len(seeds)-1; i += 2 {
		ranges[n] = []int{seeds[i], seeds[i] + seeds[i+1] - 1}
		n++
	}
	return ranges
}

func addToMap(m []int, l string) []int {
	vals := strings.Split(l, " ")
	for _, val := range vals {
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		m = append(m, v)
	}
	return m
}

func getFlag(l string) string {
	if strings.Contains(l, "to-soil") {
		return "to-soil"
	}
	if strings.Contains(l, "to-fertilizer") {
		return "to-fertilizer"
	}
	if strings.Contains(l, "to-water") {
		return "to-water"
	}
	if strings.Contains(l, "to-light") {
		return "to-light"
	}
	if strings.Contains(l, "to-temp") {
		return "to-temp"
	}
	if strings.Contains(l, "to-humidity") {
		return "to-humidity"
	}
	if strings.Contains(l, "to-location") {
		return "to-location"
	}
	return ""
}
