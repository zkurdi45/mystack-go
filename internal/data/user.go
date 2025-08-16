// internal/data/users.go
package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type UserModel struct {
	DB *sql.DB
}

// GetByID retrieves a user by their ID.
func (m *UserModel) GetByID(id int64) (*User, error) {
	query := `
		SELECT id, email, status, created_at
		FROM users
		WHERE id = $1`

	var user User

	err := m.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}

// ... (GetAll, UpdateStatus, GetOrCreateByEmail, and GetByEmail methods remain the same) ...
// GetAll retrieves all users from the database, ordered by creation date.
func (m *UserModel) GetAll() ([]*User, error) {
	query := `
		SELECT id, email, status, created_at
		FROM users
		ORDER BY created_at DESC`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Status,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateStatus changes the status of a specific user.
func (m *UserModel) UpdateStatus(id int64, status string) error {
	query := `
		UPDATE users
		SET status = $1
		WHERE id = $2`

	result, err := m.DB.Exec(query, status, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

// GetOrCreateByEmail implements robust "upsert" logic.
func (m *UserModel) GetOrCreateByEmail(email string) (*User, error) {
	query := `
		INSERT INTO users (email)
		VALUES ($1)
		ON CONFLICT (email) DO NOTHING`

	_, err := m.DB.Exec(query, email)
	if err != nil {
		return nil, err
	}

	return m.GetByEmail(email)
}

// GetByEmail retrieves a user by their email address.
func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT id, email, status, created_at
		FROM users
		WHERE email = $1`

	var user User

	err := m.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}
