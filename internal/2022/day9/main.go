package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	
	pg "github.com/Monke457/printge"
)

func main() {
	data := reader.FileToArray("data/2022/day9.txt")

	l := 10
	body := initBody(l)
	bodyMaps := initMaps(l)
	colors := initColors(l)

	for _, line := range data {
		sSteps := line[2:]
		steps, _ := strconv.Atoi(sSteps)

		dir := [2]int{0, 0}
		
		switch line[0] {
		case 'U':
			dir[0] = -1
		case 'D':
			dir[0] = 1
		case 'R':
			dir[1] = 1
		case 'L':
			dir[1] = -1
		}

		for range steps {
			body[0][0] += dir[0]
			body[0][1] += dir[1]
			bodyMaps[0][body[0]] = true

			for i := 1; i < l; i++ {
				body[i] = move(body[i], body[i-1])
				bodyMaps[i][body[i]] = true
			}
		}
	}
	temp := append(bodyMaps[0:l-1], bodyMaps[l-1])
	printMap(temp, colors)
	fmt.Println("Points visited by tail:", len(bodyMaps[l-1]))
}

func initBody(l int) [][2]int {
	body := make([][2]int, l)
	for i := range l {
		body[i] = [2]int{0,0}
	}
	return body
}

func initMaps(l int) []map[[2]int]bool {
	maps := make([]map[[2]int]bool, l)
	for i := range l {
		maps[i] = map[[2]int]bool{{0,0}:true}
	}
	return maps
}

func initColors(l int) []pg.Color {
	cols := make([]pg.Color, l)
	for i := range l {
		cols[i] = pg.GetNext()
	}
	return cols 
}

func printMap(maps []map[[2]int]bool, colors []pg.Color) {
	yOff, xOff := 0, 0
	for i := range maps {
		yo, xo := getOffsets(maps[i])
		if yo > yOff {
			yOff = yo
		}
		if xo > xOff {
			xOff = xo
		}
	}

	for i := range maps {
		maps[i] = normalizeMap(maps[i], yOff, xOff)
	}
	h, w := 0, 0
	for i := range maps {
		ht, wt := getSize(maps[i])
		if ht > h {
			h = ht
		}
		if wt > w {
			w = wt
		}
	}
	for y := range h {
		for x := range w {
			pos := [2]int{y,x}
			occupied := false
			for i := range maps {
				if maps[i][pos] {
					pg.Print("#", colors[i])
					occupied = true
					break
				}
			} 
			if !occupied {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func normalizeMap(m map[[2]int]bool, yOff, xOff int) map[[2]int]bool {
	newMap := map[[2]int]bool{}
	for k := range m {
		newKey := [2]int{k[0]+yOff, k[1]+xOff}
		newMap[newKey] = true
	}
	return newMap
}

func getSize(m map[[2]int]bool) (int, int) {
	y, x := math.MinInt, math.MinInt
	for k := range m {
		if k[0] > y {
			y = k[0]
		}
		if k[1] > x {
			x = k[1]
		}
	}
	return y+1, x+1
}

func getOffsets(m map[[2]int]bool) (int, int) {
	yMin, xMin := math.MaxInt, math.MaxInt
	for k := range m {
		if k[0] < yMin {
			yMin = k[0]
		}
		if k[1] < xMin {
			xMin = k[1]
		}
	}
	if yMin < 0 {
		yMin = -yMin
	}
	if xMin < 0 {
		xMin = -xMin
	}
	return yMin, xMin 
}

func move(tail, head [2]int) [2]int {
	distA := head[0] - tail[0]
	distB := head[1] - tail[1]

	if (distA > -2 && distA < 2)  && (distB > -2 && distB < 2) {
		return tail
	}

	if distA == 0 {
		if distB < 0 {
			distB++
		} else {
			distB--
		}
		tail[1] += distB  
		return tail
	} 

	if distB == 0 {
		if distA < 0 {
			distA++
		} else {
			distA--
		}
		tail[0] += distA
		return tail
	}

	if distA < -1 {
		distA++
	} else if distA > 1 {
		distA--
	}
	if distB < -1 {
		distB++
	} else if distB > 1 {
		distB--
	}

	tail[0] += distA
	tail[1] += distB
	return tail
}
