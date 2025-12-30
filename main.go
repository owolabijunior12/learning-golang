package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║          COMPLETE GO DEVELOPER LEARNING COURSE                 ║")
	fmt.Println("║                                                                ║")
	fmt.Println("║  A comprehensive guide to becoming a professional Go developer ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝\n")

	fmt.Println("Running all course demonstrations...\n")

	// Run all courses
	courseOne()
	courseTwo()
	courseThreeDemo()
	courseFour()
	courseFiveDemo()
	courseSixDemo()
	courseSevenDemo()
	courseEightDemo()
	courseNineDemo()
	courseTenDemo()
	courseElevenDemo()
	courseTwelveDemo()
	courseThirteenDemo()

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nCOURSE STRUCTURE:")
	fmt.Println("\nAll 13 courses have been executed above. To study individual courses:\n")

	courses := []struct {
		number      int
		name        string
		file        string
		description string
	}{
		{1, "BASICS", "01-basics.go", "Variables, types, control flow, operators"},
		{2, "FUNCTIONS & ERRORS", "02-functions-and-errors.go", "Functions, error handling, defer, panic/recover"},
		{3, "STRUCTS & INTERFACES", "03-structs-and-interfaces.go", "Structs, methods, interfaces, composition"},
		{4, "GOROUTINES & CHANNELS", "04-goroutines-and-channels.go", "Concurrency, goroutines, channels, select"},
		{5, "FILE HANDLING", "05-file-handling.go", "File I/O, directory operations, buffered reading"},
		{6, "HTTP SERVER & REST", "06-http-server.go", "HTTP servers, routing, JSON, middleware"},
		{7, "SQL DATABASES", "07-sql-database.go", "PostgreSQL, MySQL, prepared statements, transactions"},
		{8, "MONGODB", "08-mongodb-database.go", "MongoDB driver, BSON, aggregation pipelines"},
		{9, "REDIS", "09-redis-database.go", "Redis, data structures, caching, pub/sub"},
		{10, "TESTING", "10-testing.go", "Unit tests, table-driven tests, benchmarking, mocking"},
		{11, "PROJECT STRUCTURE", "11-project-structure.go", "Directory layout, packages, modules, best practices"},
		{12, "DESIGN PATTERNS", "12-design-patterns.go", "Middleware, DI, repositories, patterns"},
		{13, "ADVANCED TOPICS", "13-advanced-topics.go", "Context, profiling, reflection, optimization"},
	}

	for _, course := range courses {
		fmt.Printf("[%d]  %-20s - %s\n", course.number, course.name, course.description)
	}

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nTO RUN INDIVIDUAL COURSES:")
	fmt.Println("\n  go run 01-basics.go")
	fmt.Println("  go run 02-functions-and-errors.go")
	fmt.Println("  # ... etc for each course")

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nKEY RESOURCES:")
	fmt.Println("  • Official Go Tour: https://tour.golang.org")
	fmt.Println("  • Go by Example: https://gobyexample.com")
	fmt.Println("  • Effective Go: https://golang.org/doc/effective_go")
	fmt.Println("  • Go Package Docs: https://pkg.go.dev")

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nHAPPY LEARNING! 🚀\n")
}

	courses := []struct {
		number      int
		name        string
		file        string
		description string
	}{
		{1, "BASICS", "01-basics.go", "Variables, types, control flow, operators"},
		{2, "FUNCTIONS & ERRORS", "02-functions-and-errors.go", "Functions, error handling, defer, panic/recover"},
		{3, "STRUCTS & INTERFACES", "03-structs-and-interfaces.go", "Structs, methods, interfaces, composition"},
		{4, "GOROUTINES & CHANNELS", "04-goroutines-and-channels.go", "Concurrency, goroutines, channels, select"},
		{5, "FILE HANDLING", "05-file-handling.go", "File I/O, directory operations, buffered reading"},
		{6, "HTTP SERVER & REST", "06-http-server.go", "HTTP servers, routing, JSON, middleware"},
		{7, "SQL DATABASES", "07-sql-database.go", "PostgreSQL, MySQL, prepared statements, transactions"},
		{8, "MONGODB", "08-mongodb-database.go", "MongoDB driver, BSON, aggregation pipelines"},
		{9, "REDIS", "09-redis-database.go", "Redis, data structures, caching, pub/sub"},
		{10, "TESTING", "10-testing.go", "Unit tests, table-driven tests, benchmarking, mocking"},
		{11, "PROJECT STRUCTURE", "11-project-structure.go", "Directory layout, packages, modules, best practices"},
		{12, "DESIGN PATTERNS", "12-design-patterns.go", "Middleware, DI, repositories, patterns"},
		{13, "ADVANCED TOPICS", "13-advanced-topics.go", "Context, profiling, reflection, optimization"},
	}

	for _, course := range courses {
		fmt.Printf("[%d]  %-20s - %s\n", course.number, course.name, course.description)
	}

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nTO RUN A COURSE:")
	fmt.Println("\n  Option 1: Run individual file")
	fmt.Println("    go run 01-basics.go")
	fmt.Println("\n  Option 2: Run all files")
	fmt.Println("    go run .")
	fmt.Println("\n  Option 3: Create test file and run")
	fmt.Println("    go run . your_test.go")

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nKEY RESOURCES:")
	fmt.Println("  • Official Go Tour: https://tour.golang.org")
	fmt.Println("  • Go by Example: https://gobyexample.com")
	fmt.Println("  • Effective Go: https://golang.org/doc/effective_go")
	fmt.Println("  • Go Package Docs: https://pkg.go.dev")

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nLEARNING PATH:")
	fmt.Println("  Week 1:  Courses 1-3 (Fundamentals)")
	fmt.Println("  Week 2:  Courses 4-6 (Concurrency & Web)")
	fmt.Println("  Week 3:  Courses 7-9 (Databases)")
	fmt.Println("  Week 4:  Courses 10-13 (Advanced Topics)")

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nNEXT STEPS:")
	fmt.Println("  1. Read 00-README.md for overview")
	fmt.Println("  2. Start with 01-basics.go")
	fmt.Println("  3. Run each file: go run 01-basics.go")
	fmt.Println("  4. Modify examples to experiment")
	fmt.Println("  5. Build small projects to apply knowledge")

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("\nHAPPY LEARNING! 🚀\n")
}

// Course function stubs for compilation
func courseFiveDemo() {
	fmt.Println("\n=== COURSE 5: FILE HANDLING ===")
	fmt.Println("See 05-file-handling.go for detailed examples\n")
}

func courseSixDemo() {
	fmt.Println("\n=== COURSE 6: HTTP SERVER & REST ===")
	fmt.Println("See 06-http-server.go for detailed examples\n")
}

func courseEightDemo() {
	fmt.Println("\n=== COURSE 8: MONGODB ===")
	fmt.Println("See 08-mongodb-database.go for detailed examples\n")
}

func courseNineDemo() {
	fmt.Println("\n=== COURSE 9: REDIS ===")
	fmt.Println("See 09-redis-database.go for detailed examples\n")
}

func courseElevenDemo() {
	fmt.Println("\n=== COURSE 11: PROJECT STRUCTURE ===")
	fmt.Println("See 11-project-structure.go for detailed examples\n")
}

func courseTwelveDemo() {
	fmt.Println("\n=== COURSE 12: DESIGN PATTERNS ===")
	fmt.Println("See 12-design-patterns.go for detailed examples\n")
}

func courseThirteenDemo() {
	fmt.Println("\n=== COURSE 13: ADVANCED TOPICS ===")
	fmt.Println("See 13-advanced-topics.go for detailed examples\n")
}

