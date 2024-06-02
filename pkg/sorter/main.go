package sorter

import (
	"cmp"
	"slices"
)

func GetOrderByKey[O cmp.Ordered, T any](m map[O]T) []O {
	keys := []O{}
	for k := range m {
		if slices.Contains(keys, k) {
			continue
		}
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b O) int {
		return cmp.Compare(a, b)
	})
	return keys
}

func GetOrderByValue[O cmp.Ordered, T cmp.Ordered](m map[O]T) []O {
	keys := []O{}
	for k := range m {
		if slices.Contains(keys, k) {
			continue
		}
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b O) int {
		if m[a] < m[b] {
			return 1
		}
		if m[a] > m[b] {
			return -1
		}
		return cmp.Compare(a, b)
	})
	return keys
}
