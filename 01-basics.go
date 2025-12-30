package main

import (
	"fmt"
)

// This is the first course file - Learn Go Basics
// Topics covered:
// 1. Package declaration and imports
// 2. Variables and constants
// 3. Data types (int, string, float, bool, arrays, slices, maps)
// 4. Type conversion
// 5. Control flow (if/else, loops)
// 6. Operators

// Demonstrating constants
const (
	// Untyped constants - Go determines type when used
	Pi           = 3.14159
	Name         = "Go Programming"
	BirthdayYear = 2009 // Go was created in 2009
)

// Variables with explicit types
var (
	// Global variables can be declared this way
	globalCounter int    = 0
	globalName    string = "Global Variable"
	isProduction  bool   = true
)

func courseOne() {
	fmt.Println("\n=== COURSE 1: GO BASICS ===\n")

	// ============ 1. VARIABLES ============
	fmt.Println("1. VARIABLES")
	fmt.Println("---")

	// Method 1: Declare with var keyword
	var age int
	age = 25
	fmt.Printf("age (with var): %v (type: %T)\n", age, age)

	// Method 2: Short declaration (only inside functions)
	name := "Alice"
	fmt.Printf("name (with :=): %v (type: %T)\n", name, name)

	// Method 3: Multiple variables
	var x, y, z int = 1, 2, 3
	fmt.Printf("x=%v, y=%v, z=%v\n", x, y, z)

	// Method 4: Blank identifier (discard value)
	_, count := divideWithRemainder(10, 3)
	fmt.Printf("Remainder of 10/3: %v\n\n", count)

	// ============ 2. DATA TYPES ============
	fmt.Println("2. DATA TYPES")
	fmt.Println("---")

	// Integers - multiple sizes
	var int8Var int8 = 127 // Range: -128 to 127
	var int64Var int64 = 9223372036854775807
	var uint32Var uint32 = 4294967295 // Unsigned, range: 0 to 4294967295
	fmt.Printf("int8: %v, int64: %v, uint32: %v\n", int8Var, int64Var, uint32Var)

	// Floating point
	var floatNum float32 = 3.14
	var doubleNum float64 = 3.14159265359
	fmt.Printf("float32: %v, float64: %v\n", floatNum, doubleNum)

	// Strings
	simpleString := "Hello, Go!"
	multilineString := `This is a
raw string that
preserves formatting`
	fmt.Printf("Simple: %v\nMultiline:\n%v\n\n", simpleString, multilineString)

	// Boolean
	isProgrammer := true
	fmt.Printf("Is Programmer: %v\n\n", isProgrammer)

	// ============ 3. ARRAYS AND SLICES ============
	fmt.Println("3. ARRAYS AND SLICES")
	fmt.Println("---")

	// Arrays - fixed size
	var fruits [3]string = [3]string{"Apple", "Banana", "Orange"}
	fmt.Printf("Array: %v, length: %v\n", fruits, len(fruits))

	// Array shorthand
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers array: %v\n", numbers)

	// Slices - dynamic size (MORE commonly used than arrays)
	var colors []string = []string{"Red", "Green", "Blue"}
	fmt.Printf("Slice: %v, length: %v, capacity: %v\n", colors, len(colors), cap(colors))

	// Slice operations
	colors = append(colors, "Yellow") // Add element
	fmt.Printf("After append: %v\n", colors)

	subSlice := colors[1:3] // Get elements from index 1 to 3 (exclusive)
	fmt.Printf("Subslice [1:3]: %v\n", subSlice)

	// Create slice with make (specify length and capacity)
	emptySlice := make([]int, 5, 10) // length=5, capacity=10
	fmt.Printf("Empty slice: %v, len=%v, cap=%v\n\n", emptySlice, len(emptySlice), cap(emptySlice))

	// ============ 4. MAPS ============
	fmt.Println("4. MAPS (Key-Value Pairs)")
	fmt.Println("---")

	// Declare and initialize map
	capitals := map[string]string{
		"USA":    "Washington",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Printf("Capitals: %v\n", capitals)

	// Add to map
	capitals["Brazil"] = "Brasília"
	fmt.Printf("After adding Brazil: %v\n", capitals)

	// Access value
	fmt.Printf("Capital of France: %v\n", capitals["France"])

	// Check if key exists
	value, exists := capitals["Italy"]
	fmt.Printf("Italy capital: %v, exists: %v\n", value, exists)

	// Delete from map
	delete(capitals, "USA")
	fmt.Printf("After deleting USA: %v\n\n", capitals)

	// ============ 5. TYPE CONVERSION ============
	fmt.Println("5. TYPE CONVERSION")
	fmt.Println("---")

	intValue := 42
	floatValue := float64(intValue)
	fmt.Printf("Int to Float: %v (type: %T)\n", floatValue, floatValue)

	stringValue := "Hello"
	byteSlice := []byte(stringValue)
	fmt.Printf("String to bytes: %v\n", byteSlice)

	back := string([]byte{72, 101, 108, 108, 111})
	fmt.Printf("Bytes to string: %v\n\n", back)

	// ============ 6. CONTROL FLOW - IF/ELSE ============
	fmt.Println("6. CONTROL FLOW - IF/ELSE")
	fmt.Println("---")

	temperature := 25

	if temperature < 0 {
		fmt.Println("Freezing!")
	} else if temperature < 15 {
		fmt.Println("Cold")
	} else if temperature < 25 {
		fmt.Println("Warm")
	} else {
		fmt.Println("Hot!")
	}

	// If with initialization (variable scope limited to if block)
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or lower")
	}
	// fmt.Println(score) // ERROR: score not defined here

	fmt.Println()

	// ============ 7. LOOPS ============
	fmt.Println("7. LOOPS")
	fmt.Println("---")

	// For loop - traditional style
	fmt.Print("Traditional for loop (0-4): ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	// For loop - while style
	counter := 0
	fmt.Print("While-style loop: ")
	for counter < 5 {
		fmt.Printf("%v ", counter)
		counter++
	}
	fmt.Println()

	// For loop - infinite (with break)
	fmt.Print("Infinite loop with break: ")
	loopCount := 0
	for {
		if loopCount >= 3 {
			break
		}
		fmt.Printf("%v ", loopCount)
		loopCount++
	}
	fmt.Println()

	// Range loop - iterating over slice
	words := []string{"Go", "is", "awesome"}
	fmt.Print("Range over slice: ")
	for i, word := range words {
		fmt.Printf("[%v]=%v ", i, word)
	}
	fmt.Println()

	// Range loop - iterating over map
	fmt.Println("Range over map:")
	person := map[string]string{
		"name": "John",
		"city": "New York",
		"job":  "Developer",
	}
	for key, value := range person {
		fmt.Printf("  %v: %v\n", key, value)
	}
	fmt.Println()

	// ============ 8. OPERATORS ============
	fmt.Println("8. OPERATORS")
	fmt.Println("---")

	a, b := 10, 3

	// Arithmetic operators
	fmt.Printf("Addition: %v + %v = %v\n", a, b, a+b)
	fmt.Printf("Subtraction: %v - %v = %v\n", a, b, a-b)
	fmt.Printf("Multiplication: %v * %v = %v\n", a, b, a*b)
	fmt.Printf("Division: %v / %v = %v\n", a, b, a/b)
	fmt.Printf("Modulo: %v %% %v = %v\n", a, b, a%b)

	// Comparison operators
	fmt.Printf("Equal: %v == %v = %v\n", a, b, a == b)
	fmt.Printf("Not equal: %v != %v = %v\n", a, b, a != b)
	fmt.Printf("Greater: %v > %v = %v\n", a, b, a > b)
	fmt.Printf("Less: %v < %v = %v\n", a, b, a < b)

	// Logical operators
	x1, x2 := true, false
	fmt.Printf("AND: %v && %v = %v\n", x1, x2, x1 && x2)
	fmt.Printf("OR: %v || %v = %v\n", x1, x2, x1 || x2)
	fmt.Printf("NOT: !%v = %v\n", x1, !x1)

	fmt.Println("\n=== END OF BASICS ===")
}

// Helper function to demonstrate blank identifier usage
func divideWithRemainder(dividend, divisor int) (int, int) {
	return dividend / divisor, dividend % divisor
}

// KEY TAKEAWAYS:
// 1. Go has explicit typing (unlike Python)
// 2. Use := for short declarations inside functions
// 3. Slices are more useful than arrays for most purposes
// 4. Maps are unordered key-value stores
// 5. For loops are the ONLY loop construct in Go
// 6. Error handling is explicit (we'll cover this later)
// 7. Go is opinionated - there's usually one way to do things
// 8. Variable names should be short and descriptive
// 9. Exported names (capitalize first letter) are public globally
// 10. Unexported names (lowercase) are private to the package
