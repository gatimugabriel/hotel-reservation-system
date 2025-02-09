package services

import (
	"context"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/config"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/repository"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"time"
)

// UserService : interface for user business logic
type UserService interface {
	Authenticate(ctx context.Context, req *entity.UserLoginRequest) (string, string, error)

	Create(ctx context.Context, req *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	UpdateProfile(ctx context.Context, userID string, user *entity.User) (*entity.User, error)
	DeactivateAccount(id uuid.UUID) error
	DeleteUser(ctx context.Context, userID string) error
}

type UserServiceImpl struct {
	userRepo    repository.UserRepository
	googleOAuth *oauth2.Config
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo:    userRepo,
		googleOAuth: config.GoogleOAuthConfig,
	}
}

func (u *UserServiceImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user.PasswordHash = string(hashedPassword)
	user.ID = uuid.New()

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) Authenticate(ctx context.Context, req *entity.UserLoginRequest) (string, string, error) {
	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", "", fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", "", fmt.Errorf("invalid credentials")
	}

	accessToken, refreshToken, err := utils.GenerateTokens((user.ID).String(), user.Role)

	return accessToken, refreshToken, nil
}

func (u *UserServiceImpl) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	return u.userRepo.GetByID(ctx, id)
}

func (u *UserServiceImpl) UpdateProfile(ctx context.Context, userID string, updates *entity.User) (*entity.User, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	currentUser, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update allowed fields
	currentUser.FirstName = updates.FirstName
	currentUser.LastName = updates.LastName
	currentUser.Phone = updates.Phone
	currentUser.UpdatedAt = time.Now()

	if err := u.userRepo.Update(ctx, currentUser); err != nil {
		return nil, err
	}

	return currentUser, nil
}

func (u *UserServiceImpl) DeactivateAccount(id uuid.UUID) error {
	user, err := u.userRepo.GetByID(context.Background(), id)
	if err != nil {
		return err
	}

	user.Status = "INACTIVE"
	user.UpdatedAt = time.Now()

	return u.userRepo.Update(context.Background(), user)
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	return u.userRepo.Delete(ctx, id)
}