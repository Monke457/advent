package intcode

import (
	"fmt"
	"strings"
)

type Mode int

type Status int

const (
	parameter Mode = iota
	position
)

const (
	starting Status = iota
	running
	paused	
	Halted
)

type Computer struct {
	Data []int
	Status Status
	Index int
	mode Mode
	phase int
	Input int
	phasemode bool
	relativeBase int
}

func NewComputer(rawdata []int, phase, input int, phasemode bool) Computer {
	data := make([]int, len(rawdata))
	copy(data, rawdata)
	//add buffer
	data = append(data, make([]int, 1000)...)
	return Computer {
		Data: data, 
		Index: 0, 
		Status: starting, 
		phase: phase, 
		Input: input, 
		phasemode: phasemode,
	}
}

func (c Computer) Error(msg string) error {
	pre := "\nError:" 
	divider := strings.Repeat("-", 20)
	return fmt.Errorf(
		"\n%s%s%s%s %s\n%s\n\n", 
		divider, c.Sprint(), divider, pre, msg, divider,
	)
}

func (c Computer) Sprint() string {
	str := fmt.Sprintf(
		"\nCOMPUTER\nData: %v\nIndex: %d\nData at Index: %d\nMode: %d\n Status: %d\nPhase: %d\nInput: %d\nIn phase mode: %v\nRelative base: %d\n", 
		c.Data, c.Index, c.Data[c.Index], c.mode, c.Status, c.phase, c.Input, c.phasemode, c.relativeBase,
	)
	return str 
}
