package intcode

import (
	"fmt"
	"strings"
)

func RunDay5(raw []int, input int) error {
	data := make([]int, len(raw))
	var ok bool 
	copy(data, raw)

	i := 0
	for {
		skip := 4
		if oob(len(data), i) {
			return fmterror(raw, input) 
		}
		if data[i] == 99 {
			break
		}
		if data[i] == 0 {
			skip = 1
		} else if data[i] > 4 {
			data, skip, ok = runParameterMode(data, i, input)
			if !ok {
				return fmterror(raw, input) 
			}
		} else {
			data, skip, ok = runPositionMode(data, i, input)
			if !ok {
				return fmterror(raw, input) 
			}
		}

		i += skip
	}
	return nil
}

func runParameterMode(data []int, i, input int) ([]int, int, bool) {
	code := data[i]
	params, ok := getParams(data, i, code)
	if !ok {
		return data, 0, false
	}

	de := code % 10
	switch de {
	case 0:
		return data, 1, true
	case 1: 
		data, ok = add(data, params)
		return data, 4, ok
	case 2:
		data, ok = multiply(data, params)
		return data, 4, ok
	case 3: 
		data, ok = store(data, params[0], input)
		return data, 2, ok
	case 4:
		ok = output(data, params[0])
		return data, 2, ok
	}

	return data, 0, false 
}

func getParams(data []int, i, code int) ([3]int, bool) {
	params := [3]int{}
	if oob(len(data), i+3) {
		return params, false 
	}
	code /= 100
	for idx := 0; idx < 3; idx++ {
		if code % 10 == 0 {
			params[idx] = data[i+1+idx]
		} else {
			params[idx] = i+1+idx
		}
		code /= 10
	}
	return params, true
}

func runPositionMode(data []int, i, input int) ([]int, int, bool) {
	var ok bool
	switch data[i] {
	case 1: 
		data, ok = positionOp(add, data, i)
		return data, 4, ok
	case 2:
		data, ok = positionOp(multiply, data, i)
		return data, 4, ok
	case 3: 
		data, ok = store(data, i+1, input)
		return data, 2, ok
	case 4:
		ok = output(data, i+1)
		return data, 2, ok
	}
	return data, 0, false
}

func store(data []int, i, input int) ([]int, bool) {
	if oob(len(data), i) {
		return nil, false 
	}
	params := data[i]
	if oob(len(data), params) {
		return nil, false 
	}
	data[params] = input
	return data, true
}

func output(data []int, i int) bool {
	if oob(len(data), i) {
		return false 
	}
	params := data[i]
	if oob(len(data), params) {
		return false 
	}
	fmt.Printf("Output: %d\n", data[params])
	return true
}

func positionOp(
	op func([]int, [3]int) ([]int, bool),
	data []int, i int) ([]int, bool) {

	if oob(len(data), i+1, i+2, i+3) {
		return nil, false 
	}
	params := [3]int{
		data[i+1],
		data[i+2],
		data[i+3],
	}

	return op(data, params)
}

func multiply(data []int, params [3]int) ([]int, bool) {
	if oob(len(data), params[0], params[1], params[2]) {
		return nil, false
	}
	data[params[2]] = data[params[0]] * data[params[1]]
	return data, true 
}

func add(data []int, params [3]int) ([]int, bool) {
	if oob(len(data), params[0], params[1], params[2]) {
		return nil, false
	}
	data[params[2]] = data[params[0]] + data[params[1]]
	return data, true 
}

func RunDay2(raw []int, noun, verb int) (int, error) {
	data := make([]int, len(raw))
	copy(data, raw)
	data[1] = noun 
	data[2] = verb

	i := 0
	loop:
	for {
		if oob(len(data), i, i+1, i+2, i+3) {
			return 0, fmterror(raw, noun, verb) 
		}
		var result int
		pos1 := data[i+1]
		pos2 := data[i+2]
		pos3 := data[i+3]
		if oob(len(data), pos1, pos2, pos3) {
			return 0, fmterror(raw, noun, verb) 
		}

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
	return data[0], nil
}

func oob(l int, nums... int) bool {
	for _, i := range nums {
		if l <= i || i < 0 { 
			return true
		}
	}
	return false
} 

func fmterror(raw []int, params... int) error {
	msg := "Error: went out of bounds - check the params"
	divider := strings.Repeat("-", 20)
	return fmt.Errorf(
		"%s\n%s\ndata: %v\nparams: %v\n%s\n\n", 
		msg, divider, raw, params, divider)
}

