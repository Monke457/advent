package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

func main() {
	data := reader.FileToString("data/2022/day6.txt")

	mPacket := -1
	mMessage := -1
	for i := 0; i < len(data)-4; i++ {
		if mPacket == -1 {
			chunk := data[i:i+4]
			if allDiff(chunk) {
				mPacket = i+4
			}
		}

		if i >= len(data)-14 {
			continue
		}

		if mMessage == -1 {
			chunk := data[i:i+14]
			if allDiff(chunk) {
				mMessage = i+14
			}
		}
	}
	fmt.Println("Packet marker:", mPacket)
	fmt.Println("Message marker:", mMessage)
}

func allDiff(data string) bool {
	for i := 0; i < len(data); i++ {
		for j := i+1; j < len(data); j++ {
			if data[i] == data[j] {
				return false
			}
		}
	}
	return true
}
