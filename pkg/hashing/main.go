package hashing 

func ToASCIIValues(str string) []int {
	res := []int{}
	for _, r := range str {
		res = append(res, int(r))
	}
	return res
}

func MakeRange(a, b int) []int {
	res := []int{}
	for i := a; i < b; i++ {
		res = append(res, i)
	}
	return res
}

func DenseHash(values []int) []int {
	res := []int{}
	for i := 0; i < len(values); i += 16 {
		var hash int
		for _, v := range values[i:i+16] {
			hash ^= v
		}
		res = append(res, hash)
	}
	return res
}

func KnotHash(values, lengths []int, rounds int) []int {
	var pos, skip int
	res := make([]int, len(values)) 
	copy(res, values)
	for range rounds {
		for _, l := range lengths {
			if l > 1 {
				rev := []int{} 

				for i := 0; i < l; i++ {
					n := (pos + i) % len(res)
					rev = append(rev, res[n])
				}
				revPos := 0
				for i := l-1; i >= 0; i-- {
					n := (pos + i) % len(res)
					res[n] = rev[revPos]
					revPos++
				}
			}
			pos += (skip + l) % len(res)
			skip++
		}
	}
	return res
}
