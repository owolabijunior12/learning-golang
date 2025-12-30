package main

import (
	"fmt"
)

// COURSE 13: ADVANCED TOPICS
// Topics covered:
// 1. Context and cancellation
// 2. Performance optimization
// 3. Memory management
// 4. Reflection
// 5. Type assertions and type switches
// 6. Unsafe package (use with caution!)
// 7. Build tags
// 8. Profiling

func courseThirteen() {
	fmt.Println("=== ADVANCED TOPICS ===\n")

	fmt.Println("CONTEXT AND CANCELLATION:")
	fmt.Println("---")
	fmt.Println(`
// Context propagates cancellation and timeouts through goroutines
import "context"

// Timeout context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use in function
if err := doWork(ctx); err != nil {
	fmt.Println("Work cancelled or timed out:", err)
}

// Cancellable context
ctx, cancel := context.WithCancel(context.Background())
go func() {
	// Can cancel with cancel()
}()

// Value context
ctx := context.WithValue(context.Background(), "user_id", 123)
userID := ctx.Value("user_id")

// Always check context before blocking operations
select {
case <-ctx.Done():
	return ctx.Err()
case result := <-ch:
	// Process result
}
`)
	fmt.Println()

	fmt.Println("PERFORMANCE OPTIMIZATION:")
	fmt.Println("---")
	fmt.Println(`
// 1. Pre-allocate slices if size is known
users := make([]User, 0, 1000) // Capacity 1000
for i := 0; i < 1000; i++ {
	users = append(users, newUser())
}

// 2. Use string builder for concatenation
var sb strings.Builder
for i := 0; i < 1000; i++ {
	sb.WriteString(fmt.Sprintf("item%d ", i))
}
result := sb.String() // Much faster than +

// 3. Avoid unnecessary allocations
// Bad:  return []int{1, 2, 3}[0]
// Good: x := 1; return x

// 4. Use sync.Pool for object reuse
pool := &sync.Pool{
	New: func() interface{} {
		return &Buffer{}
	},
}
buf := pool.Get().(*Buffer)
defer pool.Put(buf)

// 5. Use io.Copy for large data
io.Copy(dst, src) // Efficient, no buffer allocation needed

// 6. Benchmark critical sections
go test -bench=. -benchmem

// 7. Profile with pprof
import _ "net/http/pprof"
// Then visit http://localhost:6060/debug/pprof
`)
	fmt.Println()

	fmt.Println("MEMORY MANAGEMENT:")
	fmt.Println("---")
	fmt.Println(`
// 1. Avoid pointer chains
// Bad:  *****User
// Good: User or *User

// 2. Close resources explicitly
file, _ := os.Open("file.txt")
defer file.Close()

// 3. Clear slices/maps when done
var cache map[string]string
// ... use cache ...
cache = nil // Allow GC to collect

// 4. Don't hold references longer than needed
func process(data []byte) {
	// Use data
	// data reference gone at function end
}

// 5. Use value types for small structs
type Point struct {
	X, Y int // 16 bytes - pass by value
}

// 6. Use pointers for large structs
type Config struct {
	// Many fields...
} // Pass *Config, not Config

// 7. Understand escape analysis
go build -gcflags="-m" // Shows escape analysis
`)
	fmt.Println()

	fmt.Println("REFLECTION:")
	fmt.Println("---")
	fmt.Println(`
import "reflect"

// Get type information at runtime
var x interface{} = "hello"
t := reflect.TypeOf(x)    // Type
v := reflect.ValueOf(x)   // Value

fmt.Println(t.Name())     // string
fmt.Println(v.String())   // hello

// Inspect struct fields
type User struct {
	Name string
	Age  int
}

u := User{"Alice", 30}
typ := reflect.TypeOf(u)

for i := 0; i < typ.NumField(); i++ {
	field := typ.Field(i)
	value := reflect.ValueOf(u).Field(i)
	fmt.Printf("%s: %v\\n", field.Name, value)
}

// Dynamic function calls
fn := reflect.ValueOf(someFunc)
results := fn.Call([]reflect.Value{
	reflect.ValueOf(arg1),
	reflect.ValueOf(arg2),
})

// Caution:
// - Reflection is slower than direct code
// - Can be hard to understand
// - Use sparingly - only when necessary
`)
	fmt.Println()

	fmt.Println("TYPE ASSERTIONS AND SWITCHES:")
	fmt.Println("---")
	fmt.Println(`
// Type assertion
var i interface{} = "hello"

// Check and convert
s := i.(string) // Panics if wrong type

// Safe check
s, ok := i.(string)
if ok {
	fmt.Println("String:", s)
}

// Type switch
switch v := i.(type) {
case string:
	fmt.Printf("String: %v\\n", v)
case int:
	fmt.Printf("Integer: %v\\n", v)
case float64:
	fmt.Printf("Float: %v\\n", v)
default:
	fmt.Printf("Unknown: %T\\n", v)
}

// Common pattern: JSON to struct
var data map[string]interface{}
json.Unmarshal([]byte(jsonStr), &data)

if name, ok := data["name"].(string); ok {
	// Use name
}
if age, ok := data["age"].(float64); ok {
	// Use age (JSON numbers are float64)
}
`)
	fmt.Println()

	fmt.Println("BUILD TAGS:")
	fmt.Println("---")
	fmt.Println(`
// At top of file, before package:
//go:build linux && amd64
// +build linux,amd64

package myapp

// This file only built on Linux AMD64

// OR
//go:build windows
// +build windows

func setupPlatformSpecific() {
	// Windows-specific code
}

// Build with tags:
go build -tags=prod

// Multiple tags:
//go:build (linux || darwin) && !debug

// In code:
// +build linux darwin
// +build !race
`)
	fmt.Println()

	fmt.Println("PROFILING:")
	fmt.Println("---")
	fmt.Println(`
// CPU profiling
import "runtime/pprof"

f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()

// ... run code to profile ...

// Memory profiling
f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)

// Analyze:
go tool pprof cpu.prof

// HTTP pprof
import _ "net/http/pprof"

go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()

// Visit http://localhost:6060/debug/pprof

// Common profiles:
// /debug/pprof/heap     - Memory allocations
// /debug/pprof/goroutine - Running goroutines
// /debug/pprof/profile  - CPU profile
// /debug/pprof/trace    - Execution trace
`)
	fmt.Println()

	fmt.Println("CACHING STRATEGIES:")
	fmt.Println("---")
	fmt.Println(`
// 1. Simple in-memory cache
type Cache struct {
	sync.RWMutex
	data map[string]interface{}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// 2. TTL cache (with expiration)
type CacheEntry struct {
	Value      interface{}
	ExpiresAt  time.Time
}

// 3. LRU cache (keep most-used items)
// Use: github.com/hashicorp/golang-lru

// 4. Distributed cache
// Use Redis for shared cache across instances
`)
	fmt.Println()

	fmt.Println("GOROUTINE MANAGEMENT:")
	fmt.Println("---")
	fmt.Println(`
// 1. Avoid goroutine leaks
// Bad:
for {
	go work()  // Infinite goroutines!
}

// Good:
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

for {
	select {
	case <-ctx.Done():
		return
	default:
		go work(ctx)
	}
}

// 2. Limit concurrent goroutines
semaphore := make(chan struct{}, 10) // Max 10

go func() {
	semaphore <- struct{}{}      // Acquire
	defer func() { <-semaphore }() // Release
	
	// Actual work
}()

// 3. Monitor goroutines
fmt.Println(runtime.NumGoroutine())

// 4. Use WaitGroup for cleanup
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Work
	}()
}
wg.Wait()
`)
	fmt.Println()

	fmt.Println("BEST PRACTICES FOR ADVANCED TOPICS:")
	fmt.Println("---")
	fmt.Println("✓ Use context for cancellation in all async operations")
	fmt.Println("✓ Profile before optimizing")
	fmt.Println("✓ Use reflection sparingly")
	fmt.Println("✓ Pre-allocate slices when size is known")
	fmt.Println("✓ Use string.Builder for string concatenation")
	fmt.Println("✓ Monitor goroutine count in production")
	fmt.Println("✓ Always close resources")
	fmt.Println("✓ Understand escape analysis")
	fmt.Println("✓ Use sync.Pool for object reuse")
	fmt.Println("✓ Build tags for platform-specific code")
	fmt.Println("✓ Regular profiling in production")
	fmt.Println("✓ Avoid unsafe package unless necessary")
	fmt.Println("✓ Cache strategically")
	fmt.Println("✓ Limit concurrent operations")
	fmt.Println("✓ Use benchmarks for critical code")
	fmt.Println()

	fmt.Println("=== END OF ADVANCED TOPICS ===")
}

// KEY TAKEAWAYS:
// 1. Context is essential for cancellation and timeouts
// 2. Always pass context through function chains
// 3. Profile before optimizing
// 4. Pre-allocate slices for better performance
// 5. Use strings.Builder for concatenation
// 6. Reflection is powerful but slow
// 7. Type assertions and switches for runtime type checking
// 8. Build tags for platform-specific code
// 9. Monitor goroutine count to avoid leaks
// 10. Use sync primitives (Mutex, WaitGroup, etc.)
// 11. Benchmark critical sections
// 12. Memory profiling helps find leaks
// 13. CPU profiling finds hot spots
// 14. Unsafe package bypasses type safety (avoid!)
// 15. Escape analysis determines allocation location
// 16. sync.Pool reuses objects to reduce GC pressure
// 17. Rate limiting prevents resource exhaustion
// 18. Caching improves performance significantly
// 19. Understand goroutine scheduling
// 20. Production requires monitoring and profiling
