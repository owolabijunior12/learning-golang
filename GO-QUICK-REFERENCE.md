# Go Developer Quick Reference

## Installation & Setup

```bash
# Download and install
https://golang.org/dl/

# Verify installation
go version

# Set up first project
mkdir myproject
cd myproject
go mod init github.com/username/myproject
```

## Essential Commands

```bash
go run main.go              # Run code
go build                    # Build binary
go test ./...              # Run tests
go test -v -cover ./...    # Verbose with coverage
go fmt ./...               # Format code
go vet ./...               # Find issues
go get package/name        # Add dependency
go mod tidy                # Clean dependencies
go mod vendor              # Create vendor directory
```

## Package Import Examples

```go
import (
    "fmt"                           // Standard library
    "github.com/redis/go-redis/v9"  // External package
)
```

## Common Patterns

### Variables & Constants
```go
var x int = 5              // Explicit type
y := 10                    // Type inference
const Pi = 3.14159         // Constant

// Multiple declarations
var (
    a int = 1
    b string = "hello"
)
```

### Functions
```go
func add(a, b int) int {
    return a + b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func multiple() (int, string, bool) {
    return 1, "hello", true
}
```

### Error Handling
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

val, err := someFunction()
if err != nil {
    log.Fatal(err)
}
```

### Structs & Methods
```go
type User struct {
    Name  string
    Email string
}

// Method with value receiver
func (u User) Display() string {
    return u.Name + " (" + u.Email + ")"
}

// Method with pointer receiver
func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail
}
```

### Interfaces
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Empty interface (any type)
var x interface{} = "hello"
```

### Loops
```go
for i := 0; i < 10; i++ { }           // Traditional
for i < 10 { }                        // While-style
for { }                               // Infinite
for i, v := range slice { }           // Range
for key, value := range map { }       // Map iteration
```

### Slices & Maps
```go
// Slice
s := []int{1, 2, 3}
s = append(s, 4)
s[0] = 10
sub := s[1:3]

// Map
m := map[string]int{"a": 1, "b": 2}
m["c"] = 3
delete(m, "a")
v, ok := m["b"]
```

### Goroutines & Channels
```go
// Goroutine
go func() {
    fmt.Println("Running concurrently")
}()

// Channel
ch := make(chan int)
ch <- 42          // Send
value := <-ch     // Receive
close(ch)         // Close

// Wait for goroutine
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Work here
}()
wg.Wait()

// Timeout
select {
case result := <-ch:
    // Got result
case <-time.After(5 * time.Second):
    // Timeout
}
```

### Defer
```go
file, err := os.Open("file.txt")
defer file.Close()  // Runs at end of function
```

### Testing
```go
func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}

// Table-driven test
tests := []struct {
    name     string
    a, b     int
    expected int
}{
    {"2+3", 2, 3, 5},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        if got := add(tt.a, tt.b); got != tt.expected {
            t.Errorf("got %d, want %d", got, tt.expected)
        }
    })
}
```

### HTTP Server
```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### JSON
```go
import "encoding/json"

// Struct to JSON
data := User{Name: "Alice", Age: 30}
jsonBytes, err := json.Marshal(data)

// JSON to struct
var user User
err := json.Unmarshal(jsonBytes, &user)

// Struct tags
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}
```

### Database
```go
import "database/sql"

db, err := sql.Open("postgres", "connection_string")
defer db.Close()

// Query
rows, err := db.Query("SELECT * FROM users WHERE age > ?", 18)
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
}

// Single row
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)

// Execute
result, err := db.Exec("INSERT INTO users (name) VALUES (?)", "John")
```

### Concurrency Patterns

**Worker Pool**
```go
jobs := make(chan Job, 100)
results := make(chan Result, 100)

for w := 0; w < 3; w++ {
    go worker(jobs, results)
}
```

**Timeout**
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case <-ctx.Done():
    fmt.Println("Timeout!")
}
```

## Performance Tips

- Pre-allocate slices: `make([]int, 0, 1000)`
- Use `strings.Builder` for concatenation
- Use `sync.Pool` for object reuse
- Benchmark: `go test -bench=.`
- Profile: `go tool pprof cpu.prof`

## Code Style

```bash
# Format code
go fmt ./...

# Import management
goimports -w .

# Lint
golangci-lint run ./...

# Find issues
go vet ./...
```

## Useful Standard Packages

| Package | Use |
|---------|-----|
| `fmt` | Formatting and printing |
| `strings` | String manipulation |
| `io` | Input/output operations |
| `os` | Operating system interaction |
| `filepath` | Path operations |
| `json` | JSON encoding/decoding |
| `http` | HTTP client/server |
| `sql` | Database connection |
| `time` | Time and duration |
| `sync` | Synchronization primitives |
| `context` | Cancellation and timeouts |
| `testing` | Testing framework |
| `log` | Logging |
| `errors` | Error creation and wrapping |

## Common External Packages

```
github.com/gorilla/mux              - HTTP routing
github.com/lib/pq                   - PostgreSQL driver
github.com/go-sql-driver/mysql      - MySQL driver
go.mongodb.org/mongo-driver         - MongoDB driver
github.com/redis/go-redis           - Redis client
github.com/go-yaml/yaml             - YAML parsing
github.com/urfave/cli               - CLI framework
github.com/spf13/cobra              - CLI command framework
go.uber.org/zap                     - Structured logging
github.com/google/uuid              - UUID generation
```

## Debugging Tips

```go
// Print values
fmt.Printf("value: %v, type: %T\n", x, x)

// Print struct with field names
fmt.Printf("%+v\n", user)

// Stack trace
runtime.PrintStack()

// Debugging with dlv (Go debugger)
dlv debug
(dlv) break main.main
(dlv) continue
(dlv) next
(dlv) print variable
```

## Memory Management

- Slices automatically grow when appended
- Maps are passed by reference
- Strings are immutable
- Pointers needed for methods to modify receiver
- Defer ensures cleanup happens

## Go Idioms

✓ Make the zero value useful
✓ Errors are values
✓ Interfaces are implicit
✓ Keep interfaces small
✓ Don't communicate by sharing memory; share memory by communicating
✓ Return errors as values, not through channels
✓ Use error wrapping (fmt.Errorf with %w)
✓ Write clear, simple code
✓ Document exported identifiers
✓ Tests are first-class citizens
