// internal/data/database.go
package data

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// OpenDB creates a new database connection pool.
func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings for performance and reliability.
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(15 * time.Minute)

	// Verify the connection to the database is still alive. This is a crucial
	// step to ensure that the database is reachable on startup.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
