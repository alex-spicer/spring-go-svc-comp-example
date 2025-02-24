# Spring Boot vs Go Service Comparison

A practical comparison of implementing the same product management API in both Spring Boot and Go. This repository
accompanies a blog series exploring how common web service patterns are implemented in both frameworks.

## Project Structure

```
.
├── product/          # Spring Boot implementation (Java 23)
├── product-go/       # Go implementation (Go 1.23.6)
└── docker-compose.yml
```

## Requirements

- Java 23+
- Go 1.23+
- Docker & Docker Compose
- Make

## Quick Start

Run both services:

```bash
make docker-run
```

GO runs on `localhost:8081` and Java on `localhost:8080`.

Run individual services:

```bash
make run-java    # Spring Boot on :8080
make run-go      # Go on :8081
```

## API Endpoints

Both services implement:

```
GET    /api/products     # List all products
GET    /api/products/:id # Get one product
POST   /api/products     # Create product
PUT    /api/products/:id # Update product
DELETE /api/products/:id # Delete product
```

## Implementation Details

- Spring Boot: Virtual Threads, Records, Pattern Matching
- Go: Goroutines, Standard Project Layout, Gin Framework
- Both use in-memory storage with thread-safe operations

## Blog Posts

- [Part 1: Dependency Injection and Composition](https://amspicer.medium.com/side-by-side-java-spring-boot-go-9e642e740e84)
- [Part 2: Routing and HTTP Handlers](https://amspicer.medium.com/side-by-side-java-spring-boot-go-e1bab8159f78)
- [Part 3: Middleware and Interceptors](https://amspicer.medium.com/side-by-side-spring-boot-go-749dafa7f285)
- Part 4: Error Handling
- Part 5: Concurrency Patterns

## License

MIT
