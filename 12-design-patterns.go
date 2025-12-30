package main

import (
	"fmt"
)

// COURSE 12: MIDDLEWARE, DESIGN PATTERNS, AND ADVANCED PATTERNS
// Topics covered:
// 1. Middleware patterns
// 2. Dependency injection
// 3. Repository pattern
// 4. Service layer pattern
// 5. Builder pattern
// 6. Observer pattern
// 7. Strategy pattern
// 8. Factory pattern

import (
	"net/http"
	"time"
)

// ============ 1. MIDDLEWARE PATTERN ============
type Middleware func(http.Handler) http.Handler

// Chain middleware in order
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// Logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s %s\n", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

// Recovery middleware
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// ============ 2. DEPENDENCY INJECTION ============
type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(msg string) {
	fmt.Println(msg)
}

type Repository interface {
	GetUser(id int) (string, error)
}

type MockRepository struct {
	logger Logger
}

func (r *MockRepository) GetUser(id int) (string, error) {
	r.logger.Log("Getting user " + fmt.Sprint(id))
	return "User", nil
}

// Service receives dependencies (injection)
type UserService struct {
	repo   Repository
	logger Logger
}

func NewUserService(repo Repository, logger Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) GetUser(id int) (string, error) {
	s.logger.Log(fmt.Sprintf("UserService.GetUser(%d)", id))
	return s.repo.GetUser(id)
}

// ============ 3. REPOSITORY PATTERN ============
type UserRepository interface {
	Create(user interface{}) error
	GetByID(id int) (interface{}, error)
	Update(id int, user interface{}) error
	Delete(id int) error
	GetAll() ([]interface{}, error)
}

type MemoryUserRepository struct {
	data map[int]interface{}
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		data: make(map[int]interface{}),
	}
}

func (r *MemoryUserRepository) Create(user interface{}) error {
	// Store user
	return nil
}

func (r *MemoryUserRepository) GetByID(id int) (interface{}, error) {
	if user, ok := r.data[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (r *MemoryUserRepository) Update(id int, user interface{}) error {
	r.data[id] = user
	return nil
}

func (r *MemoryUserRepository) Delete(id int) error {
	delete(r.data, id)
	return nil
}

func (r *MemoryUserRepository) GetAll() ([]interface{}, error) {
	var users []interface{}
	for _, user := range r.data {
		users = append(users, user)
	}
	return users, nil
}

// ============ 4. BUILDER PATTERN ============
type QueryBuilder struct {
	query  string
	params []interface{}
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Select(fields string) *QueryBuilder {
	qb.query = "SELECT " + fields
	return qb
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.query += " FROM " + table
	return qb
}

func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	qb.query += " WHERE " + condition
	qb.params = append(qb.params, args...)
	return qb
}

func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
	qb.query += fmt.Sprintf(" LIMIT %d", n)
	return qb
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	return qb.query, qb.params
}

// ============ 5. STRATEGY PATTERN ============
type PaymentStrategy interface {
	Pay(amount float64) error
}

type CreditCardPayment struct {
	cardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Printf("Paid %.2f with credit card\n", amount)
	return nil
}

type PayPalPayment struct {
	email string
}

func (p *PayPalPayment) Pay(amount float64) error {
	fmt.Printf("Paid %.2f with PayPal (%s)\n", amount, p.email)
	return nil
}

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy: strategy}
}

func (pp *PaymentProcessor) Process(amount float64) error {
	return pp.strategy.Pay(amount)
}

// ============ 6. FACTORY PATTERN ============
type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c *Car) Drive() string {
	return "Driving car"
}

type Bicycle struct{}

func (b *Bicycle) Drive() string {
	return "Riding bicycle"
}

type VehicleFactory struct{}

func (vf *VehicleFactory) Create(vehicleType string) Vehicle {
	switch vehicleType {
	case "car":
		return &Car{}
	case "bicycle":
		return &Bicycle{}
	default:
		return nil
	}
}

// ============ 7. OBSERVER PATTERN ============
type Observer interface {
	Update(message string)
}

type Subject struct {
	observers []Observer
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Subscribe(obs Observer) {
	s.observers = append(s.observers, obs)
}

func (s *Subject) Unsubscribe(obs Observer) {
	// Remove observer
}

func (s *Subject) Notify(message string) {
	for _, obs := range s.observers {
		obs.Update(message)
	}
}

type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("%s received: %s\n", co.name, message)
}

// ============ 8. SINGLETON PATTERN ============
type DatabaseConnection struct {
	connectionString string
}

var instance *DatabaseConnection

func GetDatabaseConnection(connectionString string) *DatabaseConnection {
	if instance == nil {
		instance = &DatabaseConnection{
			connectionString: connectionString,
		}
	}
	return instance
}

// ============ COURSE TWELVE MAIN FUNCTION ============
func courseTwelve() {
	fmt.Println("=== MIDDLEWARE, DESIGN PATTERNS, AND ADVANCED PATTERNS ===\n")

	fmt.Println("MIDDLEWARE PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Middleware wraps a handler to add functionality
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check authentication
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Chaining middleware
handler := http.HandlerFunc(myHandler)
handler = Chain(handler, LoggingMiddleware, AuthMiddleware, RecoveryMiddleware)

http.Handle("/api", handler)
`)
	fmt.Println()

	fmt.Println("DEPENDENCY INJECTION:")
	fmt.Println("---")
	fmt.Println(`
// Constructor injection (preferred)
func NewUserService(repo Repository, logger Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

// Benefits:
// - Explicit dependencies
// - Easy to mock for testing
// - Clear what service needs
// - Loose coupling
`)
	fmt.Println()

	fmt.Println("REPOSITORY PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Abstracts data access
type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
	Delete(id int) error
	List() ([]*User, error)
}

// Benefits:
// - Swap implementations (memory, DB, etc.)
// - Easier testing with mocks
// - Centralized data access
// - Decouple from storage layer
`)
	fmt.Println()

	fmt.Println("BUILDER PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Complex object construction
query := NewQueryBuilder().
	Select("id, name, email").
	From("users").
	Where("age > ?", 18).
	Limit(10).
	Build()

// Benefits:
// - Clear, readable object construction
// - Optional parameters without overloading
// - Can validate in each step
// - Complex queries can be built step by step
`)
	fmt.Println()

	fmt.Println("STRATEGY PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Different algorithms, same interface
strategies := []PaymentStrategy{
	&CreditCardPayment{},
	&PayPalPayment{},
	&BitcoinPayment{},
}

// Client code doesn't care which
processor := NewPaymentProcessor(strategies[0])
processor.Process(100.00)

// Benefits:
// - Runtime algorithm selection
// - Easy to add new strategies
// - Clients don't need to know implementations
// - Encapsulates algorithms
`)
	fmt.Println()

	fmt.Println("FACTORY PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Centralized object creation
factory := &VehicleFactory{}

car := factory.Create("car")
bike := factory.Create("bicycle")

// Benefits:
// - Centralized creation logic
// - Easy to change construction
// - Encapsulates creation details
// - Clients only know interface
`)
	fmt.Println()

	fmt.Println("OBSERVER PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Event notification system
subject := NewSubject()

subject.Subscribe(&ConcreteObserver{name: "Observer1"})
subject.Subscribe(&ConcreteObserver{name: "Observer2"})

subject.Notify("Event happened!")

// Benefits:
// - Loose coupling between subject and observers
// - Dynamic subscriptions
// - Multiple observers notified at once
// - Good for event-driven systems
`)
	fmt.Println()

	fmt.Println("SINGLETON PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Single instance across application
db := GetDatabaseConnection("postgres://localhost")

// Use throughout application (same instance)

// Caution:
// - Can hide dependencies
// - Hard to test
// - Consider dependency injection instead
// - Limited use in Go
`)
	fmt.Println()

	fmt.Println("SERVICE LAYER PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// Business logic separate from HTTP handling
type UserService interface {
	CreateUser(name, email string) (*User, error)
	GetUser(id int) (*User, error)
	UpdateUser(id int, name, email string) error
	DeleteUser(id int) error
}

// API handler uses service
func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	json.NewDecoder(r.Body).Decode(&req)
	
	user, err := h.userService.CreateUser(req.Name, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	json.NewEncoder(w).Encode(user)
}

// Benefits:
// - Business logic testable without HTTP
// - Can reuse service for CLI, gRPC, etc.
// - Clear separation of concerns
`)
	fmt.Println()

	fmt.Println("BEST PRACTICES:")
	fmt.Println("---")
	fmt.Println("✓ Use interfaces for abstraction")
	fmt.Println("✓ Inject dependencies in constructors")
	fmt.Println("✓ Keep middleware composable")
	fmt.Println("✓ Separate HTTP from business logic")
	fmt.Println("✓ Use repository for data access")
	fmt.Println("✓ Avoid God objects (objects doing too much)")
	fmt.Println("✓ Follow Single Responsibility Principle")
	fmt.Println("✓ Use composition over inheritance")
	fmt.Println("✓ Make zero values useful")
	fmt.Println("✓ Document expected interfaces")
	fmt.Println()

	fmt.Println("=== END OF DESIGN PATTERNS ===")
}

// KEY TAKEAWAYS:
// 1. Middleware pattern for cross-cutting concerns
// 2. Dependency injection for loose coupling
// 3. Repository pattern for data access abstraction
// 4. Service layer for business logic
// 5. Builder pattern for complex object construction
// 6. Strategy pattern for algorithm selection
// 7. Factory pattern for object creation
// 8. Observer pattern for event systems
// 9. Composition preferred over inheritance
// 10. Interfaces for abstraction and testing
// 11. Keep packages focused and single-purpose
// 12. Avoid cyclic dependencies
// 13. Use receivers for methods
// 14. Make zero values useful
// 15. Error handling as values, not exceptions
// 16. Testing should be easy with proper design
// 17. Mocking should be possible (use interfaces)
// 18. Clear, explicit code over clever code
// 19. SOLID principles apply to Go
// 20. Go's simplicity favors simple patterns
