package reader

import (
	"bufio"
	"os"
	"strconv"
)

func FileToString(fp string) (r string) {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		r += scanner.Text()
	}
	return
}

func FileToIntArray(fp string) (r []int) {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		rn := scanner.Text()
		if rn == "\n" {
			break
		}
		i, err := strconv.Atoi(string(rn))
		if err != nil {
			panic(err)
		}
		r = append(r, i)
	}
	return
}

func FileTo2DIntArray(fp string) (r [][]int) {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := []int{}
		for _, b := range scanner.Text() {
			i, err := strconv.Atoi(string(b))
			if err != nil {
				panic(err)
			}
			row = append(row, i)
		}
		r = append(r, row)
	}
	return
}

func FileTo2DArray(fp string) (r [][]rune) {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := []rune{}
		for _, b := range scanner.Text() {
			row = append(row, b)
		}
		r = append(r, row)
	}
	return
}

func FileToArray(fp string) (r []string) {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	return
}

func Reverse(s string) (r string) {
	for _, c := range s {
		r = string(c) + r
	}
	return
}
