package impl

import (
	domain "auth-service/internal/domain/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (a *UserRepositoryImpl) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	var user domain.User
	if err := a.db.WithContext(ctx).Preload("Roles").Where("id = ?", userID).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	a.db.Model(&domain.User{}).
		Select("password", "salt", "id", "status").
		Where("email = ?", email).
		Preload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("name")
		}).
		First(&user)
	return &user, nil
}

func (a *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	if err := a.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (a *UserRepositoryImpl) DeleteUser(ctx context.Context, userID string) error {
	if err := a.db.WithContext(ctx).Delete(&domain.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}

func (a *UserRepositoryImpl) GetUserDetailById(ctx context.Context, userID string) (*domain.User, error) {
	var user domain.User
	if err := a.db.WithContext(ctx).
		Preload("Profile").
		Preload("Roles").
		First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *UserRepositoryImpl) GetAllUser(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := a.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (a *UserRepositoryImpl) SetUserAccountStatus(ctx context.Context, userID string, status string) error {
	if err := a.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", userID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (a *UserRepositoryImpl) UpdateLastLogin(ctx context.Context, userID string) error {
	if err := a.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", userID).Update("last_login", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
