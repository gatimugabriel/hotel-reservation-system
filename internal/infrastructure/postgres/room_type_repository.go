package postgres

import (
	"database/sql"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room"
	"time"

	"github.com/google/uuid"
)

// RoomTypeRepositoryImpl implements the RoomTypeRepository interface
type RoomTypeRepositoryImpl struct {
	db *sql.DB
}

func NewRoomTypeRepository(db *sql.DB) *RoomTypeRepositoryImpl {
	return &RoomTypeRepositoryImpl{db: db}
}

func (r *RoomTypeRepositoryImpl) GetByID(id uuid.UUID) (*room.RoomType, error) {
	query := `
		SELECT id, name, description, base_price, max_occupancy,
		       num_beds, bed_type, square_meters,
		       created_at, updated_at, deleted_at
		FROM room_types 
		WHERE id = $1 AND deleted_at IS NULL`

	var rt room.RoomType
	var deletedAt sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&rt.ID, &rt.Name, &rt.Description, &rt.BasePrice, &rt.MaxOccupancy,
		&rt.NumBeds, &rt.BedType, &rt.SquareMeters,
		&rt.CreatedAt, &rt.UpdatedAt, &deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("room type not found: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error querying room type: %v", err)
	}

	if deletedAt.Valid {
		rt.DeletedAt = &deletedAt.Time
	}

	return &rt, nil
}

func (r *RoomTypeRepositoryImpl) Create(rt *room.RoomType) error {
	query := `
		INSERT INTO room_types (
			id, name, description, base_price, max_occupancy,
			num_beds, bed_type, square_meters,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.Exec(query,
		rt.ID, rt.Name, rt.Description, rt.BasePrice, rt.MaxOccupancy,
		rt.NumBeds, rt.BedType, rt.SquareMeters,
		rt.CreatedAt, rt.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating room type: %v", err)
	}

	return nil
}

func (r *RoomTypeRepositoryImpl) Update(rt *room.RoomType) error {
	query := `
		UPDATE room_types 
		SET name = $2, description = $3, base_price = $4,
		    max_occupancy = $5, num_beds = $6, bed_type = $7,
		    square_meters = $8, updated_at = $9
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := r.db.Exec(query,
		rt.ID, rt.Name, rt.Description, rt.BasePrice,
		rt.MaxOccupancy, rt.NumBeds, rt.BedType,
		rt.SquareMeters, time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error updating room type: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("room type not found or already deleted")
	}

	return nil
}

func (r *RoomTypeRepositoryImpl) Delete(id uuid.UUID) error {
	query := `
		UPDATE room_types 
		SET deleted_at = $2
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := r.db.Exec(query, id, time.Now())
	if err != nil {
		return fmt.Errorf("error deleting room type: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("room type not found or already deleted")
	}

	return nil
}