package reader

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
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

func FileToJsonEncoded(fp string) (r []byte) {
	data := FileToString(fp)
	val, err := json.Marshal(data)
	if err != nil {
		log.Printf("Could not marshal json %s\n", err)
	}
	return val
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
		text := scanner.Text()

		var nums []string
		if strings.Contains(text, " ") {
			nums = strings.Split(text, " ")

		} else if strings.Contains(text, "\t") {
			nums = strings.Split(text, "\t")

		} else {
			nums = strings.Split(text, "") 
		}

		for _, num := range nums {
			i := convertToInt(num)
			if i != nil {
				row = append(row, *i)
			}
		}
		r = append(r, row)
	}
	return
}

func convertToInt(str string) *int {
	res, err := strconv.Atoi(str)
	if err != nil {
		return nil
	}
	return &res
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
