// internal/data/models.go
package data

import (
	"database/sql"
	"time"
)

// Models is a wrapper for all our database models.
type Models struct {
	Admins AdminModel
	Users  UserModel
}

// NewModels is a helper function to initialize all our models at once.
func NewModels(db *sql.DB) Models {
	return Models{
		Admins: AdminModel{DB: db},
		Users:  UserModel{DB: db},
	}
}

// Admin represents an administrator with login credentials.
type Admin struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

// User represents an employee who can use the system.
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
