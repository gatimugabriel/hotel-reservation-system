package services

import (
	"context"
	"github.com/gatimugabriel/hotel-reservation-system/internal/config"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/repository"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// UserService : interface for user business logic
type UserService interface {
	Create(ctx context.Context, req *entity.User) (entity.User, error)
	CreateOrGetWthGoogleOauth(ctx context.Context, token *oauth2.Token) (*entity.User, error)

	Authenticate(ctx context.Context, req *entity.UserLoginRequest) (string, string, error)

	GetUser(ctx context.Context, userID string) (*entity.User, error)
	UpdateProfile(ctx context.Context, userID string, user *entity.User) (*entity.User, error)
	DeactivateAccount(id uuid.UUID) error
	DeleteUser(ctx context.Context, userID string) error
	GetUsers(ctx context.Context) (interface{}, interface{})
}

// UserServiceImpl implements the user repository
type UserServiceImpl struct {
	repo        repository.UserRepository
	googleOAuth *oauth2.Config
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo:        userRepo,
		googleOAuth: config.GoogleOAuthConfig,
	}
}

func (u UserServiceImpl) Create(ctx context.Context, req *entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) CreateOrGetWthGoogleOauth(ctx context.Context, token *oauth2.Token) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) Authenticate(ctx context.Context, req *entity.UserLoginRequest) (string, string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) UpdateProfile(ctx context.Context, userID string, user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) DeactivateAccount(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUsers(ctx context.Context) (interface{}, interface{}) {
	//TODO implement me
	panic("implement me")
}