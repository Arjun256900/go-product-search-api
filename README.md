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
- Example: `curl localhost:8080/search?q=kitchen`

# Benchmark
On my machine (16 GB RAM, 512GB SSD), searching “kitchen” over 1M records took ~240ms on average (using hey `-n 1000`, `-c 1`). Queries used to run tests:
- http://localhost:8080/search?q=kitchen
- http://localhost:8080/search?q=lucid
- http://localhost:8080/search?q=icy  

## Change log

| Time stamp | Progress |
| ---------- | ------ |
| 14th May 9:00 IST | Initial setup with standard Go structure |
| 14th May 11:00 IST | Added random product name generator and product struct|
| 14th May 12:00 IST | Bleve indexer and search service |
| 14th May 14:00 IST | Search route and it's handler implementation |
| 14th May 17:00 IST | Optimized O(n) lookup to O(1) using store map |

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
- [Bleve Docs](https://blevesearch.com/docs/)
- [Chi router](https://github.com/go-chi/chi)
- [Golang](https://go.dev/)
- [hey](https://github.com/rakyll/hey/)
