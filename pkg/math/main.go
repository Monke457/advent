package math

import "math"

func Max(np []*int) int {
	m := int(math.Inf(-1))
	for _, v := range np {
		if v == nil {
			continue
		}
		if *v > m {
			m = *v
		}
	}
	return m
}

func Factorial(n int) int {
	if n <= 0 {
		return 0
	}
	return n + Factorial(n-1)
}

func AbsDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff *= -1
	}
	return diff
}

func LCDM(a, b int) int {
	return a * b / GCD(a, b)
}

func LCD(args ...int) int {
	if len(args) < 2 {
		panic("not enough args")
	}
	if len(args) == 2 {
		return LCDM(args[0], args[1])
	} else {
		return LCDM(args[0], LCD(args[1:]...))
	}
}

func GCD(a, b int) int {
	c := 0
	if a&1 == 0 && b&1 == 0 {
		for a&1 == 0 && b&1 == 0 {
			a = a >> 1
			b = b >> 1
			c++
		}
	}
	if a&1 == 0 {
		for a&1 == 0 {
			a = a >> 1
		}
	}
	if b&1 == 0 {
		for b&1 == 0 {
			b = b >> 1
		}
	}
	for a != b {
		if a > b {
			a = (a - b) >> 1
			for a&1 == 0 {
				if a == b {
					break
				}
				a = a >> 1
			}
		} else {
			b = (b - a) >> 1
			for b&1 == 0 {
				if a == b {
					break
				}
				b = b >> 1
			}
		}
	}
	return 1 << c * a
}
