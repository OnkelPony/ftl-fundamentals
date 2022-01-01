package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID              string
	Title           string
	Author          string
	Count           int
	SeriesNumber    int
	PriceCents      int
	discountPercent int
}

type Catalog map[string]Book

func (catalog Catalog) GetAllBooks() []Book {
	books := []Book{}
	for _, book := range catalog {
		books = append(books, book)
	}
	return books
}

func (catalog Catalog) GetBookDetails(ID string) string {
	b := catalog[ID]
	return fmt.Sprintf("%s od %s", b.Title, b.Author)
}

func (b Book) NetPrice() int {
	saving := b.PriceCents * b.discountPercent / 100
	return b.PriceCents - saving
}

func (b Book) SalePrice() int {
	return b.PriceCents / 2
}

func (b *Book) SetPriceCents(p int) {
	b.PriceCents = p
}

func (b *Book) SetDiscountPercent(p int) error {
	if p < 0 || p > 100 {
		return errors.New("discount must be between 0 and 100 percent")
	}
	b.discountPercent = p
	return nil
}

func (b Book) GetDiscountPercent() int {
	return b.discountPercent
}

func (catalog Catalog) AddBook(id string, book Book) {
	catalog[id] = book
}
