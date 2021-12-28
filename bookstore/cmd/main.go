package main

import "fmt"

func main() {
	type Book struct {
		Title        string
		Author       string
		Count        int
		SeriesNumber int
	}
	b := Book{
		Title:        "Ikona",
		Author:       "Frederick Forsyth",
		Count:        108,
		SeriesNumber: 1,
	}
	fmt.Println(b)
}
