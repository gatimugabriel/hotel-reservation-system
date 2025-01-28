package repository

// import (
// 	repository2 "github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/repository"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"hotel-reservation/internal/domain/room"
// )

// func TestRoomRepository_Create(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewRoomRepository(db)

// 	hotelID := uuid.New()
// 	roomTypeID := uuid.New()

// 	testRoom := &room.Room{
// 		ID:          uuid.New(),
// 		RoomNumber:  "101",
// 		HotelID:     hotelID,
// 		RoomTypeID:  roomTypeID,
// 		FloorNumber: 1,
// 		IsAvailable: true,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	err := repo.Create(testRoom)
// 	assert.NoError(t, err)

// 	// Verify room was created
// 	saved, err := repo.GetByID(testRoom.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, testRoom.RoomNumber, saved.RoomNumber)
// 	assert.Equal(t, testRoom.HotelID, saved.HotelID)
// 	assert.Equal(t, testRoom.RoomTypeID, saved.RoomTypeID)
// }

// func TestRoomRepository_GetByHotelID(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewRoomRepository(db)
// 	hotelID := uuid.New()
// 	roomTypeID := uuid.New()

// 	// Create test rooms
// 	testRooms := []*room.Room{
// 		{
// 			ID:          uuid.New(),
// 			RoomNumber:  "101",
// 			HotelID:     hotelID,
// 			RoomTypeID:  roomTypeID,
// 			FloorNumber: 1,
// 			IsAvailable: true,
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 		},
// 		{
// 			ID:          uuid.New(),
// 			RoomNumber:  "102",
// 			HotelID:     hotelID,
// 			RoomTypeID:  roomTypeID,
// 			FloorNumber: 1,
// 			IsAvailable: true,
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 		},
// 	}

// 	for _, r := range testRooms {
// 		err := repo.Create(r)
// 		assert.NoError(t, err)
// 	}

// 	// Test GetByHotelID
// 	rooms, err := repo.GetByHotelID(hotelID)
// 	assert.NoError(t, err)
// 	assert.Len(t, rooms, 2)
// }

// func TestRoomRepository_GetAvailableRooms(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewRoomRepository(db)
// 	hotelID := uuid.New()
// 	roomTypeID := uuid.New()

// 	// Create test rooms
// 	room1 := &room.Room{
// 		ID:          uuid.New(),
// 		RoomNumber:  "101",
// 		HotelID:     hotelID,
// 		RoomTypeID:  roomTypeID,
// 		FloorNumber: 1,
// 		IsAvailable: true,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	err := repo.Create(room1)
// 	assert.NoError(t, err)

// 	// Test availability check
// 	now := time.Now()
// 	checkIn := now.Add(24 * time.Hour)
// 	checkOut := now.Add(48 * time.Hour)

// 	rooms, err := repo.CheckAvailability(hotelID, checkIn, checkOut)
// 	assert.NoError(t, err)
// 	assert.Len(t, rooms, 1)

// 	// Create a reservation for the room (would need reservation repository)
// 	// Then verify the room is no longer available for the same dates
// }

// func TestRoomRepository_Update(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewRoomRepository(db)
// 	hotelID := uuid.New()
// 	roomTypeID := uuid.New()

// 	testRoom := &room.Room{
// 		ID:          uuid.New(),
// 		RoomNumber:  "101",
// 		HotelID:     hotelID,
// 		RoomTypeID:  roomTypeID,
// 		FloorNumber: 1,
// 		IsAvailable: true,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	// Create test room
// 	err := repo.Create(testRoom)
// 	assert.NoError(t, err)

// 	// Update room
// 	testRoom.RoomNumber = "102"
// 	testRoom.FloorNumber = 2
// 	err = repo.Update(testRoom)
// 	assert.NoError(t, err)

// 	// Verify update
// 	updated, err := repo.GetByID(testRoom.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "102", updated.RoomNumber)
// 	assert.Equal(t, 2, updated.FloorNumber)
// }

// func TestRoomRepository_Delete(t *testing.T) {
// 	db, cleanup := repository2.setupTestDB(t)
// 	defer cleanup()

// 	repo := NewRoomRepository(db)
// 	hotelID := uuid.New()
// 	roomTypeID := uuid.New()

// 	testRoom := &room.Room{
// 		ID:          uuid.New(),
// 		RoomNumber:  "101",
// 		HotelID:     hotelID,
// 		RoomTypeID:  roomTypeID,
// 		FloorNumber: 1,
// 		IsAvailable: true,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	// Create test room
// 	err := repo.Create(testRoom)
// 	assert.NoError(t, err)

// 	// Delete room
// 	err = repo.Delete(testRoom.ID)
// 	assert.NoError(t, err)

// 	// Verify room is not found
// 	_, err = repo.GetByID(testRoom.ID)
// 	assert.Error(t, err)
// }