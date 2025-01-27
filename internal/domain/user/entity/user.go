package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleGuest   Role = "GUEST"
	RoleStaff   Role = "STAFF"
	RoleManager Role = "MANAGER"
	RoleAdmin   Role = "ADMIN"
	RoleOwner   Role = "OWNER"
)

type User struct {
	ID            uuid.UUID
	Email         string
	Phone         string
	FirstName     string
	LastName      string
	PasswordHash  string
	Role          Role
	IsActive      bool
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}