package main

import (
	"advent/internal/pkg/reader"
	"encoding/json"
	"fmt"
)

func main() {
	data := reader.FileToString("data/2015/day12.txt")

	ignore := "red"
	fmt.Println("First problem:", sumInts(data, nil))
	fmt.Println("Second problem:", sumInts(data, &ignore))
}

func sumInts(data string, ignore *string) int {
	res := []interface{}{}
	json.Unmarshal([]byte(data), &res)

	sum := sumArray(res, ignore)

	return sum
}

func sumArray(arr []interface{}, ignore *string) int {
	sum := 0
	for _, m := range arr {
		switch mv := m.(type) {
		case int:
			sum += mv
		case float64:
			sum += int(mv)
		case map[string]interface{}:
			sum += sumObject(mv, ignore)
		case []interface{}:
			sum += sumArray(mv, ignore)
		}
	}
	return sum
}

func sumObject(og map[string]interface{}, ignore *string) int { 
	sum := 0
	for key, value := range og {
		if ignore != nil && key == *ignore {
			return 0
		}
		switch v := value.(type) {
		case int:
			sum += v
		case float64:
			sum += int(v)
		case string:
			if ignore != nil && v == *ignore {
				return 0
			}
		case map[string]interface{}:
			sum += sumObject(v, ignore)
		case []interface{}:
			sum += sumArray(v, ignore)
		}
	}	
	return sum
}

