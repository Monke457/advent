package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Monke457/advent/internal/pkg/reader"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	//	fmt.Println(solveFirstProblem())
	fmt.Println(solveSecondProblem())
}

func fileToArray(fp string) []string {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func solveFirstProblem() int {
	lines := fileToArray("data/day1.txt")

	sum := 0
	for _, l := range lines {
		first, _ := getFirstDigit(l)
		last, _ := getLastDigit(l)
		val, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		if err != nil {
			panic(err)
		}
		sum += val
	}

	return sum
}

func solveSecondProblem() int {
	lines := fileToArray("data/day1.txt")

	sum := 0

	for _, l := range lines {

		first, idx := getFirstDigit(l)
	loop1:
		for i := 0; i < idx; i++ {
			for key, d := range digits {
				if len(key) > len(l)-i {
					continue
				}
				if l[i:i+len(key)] == key {
					first = fmt.Sprintf("%d", d)
					break loop1
				}
			}
		}

		last, idx := getLastDigit(l)
	loop2:
		for i := len(l) - 1; i > idx; i-- {
			for key, d := range digits {
				if len(key) > len(l)-i {
					continue
				}
				if l[i:i+len(key)] == key {
					last = fmt.Sprintf("%d", d)
					break loop2
				}
			}
		}

		str := fmt.Sprintf("%s%s", first, last)
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		sum += val
	}

	return sum
}

func getFirstDigit(line string) (string, int) {
	for i := range line {
		if line[i] > 47 && line[i] < 58 {
			return line[i : i+1], i
		}
	}
	return "", len(line)
}

func getLastDigit(line string) (string, int) {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] > 47 && line[i] < 58 {
			return line[i : i+1], i
		}
	}
	return "", 0
}
