![Golang](https://img.shields.io/badge/Go-blue.svg)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)

## Description

A Go-based backend API that performs full-text operation on 1M products with Bleve and Chi-router.

# Prerequisites
You need `go 1.21+` to run this project.

# Getting started
## Setup
1. Fork the repository
2. In the root directory, run `go mod tidy`
3. Adjust product count in `main.go` as needed
4. Run `go run cmd/server/main.go`

# Api Docs
- GET /health            → 200 {"status":"ok"}
- GET /search?q=<term>   → 200 [{"id":…, "name":"…", "category":"…"}, …]
Example: `curl localhost:8080/search?q=kitchen`

# Benchmark
On my machine (16 GB RAM, SSD), searching “kitchen” over 1M records took ~240 ms (average of 10 runs with curl_time).

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
- [Bleve Docs]([https://www.w3schools.com/go/index.php](https://blevesearch.com/docs/))
- [Chi router](https://github.com/go-chi/chi)
