package postgres

import (
	"database/sql"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room"
	"time"

	"github.com/google/uuid"
)

// RoomRepositoryImpl implements the RoomRepository interface
type RoomRepositoryImpl struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{db: db}
}

func (r *RoomRepositoryImpl) GetByID(id uuid.UUID) (*room.Room, error) {
	query := `
		SELECT id, room_number, hotel_id, room_type_id, floor_number,
		       is_available, is_maintenance, created_at, updated_at, deleted_at
		FROM rooms 
		WHERE id = $1 AND deleted_at IS NULL`

	var rm room.Room
	var deletedAt sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&rm.ID, &rm.RoomNumber, &rm.HotelID, &rm.RoomTypeID, &rm.FloorNumber,
		&rm.IsAvailable, &rm.IsMaintenance, &rm.CreatedAt, &rm.UpdatedAt, &deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("room not found: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error querying room: %v", err)
	}

	if deletedAt.Valid {
		rm.DeletedAt = &deletedAt.Time
	}

	return &rm, nil
}

func (r *RoomRepositoryImpl) GetByHotelID(hotelID uuid.UUID) ([]*room.Room, error) {
	query := `
		SELECT id, room_number, hotel_id, room_type_id, floor_number,
		       is_available, is_maintenance, created_at, updated_at, deleted_at
		FROM rooms 
		WHERE hotel_id = $1 AND deleted_at IS NULL
		ORDER BY floor_number, room_number`

	rows, err := r.db.Query(query, hotelID)
	if err != nil {
		return nil, fmt.Errorf("error querying rooms: %v", err)
	}
	defer rows.Close()

	var rooms []*room.Room

	for rows.Next() {
		var rm room.Room
		var deletedAt sql.NullTime

		err := rows.Scan(
			&rm.ID, &rm.RoomNumber, &rm.HotelID, &rm.RoomTypeID, &rm.FloorNumber,
			&rm.IsAvailable, &rm.IsMaintenance, &rm.CreatedAt, &rm.UpdatedAt, &deletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning room row: %v", err)
		}

		if deletedAt.Valid {
			rm.DeletedAt = &deletedAt.Time
		}

		rooms = append(rooms, &rm)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating room rows: %v", err)
	}

	return rooms, nil
}

func (r *RoomRepositoryImpl) GetAvailableRooms(hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*room.Room, error) {
	query := `
		SELECT r.id, r.room_number, r.hotel_id, r.room_type_id, r.floor_number,
		       r.is_available, r.is_maintenance, r.created_at, r.updated_at, r.deleted_at
		FROM rooms r
		WHERE r.hotel_id = $1
		  AND r.deleted_at IS NULL
		  AND r.is_available = true
		  AND r.is_maintenance = false
		  AND NOT EXISTS (
			SELECT 1 FROM reservations res
			WHERE res.room_id = r.id
			  AND res.status != 'CANCELLED'
			  AND (
				(res.check_in_date <= $2 AND res.check_out_date > $2)
				OR
				(res.check_in_date < $3 AND res.check_out_date >= $3)
				OR
				(res.check_in_date >= $2 AND res.check_out_date <= $3)
			  )
		  )
		ORDER BY r.floor_number, r.room_number`

	rows, err := r.db.Query(query, hotelID, checkIn, checkOut)
	if err != nil {
		return nil, fmt.Errorf("error querying available rooms: %v", err)
	}
	defer rows.Close()

	var rooms []*room.Room

	for rows.Next() {
		var rm room.Room
		var deletedAt sql.NullTime

		err := rows.Scan(
			&rm.ID, &rm.RoomNumber, &rm.HotelID, &rm.RoomTypeID, &rm.FloorNumber,
			&rm.IsAvailable, &rm.IsMaintenance, &rm.CreatedAt, &rm.UpdatedAt, &deletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning room row: %v", err)
		}

		if deletedAt.Valid {
			rm.DeletedAt = &deletedAt.Time
		}

		rooms = append(rooms, &rm)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating room rows: %v", err)
	}

	return rooms, nil
}

func (r *RoomRepositoryImpl) Create(rm *room.Room) error {
	query := `
		INSERT INTO rooms (
			id, room_number, hotel_id, room_type_id, floor_number,
			is_available, is_maintenance, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.db.Exec(query,
		rm.ID, rm.RoomNumber, rm.HotelID, rm.RoomTypeID, rm.FloorNumber,
		rm.IsAvailable, rm.IsMaintenance, rm.CreatedAt, rm.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating room: %v", err)
	}

	return nil
}

func (r *RoomRepositoryImpl) Update(rm *room.Room) error {
	query := `
		UPDATE rooms 
		SET room_number = $2, hotel_id = $3, room_type_id = $4,
		    floor_number = $5, is_available = $6, is_maintenance = $7,
		    updated_at = $8
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := r.db.Exec(query,
		rm.ID, rm.RoomNumber, rm.HotelID, rm.RoomTypeID,
		rm.FloorNumber, rm.IsAvailable, rm.IsMaintenance, time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error updating room: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("room not found or already deleted")
	}

	return nil
}

func (r *RoomRepositoryImpl) Delete(id uuid.UUID) error {
	query := `
		UPDATE rooms 
		SET deleted_at = $2, is_available = false
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := r.db.Exec(query, id, time.Now())
	if err != nil {
		return fmt.Errorf("error deleting room: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("room not found or already deleted")
	}

	return nil
}