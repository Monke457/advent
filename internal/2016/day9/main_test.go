package main

import "testing"

func TestFindCompLen(t *testing.T) {
	data := []string{
		"ADVENT",
		"A(1x5)BC",
		"(3x3)XYZ",
		"A(2x2)BCD(2x2)EFG",
		"(6x1)(1x3)A",
		"X(8x2)(3x3)ABCY",
	}

	l1 := findCompLen(data[0])
	if l1 != 6 {
		t.Errorf("Decompression failed %s, expected 6 got %d", data[0], l1)
	} else {
		t.Log("Passed", data[0], ":", l1)
	}
	l2 := findCompLen(data[1])
	if l2 != 7 {
		t.Errorf("Decompression failed, %s expected 7 got %d", data[1], l2)
	} else {
		t.Log("Passed", data[1], ":", l2)
	}
	l3 := findCompLen(data[2])
	if l3 != 9 {
		t.Errorf("Decompression failed, %s expected 9 got %d", data[2], l3)
	} else {
		t.Log("Passed", data[2], ":", l3)
	}
	l4 := findCompLen(data[3])
	if l4 != 11 {
		t.Errorf("Decompression failed, %s expected 11 got %d", data[3], l4)
	} else {
		t.Log("Passed", data[3], ":", l4)
	}
	l5 := findCompLen(data[4])
	if l5 != 6 {
		t.Errorf("Decompression failed, %s expected 6 got %d", data[4], l5)
	} else {
		t.Log("Passed", data[4], ":", l5)
	}
	l6 := findCompLen(data[5])
	if l6 != 18 {
		t.Errorf("Decompression failed, %s expected 18 got %d", data[5], l6)
	} else {
		t.Log("Passed", data[5], ":", l6)
	}
}

func TestFindCompLenRec(t *testing.T) {
	data := []string{
		"(3x3)XYZ",
		"X(8x2)(3x3)ABCY",
		"(27x12)(20x12)(13x14)(7x10)(1x12)A",
		"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
	}

	l1 := findCompLenRec(data[0])
	if l1 != 9 {
		t.Errorf("Decompression failed %s, expected 9 got %d", data[0], l1)
	} else {
		t.Log("Passed", data[0], ":", l1)
	}
	l2 := findCompLenRec(data[1])
	if l2 != 20 {
		t.Errorf("Decompression failed, %s expected 20 got %d", data[1], l2)
	} else {
		t.Log("Passed", data[1], ":", l2)
	}
	l3 := findCompLenRec(data[2])
	if l3 != 241920 {
		t.Errorf("Decompression failed, %s expected 241920 got %d", data[2], l3)
	} else {
		t.Log("Passed", data[2], ":", l3)
	}
	l4 := findCompLenRec(data[3])
	if l4 != 445 {
		t.Errorf("Decompression failed, %s expected 445 got %d", data[3], l4)
	} else {
		t.Log("Passed", data[3], ":", l4)
	}
}
