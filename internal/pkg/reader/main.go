package reader

import (
	"bufio"
	"os"
)

func FileToArray(fp string) (r []string) {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	return
}

func Reverse(s string) (r string) {
	for _, c := range s {
		r = string(c) + r
	}
	return
}
