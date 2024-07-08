package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToArray("data/2021/day3.txt")

	fmt.Println("First:", first(data))
	fmt.Println("Second:", second(data))
}

func first(data []string) int {
	com := getMap(data)
	l := len(com) - 1
	g, e := 0, 0
	for k, v := range com {
		s := l - k
		if v > 0 {
			g += 1 << s 
		} else {
			e += 1 << s 
		}
	}

	return g * e
}

func findKey(data []string, fn func(int, byte) bool) (*string, error) {
	valid := map[int]string{}
	common := getMap(data)
	for i, line := range data {
		valid[i] = line
	}

	for k := range len(common) {
		for i, str := range valid {
			if len(valid) == 1 {
				break
			}
			if fn(common[k], str[k]) {
				delete(valid, i)
			}
		}
		if len(valid) == 1 { 
			for _, v := range valid {
				return &v, nil
			}
		}
		common = updateMap(valid)
	}
	return nil, fmt.Errorf("Could not find a valid key")
}

func second(data []string) int {
	ogF, err := findKey(data, func(a int, b byte) bool {
		return (a < 0 && b == '1') || (a >= 0 && b == '0') 
	})
	if err != nil {
		panic(err)
	}

	coF, err := findKey(data, func(a int, b byte) bool {
		return (a < 0 && b == '0') || (a >= 0 && b == '1') 
	})
	if err != nil {
		panic(err)
	}

	return convertBin(*ogF) * convertBin(*coF)
}

func convertBin(bin string) int {
	res := 0
	for i, b := range bin { 
		l := len(bin) - i - 1
		if b == '1' {
			res += 1 << l 
		}
	}
	return res
}

func getMap(data []string) map[int]int {
	com := map[int]int{}
	for i := range data {
		for j := range data[i] {
			if data[i][j] == '1' {
				com[j]++
			} else {
				com[j]--
			}
		}
	}
	return com
}

func updateMap(m map[int]string) map[int]int {
	res := map[int]int{}
	for _, v := range m {
		for i, j := range v {
			if j == '1' {
				res[i]++
			} else {
				res[i]--
			}
		}
	}
	return res
}
