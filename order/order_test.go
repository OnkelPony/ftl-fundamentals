package order_test

import (
	"order"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := ""
	_, err := order.New(want)
	if err == nil {
		t.Error("wanted error for invalid id, got nil")
	}

	want = "666108"
	order, err := order.New(want)
	if err != nil {
		t.Error(err)
	}
	got := order.ID()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestSetID(t *testing.T) {
	t.Parallel()
	want := "666108"
	order, err := order.New(want)
	if err != nil {
		t.Error(err)
	}
	want = ""
	err = order.SetID(want)
	if err == nil {
		t.Error("wanted error for invalid id, got nil")
	}
	want = "108666"
	err = order.SetID(want)
	if err != nil {
		t.Error(err)
	}
	got := order.ID()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
