package main

import (
	"database/sql"
	"fmt"
)

// COURSE 7: SQL DATABASES (PostgreSQL, MySQL)
// Topics covered:
// 1. Database connection
// 2. Connection pooling
// 3. CRUD operations
// 4. Query results
// 5. Prepared statements
// 6. Transactions
// 7. Error handling
// 8. Best practices

// Note: This course demonstrates patterns. Actual DB connection requires:
// For PostgreSQL: "github.com/lib/pq"
// For MySQL: "github.com/go-sql-driver/mysql"

// ============ 1. USER MODEL ============
type DBUser struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// ============ 2. DATABASE WRAPPER ============
type SQLDatabase struct {
	conn *sql.DB
}

// ============ 3. CONNECT TO DATABASE ============
func NewSQLDatabase(dsn string) (*SQLDatabase, error) {
	// For PostgreSQL:
	// db, err := sql.Open("postgres", dsn)

	// For MySQL:
	// db, err := sql.Open("mysql", dsn)

	// For SQLite (easier for testing):
	// db, err := sql.Open("sqlite3", ":memory:")

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Set connection pool parameters
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SQLDatabase{conn: db}, nil
}

// ============ 4. CREATE TABLE ============
func (d *SQLDatabase) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER
	)`

	_, err := d.conn.Exec(query)
	return err
}

// ============ 5. INSERT USER ============
func (d *SQLDatabase) InsertUser(user DBUser) (int, error) {
	query := `INSERT INTO users (name, email, age) VALUES (?, ?, ?)`

	result, err := d.conn.Exec(query, user.Name, user.Email, user.Age)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

// ============ 6. GET USER BY ID ============
func (d *SQLDatabase) GetUserByID(id int) (*DBUser, error) {
	query := `SELECT id, name, email, age FROM users WHERE id = ?`

	var user DBUser
	err := d.conn.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ============ 7. GET ALL USERS ============
func (d *SQLDatabase) GetAllUsers() ([]DBUser, error) {
	query := `SELECT id, name, email, age FROM users ORDER BY id`

	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []DBUser

	for rows.Next() {
		var user DBUser
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// ============ 8. UPDATE USER ============
func (d *SQLDatabase) UpdateUser(id int, user DBUser) error {
	query := `UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?`

	result, err := d.conn.Exec(query, user.Name, user.Email, user.Age, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// ============ 9. DELETE USER ============
func (d *SQLDatabase) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := d.conn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// ============ 10. PREPARED STATEMENTS (PERFORMANCE) ============
func (d *SQLDatabase) GetUsersByAge(age int) ([]DBUser, error) {
	query := `SELECT id, name, email, age FROM users WHERE age = ? ORDER BY name`

	stmt, err := d.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []DBUser
	for rows.Next() {
		var user DBUser
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// ============ 11. TRANSACTIONS ============
func (d *SQLDatabase) TransferUsers(fromID, toID int, newName string) error {
	tx, err := d.conn.Begin()
	if err != nil {
		return err
	}

	// Delete first user
	_, err = tx.Exec("DELETE FROM users WHERE id = ?", fromID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update second user
	_, err = tx.Exec("UPDATE users SET name = ? WHERE id = ?", newName, toID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit if no errors
	return tx.Commit().Err()
}

// ============ 12. COUNT USERS ============
func (d *SQLDatabase) CountUsers() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users`

	err := d.conn.QueryRow(query).Scan(&count)
	return count, err
}

// ============ 13. CLOSE DATABASE ============
func (d *SQLDatabase) Close() error {
	return d.conn.Close()
}

// ============ COURSE SEVEN MAIN FUNCTION ============
func courseSeven() {
	fmt.Println("=== SQL DATABASES (PostgreSQL, MySQL) ===\n")

	fmt.Println("DATABASE SETUP EXAMPLES:")
	fmt.Println("---\n")

	fmt.Println("PostgreSQL Connection String:")
	fmt.Println(`db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/dbname?sslmode=disable")`)
	fmt.Println()

	fmt.Println("MySQL Connection String:")
	fmt.Println(`db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")`)
	fmt.Println()

	fmt.Println("SQLite Connection String:")
	fmt.Println(`db, err := sql.Open("sqlite3", "./test.db")`)
	fmt.Println()

	fmt.Println("CONNECTION POOLING:")
	fmt.Println("---")
	fmt.Println(`
db.SetMaxOpenConns(25)      // Maximum open connections
db.SetMaxIdleConns(5)       // Max idle (reusable) connections
db.SetConnMaxLifetime(...)  // Connection max lifetime
`)
	fmt.Println()

	fmt.Println("BASIC CRUD PATTERN:")
	fmt.Println("---")
	fmt.Println(`
// INSERT
result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
id, err := result.LastInsertId()

// SELECT ONE
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)

// SELECT MULTIPLE
rows, err := db.Query("SELECT id, name FROM users")
defer rows.Close()
for rows.Next() {
	var id int
	var name string
	err := rows.Scan(&id, &name)
	// process...
}

// UPDATE
result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", name, id)
rowsAffected, err := result.RowsAffected()

// DELETE
err := db.Exec("DELETE FROM users WHERE id = ?", id)
`)
	fmt.Println()

	fmt.Println("PREPARED STATEMENTS (Recommended):")
	fmt.Println("---")
	fmt.Println(`
stmt, err := db.Prepare("SELECT name FROM users WHERE id = ?")
defer stmt.Close()

var name string
err := stmt.QueryRow(userId).Scan(&name)

// Benefits:
// - Protection against SQL injection
// - Better performance (statement compiled once)
// - Reusable with different parameters
`)
	fmt.Println()

	fmt.Println("TRANSACTIONS:")
	fmt.Println("---")
	fmt.Println(`
tx, err := db.Begin()
if err != nil {
	return err
}

// Execute statements
_, err = tx.Exec("INSERT INTO...")
if err != nil {
	tx.Rollback()
	return err
}

_, err = tx.Exec("UPDATE...")
if err != nil {
	tx.Rollback()
	return err
}

// Commit if all successful
return tx.Commit().Err()
`)
	fmt.Println()

	fmt.Println("ERROR HANDLING:")
	fmt.Println("---")
	fmt.Println(`
if err == sql.ErrNoRows {
	// No row found
}

if err != nil {
	// Other database error
}

// Check rows affected
result, _ := db.Exec(...)
rowsAffected, _ := result.RowsAffected()
if rowsAffected == 0 {
	// No rows updated/deleted
}
`)
	fmt.Println()

	fmt.Println("BEST PRACTICES:")
	fmt.Println("---")
	fmt.Println("✓ Always use prepared statements")
	fmt.Println("✓ Close database connections properly")
	fmt.Println("✓ Use context for timeouts and cancellation")
	fmt.Println("✓ Defer rows.Close() to prevent resource leaks")
	fmt.Println("✓ Handle sql.ErrNoRows explicitly")
	fmt.Println("✓ Use transactions for related operations")
	fmt.Println("✓ Set connection pool limits")
	fmt.Println("✓ Add indexes for frequently queried columns")
	fmt.Println("✓ Use NULL values carefully in Go")
	fmt.Println("✓ Validate input before queries")
	fmt.Println()

	fmt.Println("COMMON LIBRARIES:")
	fmt.Println("---")
	fmt.Println("database/sql       - Standard library (basic)")
	fmt.Println("github.com/lib/pq  - PostgreSQL driver")
	fmt.Println("github.com/go-sql-driver/mysql - MySQL driver")
	fmt.Println("gorm.io/gorm       - ORM (higher level)")
	fmt.Println("sqlc               - Generate type-safe code from SQL")
	fmt.Println()

	fmt.Println("=== END OF SQL DATABASES ===")
}

// KEY TAKEAWAYS:
// 1. database/sql is the standard for database operations
// 2. Always call defer rows.Close() after queries
// 3. Use prepared statements for SQL injection protection
// 4. Check for sql.ErrNoRows explicitly
// 5. Transactions ensure consistency across multiple operations
// 6. Connection pooling improves performance
// 7. Query vs QueryRow: multiple vs single result
// 8. Scan converts database values to Go variables
// 9. LastInsertId() gets the ID of inserted row
// 10. RowsAffected() tells how many rows changed
// 11. Rollback on any error in transaction
// 12. Use context.Context for cancellation
// 13. Validate input to prevent SQL injection
// 14. NULL values in database need special handling (sql.NullString, etc.)
// 15. Keep connections open (don't create new for each query)
// 16. Index frequently queried columns
// 17. Use LIMIT for large result sets
// 18. Consider ORMs for complex applications
// 19. Test database operations thoroughly
// 20. Monitor connection pool stats in production
