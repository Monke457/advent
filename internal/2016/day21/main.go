package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	instructions := reader.FileToArray("data/2016/day21.txt")

	pass := "abcdefgh"
	scrambled := "fbgdceah"

	s := scramble(instructions, pass, false)

	fmt.Printf("[%s] -> [%s] (scrambled)\n", pass, s)
	fmt.Printf("[%s] -> [%s] (unscrambled)\n", s, scramble(instructions, s, true))
	fmt.Printf("[%s] -> [%s] (unscrambled)\n", scrambled, scramble(instructions, scrambled, true))
}

func scramble(originInst []string, origin string, rev bool) string {
	pass := []rune(origin)

	instructions := make([]string, len(originInst))
	copy(instructions, originInst)

	if rev {
		slices.Reverse(instructions)
	}

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")

		switch parts[0] {
		case "swap":
			if parts[1] == "letter" {
				pass = swapLetter(pass, parts[2], parts[5])
			} else {
				a, _ := strconv.Atoi(parts[2])
				b, _ := strconv.Atoi(parts[5])
				pass = swapPosition(pass, a, b)
			}

		case "rotate":
			if parts[1] == "left" {
				stepsStr := parts[2]
				steps, _ := strconv.Atoi(stepsStr)
				if rev {
					pass = rotateRight(pass, steps)
				} else {
					pass = rotateLeft(pass, steps)
				}

			} else if parts[1] == "right" {
				stepsStr := parts[2]
				steps, _ := strconv.Atoi(stepsStr)
				if rev {
					pass = rotateLeft(pass, steps)
				} else {
					pass = rotateRight(pass, steps)
				}

			} else {
				if rev {
					pass = reverseRotatePosition(pass, parts[6])
				} else {
					pass = rotatePosition(pass, parts[6])
				}
			}

		case "reverse":
			a, _ := strconv.Atoi(parts[2])
			b, _ := strconv.Atoi(parts[4])
			pass = reverse(pass, a, b)

		case "move":
			a, _ := strconv.Atoi(parts[2])
			b, _ := strconv.Atoi(parts[5])
			if rev {
				pass = move(pass, b, a)
			} else {
				pass = move(pass, a, b)
			}
		}
	}

	return string(pass)
}

func swapLetter(origin []rune, a, b string) []rune {
	aRune, bRune := rune(a[0]), rune(b[0])
	for i, r := range origin {
		if r == aRune {
			origin[i] = bRune
		} else if r == bRune {
			origin[i] = aRune
		}
	}
	return origin 
}

func swapPosition(origin []rune, a, b int) []rune {
	aRune, bRune := origin[a], origin[b]
	origin[a] = bRune
	origin[b] = aRune
	return origin
}

func rotateLeft(origin []rune, steps int) []rune {
	steps = steps % len(origin)
	return append(origin[steps:], origin[:steps]...)
}

func rotateRight(origin []rune, steps int) []rune {
	steps = len(origin) - steps
	return append(origin[steps:], origin[:steps]...)
}

func reverseRotatePosition(origin []rune, a string) []rune {
	aRune := rune(a[0]) 
	idx := 1
	for i, r := range origin {
		if r == aRune {
			idx += i
			if idx % 2 == 0 {
				idx = idx / 2
			} else if idx != 1 {
				idx += len(origin)
				idx /= 2
				idx++
			}
			break
		}
	}
	return rotateLeft(origin, idx)
}

func rotatePosition(origin []rune, a string) []rune {
	aRune := rune(a[0])
	idx := 1
	for i, r := range origin {
		if r == aRune {
			idx += i
			if i >= 4 {
				idx++
			}
			break
		}
	}
	idx %= len(origin)
	return rotateRight(origin, idx)
}

func reverse(origin []rune, a, b int) []rune {
	temp := origin[a:b+1]
	slices.Reverse(temp)
	temp = append(temp, origin[b+1:]...)
	return append(origin[:a], temp...)
}

func move(origin []rune, a, b int) []rune {
	moved := []rune{}
	for i := 0; i < len(origin); i++ {
		if i == a {
			continue
		} 
		if i == b {
			if len(moved) == i {
				moved = append(moved, origin[a])
				moved = append(moved, origin[i])
			} else {
				moved = append(moved, origin[i])
				moved = append(moved, origin[a])
			}
		} else {
			moved = append(moved, origin[i])
		}
	}
	return moved
}
