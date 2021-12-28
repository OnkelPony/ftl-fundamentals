package calculator_test

import (
	"calculator"
	"testing"
)

func TestAddSubMul(t *testing.T) {
	t.Parallel()
	type testCase struct {
		fn   func(float64, float64) float64
		a, b float64
		want float64
	}
	tcs := []testCase{
		{calculator.Add, 1, 2, 3},
		{calculator.Add, 4, 5, 9},
		{calculator.Add, -7, -8, -15},
		{calculator.Add, 108, 666, 774},
		{calculator.Subtract, 1, 2, -1},
		{calculator.Subtract, 8, -2, 10},
		{calculator.Multiply, 4, 5, 20},
		{calculator.Multiply, -7, -8, 56},
	}
	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("AddSubMul(%.1f, %.1f) = %.1f, want %.1f", tc.a, tc.b, got, tc.want)
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

func TestDivide(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		a, b        float64
		want        float64
		errExpected bool
	}{
		{a: 108, b: 4, want: 27, errExpected: false},
		{a: 666, b: 0, want: 999, errExpected: true},
	}
	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("\"Divide(%.1f, %.1f): unexpected error status: %v", tc.a, tc.b, err)
		}
		if !errReceived && tc.want != got {
			t.Errorf("Divide(%.1f, %.1f) = %.1f, want %.1f", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		a           float64
		want        float64
		errExpected bool
	}{
		{a: 64, want: 8, errExpected: false},
		{a: -64, want: 666, errExpected: true},
	}
	for _, tc := range tcs {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("\"Sqrt(%.1f): unexpected error status: %v", tc.a, err)
		}
		if !errReceived && tc.want != got {
			t.Errorf("Sqrt(%.1f) = %.1f, want %.1f", tc.a, got, tc.want)
		}
	}
}
