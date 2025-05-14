package product

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arjun256900/go-product-search-api/internal/models"
	"github.com/go-faker/faker/v4"
)

// // Helper function to generate 1M products
func GenerateProducts(n int) []models.Product {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	products := make([]models.Product, n)

	// Categories (better practice as an enum)
	categories := []string{"Electronics", "Footwear", "Home & Kitchen", "Fitness", "Books", "Toys"}

	// Generate n products
	for i := 0; i < n; i++ {
		// Generating random product names using go-faker
		name := faker.Word() + " " + faker.Word() // Combinining two random words for the product name

		// Generate random keywords
		keywords := []string{"prod", fmt.Sprintf("%d", rand.Intn(1000))}

		// Assign values to the product
		products[i] = models.Product{
			ID:       i + 1,
			Name:     name,
			Category: categories[rand.Intn(len(categories))],
			Keywords: keywords,
		}
	}

	return products
}
