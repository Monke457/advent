package main

import (
	"advent/internal/pkg/reader"
	"crypto/md5"
	"fmt"
	"slices"
)

const CYCLES = 2016

func main() {
	salt := reader.FileToString("data/2016/day14.txt")

	//salt = "abc"
	keys := []int{}

	var hash string
	for i := 0; len(keys) < 64; i++ {
		hash = hashData(salt, i, CYCLES)

		if t, ok := findTriple(hash, i); ok {
			if searchForQuint(salt, i, *t) {
				keys = append(keys, i)
				fmt.Printf("\r%d  ", len(keys))
			}
		}
	}

	fmt.Println("64th key found at index", keys[63])
}

var hCache = map[int]string{}
var tCache = map[int]byte{}
var qCache = map[int][]byte{}

func findTriple(hash string, n int) (*byte, bool) { 
	if t, ok := tCache[n]; ok {
		return &t, true
	}
	for i := 0; i < len(hash)-2; i++ {
		r := hash[i]
		if r == hash[i+1] && r == hash[i+2] {
			return &r, true
		}
	}
	return nil, false
} 

func searchForQuint(salt string, n int, val byte) bool { 
	var hash string
	for range 1000 {
		n++
		if arr, ok := qCache[n]; ok && slices.Contains(arr, val) {
			return true
		}

		hash = hashData(salt, n, CYCLES)

		for i := 0; i < len(hash)-4; i++ {
			if hash[i] == hash[i+1] && hash[i] == hash[i+2]  {
				if hash[i] == hash[i+3] && hash[i] == hash[i+4] {
					if _, ok := qCache[n]; !ok {
						qCache[n] = []byte{}
					}
					qCache[n] = append(qCache[n], hash[i])
					if hash[i] == val {
						return true
					}
					continue
				}
				if _, ok := tCache[n]; !ok {
					tCache[n] = hash[i]
				}
			}
		}
	}
	return false
} 

func hashData(salt string, val, n int) string {
	if hash, ok := hCache[val]; ok {
		return hash
	}
	data := fmt.Sprintf("%s%d", salt, val)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	for range n {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	hCache[val] = hash
	return hash
}

