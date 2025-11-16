package main

import(
	"advent/internal/pkg/reader"
	"fmt"
)

type track struct {
	route  []string
	height int
	width int
}

type cart struct {
	dir rune 
	turnIdx int
	ticks int
}

// turn indicies
// 0 = left
// 1 = straight
// 2 = right

func main() {
	data := reader.FileToArray("data/2018/day13.txt")

	track, carts := initializeData(data)
	var crash *[2]int
	tick := 0
	for {
		track, carts, crash = tickround(track, carts, tick)
		tick++
		if crash != nil {
			break
		}
		if len(carts) <= 1 {
			break
		}
	}
	fmt.Println(carts)
}

func tickround(track track, carts map[[2]int]cart, tick int) (track, map[[2]int]cart, *[2]int) {
	for y := range track.height {
		for x := range track.width {
			pos := [2]int{y,x}
			if cart, ok := carts[pos]; ok && cart.ticks == tick {
				coords := cart.tick(track, y, x)
				delete(carts, pos)
				cart.ticks++
				if _, ok := carts[coords]; ok {
					delete(carts, coords)
					fmt.Println(len(carts), "crash", coords)
				} else {
					carts[coords] = cart
				}
			}
		}
	}
	return track, carts, nil
}

func (t track) print(carts map[[2]int]cart, crash *[2]int) {
	fmt.Println()
	var pos [2]int
	for i := range t.route {
		pos[0] = i
		for j := range t.route[i] {
			pos[1] = j
			if crash != nil && *crash == pos {
				fmt.Printf("X")
			} else if cart, ok := carts[pos]; ok {
				fmt.Printf("%c", cart.dir)
			} else {
				fmt.Printf("%c",t.route[i][j])
			}
		}
		fmt.Println()
	}
}

func (c *cart) tick(t track, y, x int) [2]int {
	switch t.route[y][x] {
	case '+':
		c.turn()
	case '\\':
		c.curveLeft()
	case '/':
		c.curveRight()
	case '-', '|': break
	default: panic("this should not be happening...")
	}
	return [2]int{c.moveV(y),c.moveH(x)}
}

func (c cart) moveV(y int) int {
	if c.dir == 'v' { return y + 1 }
	if c.dir == '^' { return y - 1 }
	return y
}

func (c cart) moveH(x int) int {
	if c.dir == '>' { return x + 1 } 
	if c.dir == '<' { return x - 1 }
	return x
}

func (c *cart) curveLeft() {
	switch c.dir {
	case '>': c.dir = 'v'
	case 'v': c.dir = '>'
	case '<': c.dir = '^'
	case '^': c.dir = '<'
	}
}

func (c *cart) curveRight() {
	switch c.dir {
	case '>': c.dir = '^'
	case 'v': c.dir = '<'
	case '<': c.dir = 'v'
	case '^': c.dir = '>'
	}
}

func (c *cart) turn() {
	if c.turnIdx == 0 {
		switch c.dir {
		case '>': c.dir = '^'
		case 'v': c.dir = '>'
		case '<': c.dir = 'v'
		case '^': c.dir = '<'
		}
		c.turnIdx = 1

	} else if c.turnIdx == 2 {
		switch c.dir {
		case '>': c.dir = 'v'
		case 'v': c.dir = '<'
		case '<': c.dir = '^'
		case '^': c.dir = '>'
		}
		c.turnIdx = 0

	} else { c.turnIdx = 2 }
}

func initializeData(data []string) (track, map[[2]int]cart) {
	height := len(data)
	route := make([]string,height)
	carts := map[[2]int]cart{}

	var width int
	for y := range data {
		for x, r := range data[y] {
			if r == '>' || r == 'v' || r == '<' || r == '^' {
				carts[[2]int{y,x}] = cart{r, 0, 0}
				if r == '>' || r == '<'  {
					route[y] += "-"
				} else {
					route[y] += "|"
				}
			} else {
				route[y] += string(r)
			}
		}
		if len(route[y]) > width {
			width = len(route[y])
		}
	}
	return track{route, height, width}, carts
}
