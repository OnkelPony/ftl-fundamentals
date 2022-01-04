package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAddSubMul(t *testing.T) {
	t.Parallel()
	type testCase struct {
		fn      func(...float64) float64
		a, b, c float64
		want    float64
	}
	tcs := []testCase{
		{calculator.Add, 1, 2, 3, 6},
		{calculator.Add, 4, 5, 9, 18},
		{calculator.Add, -7, -8, -15, -30},
		{calculator.Add, 108, 666, 774, 1548},
		{calculator.Subtract, 1, 2, -1, 0},
		{calculator.Subtract, 8, -2, 1, 9},
		{calculator.Multiply, 4, 5, 20, 400},
		{calculator.Multiply, -7, -8, 5, 280},
	}
	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b, tc.c)
		if tc.want != got {
			t.Errorf("AddSubMul(%.1f, %.1f, %.1f) = %.1f, want %.1f", tc.a, tc.b, tc.c, got, tc.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 13
	var a, b, c float64 = 6, 2, -9
	got := calculator.Subtract(a, b, c)
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
		a, b, c     float64
		want        float64
		errExpected bool
	}{
		{a: 108, b: 4, c: 3, want: 9, errExpected: false},
		{a: 666, b: 0, c: 8, want: 999, errExpected: true},
		{666, 66.6, 0, 999, true},
	}
	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b, tc.c)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("\"Divide(%.1f, %.1f, %.1f): unexpected error status: %v", tc.a, tc.b, tc.c, err)
		}
		if !errReceived && tc.want != got {
			t.Errorf("Divide(%.1f, %.1f, %.1f) = %.1f, want %.1f", tc.a, tc.b, tc.c, got, tc.want)
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
		{a: 108, want: 10.39, errExpected: false},
		{a: -64, want: 666, errExpected: true},
	}
	for _, tc := range tcs {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("\"Sqrt(%.1f): unexpected error status: %v", tc.a, err)
		}
		const tolerance = 0.108
		if !errReceived && !closeEnough(tc.want, got, tolerance) {
			t.Errorf("Sqrt(%.1f) = %.1f, want %.1f", tc.a, got, tc.want)
		}
	}
}

func closeEnough(a float64, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
