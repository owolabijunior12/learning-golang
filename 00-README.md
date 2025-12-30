# Complete Go Developer Learning Course

This course covers everything you need to know to become a proficient Go developer, from basics to production-ready applications.

## Course Structure

### Level 1: Fundamentals
1. **01-basics.go** - Variables, types, constants, and basic operations
2. **02-functions-and-errors.go** - Function definitions, error handling, and best practices
3. **03-interfaces-and-structs.go** - Structs, interfaces, and OOP in Go

### Level 2: Concurrency & I/O
4. **04-goroutines-and-channels.go** - Goroutines, channels, and concurrent programming
5. **05-file-handling.go** - File I/O operations and stream processing

### Level 3: Web Development
6. **06-http-server.go** - HTTP servers, routing, and REST APIs

### Level 4: Databases
7. **07-sql-database.go** - PostgreSQL and MySQL with database/sql and sqlc
8. **08-mongodb-database.go** - MongoDB integration with MongoDB driver
9. **09-redis-database.go** - Redis integration for caching and sessions
10. **10-advanced-db-patterns.go** - Connection pooling, transactions, and optimization

### Level 5: Advanced Topics
11. **11-middleware-and-patterns.go** - Middleware, dependency injection, and design patterns
12. **12-testing.go** - Unit testing, table-driven tests, and mocking
13. **13-project-structure.go** - Real-world project organization and best practices
14. **14-authentication.go** - JWT, OAuth2, and security best practices
15. **15-logging-and-monitoring.go** - Structured logging and metrics

## How to Use This Course

1. Start with `01-basics.go` - Read the comments and code examples
2. Run each file: `go run filename.go`
3. Understand the output and modify examples
4. Progress sequentially through the levels
5. Build small projects after each level to reinforce learning

## Running Examples

```bash
# Run a single file
go run 01-basics.go

# Run all files (after setting up databases)
go run .

# Run with arguments
go run 02-functions-and-errors.go
```

## Prerequisites

- Go 1.19+ installed
- Basic programming knowledge
- Text editor or IDE (VS Code recommended)
- Docker (optional, for databases)

## Database Setup (Optional)

### PostgreSQL
```bash
docker run --name postgres -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres:latest
```

### MongoDB
```bash
docker run --name mongodb -d -p 27017:27017 mongo:latest
```

### Redis
```bash
docker run --name redis -d -p 6379:6379 redis:latest
```

## Key Concepts You'll Learn

✅ Go syntax and idioms
✅ Error handling patterns
✅ Concurrency and parallelism
✅ Interface design
✅ Working with multiple databases
✅ Building REST APIs
✅ Testing strategies
✅ Production-ready code organization
✅ Security best practices
✅ Performance optimization

## Go Best Practices

1. **Keep interfaces small** - Single responsibility principle
2. **Make the zero value useful** - Well-designed types work without initialization
3. **Return errors as values** - Don't panic unless truly exceptional
4. **Use goroutines wisely** - Monitor goroutine leaks
5. **Write clear code** - Readability matters more than cleverness
6. **Test your code** - Aim for >80% coverage
7. **Document exports** - Every exported name needs a doc comment

## Resources

- Official Go Tour: https://tour.golang.org
- Go by Example: https://gobyexample.com
- Effective Go: https://golang.org/doc/effective_go
- Go Database Packages: https://golang.org/pkg/database/sql/

---

**Total Learning Time**: 40-60 hours of active learning

**Next Step**: Open `01-basics.go` and start learning!
