// internal/data/admins.go
package data

import (
	"database/sql"
	"errors"
)

// AdminModel wraps a sql.DB connection pool.
type AdminModel struct {
	DB *sql.DB
}

// Insert adds a new admin record to the database.
func (m *AdminModel) Insert(username, passwordHash string) (*Admin, error) {
	query := `
		INSERT INTO admins (username, password_hash)
		VALUES ($1, $2)
		RETURNING id, created_at`

	admin := &Admin{
		Username:     username,
		PasswordHash: passwordHash,
	}

	err := m.DB.QueryRow(query, username, passwordHash).Scan(&admin.ID, &admin.CreatedAt)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

// GetByUsername retrieves an admin by their username.
func (m *AdminModel) GetByUsername(username string) (*Admin, error) {
	query := `
		SELECT id, username, password_hash, created_at
		FROM admins
		WHERE username = $1`

	var admin Admin

	err := m.DB.QueryRow(query, username).Scan(
		&admin.ID,
		&admin.Username,
		&admin.PasswordHash,
		&admin.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("admin not found")
		}
		return nil, err
	}

	return &admin, nil
}

// GetAll retrieves all admins from the database.
func (m *AdminModel) GetAll() ([]*Admin, error) {
	query := `
		SELECT id, username, created_at
		FROM admins
		ORDER BY created_at DESC`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []*Admin
	for rows.Next() {
		var admin Admin
		err := rows.Scan(
			&admin.ID,
			&admin.Username,
			&admin.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		admins = append(admins, &admin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return admins, nil
}

// --- NEW FUNCTION ---
// Delete removes an admin record from the database.
func (m *AdminModel) Delete(id int64) error {
	// Safety check: do not allow deleting the last admin.
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM admins").Scan(&count)
	if err != nil {
		return err
	}
	if count <= 1 {
		return errors.New("cannot delete the last admin account")
	}

	query := `DELETE FROM admins WHERE id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("admin record not found")
	}

	return nil
}
