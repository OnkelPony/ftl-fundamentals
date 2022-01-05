package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"testing"
)

var catalog = bookstore.Catalog{
	"1": {
		ID:           "1",
		Title:        "Ikona",
		Author:       "Frederick Forsyth",
		Count:        108,
		SeriesNumber: 1,
	},
	"2": {
		ID:           "2",
		Title:        "Křik Halidonu",
		Author:       "Robert Ludlum",
		Count:        6,
		SeriesNumber: 2,
	},
}

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		ID:           "1",
		Title:        "Ikona",
		Author:       "Frederick Forsyth",
		Count:        108,
		SeriesNumber: 1,
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	want := []bookstore.Book{
		{
			ID:           "1",
			Title:        "Ikona",
			Author:       "Frederick Forsyth",
			Count:        108,
			SeriesNumber: 1,
		},
		{
			ID:           "2",
			Title:        "Křik Halidonu",
			Author:       "Robert Ludlum",
			Count:        6,
			SeriesNumber: 9,
		},
		{
			ID:           "3",
			Title:        "Špion vypovídá",
			Author:       "Josef Frolík",
			Count:        2,
			SeriesNumber: 1,
		},
	}
	got := catalog.GetAllBooks()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()
	want := "Ikona od Frederick Forsyth"
	got := catalog.GetBookDetails("1")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNetPrice(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Jak se věci mají",
		PriceCents: 1000,
	}
	want := 730
	got := b.NetPrice()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSalePrice(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Buddha a láska",
		Author:     "Ole Nydahl",
		PriceCents: 1080,
	}
	want := 540
	got := b.SalePrice()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "O smrti a znovuzrození",
		Author:     "Ole Nydahl",
		PriceCents: 666,
	}
	b.SetPriceCents(1080)
	want := 1080
	got := b.PriceCents
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddBook(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Buddha a láska",
		Author:     "Ole Nydahl",
		PriceCents: 1080,
	}
	id := "108"
	c := bookstore.Catalog{}
	c.AddBook(id, b)
	want := "Buddha a láska od Ole Nydahl"
	got := c.GetBookDetails(id)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetDiscountPercent(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Ngöndro",
		PriceCents: 1080,
	}
	err := b.SetDiscountPercent(108)
	if err == nil {
		t.Error("SetDiscountPercent should have returned error, but it didn't")
	}
	err = b.SetDiscountPercent(6)
	want := 6
	got := b.GetDiscountPercent()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
