package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var fields = map[string]func(string) bool {
	"byr:": isByrValid, 
	"iyr:": isIyrValid, 
	"eyr:": isEyrValid, 
	"hgt:": isHgtValid, 
	"hcl:": isHclValid, 
	"ecl:": isEclValid, 
	"pid:": isPidValid,
}

var colors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func main() {
	data := reader.FileToArray("data/2020/day4.txt")

	var current string
	var count1, count2 int
	for _, line := range data {
		if len(line) > 0 {
			current += " " + line
			continue
		}
		if passIsValid(current) {
			count1++
		}
		if passIsValidStrict(current) {
			count2++
		}
		current = ""
	}

	//check last line in case it is not empty
	if passIsValid(current) {
		count1++
	}
	if passIsValidStrict(current) {
		count2++
	}

	fmt.Println("First:", count1)
	fmt.Println("Second:", count2)
}

func passIsValid(pass string) bool {
	for name := range fields {
		if !strings.Contains(pass, name) {
			return false
		}
	}
	return true
}

func passIsValidStrict(pass string) bool {
	for name, fn := range fields {
		_, val, ok := strings.Cut(pass, name)
		if !ok || !fn(strings.Split(val, " ")[0]) {
			return false
		}
	}
	return true
}

func isByrValid(s string) bool {
	return isWithinRange(s, 1920, 2002)
}

func isIyrValid(s string) bool {
	return isWithinRange(s, 2010, 2020)
}

func isEyrValid(s string) bool {
	return isWithinRange(s, 2020, 2030)
}

func isHgtValid(s string) bool {
	val, _, ok := strings.Cut(s, "cm")
	if !ok {
		val, _, ok = strings.Cut(s, "in")
		if !ok {
			return false
		}
		return isWithinRange(val, 59, 76)
	}
	return isWithinRange(val, 150, 193)
}

func isHclValid(s string) bool {
	re := regexp.MustCompile(`#(([a-f]|[0-9]){6})`)
	result := re.FindIndex([]byte(s))
	if result == nil {
		return false
	}
	return result[0] == 0 && result[1] == len(s)
}

func isEclValid(s string) bool {
	for _, color := range colors {
		if s == color {
			return true
		}
	}
	return false
}

func isPidValid(s string) bool {
	if len(s) != 9 {
		return false
	}
	if _, err := strconv.Atoi(s); err != nil {
		return false 
	}
	return true 
}


func isWithinRange(s string, lo, hi int) bool {
	val, err := strconv.Atoi(s)
	if  err != nil {
		return false
	}
	return val >= lo && val <= hi 
}

