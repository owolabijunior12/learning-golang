package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// COURSE 6: HTTP SERVERS AND REST APIs
// Topics covered:
// 1. HTTP server basics
// 2. Request and response handling
// 3. Routing
// 4. JSON encoding/decoding
// 5. Query parameters
// 6. URL parameters
// 7. Form data
// 8. Headers
// 9. Middleware patterns
// 10. Status codes

// ============ 1. REQUEST/RESPONSE TYPES ============
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// ============ 2. SIMPLE HANDLER ============
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
}

// ============ 3. JSON RESPONSE HANDLER ============
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := APIResponse{
		Success: true,
		Message: "JSON response successful",
		Data: map[string]string{
			"version": "1.0",
			"status":  "running",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// ============ 4. GET USER BY ID ============
// In-memory database for demo
var users = map[int]User{
	1: {ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30},
	2: {ID: 2, Name: "Bob", Email: "bob@example.com", Age: 25},
	3: {ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 35},
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error:   "Missing user ID",
		})
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error:   "Invalid user ID",
		})
		return
	}

	user, exists := users[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error:   "User not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "User found",
		Data:    user,
	})
}

// ============ 5. CREATE USER (POST) ============
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error:   "Only POST method allowed",
		})
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error:   "Invalid JSON",
		})
		return
	}

	// Assign new ID
	user.ID = len(users) + 1
	users[user.ID] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "User created",
		Data:    user,
	})
}

// ============ 6. LIST ALL USERS ============
func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Users retrieved",
		Data:    userList,
	})
}

// ============ 7. QUERY PARAMETERS ============
func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get query parameters
	name := r.URL.Query().Get("name")
	minAge := r.URL.Query().Get("minAge")
	maxAge := r.URL.Query().Get("maxAge")

	var minAgeInt, maxAgeInt int = 0, 150
	if minAge != "" {
		minAgeInt, _ = strconv.Atoi(minAge)
	}
	if maxAge != "" {
		maxAgeInt, _ = strconv.Atoi(maxAge)
	}

	var results []User
	for _, user := range users {
		if (name == "" || strings.Contains(strings.ToLower(user.Name), strings.ToLower(name))) &&
			user.Age >= minAgeInt && user.Age <= maxAgeInt {
			results = append(results, user)
		}
	}

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: fmt.Sprintf("Found %d users", len(results)),
		Data:    results,
	})
}

// ============ 8. FORM DATA ============
func formHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")

		json.NewEncoder(w).Encode(APIResponse{
			Success: true,
			Message: "Form received",
			Data: map[string]string{
				"name":  name,
				"email": email,
			},
		})
	} else {
		fmt.Fprintf(w, `<form method="post">
			Name: <input type="text" name="name"><br>
			Email: <input type="email" name="email"><br>
			<input type="submit" value="Submit">
		</form>`)
	}
}

// ============ 9. REQUEST HEADERS ============
func headersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	headers := make(map[string]string)
	for key, values := range r.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Request headers",
		Data:    headers,
	})
}

// ============ 10. REQUEST BODY ============
func echoBytesHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Echo",
		Data: map[string]interface{}{
			"received": string(body),
			"length":   len(body),
		},
	})
}

// ============ 11. MIDDLEWARE PATTERN ============
// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[LOG] %s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Auth middleware (simple example)
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer valid-token" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(APIResponse{
				Success: false,
				Error:   "Unauthorized",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ============ COURSE SIX MAIN FUNCTION (Demo, not executed) ============
// Note: This demonstrates setup only. To actually run a server, uncomment below.
func courseSix() {
	fmt.Println("=== HTTP SERVERS AND REST APIs ===\n")

	fmt.Println("HTTP Server Setup Example:")
	fmt.Println("---\n")

	fmt.Println(`
// To run this server, create main function:
func main() {
	// Basic handlers
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/users", listUsersHandler)
	http.HandleFunc("/users/create", createUserHandler)
	http.HandleFunc("/users/", getUserHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/headers", headersHandler)
	http.HandleFunc("/echo", echoBytesHandler)
	
	// With middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(APIResponse{
			Success: true,
			Message: "Protected resource",
		})
	})
	
	// Apply middleware
	handler := loggingMiddleware(authMiddleware(mux))
	
	// Start server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}

API ENDPOINTS:
GET  /                    - Hello world
GET  /json               - JSON response
GET  /users              - List all users
POST /users/create       - Create new user
GET  /users/{id}         - Get user by ID
GET  /search?name=...    - Search users
POST /form               - Form submission
GET  /headers            - Show request headers
POST /echo               - Echo request body
GET  /protected          - Protected endpoint (needs auth)

EXAMPLES:

1. List users:
   GET http://localhost:8080/users

2. Get specific user:
   GET http://localhost:8080/users/1

3. Create user:
   POST http://localhost:8080/users/create
   Body: {"name":"John","email":"john@example.com","age":28}

4. Search:
   GET http://localhost:8080/search?name=alice&minAge=25

5. With authentication:
   GET http://localhost:8080/protected
   Headers: Authorization: Bearer valid-token
`)

	fmt.Println("\nCommon HTTP Status Codes:")
	fmt.Println("---")
	fmt.Println("200 OK              - Request successful")
	fmt.Println("201 Created         - Resource created")
	fmt.Println("204 No Content      - Success, no response body")
	fmt.Println("400 Bad Request     - Invalid request")
	fmt.Println("401 Unauthorized    - Authentication required")
	fmt.Println("403 Forbidden       - Authenticated but not allowed")
	fmt.Println("404 Not Found       - Resource doesn't exist")
	fmt.Println("500 Internal Error  - Server error")
	fmt.Println()

	fmt.Println("Common Content Types:")
	fmt.Println("---")
	fmt.Println("application/json    - JSON data")
	fmt.Println("text/plain          - Plain text")
	fmt.Println("text/html           - HTML")
	fmt.Println("application/form-data - Form submission")
	fmt.Println("application/xml     - XML data")
	fmt.Println()

	fmt.Println("=== END OF HTTP AND REST APIs ===")
}

// KEY TAKEAWAYS:
// 1. http.HandleFunc registers handler functions
// 2. Handler signature: func(w http.ResponseWriter, r *http.Request)
// 3. ResponseWriter is used to send response back
// 4. Request contains method, URL, headers, body, etc.
// 5. Always set Content-Type header
// 6. Use json.NewEncoder(w).Encode() to send JSON responses
// 7. json.NewDecoder(r.Body).Decode(&v) to parse JSON requests
// 8. r.Method to check request type (GET, POST, etc.)
// 9. r.URL.Query() for query parameters
// 10. r.FormValue() for form data (call r.ParseForm() first)
// 11. r.Header for request headers
// 12. io.ReadAll(r.Body) to read raw request body
// 13. http.ListenAndServe(":8080", nil) starts server
// 14. Use http.NewServeMux() for more control over routing
// 15. Middleware wraps handlers for cross-cutting concerns
// 16. Check method before processing (POST vs GET)
// 17. Always handle errors appropriately
// 18. Use proper status codes (200, 201, 400, 404, 500, etc.)
// 19. For real projects, use frameworks like Echo, Gin, or Chi
// 20. Test endpoints with curl, Postman, or Go's http tests
