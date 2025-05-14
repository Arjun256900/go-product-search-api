## Description

A Go-based backend API that performs full-text operation on 1M products with Belve and Chi-router.

## Folder structure
go-product-search-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── handlers/
│   │       └── search.go
│   ├── models/
│   │   └── product.go
│   ├── product/
│   │   ├── generator.go
│   │   └── repository.go
│   └── search/
│       ├── bleve/
│       │   └── bleve.go
│       └── indexer.go
├── pkg/
│   └── graceful/
├── .gitignore
└── go.mod


## Change log

| Time stamp | Progress |
| ---------- | ------ |
| 14th May 9am | Initial setup with standard Go structure |
To be continued.

## Author

- 👨‍💻 Made with 💪 and ☕ by [Arjun](https://github.com/Arjun256900)
- ⭐️ If you find this project helpful, give it a star! and feel free to open issues or PRs.

## License

This project is licensed under the MIT License.

## Resources
TBD
