package main

import (
	"fmt"
)

// COURSE 3: STRUCTS AND INTERFACES
// Topics covered:
// 1. Struct definition and initialization
// 2. Struct fields and visibility
// 3. Receiver functions (methods)
// 4. Pointer receivers
// 5. Interfaces
// 6. Type assertion
// 7. Embedding (composition)
// 8. Value vs pointer semantics

// ============ 1. BASIC STRUCT ============
type Person struct {
	Name string
	Age  int
	City string
}

// ============ 2. STRUCT WITH METHODS ============
type Rectangle struct {
	Width  float64
	Height float64
}

// Method with value receiver (creates copy, cannot modify)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Method with pointer receiver (can modify original)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// ============ 3. INTERFACE DEFINITION ============
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ============ 4. CIRCLE STRUCT IMPLEMENTING SHAPE ============
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// ============ 5. TRIANGLE STRUCT IMPLEMENTING SHAPE ============
type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	// Heron's formula
	s := (t.SideA + t.SideB + t.SideC) / 2
	return 3.14159 * (s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// ============ 6. READER INTERFACE (common in Go) ============
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter combines multiple interfaces
type ReadWriter interface {
	Reader
	Writer
}

// ============ 7. EMBEDDING (COMPOSITION) ============
type VehicleComp struct {
	Brand string
	Year  int
}

// CarComp embeds VehicleComp (inherits its fields and methods)
type CarComp struct {
	VehicleComp
	Model string
	Doors int
}

func (v VehicleComp) Display() string {
	return fmt.Sprintf("%d %s", v.Year, v.Brand)
}

// ============ 8. EMPTY INTERFACE ============
// interface{} can hold any type - useful for generic data
type DataStore struct {
	Data map[string]interface{}
}

func (db *DataStore) Store(key string, value interface{}) {
	if db.Data == nil {
		db.Data = make(map[string]interface{})
	}
	db.Data[key] = value
}

func (db *DataStore) Retrieve(key string) (interface{}, bool) {
	value, exists := db.Data[key]
	return value, exists
}

// ============ 9. FUNCTION THAT TAKES INTERFACE ============
// This function works with ANY type that implements Reader
func ProcessData(r Reader) {
	buffer := make([]byte, 10)
	n, _ := r.Read(buffer)
	fmt.Printf("Read %d bytes\n", n)
}

// ============ 10. TYPE ASSERTION ============
// Type assertion is used to extract concrete type from interface
func PrintInterface(data interface{}) {
	switch v := data.(type) {
	case string:
		fmt.Printf("String: %s\n", v)
	case int:
		fmt.Printf("Integer: %d\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	case Person:
		fmt.Printf("Person: %s, Age: %d\n", v.Name, v.Age)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// ============ 11. CUSTOM STRINGER INTERFACE ============
type Animal struct {
	Name   string
	Type   string
	Sounds []string
}

// Implementing the Stringer interface (fmt.Stringer)
// When you print an object with %v, it calls String() if available
func (a Animal) String() string {
	return fmt.Sprintf("%s (a %s) says: %v", a.Name, a.Type, a.Sounds)
}

// ============ COURSE THREE MAIN FUNCTION ============
func courseThree() {
	fmt.Println("=== STRUCTS AND INTERFACES COURSE ===\n")

	// ============ 1. STRUCT BASICS ============
	fmt.Println("1. STRUCT BASICS")
	fmt.Println("---")

	// Method 1: Declare and initialize with field names
	person1 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Printf("Person 1: %+v\n", person1) // %+v includes field names

	// Method 2: Declare and initialize with positional args
	person2 := Person{"Bob", 25, "Los Angeles"}
	fmt.Printf("Person 2: %v\n", person2)

	// Method 3: Initialize without values (zero-initialized)
	person3 := Person{}
	fmt.Printf("Person 3 (zero values): %v\n", person3)

	// Field access
	fmt.Printf("Person 1 name: %s\n\n", person1.Name)

	// ============ 2. METHODS (RECEIVER FUNCTIONS) ============
	fmt.Println("2. METHODS - FUNCTIONS WITH RECEIVERS")
	fmt.Println("---")

	rect := Rectangle{Width: 5, Height: 10}
	fmt.Printf("Rectangle: %v x %v\n", rect.Width, rect.Height)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n\n", rect.Perimeter())

	// ============ 3. VALUE VS POINTER RECEIVERS ============
	fmt.Println("3. VALUE VS POINTER RECEIVERS")
	fmt.Println("---")

	rect2 := Rectangle{Width: 2, Height: 3}
	fmt.Printf("Original: %v x %v\n", rect2.Width, rect2.Height)

	// This creates a copy, doesn't modify original
	rect2.Scale(2)
	fmt.Printf("After Scale(2): %v x %v\n\n", rect2.Width, rect2.Height)

	// ============ 4. INTERFACES ============
	fmt.Println("4. INTERFACES")
	fmt.Println("---")

	// Different shapes implementing same interface
	circle := Circle{Radius: 3}
	rectangle := Rectangle{Width: 4, Height: 5}
	triangle := Triangle{SideA: 3, SideB: 4, SideC: 5}

	shapes := []Shape{circle, rectangle, triangle}

	fmt.Println("All shapes and their properties:")
	for i, shape := range shapes {
		fmt.Printf("[%d] Area: %.2f, Perimeter: %.2f\n", i, shape.Area(), shape.Perimeter())
	}
	fmt.Println()

	// ============ 5. EMBEDDING (COMPOSITION) ============
	fmt.Println("5. EMBEDDING (COMPOSITION)")
	fmt.Println("---")

	car := CarComp{
		VehicleComp: VehicleComp{Brand: "Toyota", Year: 2022},
		Model:       "Camry",
		Doors:       4,
	}

	fmt.Printf("Car Model: %s\n", car.Model)
	fmt.Printf("Vehicle Info: %s\n", car.Display()) // Inherited method
	fmt.Printf("Full Info: %d %s %s\n", car.Year, car.Brand, car.Model)
	fmt.Println()

	// ============ 6. EMPTY INTERFACE ============
	fmt.Println("6. EMPTY INTERFACE (STORE ANY TYPE)")
	fmt.Println("---")

	db := &DataStore{}
	db.Store("name", "Charlie")
	db.Store("age", 35)
	db.Store("salary", 75000.50)
	db.Store("active", true)

	keys := []string{"name", "age", "salary", "active"}
	fmt.Println("Database contents:")
	for _, key := range keys {
		value, _ := db.Retrieve(key)
		fmt.Printf("  %s: %v (type: %T)\n", key, value, value)
	}
	fmt.Println()

	// ============ 7. TYPE ASSERTION ============
	fmt.Println("7. TYPE ASSERTION")
	fmt.Println("---")

	testData := []interface{}{
		"Hello",
		42,
		3.14,
		Person{Name: "David", Age: 28, City: "Chicago"},
	}

	fmt.Println("Type assertion examples:")
	for _, data := range testData {
		PrintInterface(data)
	}
	fmt.Println()

	// ============ 8. STRINGER INTERFACE ============
	fmt.Println("8. STRINGER INTERFACE (CUSTOM STRING REPRESENTATION)")
	fmt.Println("---")

	dog := Animal{
		Name:   "Rex",
		Type:   "Dog",
		Sounds: []string{"Woof", "Bark", "Growl"},
	}

	cat := Animal{
		Name:   "Whiskers",
		Type:   "Cat",
		Sounds: []string{"Meow", "Purr", "Hiss"},
	}

	// When using %v with objects that implement Stringer, it uses String() method
	fmt.Printf("%v\n", dog)
	fmt.Printf("%v\n\n", cat)

	// ============ 9. INTERFACE SATISFACTION ============
	fmt.Println("9. INTERFACE SATISFACTION")
	fmt.Println("---")

	// Check if type implements interface (compile-time check)
	// This line ensures Circle implements Shape, fails at compile if it doesn't
	var _ Shape = circle
	var _ Shape = rectangle
	var _ Shape = triangle

	// You can also do this with pointer receivers
	var _ Shape = &rectangle

	fmt.Println("✓ All shapes implement Shape interface\n")

	// ============ 10. MULTIPLE INTERFACES ============
	fmt.Println("10. OBJECT SATISFYING MULTIPLE INTERFACES")
	fmt.Println("---")

	// An object can satisfy multiple interfaces
	multiShapes := []Shape{circle, rectangle}
	fmt.Printf("Multiple shapes: %d shapes satisfy Shape interface\n", len(multiShapes))

	// But they don't all implement Reader interface
	// (we don't have Read methods defined)
	fmt.Println()

	// ============ 11. COMMON GO INTERFACES ============
	fmt.Println("11. COMMON GO INTERFACES")
	fmt.Println("---")

	fmt.Println("Common interfaces in Go:")
	fmt.Println("  - fmt.Stringer: String() string")
	fmt.Println("  - io.Reader: Read(p []byte) (n int, err error)")
	fmt.Println("  - io.Writer: Write(p []byte) (n int, err error)")
	fmt.Println("  - error: Error() string")
	fmt.Println("  - json.Marshaler: MarshalJSON() ([]byte, error)")
	fmt.Println("  - json.Unmarshaler: UnmarshalJSON([]byte) error")
	fmt.Println()

	fmt.Println("=== END OF STRUCTS AND INTERFACES ===")
}

// KEY TAKEAWAYS:
// 1. Structs are collections of named fields
// 2. Methods are functions with a receiver - they attach to types
// 3. Pointer receivers can modify the receiver; value receivers cannot
// 4. Interfaces define contracts (what methods an type must have)
// 5. Any type that implements all methods of an interface satisfies it
// 6. Go uses implicit interface satisfaction (no "implements" keyword)
// 7. empty interface{} can hold any value - useful for generic code
// 8. Type assertion lets you extract the concrete type from an interface
// 9. Use type switches for different behavior on different types
// 10. Composition (embedding) is preferred over inheritance
// 11. Keep interfaces small and focused (typically 1-3 methods)
// 12. Stringer interface (String()) customizes printing
// 13. Interface values can be nil (both the interface and its value)
// 14. Reader and Writer interfaces are fundamental in Go
// 15. Error is just an interface - any type with Error() method works
