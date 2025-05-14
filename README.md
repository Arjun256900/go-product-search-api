![Golang](https://img.shields.io/badge/Go-blue.svg)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)

## Description

A Go-based backend API that performs full-text operation on 1M products with Bleve and Chi-router.

# Getting started
## Setup
1. Fork the repository
2. In the root directory, run `go mod tidy`
3. Adjust product count in `main.go` as needed
4. Run `go run cmd/server/main.go`

## Change log

| Time stamp | Progress |
| ---------- | ------ |
| 14th May 9am | Initial setup with standard Go structure |
| 14th May 11am | Added random product name generator and product struct|
| 14th May 12am | Bleve indexer and search service |
| 14th May 2am | Search route and it's handler implementation |

## 📚 What I learnt
This Golang project helped me learn:
- The native chi-router in go
- Practice Go routines
- Optimize performance-critical applications
- Bleve indexing package
- Strengthen the understanding of Go semantics

## Author

- 👨‍💻 Made with 💪 and ☕ by [Arjun](https://github.com/Arjun256900)
- ⭐️ If you find this project helpful, give it a star! and feel free to open issues or PRs.

## License

This project is licensed under the MIT License.

## Resources
- [W3schools](https://www.w3schools.com/go/index.php)
