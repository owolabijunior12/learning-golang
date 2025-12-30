package main

import (
	"errors"
	"fmt"
	"strconv"
)

// COURSE 2: FUNCTIONS AND ERROR HANDLING
// Topics covered:
// 1. Function declaration and parameters
// 2. Multiple return values
// 3. Named return values
// 4. Error handling (the Go way)
// 5. Variadic functions
// 6. Defer statement
// 7. Panic and recover
// 8. Function types and higher-order functions

// ============ 1. BASIC FUNCTION ============
// Function with parameters and single return value
func addBasics(a, b int) int {
	return a + b
}

// ============ 2. MULTIPLE RETURN VALUES ============
// This is very common in Go - especially for returning (value, error)
func divideBasics(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("division by zero")
	}
	return dividend / divisor, nil
}

// ============ 3. NAMED RETURN VALUES ============
// Return values can be named - useful for documentation
// Variables are pre-declared and zero-initialized
func calculateArea(width, height float64) (area float64, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return - returns named values (use sparingly)
}

// ============ 4. VARIADIC FUNCTIONS ============
// Accept variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// ============ 5. VARIADIC WITH MULTIPLE TYPES ============
func printAll(args ...interface{}) {
	for i, arg := range args {
		fmt.Printf("[%d] %v (type: %T)\n", i, arg, arg)
	}
}

// ============ 6. FUNCTION TYPES ============
// Functions can be assigned to variables and passed around
func multiply(a, b int) int {
	return a * b
}

// Function that takes another function as parameter
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// ============ 7. ERROR HANDLING ============
// Custom error type
type ValidationError struct {
	field   string
	message string
}

// Implement the error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in %s: %s", e.field, e.message)
}

// Function with comprehensive error handling
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			field:   "age",
			message: "age cannot be negative",
		}
	}
	if age > 150 {
		return ValidationError{
			field:   "age",
			message: "age is unrealistic",
		}
	}
	return nil
}

// ============ 8. STRING TO INT CONVERSION WITH ERROR ============
func stringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to convert '%s' to int: %w", s, err)
	}
	return num, nil
}

// ============ 9. DEFER STATEMENT ============
// Defer schedules a function to run at the end of current function
func demonstrateDefer() {
	fmt.Println("Start of function")

	defer fmt.Println("This runs last (deferred 1st)")
	defer fmt.Println("This runs second last (deferred 2nd)")
	defer fmt.Println("This runs third last (deferred 3rd)")

	fmt.Println("Middle of function")
}

// Real-world defer example - resource cleanup
func readFile(filename string) (string, error) {
	fmt.Printf("Opening file: %s\n", filename)
	// In real code, you'd open a file here

	// Defer ensures cleanup happens even if error occurs
	defer func() {
		fmt.Printf("Closing file: %s\n", filename)
	}()

	// Simulate reading file
	return "file contents", nil
}

// ============ 10. PANIC AND RECOVER ============
// Only use panic for truly exceptional circumstances!
func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("cannot divide by zero!")
	}

	return a / b
}

// ============ 11. CLOSURE EXAMPLE ============
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ============ MAIN FUNCTION ============
func courseTwo() {
	fmt.Println("=== FUNCTIONS AND ERROR HANDLING COURSE ===\n")

	// ============ 1. BASIC FUNCTIONS ============
	fmt.Println("1. BASIC FUNCTIONS")
	fmt.Println("---")
	result := addBasics(5, 3)
	fmt.Printf("addBasics(5, 3) = %v\n\n", result)

	// ============ 2. MULTIPLE RETURN VALUES ============
	fmt.Println("2. MULTIPLE RETURN VALUES")
	fmt.Println("---")
	quotient, err := divideBasics(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %v\n", quotient)
	}

	quotient, err = divideBasics(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()

	// ============ 3. NAMED RETURN VALUES ============
	fmt.Println("3. NAMED RETURN VALUES")
	fmt.Println("---")
	a, p := calculateArea(5, 4)
	fmt.Printf("Rectangle 5x4: Area = %v, Perimeter = %v\n\n", a, p)

	// ============ 4. VARIADIC FUNCTIONS ============
	fmt.Println("4. VARIADIC FUNCTIONS")
	fmt.Println("---")
	fmt.Printf("sum(1, 2, 3) = %v\n", sum(1, 2, 3))
	fmt.Printf("sum(1, 2, 3, 4, 5) = %v\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("sum() = %v\n", sum()) // Works even with no arguments

	// Passing slice as variadic
	numbers := []int{10, 20, 30}
	fmt.Printf("sum(slice...) = %v\n\n", sum(numbers...))

	// ============ 5. VARIADIC WITH MULTIPLE TYPES ============
	fmt.Println("5. VARIADIC WITH MULTIPLE TYPES")
	fmt.Println("---")
	printAll("Go", 42, true, 3.14, []string{"a", "b"})
	fmt.Println()

	// ============ 6. FUNCTION TYPES ============
	fmt.Println("6. FUNCTION TYPES AND HIGHER-ORDER FUNCTIONS")
	fmt.Println("---")

	// Assign function to variable
	var operation func(int, int) int = multiply
	fmt.Printf("operation(4, 5) = %v\n", operation(4, 5))

	// Pass function as argument
	result = applyOperation(6, 7, add)
	fmt.Printf("applyOperation(6, 7, add) = %v\n", result)

	result = applyOperation(6, 7, multiply)
	fmt.Printf("applyOperation(6, 7, multiply) = %v\n", result)

	// Return function from function
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Printf("double(5) = %v\n", double(5))
	fmt.Printf("triple(5) = %v\n\n", triple(5))

	// ============ 7. ERROR HANDLING ============
	fmt.Println("7. ERROR HANDLING")
	fmt.Println("---")

	testAges := []int{25, -5, 200, 45}
	for _, age := range testAges {
		if err := validateAge(age); err != nil {
			fmt.Printf("❌ Age %d: %v\n", age, err)
		} else {
			fmt.Printf("✓ Age %d: Valid\n", age)
		}
	}
	fmt.Println()

	// ============ 8. STRING TO INT CONVERSION ============
	fmt.Println("8. STRING CONVERSION WITH ERROR HANDLING")
	fmt.Println("---")

	testStrings := []string{"42", "abc", "-10", "0"}
	for _, str := range testStrings {
		num, err := stringToInt(str)
		if err != nil {
			fmt.Printf("❌ '%s': %v\n", str, err)
		} else {
			fmt.Printf("✓ '%s': %d\n", str, num)
		}
	}
	fmt.Println()

	// ============ 9. DEFER STATEMENT ============
	fmt.Println("9. DEFER STATEMENT")
	fmt.Println("---")
	demonstrateDefer()
	fmt.Println()

	// Real-world defer example
	content, err := readFile("data.txt")
	if err == nil {
		fmt.Printf("Read: %s\n", content)
	}
	fmt.Println()

	// ============ 10. PANIC AND RECOVER ============
	fmt.Println("10. PANIC AND RECOVER")
	fmt.Println("---")
	result = safeDivide(10, 2)
	fmt.Printf("safeDivide(10, 2) = %v\n", result)

	result = safeDivide(10, 0) // Will panic but recover
	fmt.Println()

	// ============ 11. CLOSURE ============
	fmt.Println("11. CLOSURE - FUNCTIONS CAPTURING VARIABLES")
	fmt.Println("---")
	counter1 := counter()
	fmt.Printf("counter1(): %v\n", counter1())
	fmt.Printf("counter1(): %v\n", counter1())
	fmt.Printf("counter1(): %v\n", counter1())

	counter2 := counter() // Separate counter
	fmt.Printf("counter2(): %v\n", counter2())
	fmt.Printf("counter2(): %v\n", counter2())
	fmt.Println()

	fmt.Println("=== END OF FUNCTIONS AND ERROR HANDLING ===")
}

// KEY TAKEAWAYS:
// 1. Functions can return multiple values (use for errors!)
// 2. Error handling in Go is explicit and visible
// 3. Always check for errors immediately after function calls
// 4. Named return values should be used carefully
// 5. Defer ensures cleanup code runs (like finally in other languages)
// 6. Use panic rarely - it's for exceptional situations
// 7. Custom error types allow flexible error handling
// 8. Functions are first-class citizens - assign and pass them
// 9. Closures capture variables from enclosing scope
// 10. The error interface is simple: type Error interface { Error() string }
// 11. Wrap errors with %w for error chain inspection
// 12. Use blank identifier _ to ignore unwanted return values
