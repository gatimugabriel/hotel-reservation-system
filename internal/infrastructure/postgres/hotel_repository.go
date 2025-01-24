package postgres

import (
	"database/sql"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel"
	"time"

	"github.com/google/uuid"
)

// HotelRepositoryImpl implements the hotel.Repository interface
type HotelRepositoryImpl struct {
	db *sql.DB
}

func NewHotelRepository(db *sql.DB) *HotelRepositoryImpl {
	return &HotelRepositoryImpl{db: db}
}

func (repo *HotelRepositoryImpl) GetByID(id uuid.UUID) (*hotel.Hotel, error) {
	query := `
		SELECT id, name, description, address, city, country,
		       latitude, longitude, contact_number, email,
		       owner_id, manager_id, is_active, created_at, updated_at, deleted_at
		FROM hotels 
		WHERE id = $1 AND deleted_at IS NULL`

	var h hotel.Hotel
	var deletedAt sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&h.ID, &h.Name, &h.Description, &h.Address, &h.City, &h.Country,
		&h.Latitude, &h.Longitude, &h.ContactNumber, &h.Email,
		&h.OwnerID, &h.ManagerID, &h.IsActive, &h.CreatedAt, &h.UpdatedAt, &deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("hotel not found: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error querying hotel: %v", err)
	}

	if deletedAt.Valid {
		h.DeletedAt = &deletedAt.Time
	}

	return &h, nil
}

func (repo *HotelRepositoryImpl) GetByCity(city string) ([]*hotel.Hotel, error) {
	query := `
		SELECT id, name, description, address, city, country,
		       latitude, longitude, contact_number, email,
		       owner_id, manager_id, is_active, created_at, updated_at, deleted_at
		FROM hotels 
		WHERE city ILIKE $1 AND deleted_at IS NULL`

	rows, err := r.db.Query(query, "%"+city+"%")
	if err != nil {
		return nil, fmt.Errorf("error querying hotels by city: %v", err)
	}
	defer rows.Close()

	var hotels []*hotel.Hotel

	for rows.Next() {
		var h hotel.Hotel
		var deletedAt sql.NullTime

		err := rows.Scan(
			&h.ID, &h.Name, &h.Description, &h.Address, &h.City, &h.Country,
			&h.Latitude, &h.Longitude, &h.ContactNumber, &h.Email,
			&h.OwnerID, &h.ManagerID, &h.IsActive, &h.CreatedAt, &h.UpdatedAt, &deletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning hotel row: %v", err)
		}

		if deletedAt.Valid {
			h.DeletedAt = &deletedAt.Time
		}

		hotels = append(hotels, &h)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating hotel rows: %v", err)
	}

	return hotels, nil
}

func (repo *HotelRepositoryImpl) Create(h *hotel.Hotel) error {
	query := `
		INSERT INTO hotels (
			id, name, description, address, city, country,
			latitude, longitude, contact_number, email,
			owner_id, manager_id, is_active, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	_, err := r.db.Exec(query,
		h.ID, h.Name, h.Description, h.Address, h.City, h.Country,
		h.Latitude, h.Longitude, h.ContactNumber, h.Email,
		h.OwnerID, h.ManagerID, h.IsActive, h.CreatedAt, h.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating hotel: %v", err)
	}

	return nil
}

func (repo *HotelRepositoryImpl) Update(h *hotel.Hotel) error {
	query := `
		UPDATE hotels 
		SET name = $2, description = $3, address = $4, city = $5, country = $6,
		    latitude = $7, longitude = $8, contact_number = $9, email = $10,
		    owner_id = $11, manager_id = $12, is_active = $13, updated_at = $14
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := r.db.Exec(query,
		h.ID, h.Name, h.Description, h.Address, h.City, h.Country,
		h.Latitude, h.Longitude, h.ContactNumber, h.Email,
		h.OwnerID, h.ManagerID, h.IsActive, time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error updating hotel: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("hotel not found or already deleted")
	}

	return nil
}

func (repo *HotelRepositoryImpl) Delete(id uuid.UUID) error {
	query := `
		UPDATE hotels 
		SET deleted_at = $2, is_active = false
		WHERE id = $1 AND deleted_at IS NULL`

	result, err := r.db.Exec(query, id, time.Now())
	if err != nil {
		return fmt.Errorf("error deleting hotel: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("hotel not found or already deleted")
	}

	return nil
}