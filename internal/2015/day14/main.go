package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type reindeer struct {
	name string
	speed_per_second int
	go_seconds int
	rest_seconds int
}

const SECONDS = 2503

func main() {
	data := reader.FileToArray("data/2015/day14.txt")

	reindeers := []reindeer{}

	for _, line := range data {
		reindeers = append(reindeers, parseReindeer(line))
	}

	scores := map[string]int{}
	distances := map[string]int{}

	for t := range SECONDS {
		for _, r := range reindeers {
			if r.isResting(t) {
				continue
			}
			distances[r.name] += r.speed_per_second
		}
		for _, w := range getWinners(distances) {
			scores[w]++
		}
	}

	winners := getWinners(scores)
	fmt.Println("Winning reindeer (distance) after", SECONDS, "seconds:", winners, "distance:", distances[winners[0]])
	fmt.Println("Winning reindeer (score) after", SECONDS, "seconds:", winners, "score:", scores[winners[0]])
}

func (r reindeer) isResting(t int) bool {
	d := t % (r.go_seconds + r.rest_seconds) 
	return d >= r.go_seconds
}

func getWinners(distances map[string]int) []string {
	winners := []string{}
	time := math.MinInt
	for k, v := range distances {
		if v == time {
			winners = append(winners, k)
		}
		if v > time {
			winners = []string{k}
			time = v
		}
	}
	return winners
}

func parseReindeer(data string) reindeer {
	parts := strings.Split(data, " ")

	name := parts[0]
	speed, _ := strconv.Atoi(parts[3])
	go_sec, _ := strconv.Atoi(parts[6])
	rest_sec, _ := strconv.Atoi(parts[13])

	return reindeer{
		name:name,
		speed_per_second: speed,
		go_seconds: go_sec,
		rest_seconds: rest_sec,
	}
}
