package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	var a, b float64 = 2, 2
	got := calculator.Add(a, b)
	if want != got {
		t.Errorf("Add(%.1f, %.1f) = %.1f, want %.1f", a, b, got, want)
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
