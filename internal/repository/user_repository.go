package repository

import (
	"bphn/artikel-hukum/internal/dto"
	"bphn/artikel-hukum/internal/model"
	"context"
)

type userRepository struct {
	*Repository
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	FindAll(ctx context.Context, query dto.ListQuery) (*dto.ListQueryResult[model.User], error)
	FindById(ctx context.Context, Id uint) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Delete(ctx context.Context, Id uint) error
}
