package main

import "testing"

func TestMult(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		var a, b, result int = 2, 2, 4

		realResult := Mult(a, b)

		if realResult != result {
			t.Errorf("Error in Mult(a, b): %d != %d", result, realResult)
		}
	})

	t.Run("medium", func(t *testing.T) {
		var a, b, result int = 222, 222, 49284

		realResult := Mult(a, b)

		if realResult != result {
			t.Errorf("Error in Mult(a, b): %d != %d", result, realResult)
		}
	})

	t.Run("negative", func(t *testing.T) {
		var a, b, result int = -2, 4, -8

		realResult := Mult(a, b)

		if realResult != result {
			t.Errorf("Error in Mult(a, b): %d != %d", result, realResult)
		}
	})
}
