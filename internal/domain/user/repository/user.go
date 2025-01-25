package repository

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"github.com/google/uuid"
)

// UserRepository : interface for interacting with database & user data persistence
type UserRepository interface {
	GetByID(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uuid.UUID) error
}

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *database.Service
}

func NewUserRepository(db *database.Service) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) GetByID(id uuid.UUID) (*entity.User, error) {
	var u entity.User

	return &u, nil
}

func (repo *UserRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var u entity.User
	return &u, nil
}

func (repo *UserRepositoryImpl) Create(u *entity.User) error {

	return nil
}

func (repo *UserRepositoryImpl) Update(u *entity.User) error {
	return nil
}

func (repo *UserRepositoryImpl) Delete(id uuid.UUID) error {
	return nil
}