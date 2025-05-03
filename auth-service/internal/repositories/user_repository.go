package repositories

import (
	domain "auth-service/internal/domain/models"
	"context"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserDetailById(ctx context.Context, userID string) (*domain.User, error)
	DeleteUser(ctx context.Context, userID string) error
	Create(ctx context.Context, user *domain.User) error
	GetAllUser(context.Context) ([]domain.User, error)
	UpdateLastLogin(ctx context.Context, userID string) error
	SetUserAccountStatus(ctx context.Context, userID string, status string) error
}
