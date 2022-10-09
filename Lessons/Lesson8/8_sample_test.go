package main

import "testing"

func TestMult(t *testing.T) {
	var a, b, result int = 2, 2, 4

	realResult := Mult(a, b)

	if realResult != result {
		t.Errorf("Error in Mult(a, b): %d != %d", result, realResult)
	}
}
