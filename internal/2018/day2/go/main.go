package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2018/day2.txt")

	doubles := 0
	triples := 0

	var common string
	var done bool
	for i := range data {
		counts := countLetters(data[i]) 
		if containsRepeat(counts, 2) {
			doubles++
		}
		if containsRepeat(counts, 3) {
			triples++
		}

		if done {
			continue
		}
		for j := i+1; j < len(data); j++ {
			if idx, ok := offByOne(data[i], data[j]); ok {
				common = data[i][:idx] + data[i][idx+1:]
				done = true
			}
		}
	}

	fmt.Println("First:", doubles * triples)
	fmt.Println("Common letters:", common)
}

func offByOne(str1, str2 string) (int, bool) {
	if len(str1) != len(str2) {
		return -1, false
	}
	idx := -1
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			continue
		}
		if idx > -1 {
			return -1, false
		}
		idx = i
	}
	return idx, true 
}

func countLetters(str string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range str { 
		counts[r]++
	}
	return counts 
}

func containsRepeat(letters map[rune]int, n int) bool {
	for _, count := range letters {
		if count == n {
			return true
		}
	}
	return false
}
