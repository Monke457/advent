package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

type pattern struct {
	cells [][]rune
}

func main() {
	lines := reader.FileToArray("data/day13.txt")
	patterns := parsePatterns(lines)
	sum := 0
	for _, p := range patterns {
		fmt.Printf("pattern\n")
		c, v := p.CountMirrored()
		if v {
			sum += c
		} else {
			sum += c * 100
		}
		for _, row := range p.cells {
			fmt.Printf("%c\n", row)
		}
	}
	fmt.Println(sum)
}

func parsePatterns(lines []string) (res []pattern) {
	cells := [][]rune{}
	row := 0
	for i, l := range lines {
		if i+1 == len(lines) || l == "" {
			res = append(res, pattern{cells})
			cells = [][]rune{}
			row = 0
			continue
		}
		cells = append(cells, []rune{})
		for _, r := range l {
			cells[row] = append(cells[row], r)
		}
		row++
	}
	return
}

func verticalMirrored(cells [][]rune, col, l int) map[int][2]int {
	v := map[int][2]int{}
	for i := range cells {
		for j := col; j < col+l && j < len(cells[i]); j++ {
			c := 0
			for j-c >= 0 && j+1+c < len(cells[i]) {
				if cells[i][j-c] != cells[i][j+1+c] {
					break
				}
				c++
			}
			if c >= v[i][1] {
				v[i] = [2]int{j, c}
			}
		}
	}
	return v
}

func horizontalMirrored(cells [][]rune) map[int][2]int {
	h := map[int][2]int{}
	return h
}

func (p pattern) CountMirrored() (res int, vertical bool) {
	v := verticalMirrored(p.cells, 0, len(p.cells[0]))

	vk, iv, s := getMin(v)

	for !s {
		v = verticalMirrored(p.cells, iv[0], iv[1])
		vk, iv, s = getMin(v)
	}

	fmt.Println(v, vk, iv, "same:", s)

	if iv[1] != 0 {
		fmt.Printf("vertically mirrored at col %d\n", iv[0])
		res = iv[0] + 1
		vertical = true
		return
	}

	h := horizontalMirrored(p.cells)
	fmt.Println(h)
	return
}

func getMin(m map[int][2]int) (k int, i [2]int, s bool) {
	s = true
	for key, val := range m {
		if i[1] == 0 {
			i = val
			k = key
			continue
		}
		if val[1] != i[1] {
			s = false
			if val[1] < i[1] {
				i = val
				k = key
			}
		}
	}
	return
}
