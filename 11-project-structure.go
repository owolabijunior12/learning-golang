package main

import (
	"fmt"
)

// COURSE 11: PROJECT STRUCTURE AND BEST PRACTICES
// Topics covered:
// 1. Directory organization
// 2. Package naming
// 3. Module setup (go.mod, go.sum)
// 4. Dependency management
// 5. Configuration management
// 6. Logging
// 7. Error handling patterns
// 8. Code organization patterns

func courseEleven() {
	fmt.Println("=== PROJECT STRUCTURE AND BEST PRACTICES ===\n")

	fmt.Println("RECOMMENDED DIRECTORY STRUCTURE:")
	fmt.Println("---\n")

	fmt.Println(`
myproject/
├── go.mod                    # Module definition
├── go.sum                    # Dependency checksums
├── README.md                 # Project documentation
├── LICENSE                   # License file
├── .gitignore               # Git ignore rules
├── Makefile                 # Build automation
├── main.go                  # Entry point
│
├── cmd/                     # Command-line applications
│   ├── server/
│   │   └── main.go         # HTTP server entry point
│   ├── cli/
│   │   └── main.go         # CLI tool entry point
│   └── worker/
│       └── main.go         # Background worker entry point
│
├── internal/                # Private packages (can't be imported externally)
│   ├── config/
│   │   └── config.go       # Configuration loading
│   ├── database/
│   │   ├── db.go           # Database connection
│   │   ├── user.go         # User repository
│   │   └── migrations.go    # Database migrations
│   ├── service/
│   │   ├── user.go         # User business logic
│   │   └── auth.go         # Authentication logic
│   ├── middleware/
│   │   ├── auth.go         # Auth middleware
│   │   └── logging.go      # Logging middleware
│   └── util/
│       └── helpers.go      # Utility functions
│
├── pkg/                     # Public packages (can be imported)
│   ├── api/
│   │   ├── types.go        # API types/models
│   │   └── client.go       # Public client library
│   └── logger/
│       └── logger.go       # Public logging package
│
├── api/                     # API specification and handlers
│   ├── routes.go           # Route definitions
│   ├── handler_user.go     # User API endpoints
│   ├── handler_auth.go     # Auth API endpoints
│   └── middleware.go       # HTTP middleware
│
├── migrations/              # Database migration files
│   ├── 001_create_users.sql
│   └── 002_create_posts.sql
│
├── config/                  # Configuration files
│   ├── dev.yml             # Development config
│   ├── test.yml            # Test config
│   └── prod.yml            # Production config
│
├── tests/                   # Integration tests
│   ├── api_test.go
│   └── database_test.go
│
├── scripts/                 # Utility scripts
│   ├── build.sh
│   ├── deploy.sh
│   └── migrate.sh
│
├── docs/                    # Documentation
│   ├── API.md              # API documentation
│   ├── SETUP.md            # Setup guide
│   └── ARCHITECTURE.md     # Architecture
│
├── .github/
│   ├── workflows/          # CI/CD pipelines
│   │   └── test.yml        # GitHub Actions
│   └── ISSUE_TEMPLATE/
│
└── vendor/                 # Go modules (if using vendor)
`)
	fmt.Println()

	fmt.Println("GO.MOD (Module Definition):")
	fmt.Println("---")
	fmt.Println(`
module github.com/username/myproject

go 1.21

require (
	github.com/gorilla/mux v1.8.0
	github.com/redis/go-redis/v9 v9.0.0
	github.com/lib/pq v1.10.9
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
)
`)
	fmt.Println()

	fmt.Println("PACKAGE NAMING CONVENTIONS:")
	fmt.Println("---")
	fmt.Println(`
✓ Use short, clear package names
✓ Avoid generic names like "util", "common", "helper"
✓ Use the package name in exported functions
✓ Package name matches directory name
✓ Use singular names (user, not users)
✓ Lowercase, no underscores

Examples:
✓ package user
✓ package database
✓ package http
✗ package util_helpers
✗ package UserManagement
`)
	fmt.Println()

	fmt.Println("INTERNAL VS PUBLIC PACKAGES:")
	fmt.Println("---")
	fmt.Println(`
internal/     - Private to this module
              - Cannot be imported by external projects
              - Use for business logic, database code, etc.

pkg/          - Public to external projects
              - Explicitly marked as public API
              - Exported for reuse by other modules
              - Keep APIs stable
              - Good documentation required

Example:
internal/service/user.go    - Private business logic
pkg/api/types.go            - Public API types
`)
	fmt.Println()

	fmt.Println("TYPICAL MAIN.GO:")
	fmt.Println("---")
	fmt.Println(`
package main

import (
	"context"
	"log"
	
	"github.com/username/myproject/internal/config"
	"github.com/username/myproject/internal/database"
	"github.com/username/myproject/internal/service"
	"github.com/username/myproject/api"
)

func main() {
	// Load configuration
	cfg := config.Load()
	
	// Initialize database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	// Create services
	userService := service.NewUserService(db)
	
	// Start server
	server := api.NewServer(userService)
	if err := server.Start(cfg.Port); err != nil {
		log.Fatal(err)
	}
}
`)
	fmt.Println()

	fmt.Println("CONFIGURATION MANAGEMENT:")
	fmt.Println("---")
	fmt.Println(`
// config/config.go
package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DatabaseURL string
	RedisURL    string
	LogLevel    string
	Environment string
}

func Load() *Config {
	return &Config{
		Port:        getEnvInt("PORT", 8080),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://localhost/mydb"),
		RedisURL:    getEnv("REDIS_URL", "localhost:6379"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if num, err := strconv.Atoi(value); err == nil {
			return num
		}
	}
	return defaultValue
}
`)
	fmt.Println()

	fmt.Println("STRUCTURED LOGGING:")
	fmt.Println("---")
	fmt.Println(`
// Setup: go get github.com/go-uber/zap

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	logger.Info("server started",
		zap.Int("port", 8080),
		zap.String("version", "1.0.0"),
	)
	
	logger.Error("failed to connect",
		zap.Error(err),
		zap.String("host", "localhost"),
	)
}
`)
	fmt.Println()

	fmt.Println("ERROR HANDLING PATTERNS:")
	fmt.Println("---")
	fmt.Println(`
// 1. Wrap errors with context
if err != nil {
	return fmt.Errorf("failed to create user: %w", err)
}

// 2. Custom error types
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// 3. Error checking
if err := operation(); err != nil {
	// Handle error immediately
	log.WithError(err).Error("operation failed")
	return err
}

// 4. Ignore when appropriate
_ = file.Close() // Intentionally ignore error
`)
	fmt.Println()

	fmt.Println("MAKEFILE FOR AUTOMATION:")
	fmt.Println("---")
	fmt.Println(`
.PHONY: build test run clean

build:
	go build -o bin/app ./cmd/main

test:
	go test -v -cover ./...

run:
	go run ./cmd/main

clean:
	rm -rf bin/

migrate:
	go run ./cmd/migrate/main.go

docker-build:
	docker build -t myapp:latest .

docker-run:
	docker run -p 8080:8080 myapp:latest

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...
	goimports -w .

vet:
	go vet ./...

deps:
	go mod tidy
	go mod vendor
`)
	fmt.Println()

	fmt.Println("DOCKER SETUP:")
	fmt.Println("---")
	fmt.Println(`
FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/server /usr/local/bin/
EXPOSE 8080
CMD ["server"]
`)
	fmt.Println()

	fmt.Println("GITIGNORE:")
	fmt.Println("---")
	fmt.Println(`
# Binaries
bin/
dist/
*.exe

# Go
vendor/
go.sum

# IDE
.vscode/
.idea/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Environment
.env
.env.local

# Tests
*.out
coverage.html
`)
	fmt.Println()

	fmt.Println("CODE ORGANIZATION BEST PRACTICES:")
	fmt.Println("---")
	fmt.Println("✓ Keep packages focused and single-purpose")
	fmt.Println("✓ Avoid cyclic dependencies")
	fmt.Println("✓ Use interfaces for abstraction")
	fmt.Println("✓ Keep internal implementation hidden")
	fmt.Println("✓ Document exported functions and types")
	fmt.Println("✓ Use dependency injection")
	fmt.Println("✓ Keep business logic separate from HTTP handling")
	fmt.Println("✓ Use middleware for cross-cutting concerns")
	fmt.Println("✓ Test each package independently")
	fmt.Println("✓ Handle configuration from environment")
	fmt.Println()

	fmt.Println("NAMING CONVENTIONS:")
	fmt.Println("---")
	fmt.Println("Files:        lowercase_with_underscores")
	fmt.Println("Packages:     lowercase, one word")
	fmt.Println("Functions:    PascalCase for exported, camelCase for private")
	fmt.Println("Constants:    UPPER_CASE for constants")
	fmt.Println("Interfaces:   PascalCase, usually end with 'er'")
	fmt.Println("Variables:    camelCase")
	fmt.Println()

	fmt.Println("DEPENDENCY MANAGEMENT:")
	fmt.Println("---")
	fmt.Println("go mod init module/name        - Initialize module")
	fmt.Println("go get github.com/user/repo    - Add dependency")
	fmt.Println("go get -u ./...                - Update all dependencies")
	fmt.Println("go mod tidy                    - Clean up dependencies")
	fmt.Println("go mod vendor                  - Create vendor directory")
	fmt.Println("go mod verify                  - Verify integrity")
	fmt.Println()

	fmt.Println("=== END OF PROJECT STRUCTURE ===")
}

// KEY TAKEAWAYS:
// 1. Use cmd/ for applications, internal/ for private code
// 2. Keep packages focused (single responsibility)
// 3. Use go.mod for dependency management
// 4. Document exported functions and types
// 5. Avoid generic package names (util, common)
// 6. Use interfaces for abstraction
// 7. Separate concerns (API, business logic, persistence)
// 8. Use middleware for cross-cutting concerns
// 9. Handle configuration from environment variables
// 10. Use structured logging for production
// 11. Implement proper error handling with context
// 12. Write tests alongside code
// 13. Use Makefile for build automation
// 14. Keep main.go small and simple
// 15. Use consistent naming conventions
// 16. Avoid circular dependencies
// 17. Use dependency injection
// 18. Create migration files for database changes
// 19. Use .gitignore to exclude generated files
// 20. Document architecture and setup in README
