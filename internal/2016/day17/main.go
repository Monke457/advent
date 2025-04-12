package main

import (
	"advent/internal/pkg/reader"
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte(reader.FileToString("data/2016/day17.txt"))

	calculatePaths([]byte{}, data)
	fmt.Println("shortest path: ",sprintPath(shortest))
	fmt.Println("length of longest path: ", len(longest))
}

func sprintPath(path []byte) string {
	res := ""
	for _, b := range path {
		res += fmt.Sprintf("%c", b)
	}
	return res
}

func getPos(path []byte) [2]int {
	y, x := 0, 0
	for _, p := range path {
		switch p {
		case 'U':
			y--
		case 'L':
			x--
		case 'D':
			y++
		case 'R':
			x++
		}
	}
	return [2]int{y, x}
}

var shortest = []byte{}
var longest = []byte{}

func calculatePaths(path []byte, hash []byte) {
	pos := getPos(path)
	if pos[0] == 3 && pos[1] == 3 {
		if len(shortest) == 0 || len(path) < len(shortest) {
			shortest = make([]byte, len(path))
			copy(shortest, path)
		}
		if len(path) > len(longest) {
			longest = make([]byte, len(path))
			copy(longest, path)
		}
		return 
	}
	cs := getChecksum(hash, path)
	dirs := getDirections(cs, pos)
	for _, d := range dirs {
		path = append(path, d)
		calculatePaths(path, hash)
		path = path[:len(path)-1]
	}
	return 
}

// up down left right
func getDirections(cs [2]byte, pos [2]int) []byte {
	res := []byte{}
	i := 0
	for _, b := range cs {
		if b >= 11 * 16 {
			if i == 0 && pos[0] > 0 {
				res = append(res, 'U')
			}
			if i == 2 && pos[1] > 0 {
				res = append(res, 'L')
			}
		}
		i++
		if (b % 16) / 11 > 0 {
			if i == 1 && pos[0] < 3 {
				res = append(res, 'D')
			}
			if i == 3 && pos[1] < 3 {
				res = append(res, 'R')
			}
		}
		i++
	}
	return res
}

func getChecksum(hash, path []byte) [2]byte {
	cs := md5.Sum(append(hash, path...))
	return [2]byte{cs[0], cs[1]}
}
