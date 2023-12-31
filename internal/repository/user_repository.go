package repository

import (
	"bphn/artikel-hukum/internal/model"
	"context"
)

type userRepository struct {
	*Repository
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}
