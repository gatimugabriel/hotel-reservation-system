package postgres

import (
	"database/sql"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) GetByID(id uuid.UUID) (*user.User, error) {
	query := `
		SELECT id, email, phone, first_name, last_name, password_hash, 
		       role, is_active, email_verified, created_at, updated_at, deleted_at
		FROM users 
		WHERE id = $1 AND deleted_at IS NULL`

	var u user.User
	var role string
	var deletedAt sql.NullTime

	err := repo.db.QueryRow(query, id).Scan(
		&u.ID, &u.Email, &u.Phone, &u.FirstName, &u.LastName, &u.PasswordHash,
		&role, &u.IsActive, &u.EmailVerified, &u.CreatedAt, &u.UpdatedAt, &deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error querying user: %v", err)
	}

	u.Role = user.Role(role)
	if deletedAt.Valid {
		u.DeletedAt = &deletedAt.Time
	}

	return &u, nil
}

func (repo *UserRepositoryImpl) GetByEmail(email string) (*user.User, error) {
	query := `
		SELECT id, email, phone, first_name, last_name, password_hash, 
		       role, is_active, email_verified, created_at, updated_at, deleted_at
		FROM users 
		WHERE email = $1 AND deleted_at IS NULL`

	var u user.User
	var role string
	var deletedAt sql.NullTime

	err := repo.db.QueryRow(query, email).Scan(
		&u.ID, &u.Email, &u.Phone, &u.FirstName, &u.LastName, &u.PasswordHash,
		&role, &u.IsActive, &u.EmailVerified, &u.CreatedAt, &u.UpdatedAt, &deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error querying user: %v", err)
	}

	u.Role = user.Role(role)
	if deletedAt.Valid {
		u.DeletedAt = &deletedAt.Time
	}

	return &u, nil
}

func (repo *UserRepositoryImpl) Create(u *user.User) error {
	query := `
		INSERT INTO users (
			id, email, phone, first_name, last_name, password_hash,
			role, is_active, email_verified, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := repo.db.Exec(query,
		u.ID, u.Email, u.Phone, u.FirstName, u.LastName, u.PasswordHash,
		u.Role, u.IsActive, u.EmailVerified, u.CreatedAt, u.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (repo *UserRepositoryImpl) Update(u *user.User) error {
	query := `
		UPDATE users 
		SET email = $2, phone = $3, first_name = $4, last_name = $5,
		    password_hash = $6, role = $7, is_active = $8, 
		    email_verified = $9, updated_at = $10
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := repo.db.Exec(query,
		u.ID, u.Email, u.Phone, u.FirstName, u.LastName, u.PasswordHash,
		u.Role, u.IsActive, u.EmailVerified, time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found or already deleted")
	}

	return nil
}

func (repo *UserRepositoryImpl) Delete(id uuid.UUID) error {
	query := `
		UPDATE users 
		SET deleted_at = $2, is_active = false
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := repo.db.Exec(query, id, time.Now())
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found or already deleted")
	}

	return nil
}