package main

import (
	"fmt"

	"github.com/arjun256900/go-product-search-api/internal/product"
)

func main() {
	products := product.GenerateProducts(100000)
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Category: %s, Keywords: %v\n", p.ID, p.Name, p.Category, p.Keywords)
	}
}
