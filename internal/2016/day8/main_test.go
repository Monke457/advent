package main 

import "testing"

func TestShiftRow(t *testing.T) {
	screen := screen{ 
		cells: [6][50]bool{{true, false, true, false, false}},
	}

	screen = screen.ShiftRow([2]int{0, 1})

	expected := [5]bool{false, true, false, true, false}

	for i := range expected {
		if expected[i] != screen.cells[0][i] {
			t.Errorf("result not as expected\n")
			t.Errorf("got: %v\n", screen.cells[0][:5])
			t.Errorf("expected: %v\n", expected)
			break
		}
	}
}
