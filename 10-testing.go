package main

import (
	"fmt"
)

// COURSE 10: TESTING IN GO
// Topics covered:
// 1. Unit testing basics
// 2. Table-driven tests
// 3. Subtests
// 4. Benchmarking
// 5. Mocking and stubs
// 6. Test coverage
// 7. Integration testing
// 8. Best practices

// ============ 1. FUNCTIONS TO TEST ============
func addTest(a, b int) int {
	return a + b
}

func divideTest(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func isEven(n int) bool {
	return n%2 == 0
}

// ============ 2. BASIC TEST ============
// File: mypackage_test.go
// func TestAdd(t *testing.T) {
//	result := add(2, 3)
//	expected := 5
//
//	if result != expected {
//		t.Errorf("add(2, 3) = %d, want %d", result, expected)
//	}
// }

// ============ 3. TABLE-DRIVEN TESTS (RECOMMENDED) ============
// func TestAddTableDriven(t *testing.T) {
//	tests := []struct {
//		name     string
//		a, b     int
//		expected int
//	}{
//		{"positive numbers", 2, 3, 5},
//		{"negative numbers", -2, -3, -5},
//		{"mixed", 5, -3, 2},
//		{"zero", 0, 0, 0},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result := add(tt.a, tt.b)
//			if result != tt.expected {
//				t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
//			}
//		})
//	}
// }

// ============ 4. TEST WITH ERRORS ============
// func TestDivide(t *testing.T) {
//	tests := []struct {
//		name      string
//		a, b      float64
//		expected  float64
//		shouldErr bool
//	}{
//		{"normal division", 10, 2, 5, false},
//		{"division by zero", 10, 0, 0, true},
//		{"decimal result", 10, 3, 3.333..., false},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result, err := divide(tt.a, tt.b)
//
//			if tt.shouldErr && err == nil {
//				t.Error("expected error, got nil")
//			}
//
//			if !tt.shouldErr && err != nil {
//				t.Errorf("unexpected error: %v", err)
//			}
//
//			if !tt.shouldErr && result != tt.expected {
//				t.Errorf("divide(%.0f, %.0f) = %.0f, want %.0f", tt.a, tt.b, result, tt.expected)
//			}
//		})
//	}
// }

// ============ 5. BENCHMARKING ============
// func BenchmarkAdd(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		add(2, 3)
//	}
// }
//
// func BenchmarkIsEven(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		isEven(i)
//	}
// }

// Run with: go test -bench=.

// ============ 6. MOCKING PATTERN ============
type TestDatabase interface {
	GetUser(id int) (string, error)
}

type MockDatabase struct {
	GetUserFunc func(id int) (string, error)
}

func (m *MockDatabase) GetUser(id int) (string, error) {
	return m.GetUserFunc(id)
}

func getUserName(db Database, id int) (string, error) {
	return db.GetUser(id)
}

// func TestGetUserName(t *testing.T) {
//	mock := &MockDatabase{
//		GetUserFunc: func(id int) (string, error) {
//			if id == 1 {
//				return "Alice", nil
//			}
//			return "", fmt.Errorf("user not found")
//		},
//	}
//
//	name, err := getUserName(mock, 1)
//	if name != "Alice" || err != nil {
//		t.Errorf("expected Alice, got %s", name)
//	}
// }

// ============ 7. SETUP/TEARDOWN ============
// func TestWithSetupTeardown(t *testing.T) {
//	// Setup
//	tempDir := t.TempDir() // Creates temporary directory
//
//	// Test code
//	// ... use tempDir ...
//
//	// Teardown (automatic - TempDir cleans up)
// }

// ============ 8. HELPERS ============
// func assertEqual(t *testing.T, got, want interface{}) {
//	t.Helper() // Report error at caller's line
//	if got != want {
//		t.Errorf("got %v, want %v", got, want)
//	}
// }

// ============ 9. CONCURRENT TESTING ============
// func TestConcurrent(t *testing.T) {
//	var wg sync.WaitGroup
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//			// Test code
//		}(i)
//	}
//
//	wg.Wait()
// }

// ============ COURSE 10 MAIN FUNCTION ============
func courseTenDemo() {
	fmt.Println("=== TESTING IN GO ===\n")

	fmt.Println("TEST FILE STRUCTURE:")
	fmt.Println("---\n")

	fmt.Println(`
Go has built-in testing in the testing package.

File naming convention:
- my_code.go        - Implementation
- my_code_test.go   - Tests for my_code

Test function signature:
func TestFunctionName(t *testing.T)

Benchmark function signature:
func BenchmarkFunctionName(b *testing.B)
`)
	fmt.Println()

	fmt.Println("BASIC TEST:")
	fmt.Println("---")
	fmt.Println(`
func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("add(2, 3) = %d, want %d", result, expected)
	}
}
`)
	fmt.Println()

	fmt.Println("TABLE-DRIVEN TESTS (RECOMMENDED):")
	fmt.Println("---")
	fmt.Println(`
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
		{"zero", 0, 0, 0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
`)
	fmt.Println()

	fmt.Println("TESTING ERRORS:")
	fmt.Println("---")
	fmt.Println(`
func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)
	
	if err == nil {
		t.Error("expected error, got nil")
	}
	
	if err.Error() != "division by zero" {
		t.Errorf("wrong error message: %v", err)
	}
}
`)
	fmt.Println()

	fmt.Println("BENCHMARKING:")
	fmt.Println("---")
	fmt.Println(`
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 3)
	}
}

// Run benchmarks:
go test -bench=.
go test -bench=BenchmarkAdd -benchtime=10s
go test -bench=. -benchmem  // Memory stats
`)
	fmt.Println()

	fmt.Println("MOCKING:")
	fmt.Println("---")
	fmt.Println(`
// Interface to mock
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Mock implementation
type MockReader struct {
	data []byte
	pos  int
}

func (m *MockReader) Read(p []byte) (int, error) {
	if m.pos >= len(m.data) {
		return 0, io.EOF
	}
	n := copy(p, m.data[m.pos:])
	m.pos += n
	return n, nil
}

// Test using mock
func TestFunction(t *testing.T) {
	mock := &MockReader{data: []byte("test")}
	// Use mock in test
}
`)
	fmt.Println()

	fmt.Println("TEST HELPERS:")
	fmt.Println("---")
	fmt.Println(`
func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper() // Report error at caller's line, not helper's
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestUsing Helper(t *testing.T) {
	result := add(2, 3)
	assertEqual(t, result, 5) // Error reported at this line
}
`)
	fmt.Println()

	fmt.Println("SUBTESTS:")
	fmt.Println("---")
	fmt.Println(`
func TestMain(t *testing.T) {
	t.Run("subtest 1", func(t *testing.T) {
		// Subtest 1
	})
	
	t.Run("subtest 2", func(t *testing.T) {
		// Subtest 2
	})
}

// Run specific subtest:
go test -run TestMain/subtest1
`)
	fmt.Println()

	fmt.Println("SETUP AND TEARDOWN:")
	fmt.Println("---")
	fmt.Println(`
func TestWithSetup(t *testing.T) {
	// Setup
	tempDir := t.TempDir() // Temporary directory (auto-cleaned)
	
	// Test
	file := filepath.Join(tempDir, "test.txt")
	os.WriteFile(file, []byte("test"), 0644)
	
	// Teardown (automatic)
}
`)
	fmt.Println()

	fmt.Println("COVERAGE:")
	fmt.Println("---")
	fmt.Println(`
// Run with coverage report:
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out  // View in browser

// Achieve >80% coverage for good quality
`)
	fmt.Println()

	fmt.Println("PARALLEL TESTS:")
	fmt.Println("---")
	fmt.Println(`
func TestParallel(t *testing.T) {
	t.Parallel() // Run in parallel with other parallel tests
	
	// Test code
}

// Run tests in parallel:
go test -parallel 4  // Use 4 cores
`)
	fmt.Println()

	fmt.Println("COMMANDS:")
	fmt.Println("---")
	fmt.Println("go test                         - Run all tests")
	fmt.Println("go test ./...                   - Test all packages")
	fmt.Println("go test -v                      - Verbose output")
	fmt.Println("go test -run TestName           - Run specific test")
	fmt.Println("go test -bench=.                - Run benchmarks")
	fmt.Println("go test -cover                  - Show coverage %")
	fmt.Println("go test -parallel 4             - Run in parallel")
	fmt.Println()

	fmt.Println("BEST PRACTICES:")
	fmt.Println("---")
	fmt.Println("✓ Use table-driven tests for multiple cases")
	fmt.Println("✓ Test edge cases and error conditions")
	fmt.Println("✓ Use t.Helper() in helper functions")
	fmt.Println("✓ Keep tests focused and independent")
	fmt.Println("✓ Name tests clearly (TestFunctionName_Case)")
	fmt.Println("✓ Aim for >80% code coverage")
	fmt.Println("✓ Test interfaces, not implementations")
	fmt.Println("✓ Use mocks for external dependencies")
	fmt.Println("✓ Run tests before committing")
	fmt.Println("✓ Write tests as you write code")
	fmt.Println()

	fmt.Println("=== END OF TESTING ===")
}

// Example test for documentation
func ExampleAdd() {
	result := add(2, 3)
	fmt.Println(result)
	// Output: 5
}

// KEY TAKEAWAYS:
// 1. Testing is built into Go's standard library
// 2. Test files end with _test.go
// 3. Test functions start with Test
// 4. Use table-driven tests for multiple cases
// 5. t.Errorf() for test failures with message
// 6. t.Fatalf() to stop test immediately
// 7. t.Helper() marks helper function line in error
// 8. t.Run() for subtests
// 9. Benchmark functions for performance testing
// 10. t.Parallel() runs tests in parallel
// 11. t.TempDir() creates temporary test directory
// 12. Mock interfaces for testing
// 13. TestMain() for setup/teardown
// 14. Example functions double as tests and documentation
// 15. go test -cover shows coverage percentage
// 16. go test -coverprofile=file.out shows detailed coverage
// 17. Aim for >80% coverage
// 18. Test edge cases and error paths
// 19. Keep tests focused and independent
// 20. Write tests as you write code
