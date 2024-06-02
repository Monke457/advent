package main

import "testing"

func TestIsReal(t *testing.T) {
	rooms := []Room{
		{name: "aaaa-bbb-cc-xzy", checksum: "abcxyz"},
		{name: "eee-ddd-fff-fgj", checksum: "fedgj"},
		{name: "eeeee-ddddd-fff-fagj", checksum: "degja"},
	}

	if !rooms[0].IsReal() {
		//this one should be real
		t.Error("room should be real", rooms[0])
	}

	if rooms[1].IsReal() {
		//this one should not be real
		t.Error("room should not be real", rooms[1])
	}

	if rooms[2].IsReal() {
		//this one should not be real
		t.Error("room should not be real", rooms[2])
	}
}
