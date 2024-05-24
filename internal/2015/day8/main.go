package main

import (
	"advent/pkg/reader"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2015/day8.txt")

	total := 0
	memory := 0

	for _, l := range data {
		total += len(l)
		mem := parseInMemory(l)
		memory += len(mem)
	}
	fmt.Println("First problem:", total-memory)
}

func parseInMemory(str string) string {
	mem := strings.ReplaceAll(str, `\\`, " ")
	mem = strings.ReplaceAll(mem, `\"`, " ")
	re := regexp.MustCompile(`\\x(\d|[a-f]){2}`)
	mem = re.ReplaceAllString(mem, " ")
	mem = strings.Trim(mem, "\"")
	return mem
}
