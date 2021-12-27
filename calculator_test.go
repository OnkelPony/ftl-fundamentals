package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	want := float64(4)
	var a, b float64 = 2, 2
	got := calculator.Add(a, b)
	if want != got {
		t.Errorf("Add(%.1f, %.1f) = %.1f, want %.1f", a, b, got, want)
	}
}
