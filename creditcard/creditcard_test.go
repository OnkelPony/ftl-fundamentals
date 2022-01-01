package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNewValid(t *testing.T) {
	t.Parallel()
	want := "666108"
	card, err := creditcard.New(want)
	if err != nil {
		t.Error(err)
	}
	got := card.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()
	want := ""
	_, err := creditcard.New(want)
	if err == nil {
		t.Error("want error from invalid card number, got nil")
	}
}

func TestSetNumber(t *testing.T) {
	t.Parallel()
	want := "666108"
	card, err := creditcard.New(want)
	if err != nil {
		t.Error(err)
	}
	want = ""
	err = card.SetNumber(want)
	if err == nil {
		t.Error("want error from invalid card number, got nil")
	}
	want = "108108"
	err = card.SetNumber(want)
	if err != nil {
		t.Error(err)
	}
	got := card.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
