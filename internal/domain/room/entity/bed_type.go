package entity

import "github.com/google/uuid"

type TypeOfBed string

const (
	Single TypeOfBed = "SINGLE"
	Double TypeOfBed = "DOUBLE"
	Queen  TypeOfBed = "QUEEN"
	King   TypeOfBed = "KING"
)

type BedType struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	TypeOfBed   TypeOfBed `gorm:"unique" json:"type_of_bed" validate:"required"`
	Description string    `json:"description"`
}