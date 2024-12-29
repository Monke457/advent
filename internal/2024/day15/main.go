package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2024/day15.txt")

	warehouse, robot, moves := parseData(data)
	whlarge, robotlarge := expandWarehouse(warehouse, robot)

	warehouse, robot = executeMoves(warehouse, moves, robot, false)
	solve := sumCoords(warehouse)
	fmt.Println("Sum of box GPS:", solve)

	whlarge, robotlarge = executeMoves(whlarge, moves, robotlarge, true)
	solvelarge := sumCoords(whlarge)
	fmt.Println("Sum of box GPS large:", solvelarge)
	draw(whlarge, robotlarge)
}

func executeMoves(warehouse [][]byte, moves []byte, robot [2]int, large bool) ([][]byte, [2]int) {
	for _, move := range moves {
		x, y := getNewPos(robot, move)

		if !validMove(warehouse, x, y, move, large) {
			continue
		}

		warehouse = executeMove(warehouse, x, y, move, large)
		warehouse[robot[1]][robot[0]] = '.'
		robot[0], robot[1] = x, y
	}
	return warehouse, robot
}

func getNewPos(r [2]int, m byte) (int, int) {
	x, y := r[0], r[1]
	switch m {
	case '^':
		return x, y-1
	case 'v':
		return x, y+1
	case '<':
		return x-1, y
	case '>':
		return x+1, y
	}
	panic("big problem")
}

func sumCoords(warehouse [][]byte) int {
	sum := 0
	for i := range warehouse {
		for j := range warehouse[i] {
			if warehouse[i][j] != 'O' && warehouse[i][j] != '[' {
				continue
			}
			sum += 100 * i + j
		}
	}
	return sum
}

func executeMove(warehouse [][]byte, x, y int, dir byte, large bool) [][]byte {
	switch dir {
	case '^':
		if large {
			return shiftUpLarge(warehouse, []int{x}, y)
		}
		return shiftUp(warehouse, x, y)
	case 'v':
		if large {
			return shiftDownLarge(warehouse, []int{x}, y)
		}
		return shiftDown(warehouse, x, y)
	case '<':
		return shiftLeft(warehouse, x, y)
	case '>':
		return shiftRight(warehouse, x, y)
	}
	return warehouse
}

func shiftUpLarge(warehouse [][]byte, xes []int, y int) [][]byte {
	newXes := []int{}
	done := true
	for _, x := range xes {
		if warehouse[y][x] == '[' {
			done = false
			xes = append(xes, x+1)
			newXes = append(newXes, x, x+1)
		} 
		if warehouse[y][x] == ']' {
			done = false
			xes = append(xes, x-1)
			newXes = append(newXes, x, x-1)
		}
	}
	if !done {
		warehouse = shiftUpLarge(warehouse, newXes, y-1)
	}
	for _, x := range xes {
		warehouse = shiftUp(warehouse, x, y)
	}
	return warehouse
}

func shiftDownLarge(warehouse [][]byte, xes []int, y int) [][]byte {
	newXes := []int{}
	done := true
	for _, x := range xes {
		if warehouse[y][x] == '[' {
			done = false
			xes = append(xes, x+1)
			newXes = append(newXes, x, x+1)
		} 
		if warehouse[y][x] == ']' {
			done = false
			xes = append(xes, x-1)
			newXes = append(newXes, x, x-1)
		}
	}
	if !done {
		warehouse = shiftDownLarge(warehouse, newXes, y+1)
	}
	for _, x := range xes {
		warehouse = shiftDown(warehouse, x, y)
	}
	return warehouse
}

func shiftUp(warehouse [][]byte, x, y int) [][]byte {
	for i := y; i >= 0; i-- { 
		if warehouse[i][x] == '.' {
			for j := i; j <= y; j++ {
				warehouse[j][x] = warehouse[j+1][x]
			}
			warehouse[y][x] = '.'
			break
		}
	}
	return warehouse
}

func shiftDown(warehouse [][]byte, x, y int) [][]byte {
	for i := y; i < len(warehouse); i++ { 
		if warehouse[i][x] == '.' {
			for j := i; j > y; j-- {
				warehouse[j][x] = warehouse[j-1][x]
			}
			warehouse[y][x] = '.'
			break
		}
	}
	return warehouse
}

func shiftLeft(warehouse [][]byte, x, y int) [][]byte {
	newRow := []byte{}
	skipped := false 
	for i := len(warehouse[y])-1; i >= 0; i-- {
		if i > x+1 {
			newRow = append([]byte{warehouse[y][i]}, newRow...)
		} else if i == x+1 {
			newRow = append([]byte{warehouse[y][i], '.'}, newRow...)
		} else if warehouse[y][i] == '.' && !skipped {
			skipped = true
		} else {
			newRow = append([]byte{warehouse[y][i]}, newRow...)
		}
	}
	warehouse[y] = newRow
	return warehouse
}

func shiftRight(warehouse [][]byte, x, y int) [][]byte {
	newRow := []byte{}
	skipped := false 
	for i := 0; i < len(warehouse[y]); i++ {
		if i < x-1 {
			newRow = append(newRow, warehouse[y][i])
		} else if i == x-1 {
			newRow = append(newRow, '.', warehouse[y][i])
		} else if warehouse[y][i] == '.' && !skipped {
			skipped = true
		} else {
			newRow = append(newRow, warehouse[y][i])
		}
	}
	warehouse[y] = newRow
	return warehouse
}

func validMove(warehouse [][]byte, x, y int, dir byte, large bool) bool {
	if y < 0 || y >= len(warehouse) {
		return false
	}
	if x < 0 || x >= len(warehouse[y]) {
		return false
	}
	switch dir {
	case '^':
		if large {
			return validUpLarge(warehouse, []int{x}, y)
		}
		return validUp(warehouse, x, y)
	case 'v':
		if large {
			return validDownLarge(warehouse, []int{x}, y)
		}
		return validDown(warehouse, x, y)
	case '<':
		return validLeft(warehouse, x, y)
	case '>':
		return validRight(warehouse, x, y)
	}
	panic("Big problem")
}

func validUpLarge(warehouse [][]byte, xes []int, y int) bool {
	if len(xes) == 0 {
		return true
	}
	newXes := []int{}
	for _, x := range xes {
		if warehouse[y][x] == '#' {
			return false 
		}
		if warehouse[y][x] == '[' {
			newXes = append(newXes, x, x+1)
		} 
		if warehouse[y][x] == ']' {
			newXes = append(newXes, x, x-1)
		}
	}
	return validUpLarge(warehouse, newXes, y-1)
}

func validDownLarge(warehouse [][]byte, xes []int, y int) bool {
	if len(xes) == 0 {
		return true
	}
	newXes := []int{}
	for _, x := range xes {
		if warehouse[y][x] == '#' {
			return false 
		}
		if warehouse[y][x] == '[' {
			newXes = append(newXes, x, x+1)
		} 
		if warehouse[y][x] == ']' {
			newXes = append(newXes, x, x-1)
		}
	}
	return validDownLarge(warehouse, newXes, y+1)
}

func validUp(warehouse [][]byte, x, y int) bool {
	for i := y; i >= 0; i-- { 
		if warehouse[i][x] == '.' {
			return true
		}
		if warehouse[i][x] == '#' {
			break
		}
	}
	return false
}

func validDown(warehouse [][]byte, x, y int) bool {
	for i := y; i < len(warehouse); i++ { 
		if warehouse[i][x] == '.' {
			return true
		}
		if warehouse[i][x] == '#' {
			break
		}
	}
	return false
}

func validLeft(warehouse [][]byte, x, y int) bool {
	for i := x; i >= 0; i-- {
		if warehouse[y][i] == '.' {
			return true
		}
		if warehouse[y][i] == '#' {
			break
		}
	}
	return false
}

func validRight(warehouse [][]byte, x, y int) bool {
	for _, b := range warehouse[y][x:] {
		if b == '.' {
			return true
		}
		if b == '#' {
			break
		}
	}
	return false
}

func expandWarehouse(original [][]byte, robot [2]int) ([][]byte, [2]int) {
	whlarge := [][]byte{}
	for _, row := range original {
		expanded := []byte{}
		for _, b := range row {
			if b == 'O' {
				expanded = append(expanded, '[', ']')
				continue
			}
			expanded = append(expanded, b, b)
		}
		whlarge = append(whlarge, expanded)
	}
	robot[0] = robot[0] * 2
	return whlarge, robot
}

func parseData(data []string) ([][]byte, [2]int, []byte) {
	robot := [2]int{}
	warehouse := [][]byte{}
	moves := []byte{}

	whDone := false
	for i, line := range data {
		if line == "" {
			whDone = true
		}
		if whDone {
			moves = append(moves, line...)

		} else {
			l := []byte(line)
			if j := strings.Index(line, "@"); j != -1 {
				robot[0] = j
				robot[1] = i
				l[j] = '.'
			}
			warehouse = append(warehouse, l)
		}
	}
	return warehouse, robot, moves
}

func draw(warehouse [][]byte, robot [2]int) {
	for i, row := range warehouse {
		if i == robot[1] {
			row[robot[0]] = '@'
		}
		fmt.Printf("%s\n",row)
	}
}
