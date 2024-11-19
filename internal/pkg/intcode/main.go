package intcode

import "fmt"

func (c *Computer) Run(output chan int, done chan bool) {
	if c.Status == Halted {
		fmt.Print(c.Error("Attempting to run a computer than has been shutdown"))
		return 
	}
	c.Status = running
	for {
		c.setMode()
		op := c.getOpCode()
		switch op {
		case 1: 
			if c.mode == Position {
				c.positionOperation(c.add)
			} else {
				c.add(c.getParams())
			}
			c.Index += 4
		case 2:
			if c.mode == Position {
				c.positionOperation(c.multiply)
			} else {
				c.multiply(c.getParams())
			}
			c.Index += 4
		case 3: 
			c.store()
			c.Index += 2
		case 4:
			output<-c.getOutput()
			c.Index += 2
			c.Status = paused
		case 5:
			if c.mode == Position {
				c.jump(true)
			} else {
				c.paramJump(c.getParams(), true)
			}
		case 6:
			if c.mode == Position {
				c.jump(false)
			} else {
				c.paramJump(c.getParams(), false)
			}
		case 7:
			c.less(c.getParams())
			c.Index += 4
		case 8:
			c.equal(c.getParams())
			c.Index += 4
		case 99:
			c.Status = Halted
			done<-true
		}
		if c.Status != running {
			break
		}
	}
}

func (c Computer) getOpCode() int {
	code := c.Data[c.Index]
	if c.mode == Position {
		return code
	}
	return code % 10
}

func (c *Computer) setMode() {
	if c.Data[c.Index] > 100 {
		c.mode = Parameter
		return
	}
	c.mode = Position
}

func (c Computer) getParams() [3]int {
	code := c.Data[c.Index]
	params := [3]int{}
	code /= 100
	for idx := 0; idx < 3; idx++ {
		if code % 10 == 0 {
			params[idx] = c.Data[c.Index+1+idx]
		} else {
			params[idx] = c.Index+1+idx
		}
		code /= 10
	}
	return params 
}

func (c Computer) positionOperation(operation func([3]int)) {
	params := [3]int{
		c.Data[c.Index+1],
		c.Data[c.Index+2],
		c.Data[c.Index+3],
	}
	operation(params)
}

func (c Computer) add(params [3]int) {
	c.Data[params[2]] = c.Data[params[0]] + c.Data[params[1]]
}

func (c Computer) multiply(params [3]int) {
	c.Data[params[2]] = c.Data[params[0]] * c.Data[params[1]]
}

func (c *Computer) store() {
	position := c.Data[c.Index+1]
	if c.phasemode {
		c.Data[position] = c.phase
		c.phasemode = false
		return
	}
	c.Data[position] = c.Input
}

func (c Computer) getOutput() int {
	var pos int
	if c.mode == Position {
		pos = c.Data[c.Index+1]
	} else {
		pos = c.getParams()[0]
	}
	return c.Data[pos]
}

func (c *Computer) jump(condition bool) {
	pos := c.Data[c.Index+1]
	if condition && c.Data[pos] != 0 {
		c.Index = c.Data[c.Index+2] 
		return
	}
	if !condition && c.Data[pos] == 0 {
		c.Index = c.Data[c.Index+2]
		return
	}
	c.Index += 3
}

func (c *Computer) paramJump(params [3]int, condition bool) {
	val := c.Data[params[0]]
	if (condition && val != 0) || (!condition && val == 0) {
		c.Index = c.Data[params[1]] 
		return
	}
	c.Index += 3
}

func (c Computer) less(params [3]int) {
	if c.Data[params[0]] < c.Data[params[1]] {
		c.Data[params[2]] = 1
		return
	}
	c.Data[params[2]] = 0
}

func (c Computer) equal(params [3]int) {
	if c.Data[params[0]] == c.Data[params[1]] {
		c.Data[params[2]] = 1
		return
	}
	c.Data[params[2]] = 0
}

func (c Computer) Shutdown() {
	c.Status = Halted
}
