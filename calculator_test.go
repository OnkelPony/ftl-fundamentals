package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	tcs := []testCase{
		{1, 2, 3},
		{4, 5, 9},
		{-7, -8, -15},
		{108, 666, 774},
	}
	for _, tc := range tcs {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%.1f, %.1f) = %.1f, want %.1f", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	var a, b float64 = 6, 2
	got := calculator.Subtract(a, b)
	if want != got {
		t.Errorf("Subtract(%.1f, %.1f) = %.1f, want %.1f", a, b, got, want)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 12
	var a, b float64 = 6, 2
	got := calculator.Multiply(a, b)
	if want != got {
		t.Errorf("Multiply(%.1f, %.1f) = %.1f, want %.1f", a, b, got, want)
	}
}
