package models

/*
	Product is a searchable item
	'Keywords' are used for boosting and are not returned in JSON (hidden slice of strings used for indexing)
	ID --> Sequential
	Name --> Primary search field
	Category --> Filterable attribute
*/
type Product struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Category string   `json:"category"`
}