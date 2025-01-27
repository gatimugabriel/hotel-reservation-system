package repository

// import (
// 	resitory "github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/repository"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"hotel-reservation/internal/domain/hotel"
// )

// func TestHotelRepository_Create(t *testing.T) {
// 	db, cleanup := resitory.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewHotelRepository(db)

// 	testHotel := &hotel.Hotel{
// 		ID:            uuid.New(),
// 		Name:          "Test Hotel",
// 		Description:   "A test hotel",
// 		Address:       "123 Test St",
// 		City:          "Test City",
// 		Country:       "Test Country",
// 		Latitude:      40.7128,
// 		Longitude:     -74.0060,
// 		ContactNumber: "+1234567890",
// 		Email:         "hotel@test.com",
// 		OwnerID:       uuid.New(),
// 		ManagerID:     uuid.New(),
// 		IsActive:      true,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}

// 	err := repo.Create(testHotel)
// 	assert.NoError(t, err)

// 	// Verify hotel was created
// 	saved, err := repo.GetByID(testHotel.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, testHotel.Name, saved.Name)
// 	assert.Equal(t, testHotel.City, saved.City)
// 	assert.Equal(t, testHotel.Email, saved.Email)
// }

// func TestHotelRepository_GetByCity(t *testing.T) {
// 	db, cleanup := resitory.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewHotelRepository(db)

// 	// Create test hotels in different cities
// 	testHotels := []*hotel.Hotel{
// 		{
// 			ID:        uuid.New(),
// 			Name:      "Hotel New York",
// 			City:      "New York",
// 			Country:   "USA",
// 			IsActive:  true,
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		{
// 			ID:        uuid.New(),
// 			Name:      "Hotel Paris",
// 			City:      "Paris",
// 			Country:   "France",
// 			IsActive:  true,
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 	}

// 	for _, h := range testHotels {
// 		err := repo.Create(h)
// 		assert.NoError(t, err)
// 	}

// 	// Test GetByCity
// 	nyHotels, err := repo.GetByCity("New York")
// 	assert.NoError(t, err)
// 	assert.Len(t, nyHotels, 1)
// 	assert.Equal(t, "Hotel New York", nyHotels[0].Name)

// 	// Test no results
// 	noHotels, err := repo.GetByCity("London")
// 	assert.NoError(t, err)
// 	assert.Len(t, noHotels, 0)
// }

// func TestHotelRepository_Update(t *testing.T) {
// 	db, cleanup := resitory.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewHotelRepository(db)

// 	testHotel := &hotel.Hotel{
// 		ID:        uuid.New(),
// 		Name:      "Original Name",
// 		City:      "Original City",
// 		Country:   "Original Country",
// 		IsActive:  true,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	// Create test hotel
// 	err := repo.Create(testHotel)
// 	assert.NoError(t, err)

// 	// Update hotel
// 	testHotel.Name = "Updated Name"
// 	testHotel.City = "Updated City"
// 	err = repo.Update(testHotel)
// 	assert.NoError(t, err)

// 	// Verify update
// 	updated, err := repo.GetByID(testHotel.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Updated Name", updated.Name)
// 	assert.Equal(t, "Updated City", updated.City)
// }

// func TestHotelRepository_Delete(t *testing.T) {
// 	db, cleanup := resitory.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewHotelRepository(db)

// 	testHotel := &hotel.Hotel{
// 		ID:        uuid.New(),
// 		Name:      "Test Hotel",
// 		City:      "Test City",
// 		Country:   "Test Country",
// 		IsActive:  true,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	// Create test hotel
// 	err := repo.Create(testHotel)
// 	assert.NoError(t, err)

// 	// Delete hotel
// 	err = repo.Delete(testHotel.ID)
// 	assert.NoError(t, err)

// 	// Verify hotel is not found
// 	_, err = repo.GetByID(testHotel.ID)
// 	assert.Error(t, err)
// }