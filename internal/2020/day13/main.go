package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	data := reader.FileToArray("data/2020/day13.txt")

	fmt.Println("First:", first(data))
	start := time.Now()
	fmt.Println("Second:", second(data[1]))
	fmt.Println(time.Since(start))
}

func first(data []string) int {
	ts, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	ids := parseIds(data[1])
	res := [2]int{ math.MaxInt, -1 }
	for _, id := range ids {
		diff := id - (ts % id)
		if diff < res[0] {
			res[0] = diff
			res[1] = id
		}
	}
	return res[0] * res[1]
}

var cache = [2]int{ 0, 0 }

func second(data string) int {
	ids := parseAllIds(data)
	cache = ids[0] 
	ts := cache[1] - ids[cache[0]][0]

	for !validTimes(ids, ts) {
		ts += cache[1]  
	}
	return ts
}

func validTimes(ids [][2]int, ts int) bool {
	for i, id := range ids {
		if (ts + id[0]) % id[1] != 0 {
			return false 
		}
		if i > cache[0] {
			cache[0] = i
			cache[1] = cache[1] * id[1] 
		}
	}
	return true 
}

func parseAllIds(line string) [][2]int {
	res := make([][2]int, 0)
	for i, val := range strings.Split(line, ",") {
		if val == "x" {
			continue
		}
		v, err := strconv.Atoi(val)
		if err == nil {
			res = append(res, [2]int{i, v})
		}
	}
	return res
}

func parseIds(line string) []int {
	res := make([]int, 0)
	for _, val := range strings.Split(line, ",") {
		if val == "x" {
			continue
		}
		v, err := strconv.Atoi(val)
		if err == nil {
			res = append(res, v)
		}
	}
	return res
}
