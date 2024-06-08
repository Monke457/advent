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
