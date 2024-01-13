package main

import (
	"fmt"

	"github.com/ytx-readings/reading-lists/books"
)

func main() {
	for i, item := range books.Books2024 {
		fmt.Printf("[%d] %v\n", i, item)
	}
}
