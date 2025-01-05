package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := reader.FileToArray("data/2024/day17.txt")

	regA, regB, regC, program := parseData(data)
	fmt.Printf("A: %d B: %d C: %d program: %d\n", regA, regB,regC, program)

	output := run(program, regA, regB, regC)
	fmt.Println("Output:", output) 

	val := 1
	for len(output) < len(program) {
		output = run(program, val, regB, regC)
		if len(output) == len(program) {
			break
		}
		val <<= 3
	}

	fmt.Println("looking for correct positions")
	loop:
	for {
		fmt.Println("currently", countCorrect(program, output), "correct\n", output, "\n", program)
		for i := 0; i < len(program); i++ {
			val, output = runIncrement(program, val, regB, regC, i)
			if slices.Equal(output, program) {
				break loop
			}
		}
	}

	fmt.Println("lowest A register to produce the same program:", val)

}

func countCorrect(original, output []int) int {
	correct := 0
	for i := 0; i < len(original) && i < len(output); i++ {
		if original[i] == output[i] {
			correct++
		}
	}
	return correct
}

func runIncrement(program []int, a, c, b, pos int) (int, []int) {
	inc := 1
	for i := 0; i < pos; i+=2 {
		inc *= 10
	}
	out := []int{}

	for a!= 0 {
		out = run(program, a, b, c)
		if out[pos] == program[pos] {
			break
		}
		if len(out) != len(program) {
			panic("overshot")
		}
		a += inc
	}
	return a, out
}

func run(program []int, regA, regB, regC int) []int {
	output := []int{}

	for i := 0; i < len(program); i+=2 {
		operand := program[i+1]
		combo := operand


		switch operand {
		case 4:
			combo = regA
		case 5:
			combo = regB
		case 6:
			combo = regC
		}

		switch program[i] {
		case 0:
			regA >>= combo
		case 1:
			regB = regB ^ operand 
		case 2:
			regB = combo % 8
		case 3:
			if regA == 0 {
				break
			}
			i = operand-2
		case 4:
			regB = regB ^ regC
		case 5:
			output = append(output, combo % 8)
		case 6:
			regB = regA >> combo
		case 7:
			regC = regA >> combo
		default:
			panic(fmt.Errorf("Ya dun goofed, this is not an opcode %d at index %d", program[i], i))
		}
	}

	return output
}

func parseData(data []string) (int, int, int, []int) {
	aStr := data[0][12:]
	bStr := data[1][12:]
	cStr := data[2][12:]

	pStr := strings.Split(data[4][9:], ",")

	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)
	c, _ := strconv.Atoi(cStr)

	program := make([]int, len(pStr))
	for i, str := range pStr {
		num, _ := strconv.Atoi(str)
		program[i] = num
	}

	return a, b, c, program
}
