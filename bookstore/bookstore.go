package bookstore

import "fmt"

type Book struct {
	ID              string
	Title           string
	Author          string
	Count           int
	SeriesNumber    int
	PriceCents      int
	DiscountPercent int
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
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b Book) SalePrice() int {
	return b.PriceCents / 2
}
