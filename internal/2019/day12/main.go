package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type moon struct {
	pos [3]int
	vel [3]int
}

func main() {
	data := reader.FileToArray("data/2019/day12.txt")

	//part 1
	moons := parseMoons(data)
	for s := 0; s < 1000; s++ {
		timeStep(moons)
	}
	energy := calculateEnergy(moons)
	fmt.Println("First:", energy)

	//part 2
	moons = parseMoons(data)
	lengths := map[int]bool{}
	cycles := findCycles(moons)
	for _, cycle := range cycles {
		lengths[len(cycle)] = true
	}

	nums := []int{}
	for l := range lengths {
		nums = append(nums, l)
	}
	fmt.Println(nums)
	result := lcm(nums[0], nums[1], nums[2:]...)
	fmt.Println("Second:", result)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func findCycles(raw []moon) [][]int {
	moons := copyState(raw)

	cycles := [][]int{}
	for _, moon := range moons {
		cycles = append(cycles, 
			[]int{moon.pos[0]}, 
			[]int{moon.pos[1]}, 
			[]int{moon.pos[2]}, 
			[]int{moon.vel[0]}, 
			[]int{moon.vel[1]}, 
			[]int{moon.vel[2]}, 
		)
	}

	repeatCache := map[int]int{}

	done := [][]int{}
	doneIdx := make([]bool, len(cycles))

	for {
		if len(done) == len(cycles) {
			break
		}

		timeStep(moons)

		for i := range cycles {
			if doneIdx[i] {
				continue
			}
			moonIdx := i / 6
			val := 0
			switch i % 6 {
			case 0:
				val = moons[moonIdx].pos[0]
			case 1:
				val = moons[moonIdx].pos[1]
			case 2:
				val = moons[moonIdx].pos[2]
			case 3:
				val = moons[moonIdx].vel[0]
			case 4:
				val = moons[moonIdx].vel[1]
			case 5:
				val = moons[moonIdx].vel[2]
			}
			cycles[i] = append(cycles[i], val)
			
			if seq, found := repeats(cycles[i], repeatCache[i]); found {
				done = append(done, seq)
				doneIdx[i] = true
				fmt.Println(len(done), "out of", len(cycles), "found")
			} else {
				repeatCache[i] = len(cycles[i])/2
			}
		}
	}

	return done
}

func repeats(s []int, gap int) ([]int, bool) {
	if len(s) % 2 == 1 {
		return nil, false
	}
	if gap == 0 {
		gap = 2
	}
	loop:
	for l := gap; l < len(s)/2; l++ {
		for i := 0; i <= l; i++ {
			if s[i] != s[i+l] {
				continue loop
			}
		}
		return s[:l], true
	}
	return nil, false 
}

func calculateEnergy(moons []moon) int {
	total := 0.0

	for _, moon := range moons {
		posTotal := 0.0
		velTotal := 0.0
		for i := range moon.pos {
			posTotal += math.Abs(float64(moon.pos[i]))
			velTotal += math.Abs(float64(moon.vel[i]))
		}
		total += posTotal * velTotal
	}

	return int(total)
}

func timeStep(moons []moon) {
	for i := range moons {
		for j := range moons[i:] {
			moons[i].applyGravity(&moons[i+j])
		}
	}
	for i := range moons {
		for idx := range moons[i].pos {
			moons[i].pos[idx] += moons[i].vel[idx]
		}
	}
}

func (m *moon) applyGravity(fren *moon) {
	for i := range m.pos {
		if m.pos[i] == fren.pos[i] {
			continue
		}
		if m.pos[i] < fren.pos[i] {
			m.vel[i]++
			fren.vel[i]--
			continue
		}
		m.vel[i]--
		fren.vel[i]++
	}
}

func parseMoons(data []string) []moon {
	result := []moon{}
	for _, line := range data {
		result = append(result, parseMoon(line))
	}
	return result
}

func parseMoon(line string) moon {
	line = strings.Trim(line, "<>")
	parts := strings.Split(line, ", ")

	xStr := parts[0][2:]
	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}
	yStr := parts[1][2:]
	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}
	zStr := parts[2][2:]
	z, err := strconv.Atoi(zStr)
	if err != nil {
		panic(err)
	}

	return moon{pos: [3]int{x, y, z}} 
}

func copyState(moons []moon) []moon {
	result := []moon{}
	for _, m := range moons {
		temp := moon{
			pos: [3]int{m.pos[0], m.pos[1], m.pos[2]},
			vel: [3]int{m.vel[0], m.vel[1], m.vel[2]},
		}
		result = append(result, temp)
	}
	return result
}
