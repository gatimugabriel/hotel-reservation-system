package repository

// import (
// 	"database/sql"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"hotel-reservation/internal/domain/user"
// )

// func setupTestDB(t *testing.T) (*sql.DB, func()) {
// 	// Use test database connection string
// 	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/hoteldb_test?sslmode=disable")
// 	if err != nil {
// 		t.Fatalf("Failed to connect to test database: %v", err)
// 	}

// 	// Run migrations for test database
// 	// In a real implementation, you would use proper migration tools

// 	return db, func() {
// 		// Cleanup function
// 		db.Close()
// 	}
// }

// func TestUserRepository_Create(t *testing.T) {
// 	db, cleanup := setupTestDB(t)
// 	defer cleanup()

// 	repo := NewUserRepository(db)

// 	testUser := &user.User{
// 		ID:            uuid.New(),
// 		Email:         "test@example.com",
// 		Phone:         "+1234567890",
// 		FirstName:     "Test",
// 		LastName:      "User",
// 		PasswordHash:  "hashed_password",
// 		Role:          user.RoleGuest,
// 		IsActive:      true,
// 		EmailVerified: false,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}

// 	err := repo.Create(testUser)
// 	assert.NoError(t, err)

// 	// Verify user was created
// 	saved, err := repo.GetByID(testUser.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, testUser.Email, saved.Email)
// 	assert.Equal(t, testUser.Phone, saved.Phone)
// 	assert.Equal(t, testUser.FirstName, saved.FirstName)
// 	assert.Equal(t, testUser.LastName, saved.LastName)
// }

// func TestUserRepository_GetByEmail(t *testing.T) {
// 	db, cleanup := setupTestDB(t)
// 	defer cleanup()

// 	repo := NewUserRepository(db)

// 	testUser := &user.User{
// 		ID:            uuid.New(),
// 		Email:         "test@example.com",
// 		Phone:         "+1234567890",
// 		FirstName:     "Test",
// 		LastName:      "User",
// 		PasswordHash:  "hashed_password",
// 		Role:          user.RoleGuest,
// 		IsActive:      true,
// 		EmailVerified: false,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}

// 	// Create test user
// 	err := repo.Create(testUser)
// 	assert.NoError(t, err)

// 	// Test GetByEmail
// 	found, err := repo.GetByEmail(testUser.Email)
// 	assert.NoError(t, err)
// 	assert.Equal(t, testUser.ID, found.ID)
// 	assert.Equal(t, testUser.Email, found.Email)

// 	// Test non-existent email
// 	_, err = repo.GetByEmail("nonexistent@example.com")
// 	assert.Error(t, err)
// }

// func TestUserRepository_Update(t *testing.T) {
// 	db, cleanup := setupTestDB(t)
// 	defer cleanup()

// 	repo := NewUserRepository(db)

// 	testUser := &user.User{
// 		ID:            uuid.New(),
// 		Email:         "test@example.com",
// 		Phone:         "+1234567890",
// 		FirstName:     "Test",
// 		LastName:      "User",
// 		PasswordHash:  "hashed_password",
// 		Role:          user.RoleGuest,
// 		IsActive:      true,
// 		EmailVerified: false,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}

// 	// Create test user
// 	err := repo.Create(testUser)
// 	assert.NoError(t, err)

// 	// Update user
// 	testUser.FirstName = "Updated"
// 	testUser.LastName = "Name"
// 	err = repo.Update(testUser)
// 	assert.NoError(t, err)

// 	// Verify update
// 	updated, err := repo.GetByID(testUser.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Updated", updated.FirstName)
// 	assert.Equal(t, "Name", updated.LastName)
// }

// func TestUserRepository_Delete(t *testing.T) {
// 	db, cleanup := setupTestDB(t)
// 	defer cleanup()

// 	repo := NewUserRepository(db)

// 	testUser := &user.User{
// 		ID:            uuid.New(),
// 		Email:         "test@example.com",
// 		Phone:         "+1234567890",
// 		FirstName:     "Test",
// 		LastName:      "User",
// 		PasswordHash:  "hashed_password",
// 		Role:          user.RoleGuest,
// 		IsActive:      true,
// 		EmailVerified: false,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}

// 	// Create test user
// 	err := repo.Create(testUser)
// 	assert.NoError(t, err)

// 	// Delete user
// 	err = repo.Delete(testUser.ID)
// 	assert.NoError(t, err)

// 	// Verify user is not found
// 	_, err = repo.GetByID(testUser.ID)
// 	assert.Error(t, err)
// }