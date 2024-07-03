package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2015/day8.txt")

	first := 0
	second := 0
	for _, l := range data {
		first += calculateLength(l)
		second += calculateLength(encode(l))
	}
	fmt.Println("First problem:", first)
	fmt.Println("Second problem:", second)
}

func calculateLength(data string) int {
	total := len(data)
	mem := parseInMemory(data)
	memory := len(mem)
	return total - memory
}

func encode(str string) string {
	sb := strings.Builder{}
	sb.WriteRune('"')
	for _, r := range str {
		if r == '"' || r == '\\' {
			sb.WriteRune('\\')
		}
		sb.WriteRune(r)
	}
	sb.WriteRune('"')
	return sb.String()
}

func parseInMemory(str string) string {
	mem := strings.ReplaceAll(str, `\\`, " ")
	mem = strings.ReplaceAll(mem, `\"`, " ")
	re := regexp.MustCompile(`\\x(\d|[a-f]){2}`)
	mem = re.ReplaceAllString(mem, " ")
	mem = strings.Trim(mem, "\"")
	return mem
}
