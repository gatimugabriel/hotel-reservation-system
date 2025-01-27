package repository

// import (
// 	repository2 "github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/repository"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"hotel-reservation/internal/domain/room"
// )

// func TestRoomTypeRepository_Create(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := repository2.NewRoomTypeRepository(db)

// 	testRoomType := &room.RoomType{
// 		ID:           uuid.New(),
// 		Name:         "Deluxe Room",
// 		Description:  "A luxurious room with ocean view",
// 		BasePrice:    299.99,
// 		MaxOccupancy: 2,
// 		NumBeds:      1,
// 		BedType:      "King",
// 		SquareMeters: 45.5,
// 		CreatedAt:    time.Now(),
// 		UpdatedAt:    time.Now(),
// 	}

// 	err := repo.Create(testRoomType)
// 	assert.NoError(t, err)

// 	// Verify room type was created
// 	saved, err := repo.GetByID(testRoomType.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, testRoomType.Name, saved.Name)
// 	assert.Equal(t, testRoomType.BasePrice, saved.BasePrice)
// 	assert.Equal(t, testRoomType.BedType, saved.BedType)
// }

// func TestRoomTypeRepository_Update(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := repository2.NewRoomTypeRepository(db)

// 	testRoomType := &room.RoomType{
// 		ID:           uuid.New(),
// 		Name:         "Standard Room",
// 		Description:  "A comfortable standard room",
// 		BasePrice:    199.99,
// 		MaxOccupancy: 2,
// 		NumBeds:      1,
// 		BedType:      "Queen",
// 		SquareMeters: 35.0,
// 		CreatedAt:    time.Now(),
// 		UpdatedAt:    time.Now(),
// 	}

// 	// Create test room type
// 	err := repo.Create(testRoomType)
// 	assert.NoError(t, err)

// 	// Update room type
// 	testRoomType.Name = "Premium Room"
// 	testRoomType.BasePrice = 249.99
// 	err = repo.Update(testRoomType)
// 	assert.NoError(t, err)

// 	// Verify update
// 	updated, err := repo.GetByID(testRoomType.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Premium Room", updated.Name)
// 	assert.Equal(t, 249.99, updated.BasePrice)
// }

// func TestRoomTypeRepository_Delete(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := repository2.NewRoomTypeRepository(db)

// 	testRoomType := &room.RoomType{
// 		ID:           uuid.New(),
// 		Name:         "Standard Room",
// 		Description:  "A comfortable standard room",
// 		BasePrice:    199.99,
// 		MaxOccupancy: 2,
// 		NumBeds:      1,
// 		BedType:      "Queen",
// 		SquareMeters: 35.0,
// 		CreatedAt:    time.Now(),
// 		UpdatedAt:    time.Now(),
// 	}

// 	// Create test room type
// 	err := repo.Create(testRoomType)
// 	assert.NoError(t, err)

// 	// Delete room type
// 	err = repo.Delete(testRoomType.ID)
// 	assert.NoError(t, err)

// 	// Verify room type is not found
// 	_, err = repo.GetByID(testRoomType.ID)
// 	assert.Error(t, err)
// }