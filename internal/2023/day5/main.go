package main

import (
	"advent/internal/pkg/reader"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type seeds struct {
	start  int
	length int
}

type ranger struct {
	src    int
	dest   int
	length int
}

type mapper struct {
	to      string
	rangers []ranger
}

var content []string
var seedVals []string
var maps map[string]mapper

func main() {
	content = reader.FileToArray("data/2023/day5.txt")

	seedVals = strings.SplitAfter(content[0], ":")
	seedVals = strings.Split(seedVals[1], " ")
	seedVals = cleanArray(seedVals)

	maps = parseMaps()
	maps = parseNegativeRange()

	fmt.Println(solveFirstProblem())
	fmt.Println(solveSecondProblem())
}

func solveFirstProblem() int {
	seeds := parseSeeds(seedVals)

	location := math.Inf(1)
	for _, s := range seeds {
		l, _ := convert(s, 1, "seed")
		location = min(location, l)
	}

	return int(location)
}

func solveSecondProblem() int {
	s := parseSeedRanges(seedVals)

	location := math.Inf(1)
	c := make(chan float64)
	count := 0
	for _, seed := range s {
		count++
		go func(seed seeds) {
			c <- getLowestLocation(seed)
		}(seed)
	}

	fmt.Println(count)
	for i := 0; i < count; i++ {
		fmt.Println("loop", i)
		select {
		case l := <-c:
			if l < location {
				location = l
			}
		}
	}

	return int(location)
}

func getLowestLocation(seed seeds) float64 {
	loc := float64(-1)

	remaining := seed.length
	start := seed.start

	for remaining > 0 {
		l, jump := convert(start, remaining, "seed")
		remaining -= jump
		start += jump

		if loc == -1 || l < loc {
			loc = l
		}
	}
	return loc
}

func convert(val, jump int, name string) (float64, int) {
	if len(maps[name].to) == 0 {
		return float64(val), jump
	}
	ranger := maps[name].FindRange(val)
	if ranger != nil {
		diff := val - ranger.src
		newVal := ranger.dest + diff
		return convert(newVal, min(jump, ranger.length-diff), maps[name].to)
	}
	return convert(val, 1, maps[name].to)
}

func parseMaps() map[string]mapper {
	maps := map[string]mapper{}

	var cur mapper
	var name string
	for _, line := range content[1:] {
		if strings.Contains(line, "map:") {
			names := strings.Split(line, " ")
			names = strings.Split(names[0], "-")
			name = names[0]

			cur = mapper{to: names[2], rangers: []ranger{}}
			continue
		}
		if line == "" {
			continue
		}
		cur.rangers = append(cur.rangers, parseRange(line))
		maps[name] = cur
	}
	return maps
}

func parseRange(line string) ranger {
	vals := strings.Split(line, " ")
	src, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}
	dest, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	length, err := strconv.Atoi(vals[2])
	if err != nil {
		panic(err)
	}
	return ranger{src: src, dest: dest, length: length}
}

func parseNegativeRange() map[string]mapper {
	newMaps := maps
	for name, m := range maps {
		rangers := m.rangers
		slices.SortFunc(rangers, func(a, b ranger) int {
			return cmp.Compare(a.src, b.src)
		})

		start := 0
		for i := 0; i < len(rangers); i++ {
			r := rangers[i]
			if rangers[i].src > start {
				rangers = append(rangers, ranger{})
				copy(rangers[i+1:], rangers[i:])
				rangers[i] = ranger{src: start, dest: start, length: r.length}
				i++
			}
			start = r.src + r.length
		}
		m.rangers = rangers
		newMaps[name] = m
	}
	return newMaps
}

func cleanArray(arr []string) []string {
	for i, v := range arr {
		if v == "" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func parseSeeds(vals []string) []int {
	s := []int{}
	for _, v := range vals {
		seed, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		s = append(s, seed)
	}
	return s
}

func parseSeedRanges(vals []string) []seeds {
	s := []seeds{}
	for i := 0; i < len(vals); i += 2 {
		start, err := strconv.Atoi(vals[i])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(vals[i+1])
		if err != nil {
			panic(err)
		}
		s = append(s, seeds{start, length})
	}
	return s
}

func (m mapper) FindRange(val int) *ranger {
	for _, r := range m.rangers {
		if r.src <= val && r.src+r.length > val {
			return &r
		}
	}
	return nil
}
