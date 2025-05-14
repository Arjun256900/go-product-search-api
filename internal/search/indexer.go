package search

import (
	"fmt"
	"log"

	"github.com/arjun256900/go-product-search-api/internal/models"
	"github.com/arjun256900/go-product-search-api/internal/product"
	"github.com/arjun256900/go-product-search-api/internal/search/bleve"
)

// Service ties together the Bleve indexer and in‑memory store.
type Service struct {
	index bleve.Indexer
	store []models.Product
}

// NewService generates products, creates the indexer, and indexes them.
func NewService(n int) (*Service, error) {
	// Generate products
	products := product.GenerateProducts(n) 

	// Initialize Bleve indexer
	idx, err := bleve.NewBleveIndexer()
	if err != nil {
		return nil, fmt.Errorf("init bleve indexer: %w", err)
	}

	// Index all products
	for _, p := range products {
		if err := idx.IndexProduct(p); err != nil {
			log.Printf("failed to index product ID %d: %v", p.ID, err)
		}
	}

	return &Service{index: idx, store: products}, nil
}

// Search runs a full‑text query and maps hit IDs back to full Products.
func (s *Service) Search(query string, size int) ([]models.Product, error) {
	ids, err := s.index.Search(query, size)
	if err != nil {
		return nil, err
	}

	// Map IDs → Product objects
	results := make([]models.Product, 0, len(ids))
	for _, idStr := range ids {
		for _, p := range s.store {
			if fmt.Sprintf("%d", p.ID) == idStr {
				results = append(results, p)
				break
			}
		}
	}
	return results, nil
}
