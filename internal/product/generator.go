package product

import (
	"math/rand"
	"time"

	"github.com/arjun256900/go-product-search-api/internal/models"
	"github.com/arjun256900/go-product-search-api/pkg/utils"
)

// GenerateProducts returns n synthetic products using two random adjectives as the name
func GenerateProducts(n int) []models.Product {
	rand.Seed(time.Now().UnixNano())
	products := make([]models.Product, n)
	categories := []string{"Electronics", "Footwear", "Home & Kitchen", "Fitness", "Books", "Toys"}

	for i := 0; i < n; i++ {
		adjectives := utils.Adjectives
		firstAdj := adjectives[rand.Intn(len(adjectives))]
		secondAdj := adjectives[rand.Intn(len(adjectives))]
		// Avoid duplicate adjectives (optional)
		for secondAdj == firstAdj {
			secondAdj = adjectives[rand.Intn(len(adjectives))]
		}
		name := firstAdj + " " + secondAdj
		products[i] = models.Product{
			ID:       i + 1,
			Name:     name,
			Category: categories[rand.Intn(len(categories))],
		}
	}

	return products
}
