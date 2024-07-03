package main

import (
	"advent/internal/pkg/reader"
	"advent/internal/pkg/sorter"
	"fmt"
	"strconv"
	"strings"
)

type Room struct {
	name     string
	id       int
	checksum string
}

func main() {
	data := reader.FileToArray("data/2016/day4.txt")

	rooms := []Room{}
	for _, row := range data {
		rooms = append(rooms, parseRoom(row))
	}

	solveFirstProblem(rooms)
	solveSecondProblem(rooms)
}

func solveFirstProblem(rooms []Room) {
	sum := 0
	for _, room := range rooms {
		if room.IsReal() {
			sum += room.id
		}
	}
	fmt.Println(sum)
}

func solveSecondProblem(rooms []Room) {
	var decrypted string
	idx := 0
	for _, room := range rooms {
		decrypted = room.decrypt()
		if strings.Contains(decrypted, "northpole") {
			idx = room.id
			break
		}
	}
	fmt.Println(idx)
}

func (r Room) decrypt() string {
	sb := strings.Builder{}
	for _, val := range r.name {
		if val == '-' {
			sb.WriteRune(' ')
			continue
		}
		valInt := int(val) - 97
		c := ((valInt + r.id) % 26) + 97
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

func (r Room) IsReal() bool {
	counts := map[rune]int{}
	for _, val := range r.name {
		if val == '-' {
			continue
		}
		counts[val]++
	}

	order := sorter.GetOrderByValue(counts)

	for i, val := range r.checksum {
		if order[i] != val {
			return false
		}
	}
	return true
}

func parseRoom(data string) Room {
	idx1 := strings.LastIndex(data, "-")
	idx2 := strings.Index(data, "[")

	name := data[:idx1]

	id := data[idx1+1 : idx2]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	checksum := data[idx2+1 : len(data)-1]

	return Room{name: name, id: idNum, checksum: checksum}
}
