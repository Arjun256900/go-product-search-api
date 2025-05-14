package bleve

import (
	"fmt"

	"github.com/arjun256900/go-product-search-api/internal/models"
	blevepkg "github.com/blevesearch/bleve/v2"
)

type Indexer interface {
	IndexProduct(product models.Product) error
	Search(query string, size int) ([]string, error)
}

type BleveIndexer struct {
	index blevepkg.Index
}

func NewBleveIndexer() (Indexer, error) {
	mapping := blevepkg.NewIndexMapping()
	idx, err := blevepkg.NewMemOnly(mapping)
	if err != nil {
		return nil, err
	}
	return &BleveIndexer{index: idx}, nil
}

func (b *BleveIndexer) IndexProduct(product models.Product) error {
	id := fmt.Sprintf("%d", product.ID)

	// Only index the fields that are required while fetching
	doc := map[string]interface{}{
		"name":     product.Name,
		"category": product.Category,
	}
	return b.index.Index(id, doc)
}

func (b *BleveIndexer) Search(query string, size int) ([]string, error) {
	// Use wildcard query for substring match
	wq := blevepkg.NewWildcardQuery("*" + query + "*")
	sr := blevepkg.NewSearchRequestOptions(wq, size, 0, false)

	result, err := b.index.Search(sr)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(result.Hits))
	for i, hit := range result.Hits {
		ids[i] = hit.ID
	}
	return ids, nil
}
