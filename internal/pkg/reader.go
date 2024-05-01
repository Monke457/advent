package reader

import (
	"bufio"
	"os"
)

func fileToArray(fp string) []string {
	data, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
