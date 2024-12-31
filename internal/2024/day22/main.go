package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

const (
	PRUNE_NUM int = 16777216
	ITERS = 2000
)

type results struct {
	secret int
	changes []int
	costs []int
	changeMap map[[4]int]int
}

func main() {
	secrets := reader.FileToIntArray("data/2024/day22.txt")

	resultChan := make(chan results)
	for _, secret := range secrets {
		go evolveSecret(resultChan, secret, ITERS)
	}

	sum := 0
	resultList := []results{}
	for range len(secrets) {
		select {
		case result := <-resultChan:
			sum += result.secret
			resultList = append(resultList, result)
		}
	}
	fmt.Println("Sum of new secrets after", ITERS, "iterations:", sum)

	reducedMap := map[[4]int]int{}
	for _, res := range resultList {
		for k, v := range res.changeMap {
			reducedMap[k] += v
		}
	}

	highest := 0
	cFour := [4]int{}
	for key, cost := range reducedMap {
		if cost > highest {
			highest = cost
			cFour = key
		}
	}

	fmt.Println("Highest cost:", highest, "after changes:", cFour)
}

func evolveSecret(ch chan results, sec, n int) {
	changes := []int{}
	costs := []int{sec % 10}
	for i := range n {
		sec = evolve(sec)
		cost := sec % 10
		changes = append(changes, cost - costs[i])
		costs = append(costs, cost)
	}
	res := results{secret: sec, changes: changes, costs: costs[1:]}
	res.mapCosts(n)
	ch <- res
}

func (r *results) mapCosts(n int) {
	r.changeMap = map[[4]int]int{}
	for i := 0; i < n-3; i++ {
		key := [4]int{r.changes[i], r.changes[i+1], r.changes[i+2], r.changes[i+3]}
		if _, ok := r.changeMap[key]; ok {
			continue
		}
		r.changeMap[key] = r.costs[i+3]
	}
}

func evolve(secret int) int {
	temp := bitwiseMultiply(secret, 6)
	secret = mix(secret, temp)
	secret = prune(secret, PRUNE_NUM)
	
	temp = bitwiseDivide(secret, 5)
	secret = mix(secret, temp)
	secret = prune(secret, PRUNE_NUM)

	temp = bitwiseMultiply(secret, 11)
	secret = mix(secret, temp) 
	secret = prune(secret, PRUNE_NUM)

	return secret
}

func bitwiseMultiply(x, y int) int {
	return x << y 
}

func bitwiseDivide(x, y int) int {
	return x >> y
}

func mix(x, y int) int {
	return x ^ y
}

func prune(x, pn int) int {
	return x % pn
}
