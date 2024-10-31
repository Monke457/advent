package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToIntArrayByComma("data/2019/day2.txt")

	fmt.Println("Part 1:", runProgram(data, 12, 2))

	var alarmcode int
	target := 19690720
	loop:
	for noun := 0; noun < 100; noun++ {	
		for verb := 0; verb < 100; verb++ {
			if runProgram(data, noun, verb)== target {
				alarmcode = 100 * noun + verb
				break loop
			}
		}
	}
	fmt.Println("Part 2:", alarmcode)
}

func runProgram(raw []int, noun, verb int) int {
	data := make([]int, len(raw))
	copy(data, raw)
	data[1] = noun 
	data[2] = verb

	i := 0
	loop:
	for {
		var result int
		pos1 := data[i+1]
		pos2 := data[i+2]
		pos3 := data[i+3]

		switch data[i] {
		case 1: 
			result = data[pos1] + data[pos2]
		case 2: 
			result = data[pos1] * data[pos2]
		case 99:
			break loop
		}
		data[pos3] = result

		i += 4
	}
	return data[0]
}
